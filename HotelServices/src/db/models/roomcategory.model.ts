import { CreationOptional, InferAttributes, InferCreationAttributes, Model } from "sequelize";
import sequelize from "./sequelize";
import Hotel from "./hotel.model";

enum RoomType {
  SINGLE = 'SINGLE',
  DOUBLE = 'DOUBLE',
  FAMILY = 'FAMILY',
  DELUXE = 'DELUXE',
}

class RoomCategory extends Model< InferAttributes<RoomCategory>, InferCreationAttributes<RoomCategory>> {
  declare id: CreationOptional<number>;
  declare hotel_id: number;
  declare price: number;
  declare roomType: RoomType;
  declare roomCount: number;
  declare created_at: CreationOptional<Date>;
  declare updated_at: CreationOptional<Date>;
  declare deleted_at: CreationOptional<Date> | null;
}

RoomCategory.init(
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
    price: {
      type: 'INTEGER',
      allowNull: false,
    },
    roomType: {
      type: 'ENUM',
      values: [...Object.values(RoomType)],
    },
    roomCount: {
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
  },
  {
    tableName: 'roomcategory',
    sequelize: sequelize,
    timestamps: true,
  }
);

export default RoomCategory;