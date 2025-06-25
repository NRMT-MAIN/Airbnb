import { QueryInterface } from "sequelize";


module.exports = {
  async up (queryInterface : QueryInterface) {
    await queryInterface.addConstraint("Rooms" , {
      type: 'foreign key',
      name: 'room-hid_fkey_constraint',
      fields: ['hotel_id'],
      references: {
        table: 'hotels',
        field: 'id',
      } , 
      onDelete: 'CASCADE',
      onUpdate: 'CASCADE',
    })
    await queryInterface.addConstraint("Rooms" , {
      type: 'foreign key',
      name: 'room-rcid_fkey_constraint',
      fields: ['room_category_id'],
      references: {
        table: 'roomcategory',
        field: 'id',
      } , 
      onDelete: 'CASCADE',
      onUpdate: 'CASCADE',
    })
    await queryInterface.addConstraint("Rooms" , {
      type: 'foreign key',
      name: 'room-bid_fkey_constraint',
      fields: ['booking_id'],
      references: {
        table: 'booking',
        field: 'id',
      } , 
      onDelete: 'CASCADE',
      onUpdate: 'CASCADE',
    })
  },

  async down (queryInterface : QueryInterface) {
    await queryInterface.removeConstraint("rooms" , "room-hid_fkey_constraint") ; 
    await queryInterface.removeConstraint("rooms" , "room-rcid_fkey_constraint") ; 
    await queryInterface.removeConstraint("rooms" , "room-bid_fkey_constraint") ; 
  }
};
