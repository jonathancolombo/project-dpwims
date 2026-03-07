import {useEffect, useState} from "react";
import MainLayout from "../../../core/layout/MainLayout";
import {deleteSchedule, getSchedules, type Schedule} from "../api/schedules_api.ts";
import {useNavigate} from "react-router-dom";

export default function SchedulesPage() {
    const [schedules, setSchedules] = useState<Schedule[]>([]);
    const [loading, setLoading] = useState(true);
    const [message, setMessage] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        getSchedules()
            .then((res) => setSchedules(res.data))
            .catch(() => setMessage("Errore nel caricamento degli schedule."))
            .finally(() => setLoading(false));
    }, []);

    const handleDelete = async (id: number) => {
        if (!confirm("Eliminare questo schedule?")) return;

        try {
            await deleteSchedule(id);
            setSchedules((prev) => prev.filter((s) => s.id !== id));
        } catch {
            setMessage("Errore durante l'eliminazione.");
        }
    };

    if (loading) {
        return (
            <MainLayout>
                <div className="p-6 text-center text-gray-600">Caricamento...</div>
            </MainLayout>
        );
    }

    return (
        <MainLayout>
            <div className="p-6 max-w-5xl mx-auto space-y-6">
                <div className="flex justify-between items-center">
                    <h1 className="text-3xl font-bold">Fermate</h1>
                    <button
                        onClick={() => navigate("/schedules/create")}
                        className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700"
                    >
                        + Nuova Fermata
                    </button>
                </div>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">{message}</div>
                )}

                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    {schedules.map((s) => (
                        <div
                            key={s.id}
                            className="p-4 bg-white shadow rounded-lg border border-gray-200"
                        >
                            <h2 className="text-xl font-semibold">
                                Treno: {s.train_id}
                            </h2>

                            <p className="text-gray-600 mt-1">
                                Stazione ID: {s.station_id}
                            </p>

                            <p className="text-gray-600">
                                Arrivo: {s.arrival} — Partenza: {s.departure}
                            </p>

                            <p className="text-gray-600">
                                Prezzo: €{s.price.toFixed(2)}
                            </p>

                            <div className="flex gap-3 mt-4">
                                <button
                                    onClick={() => navigate(`/schedules/${s.id}`)}
                                    className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg"
                                >
                                    Modifica
                                </button>

                                <button
                                    onClick={() => navigate(`/schedules/${s.id}/stops`)}
                                    className="flex-1 bg-green-600 hover:bg-green-700 text-white py-2 rounded-lg"
                                >
                                    Fermate
                                </button>

                                <button
                                    onClick={() => handleDelete(s.id)}
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
