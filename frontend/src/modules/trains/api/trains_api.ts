import {apiTrains} from "../../../core/api/client";
import type {CreateTrainDTO, Train} from "../types/train.ts";

export const getTrains = () => apiTrains.get<Train[]>("/trains");

export const patchTrain = (uuid: string, data: any) =>
    apiTrains.patch(`/trains/${uuid}`, data);

export const deleteTrain = (uuid: string) =>
    apiTrains.delete(`/trains/${uuid}`);

export const createTrain = (data: CreateTrainDTO) =>
    apiTrains.post("/trains", data);

