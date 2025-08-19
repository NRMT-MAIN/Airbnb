import { z } from "zod";

export const getAvailableRoomSchema = z.object({
    roomCategoryId : z.string({message : "Room Category Id required"}) , 
    checkInDate : z.string({message : "Check in Date Required"}) , 
    checkOutDate : z.string({message : "Check out Date Required"}) 
})