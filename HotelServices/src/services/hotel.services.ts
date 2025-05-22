import { createHotelDTO, updateHotelDTO } from "../dto/hotel.dto";
import { createHotel, deleteHotelById, getAllHotels, getHotelById, softDeleteHotel, updateHotel } from "../repositories/hotel.repository";
import { BadRequestError } from "../utils/Error/app.error";

const blockListedAddresses = [
    "123 Fake Lt" , 
    "438 Elm St" , 
    "SSB JSAK"
]

export function isAddressBlockListed(address: string): boolean {
    return blockListedAddresses.includes(address) ; 
}

export async function createHotelService(hotelData: createHotelDTO){
    if(isAddressBlockListed(hotelData.address)){
        throw new BadRequestError("Address is blocklisted") ; 
    }
    const hotel = await createHotel(hotelData) ; 
    return hotel ; 
}

export async function getHotelByIdServices(id : number){
    const hotel = await getHotelById(id) ; 
    return hotel ; 
}

export async function getAllHotelsService() {
    const hotels = await getAllHotels();
    return hotels;
}

export async function deleteHotelService(id : number){
    //const hotels = await deleteHotelById(id) ; 
    const hotels = await softDeleteHotel(id) ; 
    return hotels ; 
}

export async function updateHotelServices(id : number , hotelData :updateHotelDTO){
    const hotels = await updateHotel(id , hotelData) ; 
    return hotels ; 
}