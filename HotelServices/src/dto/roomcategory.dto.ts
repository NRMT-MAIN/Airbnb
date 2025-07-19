import { RoomType } from "../db/models/roomcategory.model"

export type CreateRoomCategoryDTO = { 
    hotel_id : number , 
    price : number , 
    roomType : RoomType , 
    roomCount : number
}