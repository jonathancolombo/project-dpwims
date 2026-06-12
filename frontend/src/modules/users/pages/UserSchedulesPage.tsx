import { useEffect, useState } from "react";
import { getSchedules } from "../../trains/api/schedules_api.ts";
import { useNavigate } from "react-router-dom";
import type {Schedule} from "../../trains/types/schedule.ts";
import { createTicket } from "../../tickets/api/tickets_api.ts";
import { user_authorization } from "../../../core/hooks/user_authorization";

export default function UserSchedulesPage() {
    const [schedules, setSchedules] = useState<Schedule[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState("");
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

    async function handleBuy(schedule: Schedule) {
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

    if (loading) {
        return <div className="p-6 text-center text-gray-600">Caricamento itinerari...</div>;
    }

    if (error) {
        return <div className="p-6 text-center text-red-600">{error}</div>;
    }

    return (
        <div className="max-w-4xl mx-auto p-6 space-y-6">
            <h1 className="text-3xl font-bold text-gray-900">Itinerari disponibili</h1>

            <div className="space-y-4">
                {schedules.map((schedule) => (
                    <div
                        key={schedule.id}
                        className="p-5 bg-white rounded-xl shadow border flex justify-between items-center"
                    >
                        <div className="space-y-1">
                            <h2 className="text-xl font-semibold text-gray-900">
                                {schedule.departure} → {schedule.arrival}
                            </h2>

                            <p className="text-gray-700 font-medium">
                                Prezzo: {schedule.price.toFixed(2)} €
                            </p>
                        </div>

                        <button
                            onClick={() => handleBuy(schedule)}
                            disabled={savingScheduleId === schedule.id}
                            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
                        >
                            {savingScheduleId === schedule.id ? "Acquisto..." : "Acquista"}
                        </button>
                    </div>
                ))}
            </div>
        </div>
    );
}
