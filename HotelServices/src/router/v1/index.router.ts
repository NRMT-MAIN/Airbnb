import express from "express"
import pingRouter from "./ping.router";
import hotelRouter from "./hotel.router";
import roomGenerationRouter from "./roomGeneration.router";
import schedulerRouter from "./roomSchedular";

const v1Router = express.Router() ; 

v1Router.use("/ping" , pingRouter) ; 
v1Router.use("/hotels" , hotelRouter) ; 
v1Router.use("/generate-rooms" , roomGenerationRouter)
v1Router.use("/schedule" , schedulerRouter)

export default v1Router ; 