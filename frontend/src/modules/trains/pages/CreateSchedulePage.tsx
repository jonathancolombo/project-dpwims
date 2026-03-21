import {useState} from "react";
import MainLayout from "../../../core/layout/MainLayout";
import {createSchedule} from "../api/schedules_api.ts";
import {useNavigate} from "react-router-dom";
import StationSelect from "./StationSelect.tsx";
import TrainSelect from "./TrainSelect.tsx";
import StationNameSelect from "./StationNameSelect.tsx";

export default function CreateSchedulePage() {
    const initialStateString = "";
    const initialStateNumber = 0;
    const navigate = useNavigate();
    const [trainId, setTrainId] = useState(initialStateString);
    const [stationId, setStationId] = useState(initialStateNumber);
    const [arrival, setArrival] = useState(initialStateString);
    const [departure, setDeparture] = useState(initialStateString);
    const [status, setStatus] = useState<"active" | "inactive">("active");
    const [price, setPrice] = useState(initialStateNumber);
    const [message, setMessage] = useState(initialStateString);

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
            setMessage("Errore durante la creazione della fermata.");
        }
    };

    return (
        <MainLayout>
            <div className="p-6 max-w-xl mx-auto space-y-6">
                <h1 className="text-3xl font-bold">Nuovo itinerario</h1>

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

                    <div>
                        <label className="block text-sm font-medium text-gray-700">Nome stazione di partenza</label>
                        <StationNameSelect value={departure} onChange={setDeparture} />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">Nome stazione di arrivo</label>
                        <StationNameSelect value={arrival} onChange={setArrival} />
                    </div>


                <div>
                <label className="block text-sm font-medium text-gray-700">Stato</label>
                    <select
                        className="w-full border p-2 rounded"
                        value={status}
                        onChange={(element) => setStatus(element.target.value as "active" | "inactive")}
                    >
                        <option value="active">Attivo</option>
                        <option value="inactive">Non attivo</option>
                    </select>
                </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">Prezzo</label>

                        <input
                        type="text"
                        inputMode="decimal"
                        className="w-full border p-2 rounded"
                        placeholder="Inserisci il prezzo in €"
                        value={price === 0 ? "" : price}
                        onChange={(element) => {
                            const value = element.target.value;
                            if (/^[0-9]*[.,]?[0-9]*$/.test(value)) {
                                setPrice(Number(value.replace(",", ".")));
                            }
                        }}
                    />
                    </div>
                </div>

                <div className="flex gap-3">
                <button
                    onClick={handleSave}
                    className="w-full bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700"
                >
                    Salva
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
