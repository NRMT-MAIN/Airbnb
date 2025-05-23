const dbConfig = {
    "development": {
      "username": dbConfig.DB_USER,
      "password": dbConfig.DB_PASSWORD,
      "database": dbConfig.DB_NAME,
      "host": dbConfig.DB_HOST,
      "dialect" : dbConfig.DB_DIALECT 
    },
    "test": {
      "username": "root",
      "password": null,
      "database": "database_test",
      "host": "127.0.0.1",
      "dialect": "mysql"
    },
    "production": {
      "username": "root",
      "password": null,
      "database": "database_production",
      "host": "127.0.0.1",
      "dialect": "mysql"
    }
}