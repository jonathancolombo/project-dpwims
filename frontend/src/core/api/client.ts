import axios from "axios";

export const apiTrains = axios.create({
    baseURL: import.meta.env.VITE_API_URL_TRAINS,
});

export const apiUsers = axios.create({
    baseURL: import.meta.env.VITE_API_URL_USERS,
});

export const apiTickets = axios.create({
    baseURL: import.meta.env.VITE_API_URL_TICKETS,
})

export const apiNotifications = axios.create({
    baseURL: import.meta.env.VITE_API_URL_NOTIFICATIONS,
})