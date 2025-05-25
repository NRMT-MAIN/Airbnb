import { DataTypes, QueryInterface } from "sequelize";

module.exports = {
  async up (queryInterface : QueryInterface) {
      await queryInterface.addColumn("IdompotencyKey" , "finalizedBooking" , {
          type : DataTypes.BOOLEAN , 
          defaultValue : false
      }) 
      await queryInterface.addColumn("Booking" , "idompotencyKey" , {
          type : DataTypes.INTEGER , 
          allowNull : true , 
          references : {
              model : "IdompotencyKey" , 
              key : "id"
          }
      })
  },

  async down (queryInterface : QueryInterface) {
      await queryInterface.removeColumn("IdompotencyKey" , "finalizedBooking") ; 
      await queryInterface.removeColumn("Booking" , "idompotencyKey") ; 
  }
};
