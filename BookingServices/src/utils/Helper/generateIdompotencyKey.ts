import { v4 } from "uuid"

export const generateIdompotencyKey = () : string => {
    return v4() ; 
}