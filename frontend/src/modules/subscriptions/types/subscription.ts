export type Plan = "basic" | "premium" | "full";

export interface Subscription {
    id: number;
    user_id: number;
    train_uuid: string;
    plan: Plan;
}

export interface CreateSubscriptionDTO {
    user_id: number;
    train_uuid: string;
    plan: Plan;
}

export const planLabels = {
    basic: "Base",
    premium: "Premium",
    full: "Completo"
} as const;


