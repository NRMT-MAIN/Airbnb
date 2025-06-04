import Redis from "ioredis";
import { serverConfig } from ".";
import logger from "./logger.config";
import { InternalServerError } from "../utils/errors/app.error";

function connectRedis(){
    try {
        let connection : Redis ; 

        const redisConfig = {
            host : serverConfig.REDIS_HOST , 
            port : serverConfig.REDIS_PORT , 
            maxRetriesPerRequest : null
        }

        return () => {
            if(!connection) {
                connection = new Redis(redisConfig) ; 
                return connection ; 
            }
            return connection ; 
        }
    } catch(err){
        logger.error("Connection to Redis failed") ; 
        throw new InternalServerError("Redis Connection Failed!")
    }
}

export const getRedisConnObject = connectRedis() ; 