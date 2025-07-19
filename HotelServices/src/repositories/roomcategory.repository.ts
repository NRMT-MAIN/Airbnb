import BaseRepository from "./base.repository";
import RoomCategory from "../db/models/roomcategory.model";
import { NotFoundError } from "../utils/Error/app.error";

class RoomCategoryRepository extends BaseRepository<RoomCategory>{
    constructor(){
        super(RoomCategory) ; 
    }

    async findAllByHotelId(hotel_id : number){
        const roomCategories = await this.model.findAll({
            where : {
                hotel_id : hotel_id , 
                deleted_at : null
            }
        })

        if(!roomCategories || roomCategories.length == 0){
            throw new NotFoundError(`No room category with hotelId : ${hotel_id} found !`) ;
        }
    }
}

export default RoomCategoryRepository ; 