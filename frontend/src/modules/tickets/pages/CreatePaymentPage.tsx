import { useState, useEffect } from "react";
import MainLayout from "../../../core/layout/MainLayout";
import { createPayment } from "../api/payments_api";
import { getTickets } from "../api/tickets_api.ts";
import { useNavigate } from "react-router-dom";
import type { Ticket } from "../types/ticket.ts";
import type { CreatePaymentRequest } from "../types/payment";

export default function CreatePaymentPage() {
    const navigate = useNavigate();

    const [tickets, setTickets] = useState<Ticket[]>([]);
    const [ticketId, setTicketId] = useState("");
    const [amount, setAmount] = useState(0);
    const [paymentMethod, setPaymentMethod] = useState<"credit_card" | "banknotes" | "bank_transfer">("credit_card");
    const [providerReference, setProviderReference] = useState("");

    const [message, setMessage] = useState("");

    useEffect(() => {
        getTickets()
            .then(response => {
                const bookedTickets = response.data.filter(t => t.status === "booked");
                setTickets(bookedTickets);
            })
            .catch(() => setMessage("Errore nel caricamento dei ticket."));
    }, []);


    const handleCreate = async () => {
        if (!ticketId || amount <= 0 || !providerReference) {
            setMessage("Tutti i campi sono obbligatori.");
            return;
        }

        const payload: CreatePaymentRequest = {
            ticket_id: ticketId,
            amount,
            payment_method: paymentMethod,
            provider_reference: providerReference
        };

        try {
            await createPayment(payload);
            navigate("/payments");
        } catch {
            setMessage("Errore durante la creazione del pagamento.");
        }
    };

    return (
        <MainLayout>
            <div className="p-6 max-w-xl mx-auto space-y-6">
                <h1 className="text-3xl font-bold">Nuovo Pagamento</h1>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">
                        {message}
                    </div>
                )}

                <div className="space-y-4 bg-white p-6 rounded-xl shadow border">

                    {/* Ticket */}
                    <div>
                        <label className="block text-sm font-medium">Ticket</label>
                        <select
                            className="w-full border p-2 rounded mt-1"
                            value={ticketId}
                            onChange={(element) => {
                                const id = element.target.value;
                                setTicketId(id);

                                const selected = tickets.find(ticket => ticket.uuid === id);
                                if (selected) {
                                    setAmount(selected.price);
                                }
                            }}
                        >

                        <option value="">Seleziona un ticket</option>
                            {tickets.map(ticket => (
                                <option key={ticket.uuid} value={ticket.uuid}>
                                    {ticket.uuid} — Utente {ticket.user_id} — €{ticket.price}
                                </option>
                            ))}
                        </select>
                    </div>

                    {/* Amount */}
                    <div>
                        <label className="block text-sm font-medium">Importo (€)</label>
                        <input
                            type="number"
                            className="w-full border p-2 rounded mt-1"
                            value={amount}
                            disabled
                            onChange={(element) => setAmount(Number(element.target.value))}
                        />
                    </div>

                    {/* Payment Method */}
                    <div>
                        <label className="block text-sm font-medium">Metodo di pagamento</label>
                        <select
                            className="w-full border p-2 rounded mt-1"
                            value={paymentMethod}
                            onChange={(element) =>
                                setPaymentMethod(element.target.value as "credit_card" | "banknotes" | "bank_transfer")
                            }
                        >
                            <option value="credit_card">Carta di credito</option>
                            <option value="banknotes">Contanti</option>
                            <option value="bank_transfer">Bonifico bancario</option>
                        </select>
                    </div>

                    {/* Provider Reference */}
                    <div>
                        <label className="block text-sm font-medium">Riferimento provider</label>
                        <input
                            className="w-full border p-2 rounded mt-1"
                            value={providerReference}
                            onChange={(element) => setProviderReference(element.target.value)}
                        />
                    </div>

                </div>

                <button
                    onClick={handleCreate}
                    className="w-full bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700"
                >
                    Crea Pagamento
                </button>

                <button
                    onClick={() => navigate("/payments")}
                    className="w-full bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg"
                >
                    Annulla
                </button>
            </div>
        </MainLayout>
    );
}
