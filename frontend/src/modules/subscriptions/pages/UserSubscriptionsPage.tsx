import { useEffect, useState } from "react";
import MainLayout from "../../../core/layout/MainLayout";
import { useNavigate } from "react-router-dom";
import { createSubscription, deleteSubscription, getSubscriptions } from "../api/subscriptions_api";
import { planLabels, type Plan, type Subscription } from "../types/subscription";
import { user_authorization } from "../../../core/hooks/user_authorization";
import { getSchedules } from "../../trains/api/schedules_api";
import type { Schedule } from "../../trains/types/schedule";

export default function UserSubscriptionsPage() {
    const navigate = useNavigate();
    const { user } = user_authorization();
    const userId = user?.userID ?? 0;

    const [subscriptions, setSubscriptions] = useState<Subscription[]>([]);
    const [schedules, setSchedules] = useState<Schedule[]>([]);
    const [loading, setLoading] = useState(true);
    const [saving, setSaving] = useState(false);
    const [message, setMessage] = useState("");
    const [error, setError] = useState("");
    const [selectedScheduleId, setSelectedScheduleId] = useState<number>(0);
    const [plan, setPlan] = useState<Plan>("basic");

    async function loadSubscriptions() {
        try {
            const response = await getSubscriptions(userId);
            setSubscriptions(response.data);
        } catch {
            setError("Impossibile caricare le tue sottoscrizioni.");
        } finally {
            setLoading(false);
        }
    }

    async function loadSchedules() {
        try {
            const response = await getSchedules();
            setSchedules(response.data);
        } catch {
            setError("Impossibile caricare gli itinerari disponibili.");
        }
    }

    useEffect(() => {
        if (!user?.userID) {
            return;
        }

        loadSubscriptions();
        loadSchedules();
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [userId]);

    if (!userId) {
        return (
            <MainLayout>
                <div className="p-6 max-w-3xl mx-auto space-y-4 text-center">
                    <h1 className="text-3xl font-bold">Sottoscrizioni</h1>
                    <p className="text-gray-600">
                        Non riesco a leggere il tuo profilo utente in questo momento.
                    </p>
                    <button
                        onClick={() => navigate("/login?target=user")}
                        className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg"
                    >
                        Vai al login
                    </button>
                </div>
            </MainLayout>
        );
    }

    async function handleCreate() {
        const selectedSchedule = schedules.find((schedule) => schedule.id === selectedScheduleId);

        if (!selectedSchedule) {
            setError("Seleziona un itinerario dalla lista.");
            return;
        }

        setSaving(true);
        setError("");
        setMessage("");

        try {
            await createSubscription({
                user_id: userId,
                train_uuid: selectedSchedule.train_id,
                plan,
            });

            setSelectedScheduleId(0);
            setPlan("basic");
            setMessage("Sottoscrizione creata con successo.");
            await loadSubscriptions();
        } catch {
            setError("Errore durante la creazione della sottoscrizione.");
        } finally {
            setSaving(false);
        }
    }

    async function handleDelete(id: number) {
        try {
            await deleteSubscription(id);
            setSubscriptions((current) => current.filter((subscription) => subscription.id !== id));
            setMessage("Sottoscrizione eliminata.");
        } catch {
            setError("Errore durante l'eliminazione della sottoscrizione.");
        }
    }

    if (loading) {
        return <MainLayout>Caricamento sottoscrizioni...</MainLayout>;
    }

    return (
        <MainLayout>
            <div className="p-6 space-y-6 max-w-4xl mx-auto">
                <div className="flex items-center justify-between gap-4 flex-wrap">
                    <div>
                        <h1 className="text-3xl font-bold">Le mie sottoscrizioni</h1>
                        <p className="text-gray-600 mt-1">
                            Gestisci gli avvisi sui treni che ti interessano.
                        </p>
                    </div>

                    <button
                        onClick={() => navigate("/user/schedules")}
                        className="bg-gray-100 hover:bg-gray-200 text-gray-700 px-4 py-2 rounded-lg"
                    >
                        Vai agli itinerari
                    </button>
                </div>

                {error && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">
                        {error}
                    </div>
                )}

                {message && (
                    <div className="p-3 bg-green-100 text-green-700 rounded-lg">
                        {message}
                    </div>
                )}

                <div className="bg-white p-6 rounded-xl shadow border space-y-4">
                    <h2 className="text-xl font-semibold">Nuova sottoscrizione</h2>

                    <div>
                        <label htmlFor="scheduleSelect" className="block text-sm font-medium text-gray-700">
                            Itinerario
                        </label>
                        <select
                            id="scheduleSelect"
                            value={selectedScheduleId}
                            onChange={(element) => setSelectedScheduleId(Number(element.target.value))}
                            className="mt-1 w-full border rounded-lg p-2"
                        >
                            <option value={0}>Seleziona un itinerario</option>
                            {schedules.map((schedule) => (
                                <option key={schedule.id} value={schedule.id}>
                                    #{schedule.id} {schedule.departure} → {schedule.arrival} (Treno {schedule.train_id})
                                </option>
                            ))}
                        </select>
                        {selectedScheduleId > 0 && (
                            <p className="text-sm text-gray-500 mt-1">
                                L'ID del treno verrà ricavato automaticamente dall'itinerario selezionato.
                            </p>
                        )}
                    </div>

                    <div>
                        <label htmlFor="subscriptionPlan" className="block text-sm font-medium text-gray-700">
                            Piano
                        </label>
                        <select
                            id="subscriptionPlan"
                            value={plan}
                            onChange={(element) => setPlan(element.target.value as Plan)}
                            className="mt-1 w-full border rounded-lg p-2"
                        >
                            <option value="basic">{planLabels.basic}</option>
                            <option value="premium">{planLabels.premium}</option>
                            <option value="full">{planLabels.full}</option>
                        </select>
                    </div>

                    <button
                        onClick={handleCreate}
                        disabled={saving}
                        className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg disabled:opacity-50"
                    >
                        {saving ? "Salvataggio..." : "Crea sottoscrizione"}
                    </button>
                </div>

                <div className="space-y-4">
                    <h2 className="text-2xl font-semibold">Sottoscrizioni attive</h2>

                    {subscriptions.length === 0 ? (
                        <div className="text-center text-gray-500 py-10 bg-white rounded-xl shadow border">
                            Non hai ancora sottoscritto nessun treno.
                        </div>
                    ) : (
                        subscriptions.map((subscription) => (
                            <div
                                key={subscription.id}
                                className="p-4 bg-white shadow rounded-lg border flex justify-between items-center gap-4"
                            >
                                <div>
                                    <p><strong>ID:</strong> {subscription.id}</p>
                                    <p><strong>Train UUID:</strong> {subscription.train_uuid}</p>
                                    <p><strong>Piano:</strong> {planLabels[subscription.plan]}</p>
                                </div>

                                <button
                                    onClick={() => handleDelete(subscription.id)}
                                    className="bg-red-600 hover:bg-red-700 text-white px-3 py-2 rounded-lg"
                                >
                                    Elimina
                                </button>
                            </div>
                        ))
                    )}
                </div>
            </div>
        </MainLayout>
    );
}







