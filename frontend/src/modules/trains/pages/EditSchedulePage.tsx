import {useEffect, useState} from "react";
import {useNavigate, useParams} from "react-router-dom";
import MainLayout from "../../../core/layout/MainLayout";
import {getScheduleById, updateSchedule} from "../api/schedules_api.ts";
import TrainSelect from "./TrainSelect.tsx";
import StationNameSelect from "./StationNameSelect.tsx";
import StationSelect from "./StationSelect.tsx";

export default function EditSchedulePage() {
    const { id } = useParams();
    const navigate = useNavigate();

    const [loading, setLoading] = useState(true);
    const [saving, setSaving] = useState(false);
    const [message, setMessage] = useState("");

    const [trainId, setTrainId] = useState("");
    const [stationId, setStationId] = useState(0);
    const [arrival, setArrival] = useState("");
    const [departure, setDeparture] = useState("");
    const [status, setStatus] = useState<"active" | "inactive">("active");
    const [price, setPrice] = useState(0);

    useEffect(() => {
        if (!id) return;

        getScheduleById(Number(id))
            .then((response) => {
                const schedule = response.data;
                setTrainId(schedule.train_id);
                setStationId(schedule.station_id);
                setArrival(schedule.arrival);
                setDeparture(schedule.departure);
                setStatus(schedule.status);
                setPrice(schedule.price);
            })
            .catch(() => setMessage("Errore nel caricamento dell'itinerario."))
            .finally(() => setLoading(false));
    }, [id]);

    const handleSave = async () => {
        if (!trainId || !stationId || !arrival || !departure) {
            setMessage("Tutti i campi sono obbligatori.");
            return;
        }

        setSaving(true);
        setMessage("");

        try {
            await updateSchedule(Number(id), {
                train_id: trainId,
                station_id: stationId,
                arrival,
                departure,
                status,
                price,
            });

            navigate("/schedules");
        } catch (err) {
            console.error(err);
            setMessage("Errore durante il salvataggio dell'itinerario.");
        } finally {
            setSaving(false);
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
            <div className="p-6 max-w-xl mx-auto space-y-6">
                <h1 className="text-3xl font-bold">Modifica Itinerario</h1>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">
                        {message}
                    </div>
                )}

                <div className="space-y-4 bg-white p-6 rounded-xl shadow border">

                    <div>
                        <label className="block text-sm font-medium text-gray-700">Treno</label>
                        <TrainSelect value={trainId} onChange={setTrainId} />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">Id stazione di riferimento</label>
                        <input
                            type="text"
                            value={stationId}
                            disabled
                            className="w-full mt-1 p-2 border rounded bg-gray-100"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">Stazione di partenza</label>
                        <StationSelect
                            value={stationId}
                            onChange={(id, name) => {
                                setStationId(id);
                                setDeparture(name);
                            }}
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">Nome stazione di arrivo</label>
                        <StationNameSelect value={arrival} onChange={setArrival} />
                    </div>


                    <div>
                        <label className="block text-sm font-medium text-gray-700">Stato</label>
                        <select
                            className="w-full border p-2 rounded mt-1"
                            value={status}
                            onChange={(element) => setStatus(element.target.value as "active" | "inactive")}
                        >
                            <option value="active">Attivo</option>
                            <option value="inactive">Non attivo</option>
                        </select>
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">Prezzo (€)</label>
                        <input
                            type="number"
                            className="w-full border p-2 rounded mt-1"
                            value={price}
                            onChange={(element) => setPrice(Number(element.target.value))}
                        />
                    </div>
                </div>


                <div className="flex gap-3">
                    <button
                        onClick={handleSave}
                        disabled={saving}
                        className="flex-1 bg-blue-600 hover:bg-blue-700 text-white py-2 rounded-lg transition disabled:opacity-50"
                    >
                        {saving ? "Salvataggio..." : "Salva Modifiche"}
                    </button>

                    <button
                        onClick={() => navigate("/schedules")}
                        className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg transition"
                    >
                        Annulla
                    </button>
                </div>
            </div>
        </MainLayout>
    );
}
