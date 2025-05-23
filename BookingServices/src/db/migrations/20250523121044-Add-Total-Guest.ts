import { DataTypes, QueryInterface } from "sequelize";

module.exports = {
  async up (queryInterface : QueryInterface) {
      await queryInterface.addColumn("Booking" , "totalGuest" ,{
          type : DataTypes.INTEGER , 
          allowNull : false
      })
  },

  async down (queryInterface : QueryInterface) {
      await queryInterface.removeColumn("Booking" , "totalGuest")
  }
};
