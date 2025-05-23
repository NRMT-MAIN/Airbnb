import { CreationOptional, DataTypes, EnumDataType, InferAttributes, InferCreationAttributes, Model } from "sequelize";
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
    declare createdAt : CreationOptional<Date> ; 
    declare updatedAt : CreationOptional<Date> ; 
    declare status : EnumDataType<BoookingStatus>
}

Booking.init({
        id : {
            type : "INTEGER" , 
            primaryKey : true ,
            autoIncrement : true 
        } , 
        userId: {
            type: DataTypes.INTEGER,
            allowNull: false,
        },
        hotelId: {
            type: DataTypes.INTEGER,
            allowNull: false,
        },
        status: {
            type: DataTypes.ENUM(...Object.values(BoookingStatus)),
            allowNull: false,
            defaultValue: BoookingStatus.PENDING 
        },
        createdAt: DataTypes.DATE,
        updatedAt: DataTypes.DATE,

    } ,
    {
        tableName: "Booking",
        sequelize: sequelize,
        underscored: true, // createdAt --> created_at automatically make this format
        timestamps: true,
    } 
)

export default Booking ; 