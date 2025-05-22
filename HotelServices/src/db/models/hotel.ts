import { CreationOptional, InferAttributes, InferCreationAttributes, Model } from "sequelize";
import sequelize from "./sequelize";
import { defaultValueSchemable } from "sequelize/types/utils";

class Hotel extends Model<InferAttributes<Hotel> , InferCreationAttributes<Hotel> >{
    declare id : CreationOptional<number>  ; 
    declare name  : string  ; 
    declare address : string  ; 
    declare location : string ; 
    declare createdAt : CreationOptional<Date> ; 
    declare updatedAt : CreationOptional<Date> ; 
    declare deletedAt : CreationOptional<Date | null> ; 
    declare rating? : number ; 
    declare ratingCount? : number ; 
}

Hotel.init({
    id: {
        type: "INTEGER",
        autoIncrement: true,
        primaryKey: true,
    },
    name: {
        type: "STRING",
        allowNull: false,
    },
    address: {
        type: "STRING",
        allowNull: false,
    },
    location: {
        type: "STRING",
        allowNull: false,
    },
    createdAt: {
        type: "DATE",
        defaultValue: new Date(),
    },
    updatedAt: {
        type: "DATE",
        defaultValue: new Date(),
    },
    deletedAt : {
        type : "DATE" , 
        defaultValue : null
    } , 
    rating: {
        type: "FLOAT",
        defaultValue: null,
    },
    ratingCount: {
        type: "INTEGER",
        defaultValue: null,
    }
}, {
    tableName: "hotels",
    sequelize: sequelize,
    underscored: true, // createdAt --> created_at automatically make this format
    timestamps: true, // createdAt, updatedAt
});
//Initialize a model, representing a table in the DB, with attributes 
// and options.
//Will make map to particular corresponding table in db

export default Hotel;