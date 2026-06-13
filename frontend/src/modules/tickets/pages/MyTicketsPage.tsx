import { useEffect, useState } from "react";
import MainLayout from "../../../core/layout/MainLayout";
import { useNavigate } from "react-router-dom";
import type { Ticket } from "../types/ticket";
import { apiTickets } from "../../../core/api/client";
import { user_authorization } from "../../../core/hooks/user_authorization";

export default function MyTicketsPage() {
    const [tickets, setTickets] = useState<Ticket[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState("");
    const [deleting, setDeleting] = useState<string | null>(null);

    const navigate = useNavigate();
    const { user } = user_authorization();

    const statusLabels: Record<string, string> = {
        booked: "Prenotato",
        issued: "Utilizzato",
        used: "Utilizzato",
        cancelled: "Cancellato",
        canceled: "Cancellato",
    };

    const statusClasses: Record<string, string> = {
        booked: "bg-green-600",
        issued: "bg-gray-500",
        used: "bg-gray-500",
        cancelled: "bg-red-600",
        canceled: "bg-red-600",
    };

    useEffect(() => {
        if (!user?.userID) return;

        apiTickets
            .get<Ticket[]>(`/tickets/user/${user.userID}`)
            .then((res) => setTickets(res.data))
            .catch(() => setError("Impossibile caricare i tuoi biglietti"))
            .finally(() => setLoading(false));
    }, [user?.userID]);

    async function handleDelete(uuid: string) {
        if (!confirm("Sei sicuro di voler eliminare questo biglietto?")) return;

        setDeleting(uuid);

        try {
            await apiTickets.delete(`/tickets/${uuid}`);
            setTickets((prev) => prev.filter((t) => t.uuid !== uuid));
        } catch {
            setError("Errore durante l'eliminazione del biglietto.");
        } finally {
            setDeleting(null);
        }
    }

    if (loading) {
        return <MainLayout>Caricamento...</MainLayout>;
    }

    return (
        <MainLayout>
            <div className="p-6 space-y-6 max-w-3xl mx-auto">

                {/* HEADER + BACK BUTTON */}
                <div className="flex justify-between items-center">
                    <h1 className="text-3xl font-bold">I miei biglietti</h1>

                    <button
                        onClick={() => navigate(-1)}
                        className="px-4 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300"
                    >
                        ← Torna indietro
                    </button>
                </div>

                {/* ERRORE */}
                {error && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg text-center">
                        {error}
                    </div>
                )}

                {/* EMPTY STATE */}
                {!error && tickets.length === 0 && (
                    <div className="text-center text-gray-500 py-10">
                        <p className="text-lg mb-4">Non hai ancora acquistato biglietti.</p>
                    </div>
                )}

                {/* LISTA BIGLIETTI */}
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
                                    className={`px-2 py-1 rounded text-white ${statusClasses[ticket.status] ?? "bg-gray-500"}`}
                                >
                                    {statusLabels[ticket.status] ?? ticket.status}
                                </span>
                            </div>

                            {/* Info */}
                            <div className="text-gray-700 space-y-1">
                                <div><strong>Treno:</strong> {ticket.train_id}</div>
                                <div><strong>Itinerario:</strong> {ticket.schedule_id}</div>
                                <div><strong>Posto:</strong> {ticket.seat_number}</div>
                                <div><strong>Prezzo:</strong> € {ticket.price}</div>
                            </div>

                            {/* Actions */}
                            <div className="flex gap-3 pt-2">
                                <button
                                    onClick={() => handleDelete(ticket.uuid)}
                                    disabled={deleting === ticket.uuid}
                                    className="flex-1 bg-red-600 hover:bg-red-700 text-white py-2 rounded-lg disabled:opacity-50"
                                >
                                    {deleting === ticket.uuid ? "Eliminazione..." : "Elimina"}
                                </button>
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </MainLayout>
    );
}
