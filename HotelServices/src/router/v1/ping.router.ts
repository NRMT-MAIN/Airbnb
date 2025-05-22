import express from "express" 
import { pingHandeler } from "../../controller/ping.controller";
import { validateRequestBody } from "../../validators";
import { pingSchema } from "../../validators/ping.validators";

const pingRouter = express.Router() ; 

pingRouter.get("/" , validateRequestBody(pingSchema) , pingHandeler) ; 

export default pingRouter ; 