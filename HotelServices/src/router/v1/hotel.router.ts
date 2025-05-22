import express from 'express';
import { validateRequestBody } from '../../validators';
import { createHotelHandler, deleteHotelHandler, getAllHotelsHandler, getHotelByIdHandler, updateHotelHandler } from '../../controller/hotel.controller';
import { hotelSchema } from '../../validators/hotel.validator';

const hotelRouter = express.Router();

hotelRouter.post(
    '/', 
    validateRequestBody(hotelSchema),
    createHotelHandler); 

hotelRouter.get('/:id', getHotelByIdHandler); 

hotelRouter.get('/', getAllHotelsHandler);

hotelRouter.delete("/:id" , deleteHotelHandler) ; 

hotelRouter.put("/:id" , updateHotelHandler)

export default hotelRouter;