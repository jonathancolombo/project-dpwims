import type { ReactNode } from "react";
import { Link } from "react-router-dom";
import { LogoutButton } from "./../../util/logout_button.tsx";
import { user_authorization } from "../hooks/user_authorization";
import { useEffect, useState } from "react";
import mqtt from "mqtt";

export default function MainLayout({ children }: Readonly<{ children: ReactNode }>) {
    const { role, user } = user_authorization();
    const userId = user?.userID ?? 0;
    const [notifications, setNotifications] = useState<string[]>([]);

    useEffect(() => {
        if (!userId || role !== "customer") return;

        const client = mqtt.connect("ws://localhost:9001");

        client.on("connect", () => {
            client.subscribe(`notifications/user/${userId}`);
        });

        client.on("message", (_topic, message) => {
            const payload = JSON.parse(message.toString());

            let text = "";
            if (payload.event === "schedule_updated") {
                text = `🚆 Itinerario aggiornato: ${payload.departure} → ${payload.arrival} | Stato: ${payload.status === "active" ? "Attivo" : "Non attivo"} | Prezzo: €${payload.price} | Alle ${new Date(payload.time).toLocaleTimeString("it-IT")}`;
            } else if (payload.event === "arrived") {
                text = `✅ Il treno è arrivato alle ${new Date(payload.time).toLocaleTimeString("it-IT")}`;
            } else {
                text = `🔔 Notifica: ${payload.event} alle ${new Date(payload.time).toLocaleTimeString("it-IT")}`;
            }

            setNotifications(prev => [text, ...prev]);
            setTimeout(() => {
                setNotifications(prev => prev.slice(0, -1));
            }, 5000);
        });

        return () => {
            client.end();
        };
    }, [userId, role]);

    return (
        <div className="flex h-screen bg-gray-100">

            {/* NOTIFICHE */}
            {notifications.length > 0 && (
                <div className="fixed top-4 right-4 z-50 space-y-2">
                    {notifications.map((notification, index) => (
                        <div
                            key={index}
                            className="bg-blue-600 text-white px-4 py-3 rounded-xl shadow-lg text-sm max-w-sm"
                        >
                            {notification}
                        </div>
                    ))}
                </div>
            )}

            {/* SIDEBAR */}
            <aside className="w-64 bg-white shadow-lg border-r border-gray-200 flex flex-col">
                <div className="px-6 py-5 border-b border-gray-200">
                    <h1 className="text-xl font-bold text-blue-700">Gestionale Treni</h1>
                    <p className="text-xs text-gray-500">
                        {role === "admin" ? "Pannello di Controllo Admin" : "Area Utente"}
                    </p>
                </div>

                <nav className="flex-1 px-4 py-6 space-y-2">
                    {role === "admin" && (
                        <>
                            <Link to="/admin" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">🏠 Home</Link>
                            <Link to="/trains" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">🚆 Treni</Link>
                            <Link to="/transactions" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">💳 Transazioni</Link>
                            <Link to="/users" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">👤 Utenti</Link>
                            <Link to="/subscriptions" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">🔔 Sottoscrizioni</Link>
                        </>
                    )}

                    {role === "customer" && (
                        <>
                            <Link to="/user" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">🏠 Home</Link>
                            <Link to="/user/schedules" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">🚆 Itinerari</Link>
                            <Link to="/user/tickets" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">🎫 I miei biglietti</Link>
                            <Link to="/user/subscriptions" className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium">🔔 Sottoscrizioni</Link>
                        </>
                    )}
                </nav>
            </aside>

            {/* MAIN CONTENT */}
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