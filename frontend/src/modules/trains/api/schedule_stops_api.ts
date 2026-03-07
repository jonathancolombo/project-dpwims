import axios from "axios";
import type {CreateScheduleStopRequest, ScheduleStop} from "../types/schedule_stop.ts";

const API_URL = "http://localhost:8082";

export const getStopsBySchedule = (scheduleId: number) =>
    axios.get<ScheduleStop[]>(`${API_URL}/stopschedules/schedule/${scheduleId}`);

export const createStop = (scheduleId: number, data: Omit<CreateScheduleStopRequest, "schedule_id">) =>
    axios.post(`${API_URL}/stopschedules`, {
        schedule_id: scheduleId,
        ...data,
    });

export const updateStop = (stopId: number, data: number) =>
    axios.patch(`${API_URL}/stopschedules/${stopId}`, data);

export const deleteStop = (stopId: number) =>
    axios.delete(`${API_URL}/stopschedules/${stopId}`);
