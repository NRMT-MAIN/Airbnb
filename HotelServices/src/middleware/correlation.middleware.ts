import { NextFunction, Request, Response } from "express";
import {v4 as uuidV4} from "uuid" ;
import { asyncLocalStorage } from "../utils/helper/req.helper";

export const attachCorrelationIdMiddleware = (req : Request , res : Response , next : NextFunction) => {
    const correlationId = uuidV4() ; 
    req.headers["Req-Id"] = correlationId ;

    asyncLocalStorage.run( { correlationId: correlationId } , () => {
        next();
    });
}