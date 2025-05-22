import express from "express"
import { pingHandeler } from "../../controller/ping.controller";

const pingRouter = express.Router() ; 

pingRouter.get("/" , pingHandeler) ; 

export default pingRouter ; 