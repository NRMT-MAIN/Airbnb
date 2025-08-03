import { NextFunction, Request, Response } from "express";
import { StatusCodes } from "http-status-codes";
import { scheduledRoomGeneration } from "../scheduler/roomSchedular";

export async function roomGenerationHandeler(req : Request , res: Response , next : NextFunction){
    await scheduledRoomGeneration(req.body)

    res.status(StatusCodes.CREATED).json({
        message: "Rooms generated successfully",
        success: true,
    })
}
