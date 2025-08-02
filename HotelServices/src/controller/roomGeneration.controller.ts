import { NextFunction, Request, Response } from "express";
import { generateRooms } from "../services/roomGeneration.service";
import { StatusCodes } from "http-status-codes";
import { addRoomsToQueue } from "../producers/roomGeneration.producer";

export async function roomGenerationHandeler(req : Request , res: Response , next : NextFunction){
    
    const response = await addRoomsToQueue(req.body)

    res.status(StatusCodes.CREATED).json({
        message: "Rooms generated successfully",
        data: response,
        success: true,
    })
}