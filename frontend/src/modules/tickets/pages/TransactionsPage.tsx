import { useEffect, useState } from "react";
import MainLayout from "../../../core/layout/MainLayout";
import { getTickets, deleteTicket } from "../api/tickets_api.ts";
import { getPayments, deletePayment } from "../api/payments_api";
import { useNavigate } from "react-router-dom";
import type { Ticket } from "../types/ticket.ts";
import type { Payment } from "../types/payment";

export default function TransactionsPage() {
    const [activeTab, setActiveTab] = useState<"tickets" | "payments">("tickets");
    const [tickets, setTickets] = useState<Ticket[]>([]);
    const [payments, setPayments] = useState<Payment[]>([]);
    const [loading, setLoading] = useState(true);
    const navigate = useNavigate();
    const [message, setMessage] = useState("");

    useEffect(() => {
        loadData();
    }, []);

    const loadData = async () => {
        setLoading(true);
        try {
            const [ticketsResponse, paymentsResponse] = await Promise.all([
                getTickets(),
                getPayments()
            ]);
            setTickets(ticketsResponse.data);
            setPayments(paymentsResponse.data);
        } catch {
            setMessage("Errore nel caricamento dei dati.");
        } finally {
            setLoading(false);
        }
    };

    const handleDeleteTicket = async (uuid: string) => {
        if (!confirm("Sei sicuro di voler eliminare questo biglietto?")) return;

        try {
            await deleteTicket(uuid);
            setTickets(tickets => tickets.filter(ticket => ticket.uuid !== uuid));
        } catch {
            setMessage("Errore durante l'eliminazione del biglietto.");
        }
    };

    const handleDeletePayment = async (uuid: string) => {
        if (!confirm("Sei sicuro di voler eliminare questo pagamento?")) return;

        try {
            await deletePayment(uuid);
            setPayments(payments => payments.filter(payment => payment.uuid !== uuid));
        } catch {
            setMessage("Errore durante l'eliminazione del pagamento.");
        }
    };

    const ticketStatusLabels: Record<Ticket["status"], string> = {
        booked: "Prenotato",
        issued: "Utilizzato",
        cancelled: "Cancellato"
    };

    const paymentMethodLabels: Record<Payment["payment_method"], string> = {
        credit_card: "Carta di credito",
        banknotes: "Contanti",
        bank_transfer: "Bonifico bancario"
    };

    if (loading) {
        return <MainLayout>Caricamento...</MainLayout>;
    }

    return (
        <MainLayout>
            <div className="p-6 space-y-6">
                <h1 className="text-3xl font-bold">Transazioni</h1>

                {/* Tabs */}
                <div className="flex gap-2 border-b">
                    <button
                        onClick={() => setActiveTab("tickets")}
                        className={`px-4 py-2 font-semibold border-b-2 transition ${
                            activeTab === "tickets"
                                ? "border-blue-600 text-blue-600"
                                : "border-transparent text-gray-600 hover:text-gray-900"
                        }`}
                    >
                        Biglietti ({tickets.length})
                    </button>
                    <button
                        onClick={() => setActiveTab("payments")}
                        className={`px-4 py-2 font-semibold border-b-2 transition ${
                            activeTab === "payments"
                                ? "border-blue-600 text-blue-600"
                                : "border-transparent text-gray-600 hover:text-gray-900"
                        }`}
                    >
                        Pagamenti ({payments.length})
                    </button>
                </div>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">
                        {message}
                    </div>
                )}

                {/* Tickets Tab */}
                {activeTab === "tickets" && (
                    <div className="space-y-4">
                        <button
                            onClick={() => navigate("/tickets/create")}
                            className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700"
                        >
                            + Nuovo Biglietto
                        </button>

                        {tickets.length === 0 ? (
                            <div className="p-6 text-center text-gray-500 bg-gray-50 rounded-lg">
                                Nessun biglietto disponibile
                            </div>
                        ) : (
                            <div className="space-y-4">
                                {tickets.map(ticket => (
                                    <div
                                        key={ticket.uuid}
                                        className="p-4 bg-white shadow rounded-lg border space-y-4"
                                    >
                                        {/* Header */}
                                        <div className="flex justify-between items-center">
                                            <h2 className="text-lg font-semibold">
                                                Biglietto #{ticket.uuid}
                                            </h2>

                                            <span
                                                className={`px-2 py-1 rounded text-white ${
                                                    ticket.status === "booked" ? "bg-green-600"
                                                        : ticket.status === "issued" ? "bg-gray-500" : "bg-red-600"
                                                }`}
                                            >
                                                {ticketStatusLabels[ticket.status]}
                                            </span>
                                        </div>

                                        {/* Info */}
                                        <div className="text-gray-700 space-y-1">
                                            <div><strong>User ID:</strong> {ticket.user_id}</div>
                                            <div><strong>Train UUID:</strong> {ticket.train_id}</div>
                                            <div><strong>Schedule ID:</strong> {ticket.schedule_id}</div>
                                            <div><strong>Posto:</strong> {ticket.seat_number}</div>
                                            <div><strong>Prezzo:</strong> € {ticket.price}</div>
                                        </div>

                                        {/* Actions */}
                                        <div className="flex gap-3 pt-2">
                                            <button
                                                onClick={() => navigate(`/tickets/${ticket.uuid}`)}
                                                className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg"
                                            >
                                                Modifica
                                            </button>
                                            <button
                                                onClick={() => handleDeleteTicket(ticket.uuid)}
                                                className="flex-1 bg-red-600 hover:bg-red-700 text-white py-2 rounded-lg"
                                            >
                                                Elimina
                                            </button>
                                        </div>
                                    </div>
                                ))}
                            </div>
                        )}
                    </div>
                )}

                {/* Payments Tab */}
                {activeTab === "payments" && (
                    <div className="space-y-4">
                        <button
                            onClick={() => navigate("/payments/create")}
                            className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700"
                        >
                            + Nuovo Pagamento
                        </button>

                        {payments.length === 0 ? (
                            <div className="p-6 text-center text-gray-500 bg-gray-50 rounded-lg">
                                Nessun pagamento disponibile
                            </div>
                        ) : (
                            <div className="space-y-4">
                                {payments.map(payment => (
                                    <div
                                        key={payment.uuid}
                                        className="p-4 bg-white shadow rounded-lg border space-y-4"
                                    >
                                        {/* Header */}
                                        <div className="flex justify-between items-center">
                                            <h2 className="text-lg font-semibold">
                                                Pagamento #{payment.uuid}
                                            </h2>

                                            <span className="px-2 py-1 rounded bg-gray-200 text-gray-700">
                                                {paymentMethodLabels[payment.payment_method]}
                                            </span>
                                        </div>

                                        {/* Info */}
                                        <div className="text-gray-700 space-y-1">
                                            <div><strong>Ticket ID:</strong> {payment.ticket_id}</div>
                                            <div><strong>Importo:</strong> € {payment.amount}</div>
                                            <div><strong>Riferimento provider:</strong> {payment.provider_reference}</div>
                                        </div>

                                        {/* Actions */}
                                        <div className="flex gap-3 pt-2">
                                            <button
                                                onClick={() => navigate(`/payments/${payment.uuid}`)}
                                                className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg"
                                            >
                                                Modifica
                                            </button>

                                            <button
                                                onClick={() => handleDeletePayment(payment.uuid)}
                                                className="flex-1 bg-red-600 hover:bg-red-700 text-white py-2 rounded-lg"
                                            >
                                                Elimina
                                            </button>
                                        </div>
                                    </div>
                                ))}
                            </div>
                        )}
                    </div>
                )}
            </div>
        </MainLayout>
    );
}

