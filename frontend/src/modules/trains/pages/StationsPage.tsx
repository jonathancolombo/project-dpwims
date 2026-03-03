import {useEffect, useState} from "react";
import MainLayout from "../../../core/layout/MainLayout";
import {deleteStation, getStations, type Station} from "../api/stationsApi";
import {useNavigate} from "react-router-dom";
import {StatusBadge} from "../../../util/badge.tsx";


export default function StationsPage() {
    const [stations, setStations] = useState<Station[]>([]);
    const [loading, setLoading] = useState(true);
    const [message, setMessage] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        getStations()
            .then((res) => setStations(res.data))
            .catch(() => setMessage("Errore nel caricamento delle stazioni."))
            .finally(() => setLoading(false));
    }, []);

    const handleDelete = async (id: number) => {
        if (!confirm("Sei sicuro di voler eliminare questa stazione?")) return;

        try {
            await deleteStation(id);
            setStations((prev) => prev.filter((s) => s.id !== id));
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
            <div className="p-6 max-w-4xl mx-auto space-y-6">
                <div className="flex justify-between items-center">
                    <h1 className="text-3xl font-bold">Stazioni</h1>
                    <button
                        onClick={() => navigate("/stations/create")}
                        className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700"
                    >
                        + Nuova Stazione
                    </button>
                </div>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">{message}</div>
                )}

                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    {stations.map((s) => (
                        <div
                            key={s.id}
                            className="p-4 bg-white shadow rounded-lg border border-gray-200"
                        >
                            <div className="flex justify-between">
                                <h2 className="text-xl font-semibold">{s.name}</h2>
                                <StatusBadge status={s.status} />
                            </div>

                            <p className="text-gray-600">{s.city}, {s.region}</p>

                            <div className="flex gap-3 mt-4">
                                <button
                                    onClick={() => navigate(`/stations/${s.id}`)}
                                    className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg"
                                >
                                    Modifica
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
