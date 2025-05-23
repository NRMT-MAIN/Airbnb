import logger from "../config/logger.config";
import Booking from "../db/models/booking";
import { CreateBookingDTO } from "../dto/booking.dto";


export async function createBooking(bookingInput : CreateBookingDTO){
    const booking = await Booking.create(bookingInput) ; 

    logger.info("Booking created with id :" , booking.id) ; 
    return booking ; 
}
