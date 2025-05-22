import { NextFunction, Request, Response } from "express";
import logger from "../config/logger.config";

export const pingHandeler = (req : Request , res : Response , next : NextFunction) => {
    logger.info("Ping Handeler") ; 
    res.send("Ping Handler : Pong")
}