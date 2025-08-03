import { CreationAttributes } from "sequelize";
import Rooms from "../db/models/rooms.model";
import BaseRepository from "./base.repository";

class RoomRepository extends BaseRepository<Rooms>{
    constructor(){
        super(Rooms) ; 
    }

    async bulkCreate(rooms: CreationAttributes<Rooms>[]){
        return await this.model.bulkCreate(rooms)
    }

    async findByRoomCategoryIdAndDate(room_category_id: number, currentDate: Date) {
        return await this.model.findOne({
            where: {
                room_category_id,
                date_of_availability: currentDate,
                deleted_at: null
            }
        })
    }

    async findLatestDateByRoomCategoryId(room_category_id: number): Promise<Date | null> {
        const result = await this.model.findOne({
            where: {
                room_category_id ,
                deleted_at: null
            },
            attributes: ['date_of_availability'],
            order: [['date_of_availability', 'DESC']]
        });
        
        return result ? result.date_of_availability : null;
    }

    async findLatestDatesForAllCategories(): Promise<Array<{room_category_id: number, latestDate: Date}>> {
        const results = await this.model.findAll({
            where: {
                deleted_at: null
            },
            attributes: [
                'room_category_id',
                [this.model.sequelize!.fn('MAX', this.model.sequelize!.col('date_of_availability')), 'latestDate']
            ],
            group: ['room_category_id'],
            raw: true
        });
        
        return results.map((result: any) => ({
            room_category_id: result.room_category_id,
            latestDate: new Date(result.latestDate)
        }));
    }
}

export default RoomRepository ; 