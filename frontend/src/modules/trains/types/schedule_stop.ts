export interface ScheduleStop {
    id: number;
    schedule_id: number;
    station_id: number;
    station_name: string;
    stop_order: number;
    arrival_time: string;
    departure_time: string;
}


export interface CreateScheduleStopRequest {
    station_id: number;
    arrival_time: string;
    departure_time: string;
}

