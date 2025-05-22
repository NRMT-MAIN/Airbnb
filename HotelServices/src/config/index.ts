import dotenv from "dotenv"

type ServerConfig = {
    PORT : number , 
    NODE_ENV : string
}

type DBConfig = {
    DB_HOST: string , 
    DB_USER: string ,
    DB_PASSWORD: string,
    DB_NAME: string
}

function loadConfig(){
    dotenv.config() ; 
}

loadConfig() ; 

export const serverConfig : ServerConfig = {
    PORT : Number(process.env.PORT) || 3000 , 
    NODE_ENV : process.env.NODE_ENV || "development"
}

export const dbConfig: DBConfig = {
    DB_HOST: process.env.DB_HOST || 'localhost' , 
    DB_USER: process.env.DB_USER || 'root' , 
    DB_PASSWORD: process.env.DB_PASSWORD || 'root' , 
    DB_NAME: process.env.DB_NAME || 'test_db'
}