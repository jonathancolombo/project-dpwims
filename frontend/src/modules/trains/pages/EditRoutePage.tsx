import {useEffect, useState} from "react";
import {useNavigate, useParams} from "react-router-dom";
import MainLayout from "../../../core/layout/MainLayout";
import {getRouteById, updateRoute} from "../api/routes_api.ts";

export default function EditRoutePage() {
    const { id } = useParams();
    const navigate = useNavigate();

    const [loading, setLoading] = useState(true);
    const [saving, setSaving] = useState(false);
    const [message, setMessage] = useState("");

    const [trainId, setTrainId] = useState("");
    const [departure, setDeparture] = useState("");
    const [arrival, setArrival] = useState("");
    const [distance, setDistance] = useState(0);

    useEffect(() => {
        if (!id) return;

        getRouteById(Number(id))
            .then((res) => {
                const r = res.data;
                setTrainId(r.train_id);
                setDeparture(r.departure_station);
                setArrival(r.arrival_station);
                setDistance(r.distance);
            })
            .catch(() => setMessage("Errore nel caricamento della rotta."))
            .finally(() => setLoading(false));
    }, [id]);

    const handleSave = async () => {
        if (!trainId || !departure || !arrival || distance <= 0) {
            setMessage("Tutti i campi sono obbligatori.");
            return;
        }

        setSaving(true);
        setMessage("");

        try {
            await updateRoute(Number(id), {
                train_id: trainId,
                departure_station: departure,
                arrival_station: arrival,
                distance,
            });

            navigate("/routes");
        } catch (err) {
            console.error(err);
            setMessage("Errore durante il salvataggio.");
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
                <h1 className="text-3xl font-bold">Modifica Rotta</h1>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">
                        {message}
                    </div>
                )}

                <div className="space-y-4 bg-white p-6 rounded-xl shadow border">
                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Train ID
                        </label>
                        <input
                            className="w-full border p-2 rounded mt-1"
                            value={trainId}
                            onChange={(e) => setTrainId(e.target.value)}
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Stazione di partenza
                        </label>
                        <input
                            className="w-full border p-2 rounded mt-1"
                            value={departure}
                            onChange={(e) => setDeparture(e.target.value)}
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Stazione di arrivo
                        </label>
                        <input
                            className="w-full border p-2 rounded mt-1"
                            value={arrival}
                            onChange={(e) => setArrival(e.target.value)}
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Distanza (km)
                        </label>
                        <input
                            type="number"
                            className="w-full border p-2 rounded mt-1"
                            value={distance}
                            onChange={(e) => setDistance(Number(e.target.value))}
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
                        onClick={() => navigate("/routes")}
                        className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg transition"
                    >
                        Annulla
                    </button>
                </div>
            </div>
        </MainLayout>
    );
}
