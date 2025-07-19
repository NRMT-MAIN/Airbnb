import Rooms from "../db/models/rooms.model";
import BaseRepository from "./base.repository";

class RoomRepository extends BaseRepository<Rooms>{
    constructor(){
        super(Rooms) ; 
    }
}

export default RoomRepository ; 