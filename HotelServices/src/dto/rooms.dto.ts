import { z } from "zod"

export type RoomAvailableDTO = {
    roomCategoryId : number , 
    checkInDate : Date , 
    checkOutDate : Date
} 