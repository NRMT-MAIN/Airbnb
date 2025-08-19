import axios from "axios";
import { serverConfig } from "../config";


export async function getAvailableRoom(roomCategoryId : number , checkInDate : string , checkOutDate : string){
    const response = await axios.get(serverConfig.HOTEL_SERVICE_API + "/rooms/available" , {
        params : {
            roomCategoryId , 
            checkInDate , 
            checkOutDate
        }
    }
    )
    return response ; 
}