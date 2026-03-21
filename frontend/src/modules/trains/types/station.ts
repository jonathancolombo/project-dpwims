export interface Station {
    id: number;
    name: string;
    city: string;
    region: string;
    status: "active" | "inactive";
}

export interface CreateStationRequest {
    name: string;
    city: string;
    region: string;
    status: "active" | "inactive";
}

export interface UpdateStationRequest {
    name?: string;
    city?: string;
    region?: string;
    status?: "active" | "inactive";
}