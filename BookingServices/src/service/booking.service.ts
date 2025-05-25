import { CreateBookingDTO } from "../dto/booking.dto";
import { confirmBooking, createBooking, createIdompotencyKey, finalizeIdompotencyKey, getIdompotencyKey } from "../repositories/booking.repository";
import { BadRequestError } from "../utils/Error/app.error";
import { generateIdompotencyKey } from "../utils/Helper/generateIdompotencyKey";

export async function createBookingService(bookingInput : CreateBookingDTO){
    const booking = await createBooking(bookingInput) ; 
    
    const key = generateIdompotencyKey() ; 

    await createIdompotencyKey(booking.id , key) ; 

    return {
        bookingId : booking.id , 
        idompotencyKey : key
    }
}

export async function confirmBookingService(key : string){
    const idompotencyKey = await getIdompotencyKey(key) ; 

    if(idompotencyKey.finalizedBooking) {
        throw new BadRequestError("Idompotency Key already finalized") ; 
    }

    const booking = await confirmBooking(idompotencyKey.bookingId) ; 
    await finalizeIdompotencyKey(key) ; 

    return booking ; 
}