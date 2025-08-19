import { NextFunction, Request, Response, } from "express"
import { getAvailableRoomService } from "../services/room.services"
import { StatusCodes } from "http-status-codes";


export async function getAvailableRoomHandeler(req : Request , res : Response , next : NextFunction) {
    const {roomCategoryId , checkInDate , checkOutDate , roomNo} = req.query ; 

    const parsedNumber = Number(roomCategoryId) ; 
    const parsedCheckInDate = new Date(checkInDate as string) ; 
    const parsedCheckOutDate = new Date(checkOutDate as string) ; 


    const roomsArray  = (roomNo as string).split(",")
                        .map(num => parseInt(num, 10)) ; 

    const rooms = await getAvailableRoomService(parsedNumber , parsedCheckInDate , parsedCheckOutDate , roomsArray) ; 

    res.status(StatusCodes.CREATED).json({
            message: "Hotel created successfully",
            data: rooms,
            success: true,
    })
}