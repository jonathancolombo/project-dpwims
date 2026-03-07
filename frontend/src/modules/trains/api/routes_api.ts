import axios from "axios";
import type {CreateRouteRequest, Route, UpdateRouteRequest} from "../types/route.ts";

const API_URL = "http://localhost:8082";

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
