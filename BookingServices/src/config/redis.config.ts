import IORedis from "ioredis"
import { serverConfig } from "."
import Redlock from "redlock";

const redisClient = new IORedis(serverConfig.REDIS_SERVER) ; 

export const redLock = new Redlock([redisClient] , {
    driftFactor : 0.01 , 
    retryCount : 10 , 
    retryDelay : 200 , 
    retryJitter : 200
})