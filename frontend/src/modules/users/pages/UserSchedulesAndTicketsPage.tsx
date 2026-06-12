import { useEffect, useState } from "react";
import { getSchedules } from "../../trains/api/schedules_api.ts";
import { useNavigate } from "react-router-dom";
import type { Schedule } from "../../trains/types/schedule.ts";
import { createTicket } from "../../tickets/api/tickets_api.ts";
import { user_authorization } from "../../../core/hooks/user_authorization";
import MainLayout from "../../../core/layout/MainLayout.tsx";
import { ChevronDown, ChevronUp } from "lucide-react";

export default function UserSchedulesAndTicketsPage() {
    const [schedules, setSchedules] = useState<Schedule[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState("");
    const [expandedId, setExpandedId] = useState<number | null>(null);
    const [seatNumber, setSeatNumber] = useState<Record<number, string>>({});
    const [savingScheduleId, setSavingScheduleId] = useState<number | null>(null);
    const navigate = useNavigate();
    const { user } = user_authorization();
    const userId = user?.userID ?? 0;

    useEffect(() => {
        getSchedules()
            .then((res) => setSchedules(res.data))
            .catch(() => setError("Errore nel caricamento degli itinerari."))
            .finally(() => setLoading(false));
    }, []);

    async function handleQuickBuy(schedule: Schedule) {
        if (!userId) {
            return;
        }

        setSavingScheduleId(schedule.id);
        setError("");

        try {
            const autoSeat = `A${String((schedule.id % 99) + 1).padStart(2, "0")}`;

            await createTicket({
                user_id: userId,
                schedule_id: schedule.id,
                train_id: schedule.train_id,
                seat_number: autoSeat,
                price: schedule.price,
                status: "booked",
            });

            navigate("/user/tickets");
        } catch {
            setError("Errore durante l'acquisto del biglietto.");
        } finally {
            setSavingScheduleId(null);
        }
    }

    async function handleDetailedBuy(schedule: Schedule) {
        if (!userId) {
            return;
        }

        const seat = seatNumber[schedule.id]?.trim();
        if (!seat) {
            setError("Inserisci il numero di posto.");
            return;
        }

        setSavingScheduleId(schedule.id);
        setError("");

        try {
            await createTicket({
                user_id: userId,
                schedule_id: schedule.id,
                train_id: schedule.train_id,
                seat_number: seat,
                price: schedule.price,
                status: "booked",
            });

            navigate("/user/tickets");
        } catch {
            setError("Errore durante l'acquisto del biglietto.");
        } finally {
            setSavingScheduleId(null);
        }
    }

    if (loading) {
        return (
            <MainLayout>
                <div className="p-6 text-center text-gray-600">Caricamento itinerari...</div>
            </MainLayout>
        );
    }

    if (error && schedules.length === 0) {
        return (
            <MainLayout>
                <div className="p-6 text-center text-red-600">{error}</div>
            </MainLayout>
        );
    }

    return (
        <MainLayout>
            <div className="max-w-4xl mx-auto p-6 space-y-6">
                <div className="space-y-2">
                    <h1 className="text-3xl font-bold text-gray-900">Itinerari e Acquista Biglietti</h1>
                    <p className="text-gray-600">
                        Seleziona un itinerario per acquistare il biglietto.
                    </p>
                </div>

                {error && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">
                        {error}
                    </div>
                )}

                <div className="space-y-3">
                    {schedules.map((schedule) => (
                        <div
                            key={schedule.id}
                            className="bg-white rounded-xl shadow border overflow-hidden"
                        >
                            {/* Header cliccabile */}
                            <button
                                onClick={() =>
                                    setExpandedId(expandedId === schedule.id ? null : schedule.id)
                                }
                                className="w-full p-5 flex justify-between items-center hover:bg-gray-50 transition"
                            >
                                <div className="text-left">
                                    <h2 className="text-xl font-semibold text-gray-900">
                                        {schedule.departure} → {schedule.arrival}
                                    </h2>
                                    <p className="text-gray-600 mt-1">
                                        Prezzo: € {schedule.price.toFixed(2)}
                                    </p>
                                </div>

                                {expandedId === schedule.id ? (
                                    <ChevronUp className="w-6 h-6 text-gray-600" />
                                ) : (
                                    <ChevronDown className="w-6 h-6 text-gray-600" />
                                )}
                            </button>

                            {/* Dettagli espanduti */}
                            {expandedId === schedule.id && (
                                <div className="border-t p-5 space-y-4 bg-gray-50">
                                    <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 text-gray-700">
                                        <div>
                                            <p className="text-sm text-gray-600">Train UUID</p>
                                            <p className="font-medium">{schedule.train_id}</p>
                                        </div>
                                        <div>
                                            <p className="text-sm text-gray-600">Prezzo</p>
                                            <p className="font-medium">€ {schedule.price.toFixed(2)}</p>
                                        </div>
                                        <div>
                                            <p className="text-sm text-gray-600">Partenza</p>
                                            <p className="font-medium">{schedule.departure}</p>
                                        </div>
                                        <div>
                                            <p className="text-sm text-gray-600">Arrivo</p>
                                            <p className="font-medium">{schedule.arrival}</p>
                                        </div>
                                    </div>

                                    {/* Form acquisto */}
                                    <div className="space-y-3 pt-4 border-t">
                                        <div>
                                            <label
                                                htmlFor={`seat-${schedule.id}`}
                                                className="block text-sm font-medium text-gray-700 mb-1"
                                            >
                                                Numero posto (opzionale)
                                            </label>
                                            <input
                                                id={`seat-${schedule.id}`}
                                                type="text"
                                                value={seatNumber[schedule.id] || ""}
                                                onChange={(e) =>
                                                    setSeatNumber({
                                                        ...seatNumber,
                                                        [schedule.id]: e.target.value,
                                                    })
                                                }
                                                placeholder="Es. A10 (lascia vuoto per assegnazione automatica)"
                                                className="w-full border rounded-lg p-2 text-sm"
                                            />
                                            <p className="text-xs text-gray-500 mt-1">
                                                Se lasci vuoto, assegneremo automaticamente un posto.
                                            </p>
                                        </div>

                                        <div className="flex gap-3">
                                            {seatNumber[schedule.id]?.trim() ? (
                                                // Acquisto con posto personalizato
                                                <button
                                                    onClick={() => handleDetailedBuy(schedule)}
                                                    disabled={savingScheduleId === schedule.id}
                                                    className="flex-1 bg-blue-600 hover:bg-blue-700 text-white py-2 rounded-lg disabled:opacity-50 font-medium"
                                                >
                                                    {savingScheduleId === schedule.id
                                                        ? "Acquisto in corso..."
                                                        : "Conferma acquisto"}
                                                </button>
                                            ) : (
                                                // Acquisto veloce con auto-assign
                                                <button
                                                    onClick={() => handleQuickBuy(schedule)}
                                                    disabled={savingScheduleId === schedule.id}
                                                    className="flex-1 bg-blue-600 hover:bg-blue-700 text-white py-2 rounded-lg disabled:opacity-50 font-medium"
                                                >
                                                    {savingScheduleId === schedule.id
                                                        ? "Acquisto in corso..."
                                                        : "Acquista"}
                                                </button>
                                            )}

                                            <button
                                                onClick={() => setExpandedId(null)}
                                                className="flex-1 bg-gray-200 hover:bg-gray-300 text-gray-700 py-2 rounded-lg font-medium"
                                            >
                                                Annulla
                                            </button>
                                        </div>
                                    </div>
                                </div>
                            )}
                        </div>
                    ))}
                </div>

                {schedules.length === 0 && (
                    <div className="p-6 text-center text-gray-500 bg-gray-50 rounded-lg">
                        Nessun itinerario disponibile al momento.
                    </div>
                )}
            </div>
        </MainLayout>
    );
}

