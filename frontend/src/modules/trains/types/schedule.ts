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