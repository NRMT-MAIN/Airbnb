import { DataTypes, NOW, QueryInterface, Sequelize, UUIDV4 } from "sequelize";

module.exports = {
  async up (queryInterface : QueryInterface) {
      await queryInterface.createTable("IdompotencyKey" , {
          id : {
              type : DataTypes.INTEGER , 
              primaryKey : true , 
              autoIncrement : true
          } , 
          key : {
              type : DataTypes.STRING , 
              unique : true
          } , 
          createdAt: {
            type: DataTypes.DATE,
            defaultValue: Sequelize.literal('CURRENT_TIMESTAMP'),
          },
          updatedAt: {
            type: DataTypes.DATE,
            defaultValue: Sequelize.literal('CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP'),
          },
          bookingId : {
              type : DataTypes.INTEGER , 
              allowNull : false , 
              references : {
                  model : "Booking" , 
                  key : "id"
              }
          }
      })
  },

  async down (queryInterface : QueryInterface) {
      await queryInterface.dropTable("IdompotencyKey") ; 
  }
};
