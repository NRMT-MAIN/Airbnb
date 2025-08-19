import { DataTypes, QueryInterface } from "sequelize";

module.exports = {
  async up (queryInterface : QueryInterface) {
      await queryInterface.addColumn("booking" , "roomCategoryId" , {
        type : DataTypes.INTEGER , 
        allowNull : false
      })
      await queryInterface.addColumn("booking" , "bookingAmount" , {
        type : DataTypes.INTEGER , 
        allowNull : false
      })
      await queryInterface.addColumn("booking" , "checkInDate" , {
        type : DataTypes.DATE , 
        allowNull : false
      })
      await queryInterface.addColumn("booking" , "checkOutDate" , {
        type : DataTypes.DATE , 
        allowNull : false
      })
  },

  async down (queryInterface : QueryInterface) {
      await queryInterface.removeColumn("booking" , "roomCategoryId") ; 
      await queryInterface.removeColumn("booking" , "bookingAmount") ; 
      await queryInterface.removeColumn("booking" , "checkInDate") ; 
      await queryInterface.removeColumn("booking" , "checkOutDate") ; 
  }
};
