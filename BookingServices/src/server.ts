import express from "express"
import { serverConfig } from "./config";
import v1Router from "./router/v1/index.router";
import { attachCorrelationIdMiddleware } from "./middleware/correlation.middleware";
import { appErrorHandeler, genericErrorHandler } from "./middleware/error.middleware";

const app = express() ; 

app.use("/api/v1" , v1Router) ; 
app.use(attachCorrelationIdMiddleware) ; 
app.use(appErrorHandeler) ; 
app.use(genericErrorHandler) ; 

app.listen(serverConfig.PORT , () => {
    console.log(`Server is running on Port : ${serverConfig.PORT}`) ; 
})
