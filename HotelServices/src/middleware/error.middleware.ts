import { NextFunction, Request, Response } from "express";
import { AppError } from "../utils/Error/app.error";

export const appErrorHandeler = (err : AppError ,req : Request , res : Response , next : NextFunction) => {
    console.log(err) ; 

    res.status(err.statusCode).json({
        message : err.message , 
        success : false
    })
}

export const genericErrorHandeler = (err : AppError ,req : Request , res : Response , next : NextFunction) => {
    console.log(err) ; 

    res.status(500).json({
        message : "Internal Server Error" , 
        success : false
    })
}