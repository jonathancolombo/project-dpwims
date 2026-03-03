import {useState} from "react";
import MainLayout from "../../../core/layout/MainLayout";
import {createTrain} from "../api/trainsApi";
import {useNavigate} from "react-router-dom";

export default function CreateTrainPage() {
    const navigate = useNavigate();

    const [trainNumber, setTrainNumber] = useState("");
    const [type, setType] = useState("regional");
    const [capacity, setCapacity] = useState(100);
    const [status, setStatus] = useState("active");

    const [saving, setSaving] = useState(false);
    const [message, setMessage] = useState("");

    const handleSave = async () => {
        if (!trainNumber.trim()) {
            setMessage("Il numero del treno è obbligatorio.");
            return;
        }

        setSaving(true);

        try {
            await createTrain({
                train_number: trainNumber,
                type,
                capacity,
                status,
            });

            navigate("/trains");
        } catch (err) {
            console.error(err);
            setMessage("Errore durante la creazione del treno.");
        } finally {
            setSaving(false);
        }
    };

    return (
        <MainLayout>
            <div className="p-6 max-w-2xl mx-auto space-y-6">

                <h1 className="text-3xl font-bold text-gray-900">Crea Nuovo Treno</h1>
                <p className="text-gray-600">Inserisci i dati per aggiungere un nuovo treno alla flotta.</p>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">
                        {message}
                    </div>
                )}

                <div className="space-y-4 bg-white p-6 rounded-xl shadow border border-gray-200">

                    <div>
                        <label className="block text-sm font-medium text-gray-700">Numero Treno</label>
                        <input
                            type="text"
                            value={trainNumber}
                            onChange={(e) => setTrainNumber(e.target.value)}
                            className="mt-1 w-full border rounded-lg p-2"
                            placeholder="Es. TR-500"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">Tipo</label>
                        <select
                            value={type}
                            onChange={(e) => setType(e.target.value)}
                            className="mt-1 w-full border rounded-lg p-2"
                        >
                            <option value="regional">Regionale</option>
                            <option value="intercity">Intercity</option>
                            <option value="highspeed">Alta Velocità</option>
                        </select>
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">Capacità</label>
                        <input
                            type="number"
                            value={capacity}
                            onChange={(e) => setCapacity(Number(e.target.value))}
                            className="mt-1 w-full border rounded-lg p-2"
                            min={1}
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">Stato</label>
                        <select
                            value={status}
                            onChange={(e) => setStatus(e.target.value)}
                            className="mt-1 w-full border rounded-lg p-2"
                        >
                            <option value="active">Attivo</option>
                            <option value="inactive">Non Attivo</option>
                        </select>
                    </div>
                </div>

                <div className="flex gap-3">
                    <button
                        onClick={handleSave}
                        disabled={saving}
                        className="flex-1 bg-blue-600 hover:bg-blue-700 text-white py-2 rounded-lg font-medium transition disabled:opacity-50"
                    >
                        {saving ? "Salvataggio..." : "Crea Treno"}
                    </button>

                    <button
                        onClick={() => navigate("/trains")}
                        className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg font-medium transition"
                    >
                        Annulla
                    </button>
                </div>
            </div>
        </MainLayout>
    );
}
