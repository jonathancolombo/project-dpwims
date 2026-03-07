import {useEffect, useState} from "react";
import {deleteTrain, getTrains} from "../api/trains_api.ts";
import MainLayout from "../../../core/layout/MainLayout";
import {useNavigate} from "react-router-dom";
import type {Train} from "../types/train.ts";
import {getTrainIcon} from "../../../util/train_icons.tsx";

export default function TrainsPage() {
    const [trains, setTrains] = useState<Train[]>([]);

    useEffect(() => {
        getTrains().then((response) => setTrains(response.data));
    }, []);

    const activeCount = trains.filter(train => train.status === "active").length;
    const inactiveCount = trains.filter(train => train.status !== "active").length;
    const navigate = useNavigate();
    const handleDelete = async (uuid: string) => {
        const confirmDelete = globalThis.confirm("Sei sicuro di voler cancellare questo treno?");
        if (!confirmDelete) return;

        try {
            await deleteTrain(uuid);
            setTrains((previousTrains) => previousTrains.filter((train) => train.uuid !== uuid));
        } catch (error) {
            console.error("Errore durante la cancellazione:", error);
        }
    };

    return (
        <MainLayout>
            <div className="p-6 space-y-8">

                {/* HEADER ADMIN */}
                <div className="flex justify-between items-center">
                    <div>
                        <h1 className="text-3xl font-bold text-gray-900">Pannello Treni (Admin)</h1>
                        <p className="text-gray-600 mt-1">Gestione e monitoraggio della flotta ferroviaria</p>
                    </div>

                    <button
                        onClick={() => navigate("/trains/create")}
                        className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg font-medium shadow transition"
                    >
                        + Crea Treno
                    </button>
                </div>


                <div className="grid grid-cols-1 sm:grid-cols-3 gap-6">
                    <div className="bg-white shadow rounded-xl p-5 border border-gray-200">
                        <p className="text-sm text-gray-500">Totale Treni</p>
                        <p className="text-3xl font-bold text-gray-900">{trains.length}</p>
                    </div>

                    <div className="bg-white shadow rounded-xl p-5 border border-gray-200">
                        <p className="text-sm text-gray-500">Attivi</p>
                        <p className="text-3xl font-bold text-green-600">{activeCount}</p>
                    </div>

                    <div className="bg-white shadow rounded-xl p-5 border border-gray-200">
                        <p className="text-sm text-gray-500">Non Attivi</p>
                        <p className="text-3xl font-bold text-red-600">{inactiveCount}</p>
                    </div>
                </div>

                <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                    {trains.map((train) => (
                        <div
                            key={train.uuid}
                            className="bg-white rounded-xl shadow-md border border-gray-200 hover:shadow-xl transition-all p-6"
                        >
                            {/* HEADER CARD */}
                            <div className="flex justify-between items-start mb-4">
                                <div>
                                    <div className="flex items-center gap-2">
                                        <span className="text-3xl">{getTrainIcon(train.type)}</span>
                                        <h2 className="text-2xl font-semibold text-blue-700">{train.train_number}</h2>
                                    </div>

                                    <p className="text-sm text-gray-500">UUID: {train.uuid.slice(0, 8)}…</p>
                                </div>

                                <span
                                    className={`px-3 py-1 text-xs font-semibold rounded-full ${
                                        train.status === "active"
                                            ? "bg-green-100 text-green-700"
                                            : "bg-red-100 text-red-700"
                                    }`}
                                >
                  {train.status === "active" ? "Attivo" : "Non attivo"}
                </span>
                            </div>

                            <div className="space-y-3 text-gray-700">
                                <p><span className="font-semibold">Tipo:</span> {train.type}</p>
                                <p><span className="font-semibold">Capacità:</span> {train.capacity} posti</p>
                            </div>

                            <div className="mt-5 flex gap-3">
                                <button
                                    className="flex-1 bg-red-600 hover:bg-red-700 text-white py-2 rounded-lg font-medium transition"
                                    onClick={() => handleDelete(train.uuid)}
                                >
                                    Cancella
                                </button>

                                <button className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg font-medium transition" onClick={() => navigate(`/admin/trains/${train.uuid}`)} >
                                    Modifica </button>
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </MainLayout>
    );
}
