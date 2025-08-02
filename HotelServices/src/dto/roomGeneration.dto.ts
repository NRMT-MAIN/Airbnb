import { z } from "zod";


export const RoomGenerationRequestSchema = z.object({
    roomCategoryId : z.number().positive() , 
    startDate : z.string().datetime() , 
    endDate : z.string().datetime() , 
    scheduledType : z.enum(["immediate" , "scheduled"]).default("immediate") , 
    scheduledAt : z.string().datetime().optional() ,
    priceOveride : z.number().positive().optional()
}) 


export const RoomGenerationJobSchema = z.object({
    roomCategoryId : z.number().positive() , 
    startDate: z.string().datetime(),
    endDate: z.string().datetime(),
    priceOverride: z.number().positive().optional(),
    batchSize: z.number().positive().optional().default(100),
})

export type RoomGenerationRequest = z.infer<typeof RoomGenerationRequestSchema>
export type RoomGenerationJob  = z.infer<typeof RoomGenerationJobSchema>

export type RoomAvailableDTO = {
    room_no : number , 
    dateOfAvailability : string
}
