import { CreateRoomCategoryDTO } from "../dto/roomcategory.dto";
import HotelRepository from "../repositories/hotel.repository";
import RoomCategoryRepository from "../repositories/roomcategory.repository";
import { NotFoundError } from "../utils/Error/app.error";

const roomCategoryRepository = new RoomCategoryRepository() ; 
const hotelRepository = new HotelRepository() ; 

export async function createRoomCategoryService(createRoomCategoryDTO : CreateRoomCategoryDTO){
    const roomCategory = await roomCategoryRepository.create(createRoomCategoryDTO) ; 
    return roomCategory ; 
}

export async function getRoomCategoryByIdService(id : number){
    const roomCategory = await roomCategoryRepository.findById(id) ; 
    return roomCategory ; 
}

export async function getAllRoomCategoryByHotelIdService(id : number){
    const hotel = await hotelRepository.findById(id) ; 

    if(!hotel){
        throw new NotFoundError(`Hotel with id : ${id} not found!`) ; 
    }

    const roomCategories  = await roomCategoryRepository.findAllByHotelId(id) ; 
    return roomCategories ; 
}

export async function deleteRoomCategoryService(id : number){
    const roomCategory = await roomCategoryRepository.findById(id) ; 

    if(!roomCategory){
        throw new NotFoundError(`Room Category with id : ${id} not found`) ; 
    }

    await roomCategoryRepository.delete({id}) ;
    return true ;  
}