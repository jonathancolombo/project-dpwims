import {useState} from "react";
import MainLayout from "../../../core/layout/MainLayout.tsx";
import {createUser} from "../api/usersApi.ts";
import {useNavigate} from "react-router-dom";

export default function CreateUserPage() {
    const navigate = useNavigate();

    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [telephone, setTelephone] = useState("");
    const [fiscalCode, setFiscalCode] = useState("");
    const [role, setRole] = useState<"admin" | "customer">("customer");
    const [password, setPassword] = useState("");

    const [saving, setSaving] = useState(false);
    const [message, setMessage] = useState("");

    const handleSave = async () => {
        if (!username.trim() || !email.trim() || !password.trim()) {
            setMessage("Username, email e password sono obbligatori.");
            return;
        }

        setSaving(true);
        setMessage("");

        try {
            await createUser({
                username,
                email,
                telephone,
                fiscal_code: fiscalCode,
                role,
                password,
            });


            navigate("/users");
        } catch (err) {
            console.error(err);
            setMessage("Errore durante la creazione dell'utente.");
        } finally {
            setSaving(false);
        }
    };

    return (
        <MainLayout>
            <div className="p-6 max-w-2xl mx-auto space-y-6">
                <h1 className="text-3xl font-bold text-gray-900">Crea Nuovo Utente</h1>
                <p className="text-gray-600">
                    Inserisci i dati per aggiungere un nuovo account al sistema.
                </p>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">
                        {message}
                    </div>
                )}

                <div className="space-y-4 bg-white p-6 rounded-xl shadow border border-gray-200">
                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Username
                        </label>
                        <input
                            type="text"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            className="mt-1 w-full border rounded-lg p-2"
                            placeholder="Es. j.rossi"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Email
                        </label>
                        <input
                            type="email"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            className="mt-1 w-full border rounded-lg p-2"
                            placeholder="esempio@mail.com"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Telefono
                        </label>
                        <input
                            type="text"
                            value={telephone}
                            onChange={(e) => setTelephone(e.target.value)}
                            className="mt-1 w-full border rounded-lg p-2"
                            placeholder="Es. +39 333 1234567"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Codice Fiscale
                        </label>
                        <input
                            type="text"
                            value={fiscalCode}
                            onChange={(e) => setFiscalCode(e.target.value.toUpperCase())}
                            className="mt-1 w-full border rounded-lg p-2"
                            placeholder="Es. RSSMRA80A01H501U"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Ruolo
                        </label>
                        <select
                            value={role}
                            onChange={(e) => setRole(e.target.value as "admin" | "customer")}
                            className="mt-1 w-full border rounded-lg p-2"
                        >
                            <option value="admin">Admin</option>
                            <option value="customer">Cliente</option>
                        </select>


                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Password
                        </label>
                        <input
                            type="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            className="mt-1 w-full border rounded-lg p-2"
                        />
                    </div>
                </div>

                <div className="flex gap-3">
                    <button
                        onClick={handleSave}
                        disabled={saving}
                        className="flex-1 bg-blue-600 hover:bg-blue-700 text-white py-2 rounded-lg font-medium transition disabled:opacity-50"
                    >
                        {saving ? "Salvataggio..." : "Crea Utente"}
                    </button>

                    <button
                        onClick={() => navigate("/users")}
                        className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg font-medium transition"
                    >
                        Annulla
                    </button>
                </div>
            </div>
        </MainLayout>
    );
}
