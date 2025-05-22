import { Request, Response } from "express";

export const pingHandeler = (req : Request , res : Response) => {
    res.send("Pong") ; 
}