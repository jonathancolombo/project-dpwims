import axios from "axios";

const API_URL = "http://localhost:8083";

export interface Ticket {
    uuid: string;
    user_id: number;
    train_id: string;
    schedule_id: number;
    seat_number: number;
    price: number;
    status: "booked" | "canceled" | "used";
}

export const getTickets = () =>
    axios.get<Ticket[]>(`${API_URL}/tickets`);
