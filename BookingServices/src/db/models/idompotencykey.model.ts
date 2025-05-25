import { CreationOptional, DataTypes, InferAttributes, InferCreationAttributes, Model } from "sequelize";
import sequelize from "./sequelize";

class IdompotencyKey extends Model<InferAttributes<IdompotencyKey> , InferCreationAttributes<IdompotencyKey>>{
    declare id : CreationOptional<number> ; 
    declare key : string  ; 
    declare createdAt : CreationOptional<Date> ; 
    declare updatedAt : CreationOptional<Date> ; 
    declare bookingId : number ; 
    declare finalizedBooking : CreationOptional<Boolean> ; 
}

IdompotencyKey.init({
    id : {
        type : DataTypes.INTEGER , 
        primaryKey : true , 
        autoIncrement : true
    } , 
    key : {
        type : DataTypes.STRING , 
        allowNull : false
    } , 
    createdAt: DataTypes.DATE,
    updatedAt: DataTypes.DATE,
    bookingId : {
        type : DataTypes.INTEGER , 
        allowNull : false
    } , 
    finalizedBooking : {
        type : DataTypes.STRING , 
        allowNull : true
    }
} , 
    {
        tableName: "IdompotencyKey",
        sequelize: sequelize,
        timestamps: true,
    } 
)

export default IdompotencyKey ; 
