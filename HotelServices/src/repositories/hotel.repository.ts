import logger from "../config/logger.config";
import Hotel from "../db/models/hotel.model";
import { NotFoundError } from "../utils/Error/app.error";
import BaseRepository from "./base.repository";


class HotelRepository extends BaseRepository<Hotel> { 
    constructor(){
        super(Hotel) ; 
    }

    async getAll() {
        const hotels = await Hotel.findAll({
            where : {
                deletedAt : null
            }
        });

        if(!hotels) {
            logger.error(`No hotels found`);
            throw new NotFoundError(`No hotels found`);
        }

        logger.info(`Hotels found: ${hotels.length}`);
        return hotels;
    }

    async softDelete(id : number) {
        const hotel = await Hotel.findByPk(id) ; 

        if(!hotel) {
            logger.error(`Hotels not found ${id}`);
            throw new NotFoundError(`Hotel with id ${id} not found`);
        }

        hotel.deletedAt = new Date() ; 
        await hotel.save() ; //save the changes to the database
        logger.info(`Hotel Soft Deleted : ${hotel.id}`) ; 
        return true ; 
    }
}

export default HotelRepository ; 