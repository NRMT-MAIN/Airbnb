import winston from "winston"
import { getCorrelationId } from "../utils/Helper/req.helper"
import DailyRotateFile from "winston-daily-rotate-file"

const logger = winston.createLogger({
    format : winston.format.combine(
        winston.format.timestamp({format : "DD:MM:YYYY HH:mm:ss"}) , 
        winston.format.json() , 
        winston.format.printf(({level , timestamp , message , ...data}) => {
            const output = {
                level , 
                message , 
                timestamp , 
                correlationId : getCorrelationId() , 
                data
            }
            return JSON.stringify(output) ; 
        })
    ) , 
    transports : [
        new winston.transports.Console() , 
        new DailyRotateFile({
            filename : "logs/%DATE%.log" , 
            datePattern: 'YYYY-MM-DD',
            zippedArchive: true,
            maxSize: '20m',
            maxFiles: '14d'
        })
    ]
})

export default logger ; 