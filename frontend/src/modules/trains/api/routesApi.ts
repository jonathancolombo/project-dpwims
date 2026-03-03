import axios from "axios";

const API_URL = "http://localhost:8082";

export interface Route {
    id: number;
    train_id: string;
    departure_station: string;
    arrival_station: string;
    distance: number;
}

export interface CreateRouteRequest {
    train_id: string;
    departure_station: string;
    arrival_station: string;
    distance: number;
}

export interface UpdateRouteRequest {
    train_id?: string;
    departure_station?: string;
    arrival_station?: string;
    distance?: number;
}

export const getRoutes = () =>
    axios.get<Route[]>(`${API_URL}/routes`);

export const getRouteById = (id: number) =>
    axios.get<Route>(`${API_URL}/routes/${id}`);

export const createRoute = (data: CreateRouteRequest) =>
    axios.post(`${API_URL}/routes`, data);

export const updateRoute = (id: number, data: UpdateRouteRequest) =>
    axios.patch(`${API_URL}/routes/${id}`, data);

export const deleteRoute = (id: number) =>
    axios.delete(`${API_URL}/routes/${id}`);
