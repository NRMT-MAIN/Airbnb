import { CreationOptional, DataTypes, EnumDataType, InferAttributes, InferCreationAttributes, Model, NOW } from "sequelize";
import sequelize from "./sequelize";

export enum BoookingStatus { 
    PENDING = 'PENDING',
    CONFIRMED = 'CONFIRMED',
    CANCELLED = 'CANCELLED'
}

class Booking extends Model<InferAttributes<Booking> , InferCreationAttributes<Booking>> {
    declare id : CreationOptional<number>  ; 
    declare userId : number ; 
    declare hotelId : number ; 
    declare totalGuest : number ; 
    declare createdAt : CreationOptional<Date> ; 
    declare updatedAt : CreationOptional<Date> ; 
    declare status : CreationOptional<String> ;
}

Booking.init({
        id : {
            type : DataTypes.INTEGER , 
            primaryKey : true ,
            autoIncrement : true 
        } , 
        userId: {
            type: DataTypes.INTEGER,
            allowNull: false,
            field: 'userId'
        },
        hotelId: {
            type: DataTypes.INTEGER,
            allowNull: false,
            field : 'hotelId'
        },
        status: {
            type: DataTypes.STRING,
            allowNull: false,
            defaultValue: BoookingStatus.PENDING 
        },
        createdAt : {
            type :  DataTypes.DATE,
            defaultValue : NOW
        } , 
        updatedAt: {
            type :  DataTypes.DATE,
            defaultValue : NOW
        } , 
        totalGuest : {
            type : DataTypes.INTEGER , 
            allowNull : false
        }

    } ,
    {
        tableName: "Booking",
        sequelize: sequelize,
        timestamps: true,
    } 
)

export default Booking ; 