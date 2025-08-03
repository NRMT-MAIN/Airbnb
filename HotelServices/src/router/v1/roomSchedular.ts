import { Router } from "express";
import { getSchedulerStatusHandler, manualExtendAvailabilityHandler, startSchedulerHandler, stopSchedulerHandler } from "../../controller/roomScheduler.controller";


const schedulerRouter = Router();

schedulerRouter.post("/start", startSchedulerHandler);

schedulerRouter.post("/stop", stopSchedulerHandler);

schedulerRouter.get("/status", getSchedulerStatusHandler);

schedulerRouter.post("/extend", manualExtendAvailabilityHandler);

export default schedulerRouter; 
