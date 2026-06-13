import {apiTrains} from "../../../core/api/client";
import type {CreateScheduleRequest, Schedule, UpdateScheduleRequest} from "../types/schedule.ts";

export const getSchedules = () =>
    apiTrains.get<Schedule[]>(`/schedules`);

export const getScheduleById = (id: number) =>
    apiTrains.get<Schedule>(`/schedules/${id}`);

export const createSchedule = (data: CreateScheduleRequest) =>
    apiTrains.post(`/schedules`, data);

export const updateSchedule = (id: number, data: UpdateScheduleRequest) =>
    apiTrains.patch(`/schedules/${id}`, data);

export const deleteSchedule = (id: number) =>
    apiTrains.delete(`/schedules/${id}`);
