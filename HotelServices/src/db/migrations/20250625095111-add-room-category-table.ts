import { DataTypes, QueryInterface } from "sequelize";

module.exports = {
  async up (queryInterface : QueryInterface) {
      await queryInterface.createTable("roomCategory" , {
        id : { 
          type : DataTypes.INTEGER ,
          primaryKey : true , 
          allowNull : false , 
          autoIncrement : true
        } ,
        hotel_id : {
          type : DataTypes.INTEGER , 
          allowNull : false , 
          references : {
            model :  "Hotels" , 
            key : "id"
          }
        } , 
        price :  {
          type: DataTypes.INTEGER,
          allowNull: false,
        },
         roomType: {
          type: DataTypes.ENUM("SINGLE" , "DOUBLE" , "DELUXE" , "FAMILY"),
          allowNull : false
        },
        roomCount: {
          type: DataTypes.INTEGER , 
          allowNull: false,
        },
        created_at : {
          type : DataTypes.DATE , 
          allowNull : false , 
          defaultValue : DataTypes.NOW
        } ,
        updated_at : {
          type : DataTypes.DATE , 
          allowNull : false , 
          defaultValue : DataTypes.NOW
        } , 
        deleted_at  : {
          type : DataTypes.DATE , 
          allowNull : true
        }
      })
  },

  async down (queryInterface : QueryInterface) {
    await queryInterface.dropTable("roomCategory") ; 
  }
};
