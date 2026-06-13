import {apiTrains} from "../../../core/api/client";

import type {CreateScheduleStopRequest, ScheduleStop} from "../types/schedule_stop.ts";

export const getStopsBySchedule = (scheduleId: number) =>
    apiTrains.get<ScheduleStop[]>(`/stopschedules/schedule/${scheduleId}`);

export const createStop = (scheduleId: number, data: Omit<CreateScheduleStopRequest, "schedule_id">) =>
    apiTrains.post(`/stopschedules`, {
        schedule_id: scheduleId,
        ...data,
    });

export const updateStop = (stopId: number, data: number) =>
    apiTrains.patch(`/stopschedules/${stopId}`, data);

export const deleteStop = (stopId: number) =>
    apiTrains.delete(`/stopschedules/${stopId}`);
