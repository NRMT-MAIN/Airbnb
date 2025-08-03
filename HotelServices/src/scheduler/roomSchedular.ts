import * as cron from 'node-cron';
import { RoomGenerationJob, RoomGenerationRequest } from '../dto/roomGeneration.dto';
import logger from '../config/logger.config';
import { serverConfig } from '../config';
import RoomRepository from '../repositories/room.repository';
import RoomCategoryRepository from '../repositories/roomcategory.repository';
import { addRoomsToQueue } from '../producers/roomGeneration.producer';
import { ForbiddenError, NotFoundError } from '../utils/Error/app.error';
import { z } from 'zod';

const roomRepository = new RoomRepository();
const roomCategoryRepository = new RoomCategoryRepository();

let cronJob: cron.ScheduledTask | null = null;

export const startScheduler = (): void => {
    if (cronJob) {
        logger.warn('Room scheduler is already running');
        return;
    }

    cronJob = cron.schedule(serverConfig.ROOM_CRON, async () => {
        try {
            logger.info('Starting room availability extension check');
            await extendRoomAvailability();
            logger.info('Room availability extension check completed');
        } catch (error) {
            logger.error('Error in room availability extension scheduler:', error);
        }
    }, {
        // scheduled: false,
        timezone: 'UTC'
    });

    cronJob.start();
    logger.info(`Room availability extension scheduler started - running every ${serverConfig.ROOM_CRON}`);
};

export const stopScheduler = (): void => {
    if (cronJob) {
        cronJob.stop();
        cronJob = null;
        logger.info('Room availability extension scheduler stopped');
    }
};

export const getSchedulerStatus = (): { isRunning: boolean } => {
    return {
        isRunning: cronJob !== null && cronJob.getStatus() === 'scheduled'
    };
};

const extendRoomAvailability = async (): Promise<void> => {
    try {
        // Get all room categories with their latest availability dates
        const roomCategoriesWithLatestDates = await roomRepository.findLatestDatesForAllCategories();
        
        if (roomCategoriesWithLatestDates.length === 0) {
            logger.info('No room categories found with availability dates');
            return;
        }

        logger.info(`Found ${roomCategoriesWithLatestDates.length} room categories to extend`);

        // Process each room category
        for (const categoryData of roomCategoriesWithLatestDates) {
            await extendCategoryAvailability(categoryData);
        }

    } catch (error) {
        logger.error('Error extending room availability:', error);
        throw error;
    }
};

/**
 * Extend availability for a specific room category
 */
const extendCategoryAvailability = async (categoryData: { room_category_id: number, latestDate: Date }): Promise<void> => {
    try {
        const { room_category_id, latestDate } = categoryData;

        // Calculate the next date (one day after the latest date)
        const nextDate = new Date(latestDate);
        nextDate.setDate(nextDate.getDate() + 1);

        // Check if the room category still exists
        const roomCategory = await roomCategoryRepository.findById(room_category_id);
        if (!roomCategory) {
            logger.warn(`Room category ${room_category_id} not found, skipping extension`);
            return;
        }

        // Check if room for next date already exists
        const existingRoom = await roomRepository.findByRoomCategoryIdAndDate(room_category_id, nextDate);
        if (existingRoom) {
            logger.debug(`Room for category ${room_category_id} on ${nextDate.toISOString()} already exists, skipping`);
            return;
        }

        const endDate = new Date(nextDate);
        endDate.setDate(endDate.getDate() + 1);

        // Create job to generate room for the next date
        const jobData: RoomGenerationJob = {
            roomCategoryId: room_category_id,
            startDate: nextDate.toISOString(),
            endDate: endDate.toISOString(),
            priceOverride: roomCategory.price,
            batchSize: 1
        };

        // Add job to queue
        await addRoomsToQueue(jobData);
        
        logger.info(`Added room generation job for category ${room_category_id} on ${nextDate.toISOString()}`);

    } catch (error) {
        logger.error(`Error extending availability for room category ${categoryData.room_category_id}:`, error);
        // Don't throw here to avoid stopping the entire scheduler
    }
};

/**
 * Manually trigger room availability extension (for testing or manual execution)
 */
export const manualExtendAvailability = async (): Promise<void> => {
    logger.info('Manual room availability extension triggered');
    await extendRoomAvailability();
}; 

export const scheduledRoomGeneration = async (categoryData: RoomGenerationRequest) => {
    const roomCategory = await roomCategoryRepository.findById(categoryData.roomCategoryId);
    if (!roomCategory) {
        logger.warn(`Room category ${categoryData.roomCategoryId} not found, skipping extension`);
        throw new NotFoundError("Category Id not found");
    }

    const jobData: RoomGenerationJob = {
        roomCategoryId: categoryData.roomCategoryId,
        startDate: categoryData.startDate,
        endDate: categoryData.endDate,
        priceOverride: categoryData.priceOveride || roomCategory.price,
        batchSize: Number(process.env.BATCH_SIZE) || 100,
    };

    console.log(categoryData.scheduledAt)

    

    const scheduledTime = categoryData.scheduledAt? new Date(categoryData.scheduledAt) : new Date(Date.now() + 60000);


    if (scheduledTime.getTime() < Date.now()) {
        logger.warn("Scheduled time is in the past");
        throw new ForbiddenError("Scheduled At cannot be at past")
    } 
    
    const dateToCron = (date: Date): string => {
        return [
            date.getUTCMinutes(), 
            date.getUTCHours(),    
            date.getUTCDate(),     
            date.getUTCMonth() + 1,
            date.getUTCDay()     
        ].join(' ');
    };

    const cronExpression = dateToCron(scheduledTime);

    const job = cron.schedule(cronExpression, async () => {
        await addRoomsToQueue(jobData);
        logger.info("Room generation job executed at:");
    });

    logger.info(`Scheduled room generation for cron: ${cronExpression}`);
    return job; 
};