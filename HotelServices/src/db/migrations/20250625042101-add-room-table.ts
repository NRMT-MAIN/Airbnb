import { DataTypes, QueryInterface, Sequelize} from "sequelize";

module.exports = {
  async up (queryInterface : QueryInterface) {
      await queryInterface.createTable("Rooms" , {
          id : {
            allowNull : false , 
            autoIncrement : true , 
            primaryKey : true , 
            type : DataTypes.INTEGER
          } , 
          room_category_id : {
            type : DataTypes.INTEGER
          } , 
          hotel_id : {
            type : DataTypes.INTEGER
          } , 
          room_no : {
            type : DataTypes.INTEGER , 
            allowNull : false
          } , 
          date_of_availability : {
            type : DataTypes.DATE , 
            allowNull : false
          } , 
          booking_id : {
            type : DataTypes.INTEGER
          } , 
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
      await queryInterface.dropTable("Rooms")
  }
};
