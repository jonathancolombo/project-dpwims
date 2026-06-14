import {useEffect, useState} from "react";
import {useNavigate, useParams} from "react-router-dom";
import MainLayout from "../../../core/layout/MainLayout";
import {getTrains, patchTrain} from "../api/trains_api.ts";
import type {Train} from "../types/train.ts";

export default function TrainDetailPage() {
    const { uuid } = useParams();
    const navigate = useNavigate();

    const [train, setTrain] = useState<Train | null>(null);

    const [trainNumber, setTrainNumber] = useState("");
    const [type, setType] = useState("");
    const [capacity, setCapacity] = useState(0);
    const [status, setStatus] = useState("");

    const [saving, setSaving] = useState(false);
    const [message, setMessage] = useState("");

    useEffect(() => {
        getTrains().then((response) => {
            const trainFound = response.data.find((train) => train.uuid === uuid);
            if (trainFound) {
                setTrain(trainFound);
                setTrainNumber(trainFound.train_number);
                setType(trainFound.type);
                setCapacity(trainFound.capacity);
                setStatus(trainFound.status);
            }
        });
    }, [uuid]);

    const handleSave = async () => {
        if (!train) return;

        setSaving(true);
        setMessage("");

        const timeout = 1000;
        try {
            await patchTrain(train.uuid, {
                train_number: trainNumber,
                type,
                capacity,
                status,
            });

            setMessage("Modifiche salvate con successo!");
            setSaving(false);

            setTimeout(() => navigate("/trains"), timeout);
        } catch (err) {
            setMessage("Errore durante il salvataggio.");
            setSaving(false);
        }
    };

    if (!train) {
        return (
            <MainLayout>
                <div className="p-6 text-gray-700">Caricamento treno...</div>
            </MainLayout>
        );
    }

    return (
        <MainLayout>
            <div className="p-6 space-y-6">
                <h1 className="text-3xl font-bold text-gray-900">
                    Modifica Treno {train.train_number}
                </h1>

                {message && (
                    <div className="p-3 bg-green-100 text-green-700 rounded-lg border border-green-300">
                        {message}
                    </div>
                )}

                <div className="bg-white shadow rounded-xl p-6 border border-gray-200 space-y-4">
                    <label className="block">
                        <span className="text-gray-700 font-medium">Numero treno</span>
                        <input
                            type="text"
                            value={trainNumber}
                            onChange={(element) => setTrainNumber(element.target.value)}
                            className="mt-1 w-full border rounded-lg px-3 py-2"
                        />
                    </label>

                    <label className="block">
                        <span className="text-gray-700 font-medium">Tipo</span>
                        <select
                            value={type}
                            onChange={(element) => setType(element.target.value)}
                            className="mt-1 w-full border rounded-lg px-3 py-2"
                        >
                            <option value="regional">Regionale</option>
                            <option value="intercity">Intercity</option>
                            <option value="highspeed">Alta Velocità</option>
                        </select>
                    </label>

                    <label className="block">
                        <span className="text-gray-700 font-medium">Capacità</span>
                        <input
                            type="number"
                            value={capacity}
                            onChange={(element) => setCapacity(Number(element.target.value))}
                            className="mt-1 w-full border rounded-lg px-3 py-2"
                        />
                    </label>

                    <label className="block">
                        <span className="text-gray-700 font-medium">Stato</span>
                        <select
                            value={status}
                            onChange={(element) => setStatus(element.target.value)}
                            className="mt-1 w-full border rounded-lg px-3 py-2"
                        >
                            <option value="active">Attivo</option>
                            <option value="inactive">Non Attivo</option>
                        </select>
                    </label>


                    <div className="flex gap-3 pt-2">
                        <button
                            onClick={() => navigate(-1)}
                            className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg font-medium transition"
                        >
                            ← Torna indietro
                        </button>
                        <button
                            onClick={handleSave}
                            disabled={saving}
                            className="flex-1 bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg disabled:opacity-50"
                        >
                            {saving ? "Salvataggio..." : "Salva modifiche"}
                        </button>
                    </div>


                </div>
            </div>
        </MainLayout>
    );
}
