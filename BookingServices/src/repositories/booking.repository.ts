import logger from "../config/logger.config";
import Booking, { BoookingStatus } from "../db/models/booking.model";
import IdompotencyKey from "../db/models/idompotencykey.model";
import { CreateBookingDTO } from "../dto/booking.dto";
import { NotFoundError } from "../utils/Error/app.error";


export async function createBooking(bookingInput : CreateBookingDTO){
    const booking = await Booking.create(bookingInput) ; 

    logger.info("Booking created with id :" , booking.id) ; 
    return booking ; 
}

export async function createIdompotencyKey(bookingId : number , key : string){
    const booking = await Booking.findByPk(bookingId) ; 
    if(!booking){
        logger.error("Booking doesn't found with this id :" , bookingId) ; 
        throw new NotFoundError("Booking Not Found") ; 
    }

    const idompotencyKey = await IdompotencyKey.create({
        key : key , 
        bookingId : bookingId
    })

    logger.info("Idompotency Key created with id :" , idompotencyKey.id)

    return idompotencyKey ; 
}

export async function getIdompotencyKey(key : string){
    const idompotencyKey = await IdompotencyKey.findOne({
        where : {
            key
        } 
    }) ; 
    if(!idompotencyKey){
        logger.error("Idompotency Key not found with :" , key) ; 
        throw new NotFoundError("Idompotency Key Not Found") ; 
    }

    logger.info("Idompotency Key found with key:" , key) ; 
    return idompotencyKey ; 
}

export async function getBookingById(bookingId : number){
    const booking = await Booking.findByPk(bookingId) ; 

    if(!booking){
        logger.error("No Booking found with id :" , bookingId) ; 
        throw new NotFoundError("Booking Not Found") ; 
    }

    logger.info("Booking Key Found with id:" , bookingId) ; 
    return booking ; 
}

export async function confirmBooking(bookingId : number){
    const booking = await Booking.update({
        status : BoookingStatus.CONFIRMED
    } , {
        where : {
            id : bookingId
        } , 
    })

    if(!booking){
        logger.error("Booking status not updated yet") ; 
        throw new Error("Booking status not updated yet") ; 
    }

    logger.info("Booking status updated") ; 
    return bookingId ; 
}

export async function cancelBooking(bookingId : number){
    const booking = await Booking.update({
        status : BoookingStatus.CANCELLED
    } , {
        where : {
            id : bookingId
        } 
    })

    if(!booking){
        logger.error("Booking status not updated yet") ; 
        throw new Error("Booking status not updated yet") ; 
    }

    logger.info("Booking status updated") ; 
    return bookingId ; 
}

export async function finalizeIdompotencyKey(key : string){
    const idompotencyKey = await IdompotencyKey.update({
        finalizedBooking : true
    } , {
        where : {
            key : key
        } 
    })

    if(!idompotencyKey){
        logger.error("Idompotency Key is not finalized with key:" , key) ; 
        throw new Error("Idompotency Key is not finalized") ; 
    }

    logger.info("Idompotency Key is finalized with key:" , key) ; 
    return idompotencyKey ; 
}