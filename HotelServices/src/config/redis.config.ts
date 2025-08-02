import Redis from "ioredis";
import { serverConfig } from ".";
import logger from "./logger.config";
import { IntervalServerError } from "../utils/Error/app.error";

function connectRedis(){
    try{
        let connection : Redis ; 

        return () => {
            if(!connection){
                connection = new Redis({
                    host : serverConfig.REDIS_HOST , 
                    port : serverConfig.REDIS_PORT , 
                    maxRetriesPerRequest : null
                })
                return connection ; 
            }
            return connection ; 
        }
    } catch {
        logger.error("Connection to Redis failed") ; 
        throw new IntervalServerError("Redis Connection Failed!")
    }
}

export const getRedisConObject = connectRedis()