import { createHotelDTO, updateHotelDTO } from "../dto/hotel.dto";
import HotelRepository from "../repositories/hotel.repository";
import { BadRequestError } from "../utils/Error/app.error";

const blockListedAddresses = [
    "123 Fake Lt" , 
    "438 Elm St" , 
    "SSB JSAK"
]

const hotelRepository = new HotelRepository() ; 

export function isAddressBlockListed(address: string): boolean {
    return blockListedAddresses.includes(address) ; 
}

export async function createHotelService(hotelData: createHotelDTO){
    if(isAddressBlockListed(hotelData.address)){
        throw new BadRequestError("Address is blocklisted") ; 
    }
    const hotel = await hotelRepository.create(hotelData) ; 
    return hotel ; 
}

export async function getHotelByIdServices(id : number){
    const hotel = await hotelRepository.findById(id) ; 
    return hotel ; 
}

export async function getAllHotelsService() {
    const hotels = await hotelRepository.findAll();
    return hotels;
}

export async function deleteHotelService(id : number){
    //const hotels = await deleteHotelById(id) ; 
    const hotels = await hotelRepository.softDelete(id) ; 
    return hotels ; 
}

export async function updateHotelServices(id : number , hotelData :updateHotelDTO){
    const hotels = await hotelRepository.update(id , hotelData) ; 
    return hotels ; 
}