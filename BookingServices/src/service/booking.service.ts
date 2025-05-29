import { Transaction } from "sequelize";
import sequelize from "../db/models/sequelize";
import { CreateBookingDTO } from "../dto/booking.dto";
import { confirmBooking, createBooking, createIdompotencyKey, finalizeIdompotencyKey, getIdompotencyKey } from "../repositories/booking.repository";
import { BadRequestError, InternalSeverError } from "../utils/Error/app.error";
import { generateIdompotencyKey } from "../utils/Helper/generateIdompotencyKey";
import { serverConfig } from "../config";
import { redLock } from "../config/redis.config";
import logger from "../config/logger.config";

export async function createBookingService(bookingInput : CreateBookingDTO){
    const ttl = serverConfig.TTL ; 
    const bookingResource = `hotel:${bookingInput.hotelId}` ; 

    try {
        await redLock.acquire([bookingResource] , ttl) ; 
        const booking = await createBooking(bookingInput) ; 
    
        const key = generateIdompotencyKey() ; 

        await createIdompotencyKey(booking.id , key) ; 

        return {
            bookingId : booking.id , 
            idompotencyKey : key
        }
    } catch(err) {
        logger.error(`Failed to acquire lock for hotelId : ${bookingInput.hotelId}`) ; 
        throw new InternalSeverError("Failed to acquire lock.") ; 
    }
}

export async function confirmBookingService(key : string){
    const result = await sequelize.transaction( {
        isolationLevel : Transaction.ISOLATION_LEVELS.REPEATABLE_READ
    } , async (txn) => {
        const idompotencyKey = await getIdompotencyKey(txn , key) ; 

        if(idompotencyKey.finalizedBooking) {
            throw new BadRequestError("Idompotency Key already finalized") ; 
        }

        const booking = await confirmBooking(txn , idompotencyKey.bookingId) ; 
        await finalizeIdompotencyKey(txn , key) ; 

        return booking ; 
    }) ; 

    if(!result){
        throw new InternalSeverError("Transactions failed to Complete") ; 
    }
}