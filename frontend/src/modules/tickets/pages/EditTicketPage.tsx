import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import MainLayout from "../../../core/layout/MainLayout.tsx";
import { getTicket, updateTicket } from "../api/tickets_api.ts";
import type {Ticket} from "../types/ticket.ts";

export default function EditTicketPage() {
    const { uuid } = useParams<{ uuid: string }>(); // <-- deve chiamarsi uuid
    const navigate = useNavigate();

    const [ticket, setTicket] = useState<Ticket | null>(null);
    const [loading, setLoading] = useState(true);
    const [message, setMessage] = useState("");

    useEffect(() => {
        if (!uuid) {
            setMessage("UUID non presente nell'URL.");
            setLoading(false);
            return;
        }

        getTicket(uuid)
            .then(response => {
                setTicket(response.data);
                setLoading(false);
            })
            .catch(() => {
                setMessage("Errore nel caricamento del biglietto.");
                setLoading(false);
            });
    }, [uuid]);

    async function handleSave() {
        if (!ticket) return;

        try {
            await updateTicket(ticket.uuid, {
                seat_number: ticket.seat_number,
                price: ticket.price,
                status: ticket.status,
            });
            navigate("/tickets");
        } catch {
            setMessage("Errore durante il salvataggio.");
        }
    }

    if (loading) return <MainLayout>Caricamento...</MainLayout>;
    if (!ticket) return <MainLayout>{message || "Biglietto non trovato."}</MainLayout>;

    return (
        <MainLayout>
            <div className="p-6 space-y-6 max-w-xl mx-auto">
                <h1 className="text-3xl font-bold">Modifica Biglietto</h1>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">{message}</div>
                )}

                <div className="p-4 bg-white shadow rounded-lg border space-y-4">
                    <div>
                        <label className="block text-sm font-medium">UUID del biglietto</label>
                        <input
                            type="text"
                            value={ticket.uuid}
                            disabled
                            className="w-full mt-1 p-2 border rounded bg-gray-100"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">ID utente</label>
                        <input
                            type="number"
                            value={ticket.user_id}
                            disabled
                            className="w-full mt-1 p-2 border rounded bg-gray-100"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">UUID del treno</label>
                        <input
                            type="text"
                            value={ticket.train_id}
                            disabled
                            className="w-full mt-1 p-2 border rounded bg-gray-100"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">ID della fermata</label>
                        <input
                            type="number"
                            value={ticket.schedule_id}
                            disabled
                            className="w-full mt-1 p-2 border rounded bg-gray-100"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Numero Posto</label>
                        <input
                            type="text"
                            placeholder={"Es. A10"}
                            value={ticket.seat_number}
                            onChange={(element) =>
                                setTicket({ ...ticket, seat_number: element.target.value })
                            }
                            className="w-full mt-1 p-2 border rounded"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Prezzo (€)</label>
                        <input
                            type="number"
                            step="0.01"
                            value={ticket.price}
                            onChange={(element) =>
                                setTicket({ ...ticket, price: Number(element.target.value) })
                            }
                            className="w-full mt-1 p-2 border rounded"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Stato del biglietto</label>
                        <select
                            value={ticket.status}
                            onChange={(element) =>
                                setTicket({ ...ticket, status: element.target.value as Ticket["status"] })
                            }
                            className="w-full mt-1 p-2 border rounded"
                        >
                            <option value="booked">Prenotato</option>
                            <option value="issued">Utilizzato</option>
                            <option value="canceled">Cancellato</option>
                        </select>
                    </div>

                    <div className="flex gap-3 pt-4">
                        <button
                            onClick={handleSave}
                            className="flex-1 bg-blue-600 hover:bg-blue-700 text-white py-2 rounded-lg"
                        >
                            Salva
                        </button>

                        <button
                            onClick={() => navigate("/tickets")}
                            className="flex-1 bg-gray-200 hover:bg-gray-300 text-gray-700 py-2 rounded-lg"
                        >
                            Annulla
                        </button>
                    </div>
                </div>
            </div>
        </MainLayout>
    );
}
