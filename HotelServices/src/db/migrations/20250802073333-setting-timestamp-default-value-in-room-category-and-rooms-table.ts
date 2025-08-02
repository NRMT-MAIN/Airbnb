'use strict';

import { QueryInterface } from "sequelize";

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up (queryInterface : QueryInterface) {
      await queryInterface.sequelize.query(`
        ALTER TABLE ROOMCATEGORY 
        MODIFY CREATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        MODIFY UPDATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        MODIFY DELETED_AT TIMESTAMP NULL
      `);
      await queryInterface.sequelize.query(`
        ALTER TABLE ROOMS 
        MODIFY CREATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        MODIFY UPDATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        MODIFY DELETED_AT TIMESTAMP NULL
      `);
  },

  async down (queryInterface : QueryInterface) {
      await queryInterface.sequelize.query(`
        ALTER TABLE ROOMCATEGORY 
        MODIFY CREATED_AT TIMESTAMP NULL DEFAULT NULL,
        MODIFY UPDATED_AT TIMESTAMP NULL DEFAULT NULL,
        MODIFY DELETED_AT TIMESTAMP NULL DEFAULT NULL
      `);
      await queryInterface.sequelize.query(`
        ALTER TABLE ROOMS 
        MODIFY CREATED_AT TIMESTAMP NULL DEFAULT NULL,
        MODIFY UPDATED_AT TIMESTAMP NULL DEFAULT NULL,
        MODIFY DELETED_AT TIMESTAMP NULL DEFAULT NULL
      `);
  }
};
