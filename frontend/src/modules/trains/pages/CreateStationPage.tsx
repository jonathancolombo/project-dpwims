import {useState} from "react";
import MainLayout from "../../../core/layout/MainLayout";
import {createStation} from "../api/stationsApi";
import {useNavigate} from "react-router-dom";

export default function CreateStationPage() {
    const navigate = useNavigate();

    const [name, setName] = useState("");
    const [city, setCity] = useState("");
    const [region, setRegion] = useState("");
    const [status, setStatus] = useState<"active" | "inactive">("active");
    const [message, setMessage] = useState("");

    const handleSave = async () => {
        if (!name.trim() || !city.trim() || !region.trim()) {
            setMessage("Tutti i campi sono obbligatori.");
            return;
        }

        try {
            await createStation({ name, city, region, status });
            navigate("/stations");
        } catch {
            setMessage("Errore durante la creazione.");
        }
    };

    return (
        <MainLayout>
            <div className="p-6 max-w-xl mx-auto space-y-6">
                <h1 className="text-3xl font-bold">Nuova Stazione</h1>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">{message}</div>
                )}

                <div className="space-y-4 bg-white p-6 rounded-xl shadow border">
                    <input
                        className="w-full border p-2 rounded"
                        placeholder="Nome"
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                    />

                    <input
                        className="w-full border p-2 rounded"
                        placeholder="Città"
                        value={city}
                        onChange={(e) => setCity(e.target.value)}
                    />

                    <input
                        className="w-full border p-2 rounded"
                        placeholder="Regione"
                        value={region}
                        onChange={(e) => setRegion(e.target.value)}
                    />

                    <select
                        className="w-full border p-2 rounded"
                        value={status}
                        onChange={(e) => setStatus(e.target.value as "active" | "inactive")}
                    >
                        <option value="active">Attiva</option>
                        <option value="inactive">Non attiva</option>
                    </select>
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
