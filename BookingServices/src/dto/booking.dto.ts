import { EnumDataType } from "sequelize";
import { BoookingStatus } from "../db/models/booking.model";

export type CreateBookingDTO = {
    userId: number;
    hotelId: number;
    totalGuest: number;
    roomCategoryId : number ; 
    bookingAmount : number ; 
    checkInDate : string  ; 
    checkOutDate : string ; 
};
