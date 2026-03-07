import {useEffect, useState} from "react";
import MainLayout from "../../../core/layout/MainLayout";
import {deleteRoute, getRoutes} from "../api/routes_api.ts";
import {useNavigate} from "react-router-dom";
import type {Route} from "../types/route.ts";

export default function RoutesPage() {
    const [routes, setRoutes] = useState<Route[]>([]);
    const [loading, setLoading] = useState(true);
    const [message, setMessage] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        getRoutes()
            .then((response) => setRoutes(response.data))
            .catch(() => setMessage("Errore nel caricamento delle rotte."))
            .finally(() => setLoading(false));
    }, []);

    const handleDelete = async (id: number) => {
        if (!confirm("Eliminare questa rotta?")) return;

        try {
            await deleteRoute(id);
            setRoutes((routes) => routes.filter((route) => route.id !== id));
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
                    <h1 className="text-3xl font-bold">Rotte</h1>
                    <button
                        onClick={() => navigate("/routes/create")}
                        className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700"
                    >
                        + Nuova Rotta
                    </button>
                </div>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">{message}</div>
                )}

                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    {routes.map((route) => (
                        <div
                            key={route.id}
                            className="p-4 bg-white shadow rounded-lg border border-gray-200"
                        >
                            <h2 className="text-xl font-semibold">
                                {route.departure_station} → {route.arrival_station}
                            </h2>

                            <p className="text-gray-600 mt-1">
                                Treno: {route.train_id}
                            </p>

                            <p className="text-gray-600">
                                Distanza: {route.distance} km
                            </p>

                            <div className="flex gap-3 mt-4">
                                <button
                                    onClick={() => navigate(`/routes/${route.id}`)}
                                    className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg"
                                >
                                    Modifica
                                </button>
                                <button
                                    onClick={() => handleDelete(route.id)}
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
