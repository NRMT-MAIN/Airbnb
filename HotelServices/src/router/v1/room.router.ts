import express from 'express';
import { validateQueryParam, validateRequestBody } from '../../validators';
import { getAvailableRoomSchema } from '../../validators/room.validator';
import { getAvailableRoomHandeler } from '../../controller/rooms.controller';


const roomRouter = express.Router();

roomRouter.get("/available" , validateQueryParam(getAvailableRoomSchema) , getAvailableRoomHandeler) ; 


export default roomRouter ; 