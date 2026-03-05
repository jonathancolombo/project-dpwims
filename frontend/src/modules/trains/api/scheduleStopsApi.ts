import axios from "axios";

const API_URL = "http://localhost:8082";

export interface ScheduleStop {
    id: number;
    schedule_id: number;
    station_id: number;
    station_name: string;
    stop_order: number;
    arrival_time: string;
    departure_time: string;
}

export interface CreateScheduleStopRequest {
    schedule_id: number;
    station_id: number;
    arrival_time: string;
    departure_time: string;
}

export interface UpdateScheduleStopRequest {
    station_id?: number;
    arrival_time?: string;
    departure_time?: string;
    stop_order?: number;
}

// GET fermate di uno schedule
export const getStopsBySchedule = (scheduleId: number) =>
    axios.get<ScheduleStop[]>(`${API_URL}/stopschedules/schedule/${scheduleId}`);

// CREATE fermata
export const createStop = (scheduleId: number, data: Omit<CreateScheduleStopRequest, "schedule_id">) =>
    axios.post(`${API_URL}/stopschedules`, {
        schedule_id: scheduleId,
        ...data,
    });

// UPDATE fermata
export const updateStop = (stopId: number, data: number) =>
    axios.patch(`${API_URL}/stopschedules/${stopId}`, data);

// DELETE fermata
export const deleteStop = (stopId: number) =>
    axios.delete(`${API_URL}/stopschedules/${stopId}`);
