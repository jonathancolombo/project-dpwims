// src/modules/trains/api/trainsApi.ts
import {apiTrains} from "../../../core/api/client";
import type {Train} from "../types/Train.ts";

export const getTrains = () => apiTrains.get<Train[]>("/trains");
export const patchTrain = (uuid: string, data: any) =>
    apiTrains.patch(`/trains/${uuid}`, data);
export const deleteTrain = (uuid: string) =>
    apiTrains.delete(`/trains/${uuid}`);
export const createTrain = (data: {
    train_number: string;
    type: string;
    capacity: number;
    status: string;
}) => apiTrains.post("/trains", data);

