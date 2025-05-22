'use strict';

import { QueryInterface } from "sequelize";

module.exports = {
  async up (queryInterface : QueryInterface) {
      await queryInterface.sequelize.query(`
          ALTER TABLE HOTELS
          ADD COLUMN RATING DECIMAL(3 , 2) DEFAULT NULL , 
          ADD COLUMN RATING_COUNT INT DEFAULT NULL
      `) ;
  },

  async down (queryInterface : QueryInterface) {
      await queryInterface.sequelize.query(`
          ALTER TABLE HOTELS
          DROP COLUMN RATING , 
          DROP COLUMN RATING_COUNT
      `) ;
  }
};
