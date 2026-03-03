import {useEffect, useState} from "react";
import {useNavigate, useParams} from "react-router-dom";
import MainLayout from "../../../core/layout/MainLayout";
import {getStationById, updateStation} from "../api/stationsApi";

export default function EditStationPage() {
    const { id } = useParams();
    const navigate = useNavigate();

    const [loading, setLoading] = useState(true);
    const [saving, setSaving] = useState(false);
    const [message, setMessage] = useState("");

    const [name, setName] = useState("");
    const [city, setCity] = useState("");
    const [region, setRegion] = useState("");
    const [status, setStatus] = useState<"active" | "inactive">("active");

    useEffect(() => {
        if (!id) return;

        getStationById(Number(id))
            .then((res) => {
                const s = res.data;
                setName(s.name);
                setCity(s.city);
                setRegion(s.region);
                setStatus(s.status);
            })
            .catch(() => setMessage("Errore nel caricamento della stazione."))
            .finally(() => setLoading(false));
    }, [id]);

    const handleSave = async () => {
        if (!name.trim() || !city.trim() || !region.trim()) {
            setMessage("Tutti i campi sono obbligatori.");
            return;
        }

        setSaving(true);
        setMessage("");

        try {
            await updateStation(Number(id), {
                name,
                city,
                region,
                status,
            });

            navigate("/stations");
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
                <h1 className="text-3xl font-bold">Modifica Stazione</h1>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">
                        {message}
                    </div>
                )}

                <div className="space-y-4 bg-white p-6 rounded-xl shadow border">
                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Nome
                        </label>
                        <input
                            className="w-full border p-2 rounded mt-1"
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Città
                        </label>
                        <input
                            className="w-full border p-2 rounded mt-1"
                            value={city}
                            onChange={(e) => setCity(e.target.value)}
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Regione
                        </label>
                        <input
                            className="w-full border p-2 rounded mt-1"
                            value={region}
                            onChange={(e) => setRegion(e.target.value)}
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Stato
                        </label>
                        <select
                            className="w-full border p-2 rounded mt-1"
                            value={status}
                            onChange={(e) =>
                                setStatus(e.target.value as "active" | "inactive")
                            }
                        >
                            <option value="active">Attiva</option>
                            <option value="inactive">Non attiva</option>
                        </select>
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
                        onClick={() => navigate("/stations")}
                        className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg transition"
                    >
                        Annulla
                    </button>
                </div>
            </div>
        </MainLayout>
    );
}
