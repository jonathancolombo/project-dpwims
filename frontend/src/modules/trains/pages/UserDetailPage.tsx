import {useEffect, useState} from "react";
import {useNavigate, useParams} from "react-router-dom";
import MainLayout from "../../../core/layout/MainLayout";
import {getUserById, patchUser} from "../api/usersApi";

export default function UserDetailPage() {
    const { id } = useParams();
    const navigate = useNavigate();

    const [loading, setLoading] = useState(true);
    const [saving, setSaving] = useState(false);
    const [message, setMessage] = useState("");

    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [telephone, setTelephone] = useState("");
    const [fiscalCode, setFiscalCode] = useState("");
    const [role, setRole] = useState<"admin" | "customer">("customer");
    const [password, setPassword] = useState("");

    useEffect(() => {
        if (!id) return;

        getUserById(Number(id))
            .then((res) => {
                const u = res.data;
                setUsername(u.username);
                setEmail(u.email);
                setTelephone(u.telephone);
                setFiscalCode(u.fiscal_code);
                setRole(u.role as "admin" | "customer");
            })
            .catch(() => setMessage("Errore nel caricamento dell'utente."))
            .finally(() => setLoading(false));
    }, [id]);

    const handleSave = async () => {
        setSaving(true);
        setMessage("");

        try {
            await patchUser(Number(id), {
                username,
                email,
                telephone,
                fiscal_code: fiscalCode,
                role,
                password: password || undefined,
            });

            console.log("PATCH DATA:", {
                username,
                email,
                telephone,
                fiscal_code: fiscalCode,
                role,
                password: password || undefined,
            });


            navigate("/users");
        } catch (err) {
            console.error(err);
            setMessage("Errore durante il salvataggio.");
        } finally {
            setSaving(false);
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
            <div className="p-6 max-w-2xl mx-auto space-y-6">
                <h1 className="text-3xl font-bold text-gray-900">Modifica Utente</h1>
                <p className="text-gray-600">Aggiorna i dati dell'utente selezionato.</p>

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
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Ruolo
                        </label>
                        <select
                            value={role}
                            onChange={(e) => setRole(e.target.value.toLowerCase() as "admin" | "customer")}
                            className="mt-1 w-full border rounded-lg p-2"
                        >
                            <option value="admin">Admin</option>
                            <option value="customer">Cliente</option>
                        </select>
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Nuova Password (opzionale)
                        </label>
                        <input
                            type="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            className="mt-1 w-full border rounded-lg p-2"
                            placeholder="Lascia vuoto per non cambiarla"
                        />
                    </div>
                </div>

                <div className="flex gap-3">
                    <button
                        onClick={handleSave}
                        disabled={saving}
                        className="flex-1 bg-blue-600 hover:bg-blue-700 text-white py-2 rounded-lg font-medium transition disabled:opacity-50"
                    >
                        {saving ? "Salvataggio..." : "Salva Modifiche"}
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
