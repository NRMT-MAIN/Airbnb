import { CreationAttributes } from "sequelize";
import Rooms from "../db/models/rooms.model";
import BaseRepository from "./base.repository";

class RoomRepository extends BaseRepository<Rooms>{
    constructor(){
        super(Rooms) ; 
    }

    async bulkCreate(rooms: CreationAttributes<Rooms>[]){
        return await this.model.bulkCreate(rooms)
    }
}

export default RoomRepository ; 