import {useState} from "react";
import MainLayout from "../../../core/layout/MainLayout";
import {createRoute} from "../api/routes_api.ts";
import {useNavigate} from "react-router-dom";

export default function CreateRoutePage() {
    const initialStateString = "";
    const initialStateNumber = 0;
    const navigate = useNavigate();
    const [trainId, setTrainId] = useState(initialStateString);
    const [departure, setDeparture] = useState(initialStateString);
    const [arrival, setArrival] = useState(initialStateString);
    const [distance, setDistance] = useState(initialStateNumber);
    const [message, setMessage] = useState(initialStateString);

    const handleSave = async () => {
        if (!trainId || !departure || !arrival || distance <= initialStateNumber) {
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
                        onChange={(element) => setTrainId(element.target.value)}
                    />

                    <input
                        className="w-full border p-2 rounded"
                        placeholder="Stazione di partenza"
                        value={departure}
                        onChange={(element) => setDeparture(element.target.value)}
                    />

                    <input
                        className="w-full border p-2 rounded"
                        placeholder="Stazione di arrivo"
                        value={arrival}
                        onChange={(element) => setArrival(element.target.value)}
                    />

                    <input
                        type="number"
                        className="w-full border p-2 rounded"
                        placeholder="Distanza (km)"
                        value={distance}
                        onChange={(element) => setDistance(Number(element.target.value))}
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
