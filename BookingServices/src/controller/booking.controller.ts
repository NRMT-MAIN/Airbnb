import { Request, Response } from "express"
import { confirmBookingService, createBookingService } from "../service/booking.service"
import { StatusCodes } from "http-status-codes";

export const createBookingHandeler = async (req : Request , res : Response) => {
    const booking = await createBookingService(req.body) ; 

    res.status(StatusCodes.CREATED).json({
        bookingId : booking.bookingId , 
        key : booking.idompotencyKey
    })
}

export const confirmBookingHandeler = async (req : Request , res:  Response) => {
    const booking = await confirmBookingService(req.params.id) ; 

    res.status(StatusCodes.ACCEPTED).json({
        bookingId : booking , 
        status : "CONFIRMED"
    })
}