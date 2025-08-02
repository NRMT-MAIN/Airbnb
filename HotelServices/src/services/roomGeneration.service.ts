import { CreationAttributes, QueryTypes } from "sequelize";
import RoomCategory from "../db/models/roomcategory.model";
import sequelize from "../db/models/sequelize";
import RoomRepository from "../repositories/room.repository";
import Rooms from "../db/models/rooms.model";
import { RoomGenerationJob } from "../dto/roomGeneration.dto";
import RoomCategoryRepository from "../repositories/roomcategory.repository";
import { ForbiddenError, NotFoundError } from "../utils/Error/app.error";
import logger from "../config/logger.config";


const roomRepository = new RoomRepository() ; 
const roomCategoryRepository = new RoomCategoryRepository() ; 


export async function generateRooms(jobData : RoomGenerationJob){
    let totalRoomCreated = 0 ; 
    let totalDatesCreated = 0 ; 
    
    const roomCategory = await roomCategoryRepository.findById(jobData.roomCategoryId)

    if(!roomCategory){
        throw new NotFoundError("Category not found")
    }

    const startDate = new Date(jobData.startDate);
    const endDate = new Date(jobData.endDate);

    if(startDate >= endDate){
        throw new ForbiddenError("Start date cannot be greater than end date")
    }

    if(startDate < new Date()){
        throw new ForbiddenError("Start date must be in the future")
    }
    const totalDays = Math.ceil((endDate.getTime() - startDate.getTime()) / (1000*60*60*24));

    logger.info(`Generating rooms for ${totalDays} days`);

    const currentDate = new Date(startDate)
    while(currentDate < endDate){
        const batchEndDate = new Date(currentDate)

        batchEndDate.setDate(currentDate.getDate() + jobData.batchSize)

        if(batchEndDate > endDate ) {
            batchEndDate.setTime(endDate.getTime());
        }

        const batchResult = await processBatchDate(roomCategory , currentDate ,batchEndDate ,jobData.priceOverride)

        totalRoomCreated += batchResult.roomsCreated ; 
        totalDatesCreated += batchResult.datesProcessed ; 

        currentDate.setTime(batchEndDate.getTime()) ; 
    }

    return {
        totalRoomCreated , 
        totalDatesCreated
    }
}



export async function processBatchDate(
    roomCategory : RoomCategory , startDate : Date , endDate : Date , priceOverride? : number){
        const query : string = `SELECT * FROM ROOMS
                                    WHERE ROOM_CATEGORY_ID = ?
                                    AND HOTEL_ID = ? 
                                    AND DATE_OF_AVAILABILITY BETWEEN ? AND ?
                                    AND ROOM_NO BETWEEN 1 AND ?;`

        const results = await sequelize.query(query , {
            replacements : [roomCategory.id , roomCategory.hotel_id , startDate , endDate , roomCategory.roomCount] , 
            type : QueryTypes.SELECT , 
            model : Rooms , 
            mapToModel : true
        })

        const roomsToCreate : CreationAttributes<Rooms>[] = []
        let roomsCreated = 0;
        let datesProcessed = 0;
    
        const currentDate = new Date(startDate)
        while(currentDate <= endDate){
            for(let room_no = 1 ; room_no <= roomCategory.roomCount ; room_no++){
                const existingRoom = results.find(room => 
                    room.date_of_availability == currentDate && room.room_no == room_no
                )

                if(!existingRoom){
                    roomsToCreate.push({
                        hotel_id : roomCategory.hotel_id , 
                        room_category_id: roomCategory.id,
                        date_of_availability: new Date(currentDate),
                        price: priceOverride || roomCategory.price,
                        created_at: new Date(),
                        updated_at: new Date(),
                        deleted_at: null,
                        room_no : room_no
                    })
                }
            }
            currentDate.setDate(currentDate.getDate() + 1)
            datesProcessed++ ; 
        }

        if(roomsToCreate.length > 0){
            await roomRepository.bulkCreate(roomsToCreate)
            roomsCreated += roomsToCreate.length
        }
    return {
        datesProcessed , 
        roomsCreated
    }
}