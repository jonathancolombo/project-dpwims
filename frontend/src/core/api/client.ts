// src/core/api/client.ts
import axios from "axios";

export const apiTrains = axios.create({
    baseURL: import.meta.env.VITE_API_URL_TRAINS,
});

export const apiUsers = axios.create({
    baseURL: import.meta.env.VITE_API_URL_USERS,
});