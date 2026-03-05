import axios from "axios";

const API_URL = "http://localhost:8082";

export interface Schedule {
    id: number;
    train_id: string;
    station_id: number;
    departure: string;
    arrival: string;
    status: "active" | "inactive";
    price: number;
}

export interface CreateScheduleRequest {
    train_id: string;
    station_id: number;
    departure: string;
    arrival: string;
    status: "active" | "inactive";
    price: number;
}

export interface UpdateScheduleRequest {
    train_id?: string;
    station_id?: number;
    departure?: string;
    arrival?: string;
    status?: "active" | "inactive";
    price?: number;
}

export const getSchedules = () =>
    axios.get<Schedule[]>(`${API_URL}/schedules`);

export const getScheduleById = (id: number) =>
    axios.get<Schedule>(`${API_URL}/schedules/${id}`);

export const createSchedule = (data: CreateScheduleRequest) =>
    axios.post(`${API_URL}/schedules`, data);

export const updateSchedule = (id: number, data: UpdateScheduleRequest) =>
    axios.patch(`${API_URL}/schedules/${id}`, data);

export const deleteSchedule = (id: number) =>
    axios.delete(`${API_URL}/schedules/${id}`);
