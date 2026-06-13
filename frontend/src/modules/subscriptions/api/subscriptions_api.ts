import { apiNotifications } from "../../../core/api/client";
import type { Subscription, CreateSubscriptionDTO } from "../types/subscription";

export const getSubscriptions = (userId?: number) =>
    apiNotifications.get<Subscription[]>(
        userId ? `/subscriptions?user_id=${userId}` : "/subscriptions"
    );

export const getSubscriptionsByTrain = (trainUUID: string) =>
    apiNotifications.get<Subscription[]>(`/subscriptions/train/${trainUUID}`);

export const getSubscriptionsBySchedule = (scheduleID: number) =>
    apiNotifications.get<Subscription[]>(`/subscriptions/schedule/${scheduleID}`);

export const createSubscription = (data: CreateSubscriptionDTO) =>
    apiNotifications.post("/subscriptions", data);

export const deleteSubscription = (id: number) =>
    apiNotifications.delete(`/subscriptions/${id}`);