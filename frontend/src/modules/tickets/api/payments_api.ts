import type { Payment, CreatePaymentRequest, UpdatePaymentRequest } from "../types/payment";
import {apiTickets} from "../../../core/api/client.ts";

export const getPayments = () =>
    apiTickets.get<Payment[]>("/payments");

export const getPayment = (uuid: string) =>
    apiTickets.get<Payment>(`/payments/${uuid}`);

export const createPayment = (data: CreatePaymentRequest) =>
    apiTickets.post("/payments", data);

export const updatePayment = (uuid: string, data: UpdatePaymentRequest) =>
    apiTickets.patch(`/payments/${uuid}`, data);

export const deletePayment = (uuid: string) =>
    apiTickets.delete(`/payments/${uuid}`);
