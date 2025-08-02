import logger from "../config/logger.config"
import { RoomGenerationJob } from "../dto/roomGeneration.dto"
import { RG_QUEUE, roomGenrationQueue } from "../queue/roomGeneration.queue"


export const RG_PRODUCER = "producer-rg"

export const addRoomsToQueue = async (payload : RoomGenerationJob) => {
    await roomGenrationQueue.add(RG_PRODUCER , payload)
    logger.info("Added Job to the Queue :" ,JSON.stringify(payload)); 
}