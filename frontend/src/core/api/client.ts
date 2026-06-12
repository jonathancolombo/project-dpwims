import axios, {type AxiosInstance} from "axios";

export const apiTrains : AxiosInstance = axios.create({
    baseURL: import.meta.env.VITE_API_URL_TRAINS,
});

apiTrains.interceptors.request.use((config) => {
    const token = localStorage.getItem("token");
    if (token) {
        config.headers["Authorization"] = `Bearer ${token}`;
    }
    return config;
});

export const apiUsers : AxiosInstance = axios.create({
    baseURL: import.meta.env.VITE_API_URL_USERS,
});

apiUsers.interceptors.request.use((config) => {
    const token = localStorage.getItem("token");
    if (token) {
        config.headers["Authorization"] = `Bearer ${token}`;
    }
    return config;
});

export const apiTickets : AxiosInstance = axios.create({
    baseURL: import.meta.env.VITE_API_URL_TICKETS,
})

apiTickets.interceptors.request.use((config) => {
    const token = localStorage.getItem("token");
    if (token) {
        config.headers["Authorization"] = `Bearer ${token}`;
    }
    return config;
});

export const apiNotifications : AxiosInstance = axios.create({
    baseURL: import.meta.env.VITE_API_URL_NOTIFICATIONS,
})

apiNotifications.interceptors.request.use((config) => {
    const token = localStorage.getItem("token");
    if (token) {
        config.headers["Authorization"] = `Bearer ${token}`;
    }
    return config;
});