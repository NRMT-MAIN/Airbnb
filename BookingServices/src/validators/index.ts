import { NextFunction, Request, Response } from "express";
import { AnyZodObject } from "zod";

export const validateRequestBody = (schema : AnyZodObject) => {
    return async (req : Request , res : Response , next : NextFunction) => {
        try {
            await schema.parseAsync(req.body) ; 
            next() ; 
        } catch(err){
            res.status(400).json({
                message : "Request Body is invalid!" , 
                success : false , 
                error : err
            })
        }
    }
}

export const validateQueryParam = (schema : AnyZodObject) => {
    return async (req : Request , res : Response , next : NextFunction) => {
        try {
            await schema.parseAsync(req.body) ; 
            next() ; 
        } catch(err){
            res.status(400).json({
                message : "Request Body is invalid!" , 
                success : false , 
                error : err
            })
        }
    }
}