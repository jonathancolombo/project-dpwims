export interface Ticket {
    uuid: string;
    user_id: number;
    train_id: string;
    schedule_id: number;
    seat_number: string;
    price: number;
    status: "booked" | "cancelled" | "issued";
}

