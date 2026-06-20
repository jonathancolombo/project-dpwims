import React, { type ReactNode } from "react";
import { Link } from "react-router-dom";
import { LogoutButton } from "./../../util/logout_button.tsx";
import { user_authorization } from "../hooks/user_authorization";
import { useEffect, useState } from "react";
import mqtt from "mqtt";

const NOTIFICATION_TIMEOUT_MS = 5000;

function formatNotification(payload: any): string {
    const time = new Date(payload.time).toLocaleTimeString("it-IT");

    if (payload.event === "schedule_updated") {
        const status = payload.status === "active" ? "Attivo" : "Non attivo";
        return `🚆 Itinerario aggiornato: ${payload.departure} → ${payload.arrival} | Stato: ${status} | Prezzo: €${payload.price} | Alle ${time}`;
    }

    if (payload.event === "arrived") {
        return `✅ Il treno è arrivato alle ${time}`;
    }

    return `🔔 Notifica: ${payload.event} alle ${time}`;
}

interface Notification {
    id: string;
    text: string;
}


function addTemporaryNotification(
    setNotifications: React.Dispatch<React.SetStateAction<Notification[]>>,
    text: string
) {
    const id = crypto.randomUUID();
    setNotifications(prev => [{ id, text }, ...prev]);

    setTimeout(() => {
        setNotifications(prev => prev.filter(notification => notification.id !== id));
    }, NOTIFICATION_TIMEOUT_MS);
}

function useMqttNotifications(userId: number, enabled: boolean) {
    const [notifications, setNotifications] = useState<Notification[]>([]);

    useEffect(() => {
        if (!userId || !enabled) return;

        const client = mqtt.connect("ws://localhost:9001");
        const topic = `notifications/user/${userId}`;

        const handleConnect = () => client.subscribe(topic);

        const handleMessage = (_topic: string, message: Buffer) => {
            const payload = JSON.parse(message.toString());
            const text = formatNotification(payload);
            addTemporaryNotification(setNotifications, text);
        };

        client.on("connect", handleConnect);
        client.on("message", handleMessage);

        return () => {
            client.end();
        };
    }, [userId, enabled]);

    return notifications;
}

function NotificationToasts({ notifications }: Readonly<{ notifications: Notification[] }>) {
    if (notifications.length === 0) return null;

    return (
        <div className="fixed top-4 right-4 z-50 space-y-2">
            {notifications.map(notification => (
                <div
                    key={notification.id}
                    className="bg-blue-600 text-white px-4 py-3 rounded-xl shadow-lg text-sm max-w-sm"
                >
                    {notification.text}
                </div>
            ))}
        </div>
    );
}

function AdminNav() {
    return (
        <>
            <Link to="/admin" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">🏠 Home</Link>
            <Link to="/trains" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">🚆 Treni</Link>
            <Link to="/transactions" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">💳 Transazioni</Link>
            <Link to="/users" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">👤 Utenti</Link>
            <Link to="/subscriptions" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">🔔 Sottoscrizioni</Link>
        </>
    );
}

function CustomerNav() {
    return (
        <>
            <Link to="/user" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">🏠 Home</Link>
            <Link to="/user/schedules" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">🚆 Itinerari</Link>
            <Link to="/user/tickets" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">🎫 I miei biglietti</Link>
            <Link to="/user/subscriptions" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">🔔 Sottoscrizioni</Link>
        </>
    );
}

export default function MainLayout({ children }: Readonly<{ children: ReactNode }>) {
    const { role, user } = user_authorization();
    const userId = user?.userID ?? 0;
    const notifications = useMqttNotifications(userId, role === "customer");

    return (
        <div className="flex h-screen bg-gray-100">
            <NotificationToasts notifications={notifications} />

            <aside className="w-64 bg-white shadow-lg border-r border-gray-200 flex flex-col">
                <div className="px-6 py-5 border-b border-gray-200">
                    <h1 className="text-xl font-bold text-blue-700">Gestionale Treni</h1>
                    <p className="text-xs text-gray-500">
                        {role === "admin" ? "Pannello di Controllo Admin" : "Area Utente"}
                    </p>
                </div>

                <nav className="flex-1 px-4 py-6 space-y-2">
                    {role === "admin" && <AdminNav />}
                    {role === "customer" && <CustomerNav />}
                </nav>
            </aside>

            <div className="flex-1 flex flex-col">
                <header className="h-16 bg-white shadow-sm border-b border-gray-200 flex items-center justify-between px-6">
                    <h2 className="text-lg font-semibold text-gray-800">Dashboard</h2>
                    <div className="flex items-center gap-4">
                        <LogoutButton />
                    </div>
                </header>

                <main className="flex-1 overflow-y-auto p-6">
                    {children}
                </main>
            </div>
        </div>
    );
}