import { useEffect, useState } from "react";
import MainLayout from "../../../core/layout/MainLayout";
import { useNavigate } from "react-router-dom";
import type {Ticket} from "../types/ticket";
import { apiTickets } from "../../../core/api/client";
import { user_authorization } from "../../../core/hooks/user_authorization";

export default function MyTicketsPage() {
    const [tickets, setTickets] = useState<Ticket[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState("");

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
        if (!user?.userID) {
            return;
        }

        apiTickets
            .get<Ticket[]>(`/tickets/user/${user.userID}`)
            .then((response) => setTickets(response.data))
            .catch((error) => {
                console.error(error);
                setError("Impossibile caricare i tuoi biglietti");
            })
            .finally(() => setLoading(false));
    }, [user?.userID]);

    if (!user?.userID) {
        return null;
    }

    if (loading) {
        return <MainLayout>Caricamento...</MainLayout>;
    }

    return (
        <MainLayout>
            <div className="p-6 space-y-6 max-w-3xl mx-auto">

                <h1 className="text-3xl font-bold text-center">I miei biglietti</h1>

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

                        <button
                            onClick={() => navigate("/user/schedules")}
                            className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
                        >
                            Cerca un itinerario e acquista il tuo primo biglietto
                        </button>
                    </div>
                )}

                {/* LISTA BIGLIETTI */}
                <div className="space-y-4">
                    {tickets.map(ticket => (
                        <div
                            key={ticket.uuid}
                            className="p-4 bg-white shadow rounded-lg border space-y-4"
                        >
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

                            <div className="text-gray-700 space-y-1">
                                <div><strong>Treno:</strong> {ticket.train_id}</div>
                                <div><strong>Itinerario:</strong> {ticket.schedule_id}</div>
                                <div><strong>Posto:</strong> {ticket.seat_number}</div>
                                <div><strong>Prezzo:</strong> € {ticket.price}</div>
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </MainLayout>
    );
}
