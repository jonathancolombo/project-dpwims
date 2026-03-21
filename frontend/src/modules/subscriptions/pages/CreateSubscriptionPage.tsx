import { useState } from "react";
import MainLayout from "../../../core/layout/MainLayout";
import { createSubscription } from "../api/subscriptions_api";
import { useNavigate } from "react-router-dom";
import type { Plan } from "../types/subscription";

export default function CreateSubscriptionPage() {
    const navigate = useNavigate();

    const [userId, setUserId] = useState<number>(0);
    const [trainUUID, setTrainUUID] = useState("");
    const [plan, setPlan] = useState<Plan>("basic");

    const [saving, setSaving] = useState(false);
    const [message, setMessage] = useState("");

    async function handleSave() {
        if (!userId || !trainUUID.trim()) {
            setMessage("Compila tutti i campi.");
            return;
        }

        setSaving(true);

        try {
            await createSubscription({
                user_id: userId,
                train_uuid: trainUUID,
                plan
            });

            navigate("/subscriptions");
        } catch (err) {
            console.error(err);
            setMessage("Errore durante la creazione della sottoscrizione.");
        } finally {
            setSaving(false);
        }
    }

    return (
        <MainLayout>
            <div className="p-6 max-w-xl mx-auto space-y-6">

                <h1 className="text-3xl font-bold">Crea Sottoscrizione</h1>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">
                        {message}
                    </div>
                )}

                <div className="space-y-4 bg-white p-6 rounded-xl shadow border">

                    <div>
                        <label className="block text-sm font-medium">User ID</label>
                        <input
                            type="number"
                            value={userId}
                            onChange={(element) => setUserId(Number(element.target.value))}
                            className="mt-1 w-full border rounded-lg p-2"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Train UUID</label>
                        <input
                            type="text"
                            value={trainUUID}
                            onChange={(element) => setTrainUUID(element.target.value)}
                            className="mt-1 w-full border rounded-lg p-2"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Piano</label>
                        <select
                            value={plan}
                            onChange={(element) => setPlan(element.target.value as Plan)}
                            className="mt-1 w-full border rounded-lg p-2"
                        >
                            <option value="basic">Basic</option>
                            <option value="premium">Premium</option>
                            <option value="full">Full</option>
                        </select>
                    </div>
                </div>

                <div className="flex gap-3">
                    <button
                        onClick={handleSave}
                        disabled={saving}
                        className="flex-1 bg-blue-600 hover:bg-blue-700 text-white py-2 rounded-lg disabled:opacity-50"
                    >
                        {saving ? "Salvataggio..." : "Crea"}
                    </button>

                    <button
                        onClick={() => navigate("/subscriptions")}
                        className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg"
                    >
                        Annulla
                    </button>
                </div>
            </div>
        </MainLayout>
    );
}
