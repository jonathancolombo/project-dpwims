export interface Payment {
    uuid: string;
    ticket_id: string;
    amount: number;
    payment_method: "credit_card" | "banknotes" | "bank_transfer";
    provider_reference: string;
}

export interface CreatePaymentRequest {
    ticket_id: string;
    amount: number;
    payment_method: "credit_card" | "banknotes" | "bank_transfer";
    provider_reference: string;
}

export interface UpdatePaymentRequest extends CreatePaymentRequest {}
