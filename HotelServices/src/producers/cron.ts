import nodeCron from "node-cron";
import { generateRooms } from "../services/roomGeneration.service";


nodeCron.schedule("*****" , async () => {
    const payload = {
        "roomCategoryId": 1,
        "startDate": "2025-08-03T00:00:00Z",
        "endDate": "2025-08-04T00:00:00Z",
        "priceOverride": 199.99,
        "batchSize": 50
    }
    console.log("Create Rooms through cron")
    await generateRooms(payload) ; 
})