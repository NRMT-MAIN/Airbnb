import express from "express"
import { serverConfig } from "./config";
import v1Router from "./router/v1/index.router";
import { attachCorrelationIdMiddleware } from "./middleware/correlation.middleware";
import { appErrorHandeler, genericErrorHandler } from "./middleware/error.middleware";
import logger from "./config/logger.config";
import sequelize from "./db/models/sequelize";
import { InternalSeverError } from "./utils/Error/app.error";

const app = express() ; 

app.use(express.json()) ; 
app.use("/api/v1" , v1Router) ; 
app.use(attachCorrelationIdMiddleware) ; 
app.use(appErrorHandeler) ; 
app.use(genericErrorHandler) ; 

app.listen(serverConfig.PORT , async () => {
    console.log(`Server is running on Port : ${serverConfig.PORT}`) ; 
    logger.info("To exist , press CTRL + C") ; 

    try {
        await sequelize.authenticate() ; 
        logger.info("Database Connected Succesfully") ; 
    } catch(err){
        logger.error("Error in connecting to Database") ; 
        throw new InternalSeverError("Issue in connecting to the database.")
    }
})
