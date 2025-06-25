import { CreationOptional, DataTypes, InferAttributes, InferCreationAttributes, Model } from "sequelize";
import sequelize from "./sequelize";
import Hotel from "./hotel.model";
import RoomCategory from "./roomcategory.model";


class Rooms extends Model<InferAttributes<Rooms> , InferCreationAttributes<Rooms>>{
    declare id : CreationOptional<number>  ; 
    declare room_category_id : number ; 
    declare hotel_id : number ; 
    declare room_no : number ; 
    declare date_of_availability : Date ; 
    declare booking_id? : number | null ; 
    declare price : number ; 
    declare created_at : CreationOptional<Date> ; 
    declare updated_at : CreationOptional<Date> ; 
    declare deleted_at : CreationOptional<Date> | null ; 
}

Rooms.init(
  {
    id: {
      type: 'INTEGER',
      autoIncrement: true,
      primaryKey: true,
    },
    hotel_id: {
      type: 'INTEGER',
      allowNull: false,
      references: {
        model: Hotel,
        key: 'id',
      },
    },
    room_category_id: {
      type: 'INTEGER',
      allowNull: false,
      references: {
        model: RoomCategory,
        key: 'id',
      },
    },
    date_of_availability: {
      type: 'DATE',
      allowNull: false,
    },
    price: {
      type: 'INTEGER',
      allowNull: false,
    },
    created_at: {
      type: 'DATE',
      defaultValue: new Date(),
    },
    updated_at: {
      type: 'DATE',
      defaultValue: new Date(),
    },
    deleted_at: {
      type: 'DATE',
      defaultValue: null,
    },
    booking_id: {
      type : "INTEGER" , 
      defaultValue : null
    },
    room_no : {
      type : "INTEGER" , 
      allowNull : false
    }
  },
  {
    tableName: 'rooms',
    sequelize: sequelize,
    timestamps: true,
  }
);

export default Rooms;