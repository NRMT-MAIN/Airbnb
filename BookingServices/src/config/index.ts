import dotenv from "dotenv"

type ServerConfig = {
    PORT : number , 
    REDIS_SERVER : string , 
    TTL : number
}

type DBConfig = {
    DB_USER : string , 
    DB_PASSWORD : string , 
    DB_HOST : string , 
    DB_NAME : string , 
    DB_DIALECT : string
}

function loadEnv(){
    dotenv.config() ; 
}

loadEnv() ; 

export const serverConfig : ServerConfig = {
    PORT : Number(process.env.PORT) || 3001 , 
    REDIS_SERVER : process.env.REDIS_SERVER || "redis://localhost:6379" , 
    TTL : Number(process.env.TTL) || 10000
}

export const dbConfig : DBConfig = {
    DB_HOST: process.env.DB_HOST || 'localhost' , 
    DB_USER: process.env.DB_USER || 'root' , 
    DB_PASSWORD: process.env.DB_PASSWORD || 'root' , 
    DB_NAME: process.env.DB_NAME || 'test_db' , 
    DB_DIALECT : process.env.DB_DIALECT || "mysql"
}