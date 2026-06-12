import { useEffect, useState } from "react";
import MainLayout from "../../../core/layout/MainLayout";
import {
    getSubscriptions,
    getSubscriptionsByTrain,
    deleteSubscription
} from "../api/subscriptions_api";
import {planLabels, type Subscription} from "../types/subscription";
import {useNavigate} from "react-router-dom";

export default function SubscriptionsPage() {
    const [subscriptions, setSubscriptions] = useState<Subscription[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState("");
    const [trainFilter, setTrainFilter] = useState("");
    const navigate = useNavigate();

    async function loadAll() {
        try {
            const response = await getSubscriptions();
            setSubscriptions(response.data);
        } catch (err) {
            console.error(err);
            setError("Impossibile caricare le sottoscrizioni.");
        } finally {
            setLoading(false);
        }
    }

    useEffect(() => {
        loadAll();
    }, []);

    async function handleFilter() {
        if (!trainFilter.trim()) {
            await loadAll();
            return;
        }

        try {
            const response = await getSubscriptionsByTrain(trainFilter);
            setSubscriptions(response.data);
            setError("");
        } catch (err) {
            console.error(err);
            setError("Impossibile filtrare le sottoscrizioni.");
        }
    }

    async function handleDelete(id: number) {
        try {
            await deleteSubscription(id);
            setSubscriptions(subscriptions => subscriptions.filter(subscription => subscription.id !== id));
        } catch (err) {
            console.error(err);
            setError("Impossibile eliminare la sottoscrizione.");
        }
    }

    if (loading) return <MainLayout>Caricamento sottoscrizioni...</MainLayout>;

    return (
        <MainLayout>
            <div className="p-6 space-y-6">
                <h1 className="text-3xl font-bold">Sottoscrizioni</h1>
                {error && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">
                        {error}
                    </div>
                )}
                <button
                    onClick={() => navigate("/subscriptions/create")}
                    className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg"
                >
                    + Nuova Sottoscrizione
                </button>

                <div className="flex gap-3">
                    <input
                        type="text"
                        placeholder="UUID treno"
                        value={trainFilter}
                        onChange={(element) => setTrainFilter(element.target.value)}
                        className="border rounded-lg p-2 flex-1"
                    />
                    <button
                        onClick={handleFilter}
                        className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg"
                    >
                        Filtra
                    </button>
                </div>

                {/* LISTA */}
                <div className="space-y-4">
                    {subscriptions.map(subscription => (
                        <div
                            key={subscription.id}
                            className="p-4 bg-white shadow rounded-lg border"
                        >
                            <div className="flex justify-between items-center">
                                <div>
                                    <p><strong>ID:</strong> {subscription.id}</p>
                                    <p><strong>User:</strong> {subscription.user_id}</p>
                                    <p><strong>Train:</strong> {subscription.train_uuid}</p>
                                    <p><strong>Plan:</strong> {planLabels[subscription.plan]}</p>
                                </div>

                                <button
                                    onClick={() => handleDelete(subscription.id)}
                                    className="bg-red-600 hover:bg-red-700 text-white px-3 py-1 rounded"
                                >
                                    Elimina
                                </button>
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </MainLayout>
    );
}
