import express from "express"
import { roomGenerationHandeler } from "../../controller/roomGeneration.controller";
import { validateRequestBody } from "../../validators";
import { RoomGenerationJobSchema, RoomGenerationRequestSchema } from "../../dto/roomGeneration.dto";

const roomGenerationRouter = express.Router()


roomGenerationRouter.post("/" , validateRequestBody(RoomGenerationRequestSchema) ,  roomGenerationHandeler)


export default roomGenerationRouter ; 