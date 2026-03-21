export interface Train {
    uuid: string;
    train_number: string;
    type: string;
    capacity: number;
    status: string;
}

export type CreateTrainDTO = Partial<Omit<Train, "uuid">>;
