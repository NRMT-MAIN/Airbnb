import path from "path";
import fs from "fs/promises"
import Handlebars from "handlebars";
import { InternalServerError } from "../utils/errors/app.error";
import logger from "../config/logger.config";


export async function renderEmailTemplate(templateId : string , params : Record<string , any>) : Promise<string> {
    const templatePath = path.join(__dirname , "mailer" , `${templateId}.hbs`) ; 
    console.log(templatePath)
    try {
        const content = await fs.readFile(templatePath , "utf8") ; 
        const finalTemplate = Handlebars.compile(content) ; 
        logger.info("Template Compiled") ;
        return finalTemplate(params) ; 
    } catch (error) {
        throw new InternalServerError(`Template with id - ${templateId} not found!`) ; 
    }
}