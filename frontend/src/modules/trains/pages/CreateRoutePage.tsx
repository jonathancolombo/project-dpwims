import {useState} from "react";
import MainLayout from "../../../core/layout/MainLayout";
import {createRoute} from "../api/routesApi";
import {useNavigate} from "react-router-dom";

export default function CreateRoutePage() {
    const navigate = useNavigate();

    const [trainId, setTrainId] = useState("");
    const [departure, setDeparture] = useState("");
    const [arrival, setArrival] = useState("");
    const [distance, setDistance] = useState(0);
    const [message, setMessage] = useState("");

    const handleSave = async () => {
        if (!trainId || !departure || !arrival || distance <= 0) {
            setMessage("Tutti i campi sono obbligatori.");
            return;
        }

        try {
            await createRoute({
                train_id: trainId,
                departure_station: departure,
                arrival_station: arrival,
                distance,
            });

            navigate("/routes");
        } catch {
            setMessage("Errore durante la creazione.");
        }
    };

    return (
        <MainLayout>
            <div className="p-6 max-w-xl mx-auto space-y-6">
                <h1 className="text-3xl font-bold">Nuova Rotta</h1>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">{message}</div>
                )}

                <div className="space-y-4 bg-white p-6 rounded-xl shadow border">
                    <input
                        className="w-full border p-2 rounded"
                        placeholder="Train ID"
                        value={trainId}
                        onChange={(e) => setTrainId(e.target.value)}
                    />

                    <input
                        className="w-full border p-2 rounded"
                        placeholder="Stazione di partenza"
                        value={departure}
                        onChange={(e) => setDeparture(e.target.value)}
                    />

                    <input
                        className="w-full border p-2 rounded"
                        placeholder="Stazione di arrivo"
                        value={arrival}
                        onChange={(e) => setArrival(e.target.value)}
                    />

                    <input
                        type="number"
                        className="w-full border p-2 rounded"
                        placeholder="Distanza (km)"
                        value={distance}
                        onChange={(e) => setDistance(Number(e.target.value))}
                    />
                </div>

                <button
                    onClick={handleSave}
                    className="w-full bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700"
                >
                    Salva
                </button>
            </div>
        </MainLayout>
    );
}
