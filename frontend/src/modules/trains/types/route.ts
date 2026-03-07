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