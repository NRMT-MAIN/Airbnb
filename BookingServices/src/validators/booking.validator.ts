import { z } from "zod";

export const createBookingSchema = z.object({
    userId: z.number({ message: "User ID must be present" }),
    hotelId: z.number({ message: "Hotel ID must be present" }),
    totalGuest : z.number({ message: "Total guests must be present" }).min(1, { message: "Total guests must be at least 1" }),
})