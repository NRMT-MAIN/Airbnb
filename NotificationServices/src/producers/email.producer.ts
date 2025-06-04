import logger from "../config/logger.config";
import { NotificationDTO } from "../dtos/notification.dto"
import { mailerQueue } from "../queue/email.queue"



export const MAILER_PAYLOAD = "payload-mailer"

export const addEmailToQueue = async (payload : NotificationDTO) => {
    await mailerQueue.add(MAILER_PAYLOAD , payload) ; 
    logger.info("Added Email to the Queue :" , JSON.stringify(payload)) ; 
}

