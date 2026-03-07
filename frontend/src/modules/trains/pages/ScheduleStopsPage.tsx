import {useEffect, useState} from "react";
import {useNavigate, useParams} from "react-router-dom";
import MainLayout from "../../../core/layout/MainLayout";
import {createStop, deleteStop, getStopsBySchedule,} from "../api/schedule_stops_api.ts";
import StationSelect from "./StationSelect.tsx";
import type {ScheduleStop} from "../types/schedule_stop.ts";

export default function ScheduleStopsPage() {
    const { id } = useParams();
    const navigate = useNavigate();

    const [stops, setStops] = useState<ScheduleStop[]>([]);
    const [loading, setLoading] = useState(true);
    const [message, setMessage] = useState("");

    const [newStationId, setNewStationId] = useState(0);
    const [newArrival, setNewArrival] = useState("");
    const [newDeparture, setNewDeparture] = useState("");

    useEffect(() => {
        if (!id) return;

        getStopsBySchedule(Number(id))
            .then((response) => {
                const sorted = response.data.sort((firstScheduleStop, secondScheduleStop) => firstScheduleStop.stop_order - secondScheduleStop.stop_order);
                setStops(sorted);
            })
            .catch(() => setMessage("Errore nel caricamento delle fermate."))
            .finally(() => setLoading(false));
    }, [id]);

    const handleAddStop = async () => {
        if (!newStationId || !newArrival || !newDeparture) {
            setMessage("Tutti i campi sono obbligatori.");
            return;
        }

        try {
            await createStop(Number(id), {
                station_id: newStationId,
                arrival_time: newArrival,
                departure_time: newDeparture,
            });

            const response = await getStopsBySchedule(Number(id));
            setStops(response.data.sort((firstScheduleStop, secondScheduleStop) => firstScheduleStop.stop_order - secondScheduleStop.stop_order));

            // reset
            setNewStationId(0);
            setNewArrival("");
            setNewDeparture("");
        } catch {
            setMessage("Errore durante l'aggiunta della fermata.");
        }
    };

    const handleDelete = async () => {
        if (!confirm("Eliminare questa fermata?")) return;

        try {
            await deleteStop(Number(id));

            const response = await getStopsBySchedule(Number(id));
            setStops(response.data.sort((firstScheduleStop, secondScheduleStop) => firstScheduleStop.stop_order - secondScheduleStop.stop_order));
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
            <div className="p-6 max-w-3xl mx-auto space-y-6">
                <h1 className="text-3xl font-bold">Fermate dello Schedule</h1>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">{message}</div>
                )}

                {/* Lista fermate */}
                <div className="space-y-4">
                    {stops.map((stop) => (
                        <div
                            key={stop.id}
                            className="p-4 bg-white shadow rounded-lg border border-gray-200"
                        >
                            <div className="flex justify-between items-center">
                                <h2 className="text-lg font-semibold">
                                    {stop.stop_order}. {stop.station_name}
                                </h2>

                                <div className="flex gap-2">
                                    <button
                                        onClick={() => handleDelete()}
                                        className="px-2 py-1 bg-red-600 hover:bg-red-700 text-white rounded"
                                    >
                                        Elimina
                                    </button>
                                </div>
                            </div>

                            <div className="mt-2 text-gray-700">
                                Arrivo: {stop.arrival_time} — Partenza: {stop.departure_time}
                            </div>
                        </div>
                    ))}
                </div>

                {/* Aggiungi fermata */}
                <div className="bg-white p-6 rounded-xl shadow border space-y-4">
                    <h2 className="text-xl font-semibold">Aggiungi Fermata</h2>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">Stazione</label>
                        <StationSelect value={newStationId} onChange={setNewStationId} />
                    </div>


                    <input
                        className="w-full border p-2 rounded"
                        placeholder="Arrivo (es. 08:30)"
                        value={newArrival}
                        onChange={(element) => setNewArrival(element.target.value)}
                    />

                    <input
                        className="w-full border p-2 rounded"
                        placeholder="Partenza (es. 08:45)"
                        value={newDeparture}
                        onChange={(element) => setNewDeparture(element.target.value)}
                    />

                    <button
                        onClick={handleAddStop}
                        className="w-full bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700"
                    >
                        Aggiungi Fermata
                    </button>
                </div>

                <button
                    onClick={() => navigate("/schedules")}
                    className="w-full bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg"
                >
                    Torna alle Fermate
                </button>
            </div>
        </MainLayout>
    );
}
