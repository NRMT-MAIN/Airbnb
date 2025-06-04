import { Job, Worker } from "bullmq"
import { NotificationDTO } from "../dtos/notification.dto"
import { MAILER_QUEUE } from "../queue/email.queue"
import { MAILER_PAYLOAD } from "../producers/email.producer"
import { NotFoundError } from "../utils/errors/app.error"
import { getRedisConnObject } from "../config/redis.config"

export const setupMailerWorker = () => {
    const emailProcessor = new Worker<NotificationDTO>(
        MAILER_QUEUE ,
        async (job : Job) => {
            if(job.name != MAILER_PAYLOAD) {
                throw new NotFoundError("Invalid Job Name") ; 
            }

            const payload = job.data ; 
            console.log("Email is processing for :" , JSON.stringify(payload)) ; 
        } , 
        {
            connection : getRedisConnObject()
        }
    ) ; 
    emailProcessor.on("failed" , () => {
        console.error("Email processing failed!") ; 
    }) ; 

    emailProcessor.on("completed" , () => {
        console.log("Email Processing completed.")
    })
}