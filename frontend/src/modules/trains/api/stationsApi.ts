import axios from "axios";

const API_URL = "http://localhost:8082";

export interface Station {
    id: number;
    name: string;
    city: string;
    region: string;
    status: "active" | "inactive";
}

export interface CreateStationRequest {
    name: string;
    city: string;
    region: string;
    status: "active" | "inactive";
}

export interface UpdateStationRequest {
    name?: string;
    city?: string;
    region?: string;
    status?: "active" | "inactive";
}

export const getStations = () =>
    axios.get<Station[]>(`${API_URL}/stations`);

export const getStationById = (id: number) =>
    axios.get<Station>(`${API_URL}/stations/${id}`);

export const createStation = (data: CreateStationRequest) =>
    axios.post(`${API_URL}/stations`, data);

export const updateStation = (id: number, data: UpdateStationRequest) =>
    axios.patch(`${API_URL}/stations/${id}`, data);

export const deleteStation = (id: number) =>
    axios.delete(`${API_URL}/stations/${id}`);

