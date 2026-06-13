export interface Subscription {
    id: number;
    user_id: number;
    train_uuid: string;
    schedule_id: number;
}

export interface CreateSubscriptionDTO {
    user_id: number;
    train_uuid: string;
    schedule_id: number;
}