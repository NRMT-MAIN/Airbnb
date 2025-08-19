import RoomRepository from "../repositories/room.repository";

const roomRepository = new RoomRepository() ; 

export async function getAvailableRoomService(roomCategoryId : number , checkInDate : Date , checkOutDate : Date , roomNo : number[]) {
    const room = await roomRepository.findRoomByIdAndDateRange(roomCategoryId ,roomNo , checkInDate , checkOutDate) ; 
    return room ; 
}