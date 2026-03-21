import {useEffect, useState} from "react";
import MainLayout from "../../../core/layout/MainLayout.tsx";
import {deleteTicket, getTickets} from "../api/tickets_api.ts";
import {useNavigate} from "react-router-dom";
import type {Ticket} from "../types/ticket.ts";

export default function TicketsPage() {
    const [tickets, setTickets] = useState<Ticket[]>([]);
    const [loading, setLoading] = useState(true);
    const navigate = useNavigate();
    const [message, setMessage] = useState("");

    useEffect(() => {
        getTickets().then(response => {
            setTickets(response.data);
            setLoading(false);
        });
    }, []);

    if (loading) return <MainLayout>Caricamento...</MainLayout>;

    async function handleDelete(uuid: string) {
        if (!confirm("Sei sicuro di voler eliminare questo biglietto?")) return;

        try {
            await deleteTicket(uuid);
            setTickets(tickets => tickets.filter(ticket => ticket.uuid !== uuid));
        } catch {
            setMessage("Errore durante l'eliminazione.");
        }
    }

    const statusLabels: Record<Ticket["status"], string> = {
        booked: "Prenotato",
        issued: "Utilizzato",
        cancelled: "Cancellato"
    };

    return (
        <MainLayout>
            <div className="p-6 space-y-6">
                <h1 className="text-3xl font-bold">Biglietti</h1>
                <button
                    onClick={() => navigate("/tickets/create")}
                    className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700"
                >
                    + Nuovo Biglietto
                </button>
                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">
                        {message}
                    </div>
                )}

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
                                    {statusLabels[ticket.status]}
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
                                    onClick={() => handleDelete(ticket.uuid)}
                                    className="flex-1 bg-red-600 hover:bg-red-700 text-white py-2 rounded-lg"
                                >
                                    Elimina
                                </button>
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </MainLayout>
    );
}
