import { Queue } from "bullmq";
import { getRedisConObject } from "../config/redis.config";


export const RG_QUEUE = "queue-rg"

export const roomGenrationQueue = new Queue(RG_QUEUE, {
    connection : getRedisConObject()
})