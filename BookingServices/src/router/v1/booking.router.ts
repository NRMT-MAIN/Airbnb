import express from "express"
import { confirmBookingHandeler, createBookingHandeler } from "../../controller/booking.controller";
import { validateRequestBody } from "../../validators";
import { createBookingSchema } from "../../validators/booking.validator";

const bookingRouter = express.Router() ; 

bookingRouter.post("/" ,validateRequestBody(createBookingSchema)  , createBookingHandeler) ;
bookingRouter.post("/confirm/:id" , confirmBookingHandeler) ;  

export default bookingRouter ; 