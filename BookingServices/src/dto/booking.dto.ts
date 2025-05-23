import { EnumDataType } from "sequelize";
import { BoookingStatus } from "../db/models/booking";

export type CreateBookingDTO = {
  userId: number;
  hotelId: number;
  totalGuest: number;
  status : EnumDataType<BoookingStatus> ; 
};
