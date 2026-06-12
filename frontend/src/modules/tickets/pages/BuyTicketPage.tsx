import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import MainLayout from "../../../core/layout/MainLayout";
import { createTicket } from "../api/tickets_api";
import { getScheduleById } from "../../trains/api/schedules_api";
import type { Schedule } from "../../trains/types/schedule";
import { user_authorization } from "../../../core/hooks/user_authorization";

export default function BuyTicketPage() {
    const { scheduleId } = useParams();
    const navigate = useNavigate();
    const { user } = user_authorization();
    const userId = user?.userID ?? 0;

    const [schedule, setSchedule] = useState<Schedule | null>(null);
    const [seatNumber, setSeatNumber] = useState("");
    const [loading, setLoading] = useState(true);
    const [saving, setSaving] = useState(false);
    const [error, setError] = useState("");
    const [message, setMessage] = useState("");

    useEffect(() => {
        async function loadSchedule() {
            if (!scheduleId) {
                setError("Itinerario non valido.");
                setLoading(false);
                return;
            }

            const numericId = Number(scheduleId);
            if (Number.isNaN(numericId) || numericId <= 0) {
                setError("Itinerario non valido.");
                setLoading(false);
                return;
            }

            try {
                const response = await getScheduleById(numericId);
                setSchedule(response.data);
            } catch {
                setError("Impossibile caricare i dettagli dell'itinerario.");
            } finally {
                setLoading(false);
            }
        }

        loadSchedule();
    }, [scheduleId]);

    if (!userId) {
        return (
            <MainLayout>
                <div className="p-6 max-w-2xl mx-auto space-y-4 text-center">
                    <h1 className="text-3xl font-bold">Acquista biglietto</h1>
                    <p className="text-gray-600">
                        Non riesco a leggere il tuo profilo utente in questo momento.
                    </p>
                    <button
                        onClick={() => navigate("/login?target=user")}
                        className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg"
                    >
                        Vai al login
                    </button>
                </div>
            </MainLayout>
        );
    }

    async function handlePurchase() {
        if (!schedule) {
            setError("Itinerario non disponibile.");
            return;
        }

        if (!seatNumber.trim()) {
            setError("Inserisci il numero di posto.");
            return;
        }

        setSaving(true);
        setError("");
        setMessage("");

        try {
            await createTicket({
                user_id: userId,
                schedule_id: schedule.id,
                train_id: schedule.train_id,
                seat_number: seatNumber.trim(),
                price: schedule.price,
                status: "booked",
            });

            setMessage("Biglietto acquistato con successo.");
            navigate("/user/tickets");
        } catch {
            setError("Errore durante l'acquisto del biglietto.");
        } finally {
            setSaving(false);
        }
    }

    if (loading) {
        return <MainLayout>Caricamento itinerario...</MainLayout>;
    }

    return (
        <MainLayout>
            <div className="p-6 max-w-2xl mx-auto space-y-6">
                <div className="space-y-2">
                    <h1 className="text-3xl font-bold">Acquista biglietto</h1>
                    <p className="text-gray-600">
                        Completa i dati per acquistare il biglietto dell'itinerario selezionato.
                    </p>
                </div>

                {error && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">
                        {error}
                    </div>
                )}

                {message && (
                    <div className="p-3 bg-green-100 text-green-700 rounded-lg">
                        {message}
                    </div>
                )}

                {schedule && (
                    <div className="bg-white p-6 rounded-xl shadow border space-y-4">
                        <div>
                            <p className="text-sm uppercase text-gray-500">Itinerario</p>
                            <h2 className="text-2xl font-semibold text-gray-900">
                                {schedule.departure} → {schedule.arrival}
                            </h2>
                        </div>

                        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 text-gray-700">
                            <div>
                                <strong>Train UUID:</strong> {schedule.train_id}
                            </div>
                            <div>
                                <strong>Prezzo:</strong> € {schedule.price.toFixed(2)}
                            </div>
                            <div>
                                <strong>Partenza:</strong> {schedule.departure}
                            </div>
                            <div>
                                <strong>Arrivo:</strong> {schedule.arrival}
                            </div>
                        </div>

                        <div>
                            <label htmlFor="seatNumber" className="block text-sm font-medium text-gray-700">
                                Numero posto
                            </label>
                            <input
                                id="seatNumber"
                                type="text"
                                value={seatNumber}
                                onChange={(element) => setSeatNumber(element.target.value)}
                                placeholder="Es. A10"
                                className="mt-1 w-full border rounded-lg p-2"
                            />
                        </div>

                        <div className="flex gap-3">
                            <button
                                onClick={handlePurchase}
                                disabled={saving}
                                className="flex-1 bg-blue-600 hover:bg-blue-700 text-white py-2 rounded-lg disabled:opacity-50"
                            >
                                {saving ? "Acquisto in corso..." : "Conferma acquisto"}
                            </button>

                            <button
                                onClick={() => navigate("/user/schedules")}
                                className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg"
                            >
                                Annulla
                            </button>
                        </div>

                        <button
                            onClick={() => navigate("/user/schedules")}
                            className="w-full text-sm text-blue-700 hover:text-blue-800 underline text-center"
                        >
                            Torna agli itinerari senza acquistare
                        </button>
                    </div>
                )}
            </div>
        </MainLayout>
    );
}






