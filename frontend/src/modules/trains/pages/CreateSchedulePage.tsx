import {useState} from "react";
import MainLayout from "../../../core/layout/MainLayout";
import {createSchedule} from "../api/schedulesApi";
import {useNavigate} from "react-router-dom";
import StationSelect from "./StationSelect.tsx";
import TrainSelect from "./TrainSelect.tsx";

export default function CreateSchedulePage() {
    const navigate = useNavigate();

    const [trainId, setTrainId] = useState("");
    const [stationId, setStationId] = useState(0);
    const [arrival, setArrival] = useState("");
    const [departure, setDeparture] = useState("");
    const [status, setStatus] = useState<"active" | "inactive">("active");
    const [price, setPrice] = useState(0);
    const [message, setMessage] = useState("");

    const handleSave = async () => {
        if (!trainId || !stationId || !arrival || !departure) {
            setMessage("Tutti i campi sono obbligatori.");
            return;
        }

        try {
            await createSchedule({
                train_id: trainId,
                station_id: stationId,
                arrival,
                departure,
                status,
                price,
            });

            navigate("/schedules");
        } catch {
            setMessage("Errore durante la creazione.");
        }
    };

    return (
        <MainLayout>
            <div className="p-6 max-w-xl mx-auto space-y-6">
                <h1 className="text-3xl font-bold">Nuova fermata</h1>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">{message}</div>
                )}

                <div className="space-y-4 bg-white p-6 rounded-xl shadow border">

                    <div>
                        <label className="block text-sm font-medium text-gray-700">Treno</label>
                        <TrainSelect value={trainId} onChange={setTrainId} />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">Stazione</label>
                        <StationSelect value={stationId} onChange={setStationId} />
                    </div>

                    <input
                        className="w-full border p-2 rounded"
                        placeholder="Arrivo (es. 08:30)"
                        value={arrival}
                        onChange={(e) => setArrival(e.target.value)}
                    />

                    <input
                        className="w-full border p-2 rounded"
                        placeholder="Partenza (es. 08:45)"
                        value={departure}
                        onChange={(e) => setDeparture(e.target.value)}
                    />

                    <select
                        className="w-full border p-2 rounded"
                        value={status}
                        onChange={(e) => setStatus(e.target.value as "active" | "inactive")}
                    >
                        <option value="active">Attivo</option>
                        <option value="inactive">Non attivo</option>
                    </select>

                    <input
                        type="number"
                        className="w-full border p-2 rounded"
                        placeholder="Prezzo"
                        value={price}
                        onChange={(e) => setPrice(Number(e.target.value))}
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
