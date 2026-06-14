import {useEffect, useState} from "react";
import MainLayout from "../../../core/layout/MainLayout";
import {deleteSchedule, getSchedules} from "../api/schedules_api.ts";
import {useNavigate} from "react-router-dom";
import type {Schedule} from "../types/schedule.ts";

export default function SchedulesPage() {
    const [schedules, setSchedules] = useState<Schedule[]>([]);
    const [loading, setLoading] = useState(true);
    const [message, setMessage] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        getSchedules()
            .then((response) => setSchedules(response.data))
            .catch(() => setMessage("Errore nel caricamento degli itinerari."))
            .finally(() => setLoading(false));
    }, []);

    const handleDelete = async (id: number) => {
        if (!confirm("Sei sicuro di voler eliminare questo itinerario?")) return;

        try {
            await deleteSchedule(id);
            setSchedules((schedules) => schedules.filter((schedule) => schedule.id !== id));
        } catch {
            setMessage("Errore durante l'eliminazione dell'itinerario.");
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
                    <h1 className="text-3xl font-bold">Itinerari</h1>

                    <div className="flex gap-3">
                        <button
                            onClick={() => navigate(-1)}
                            className="bg-gray-200 hover:bg-gray-300 text-gray-700 px-4 py-2 rounded-lg font-medium transition"
                        >
                            ← Torna indietro
                        </button>
                        <button
                            onClick={() => navigate("/schedules/create")}
                            className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700"
                        >
                            + Nuovo itinerario
                        </button>
                    </div>
                </div>
                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">{message}</div>
                )}

                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    {schedules.map((schedule) => (
                        <div
                            key={schedule.id}
                            className="p-4 bg-white shadow rounded-lg border border-gray-200"
                        >
                            <h2 className="text-xl font-semibold">
                                UUID Treno: {schedule.train_id}
                            </h2>

                            <p className="text-gray-600 mt-1">
                                ID Stazione di partenza: {schedule.station_id}
                            </p>

                            <p className="text-gray-600">
                                Stazione di partenza: {schedule.departure} -  Stazione di arrivo: {schedule.arrival}
                            </p>

                            <p className="text-gray-600">
                                Prezzo: {schedule.price.toFixed(2)} €
                            </p>

                            <div className="flex gap-3 mt-4">
                                <button
                                    onClick={() => navigate(`/schedules/${schedule.id}`)}
                                    className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg"
                                >
                                    Modifica
                                </button>

                                <button
                                    onClick={() => navigate(`/schedules/${schedule.id}/stops`)}
                                    className="flex-1 bg-green-600 hover:bg-green-700 text-white py-2 rounded-lg"
                                >
                                    Fermate
                                </button>

                                <button
                                    onClick={() => handleDelete(schedule.id)}
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
