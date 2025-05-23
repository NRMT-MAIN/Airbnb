import { Sequelize } from "sequelize";
import { dbConfig } from "../../config";

const sequelize = new Sequelize({
    dialect : "mysql" , 
    host : dbConfig.DB_HOST , 
    username : dbConfig.DB_USER , 
    database : dbConfig.DB_NAME , 
    password : dbConfig.DB_PASSWORD , 
    logging : true
})

export default sequelize ; 