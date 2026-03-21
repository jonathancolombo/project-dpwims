import {apiTrains} from "../../../core/api/client";
import type {CreateStationRequest, Station, UpdateStationRequest} from "../types/station.ts";

export const getStations = () =>
    apiTrains.get<Station[]>(`/stations`);

export const getStationById = (id: number) =>
    apiTrains.get<Station>(`/stations/${id}`);

export const createStation = (data: CreateStationRequest) =>
    apiTrains.post(`/stations`, data);

export const updateStation = (id: number, data: UpdateStationRequest) =>
    apiTrains.patch(`/stations/${id}`, data);

export const deleteStation = (id: number) =>
    apiTrains.delete(`/stations/${id}`);

