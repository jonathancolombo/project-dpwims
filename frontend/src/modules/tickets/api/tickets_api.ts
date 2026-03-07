import {apiTickets} from "../../../core/api/client";
import type {Ticket} from "../types/ticket.ts";


export const getTicket = (uuid: string) =>
    apiTickets.get<Ticket>(`/tickets/${uuid}`);

export const getTickets = () =>
    apiTickets.get<Ticket[]>(`/tickets`);

export const deleteTicket = (uuid: string) =>
    apiTickets.delete<Ticket>(`/tickets/${uuid}`)

export const updateTicket = (uuid: string, data: Partial<Ticket>) =>
    apiTickets.patch(`/tickets/${uuid}`, data);