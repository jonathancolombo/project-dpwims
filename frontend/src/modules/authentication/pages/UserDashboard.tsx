import { useNavigate } from "react-router-dom";
import { Train, Ticket, AlarmClockCheck } from "lucide-react";
import MainLayout from "../../../core/layout/MainLayout.tsx";
import {useAuth} from "../../../core/hooks/useAuth.ts";

export default function UserDashboard() {
    const navigate = useNavigate();
    const { user } = useAuth();

    return (
        <MainLayout>
            <div className="p-10 max-w-4xl mx-auto space-y-10">

                {/* HEADER */}
                <div className="text-center space-y-2">
                    <h1 className="text-3xl font-bold text-gray-900">
                        Benvenuto, {user?.email || "Utente"}
                    </h1>
                    <p className="text-gray-600">
                        Gestisci i tuoi viaggi e le tue notifiche.
                    </p>
                </div>

                {/* GRID DELLE FUNZIONALITÀ */}
                <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">

                    {/* TRENI */}
                    <div
                        onClick={() => navigate("/trains")}
                        className="cursor-pointer bg-white p-6 rounded-xl shadow hover:shadow-xl border border-gray-200 transition"
                    >
                        <div className="flex items-center gap-4 mb-4">
                            <Train className="w-10 h-10 text-blue-600" />
                            <h2 className="text-xl font-semibold text-gray-900">Treni</h2>
                        </div>
                        <p className="text-gray-600">
                            Consulta i treni disponibili e gli orari.
                        </p>
                    </div>

                    {/* BIGLIETTI */}
                    <div
                        onClick={() => navigate("/my-tickets")}
                        className="cursor-pointer bg-white p-6 rounded-xl shadow hover:shadow-xl border border-gray-200 transition"
                    >
                        <div className="flex items-center gap-4 mb-4">
                            <Ticket className="w-10 h-10 text-green-600" />
                            <h2 className="text-xl font-semibold text-gray-900">I miei biglietti</h2>
                        </div>
                        <p className="text-gray-600">
                            Visualizza i biglietti che hai acquistato.
                        </p>
                    </div>

                    {/* SOTTOSCRIZIONI */}
                    <div
                        onClick={() => navigate("/subscriptions")}
                        className="cursor-pointer bg-white p-6 rounded-xl shadow hover:shadow-xl border border-gray-200 transition"
                    >
                        <div className="flex items-center gap-4 mb-4">
                            <AlarmClockCheck className="w-10 h-10 text-purple-600" />
                            <h2 className="text-xl font-semibold text-gray-900">Sottoscrizioni</h2>
                        </div>
                        <p className="text-gray-600">
                            Gestisci le notifiche sui treni che ti interessano.
                        </p>
                    </div>

                </div>
            </div>
        </MainLayout>
    );
}
