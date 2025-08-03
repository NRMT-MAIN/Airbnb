import { Job, Worker } from "bullmq"
import { RoomGenerationJob } from "../dto/roomGeneration.dto"
import { RG_QUEUE } from "../queue/roomGeneration.queue"
import { RG_PRODUCER } from "../producers/roomGeneration.producer"
import { NotFoundError } from "../utils/Error/app.error"
import { getRedisConObject } from "../config/redis.config"
import { generateRooms } from "../services/roomGeneration.service"
import { error } from "console"


export const setupRoomGeneratorWorker = () => {
    const roomProcessor = new Worker<RoomGenerationJob>(
        RG_QUEUE , 
        async (job : Job) => {
            if(job.name != RG_PRODUCER) {
                throw new NotFoundError("Invalid Job Name") ; 
            }

            const payload = job.data ; 
            console.log(payload)
            console.log("Job is processing for :" , JSON.stringify(payload)) ; 
            
            await generateRooms(payload)
        } , 
        {
            connection : getRedisConObject()
        }
    )

    roomProcessor.on("failed" , () => {
        console.error("Room processing failed!") ; 
    }) ; 

    roomProcessor.on("completed" , () => {
        console.log("Room Processing completed.")
    })
}