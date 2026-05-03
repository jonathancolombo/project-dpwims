import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import MainLayout from "../../../core/layout/MainLayout";
import { user_authorization } from "../../../core/hooks/user_authorization";

import {
    AlarmClockCheck,
    Banknote,
    BellElectric,
    StopCircle,
    Ticket,
    Train,
    UserIcon
} from "lucide-react";

export default function AdminHomePage() {
    const navigate = useNavigate();
    const { isLoggedIn, role } = user_authorization();

    useEffect(() => {
        if (role !== null && (!isLoggedIn || role !== "admin")) {
            navigate("/login?target=admin");
        }
    }, [isLoggedIn, role, navigate]);

    if (role === null) return null;

    if (!isLoggedIn || role !== "admin") return null;

    return (
        <MainLayout>
            <div className="p-10 max-w-5xl mx-auto space-y-12">

                <div className="text-center space-y-4">
                    <h2 className="text-3xl font-bold text-gray-900">
                        Pannello di Controllo Admin
                    </h2>
                    <p className="text-gray-600 text-lg">
                        Seleziona una sezione per iniziare.
                    </p>
                </div>

                <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">

                    <div
                        onClick={() => navigate("/trains")}
                        className="cursor-pointer bg-white p-6 rounded-xl shadow hover:shadow-xl border border-gray-200 transition"
                    >
                        <div className="flex items-center gap-4 mb-4">
                            <Train className="w-10 h-10 text-blue-600" />
                            <h2 className="text-xl font-semibold text-gray-900">Treni</h2>
                        </div>
                        <p className="text-gray-600">
                            Gestisci la flotta ferroviaria.
                        </p>
                    </div>

                    <div
                        onClick={() => navigate("/users")}
                        className="cursor-pointer bg-white p-6 rounded-xl shadow hover:shadow-xl border border-gray-200 transition"
                    >
                        <div className="flex items-center gap-4 mb-4">
                            <UserIcon className="w-10 h-10 text-purple-600" />
                            <h2 className="text-xl font-semibold text-gray-900">Utenti</h2>
                        </div>
                        <p className="text-gray-600">
                            Gestisci gli account del sistema.
                        </p>
                    </div>

                    <div
                        onClick={() => navigate("/stations")}
                        className="cursor-pointer bg-white p-6 rounded-xl shadow hover:shadow-xl border border-gray-200 transition"
                    >
                        <div className="flex items-center gap-4 mb-4">
                            <BellElectric className="w-10 h-10 text-blue-600" />
                            <h2 className="text-xl font-semibold text-gray-900">Stazioni</h2>
                        </div>
                        <p className="text-gray-600">
                            Gestisci le stazioni ferroviarie.
                        </p>
                    </div>

                    <div
                        onClick={() => navigate("/schedules")}
                        className="cursor-pointer p-6 bg-white shadow rounded-xl border hover:shadow-lg transition"
                    >
                        <div className="flex items-center gap-4 mb-4">
                            <StopCircle className="w-10 h-10 text-blue-600" />
                            <h2 className="text-xl font-semibold text-gray-900">Itinerari</h2>
                        </div>
                        <p className="text-gray-600 mt-2">
                            Gestisci orari e fermate.
                        </p>
                    </div>

                    <div
                        onClick={() => navigate("/tickets")}
                        className="cursor-pointer bg-white p-6 rounded-xl shadow hover:shadow-xl border border-gray-200 transition"
                    >
                        <div className="flex items-center gap-4 mb-4">
                            <Ticket className="w-10 h-10 text-green-600" />
                            <h2 className="text-xl font-semibold text-gray-900">Biglietti</h2>
                        </div>
                        <p className="text-gray-600">
                            Gestisci i biglietti acquistati.
                        </p>
                    </div>

                    <div
                        onClick={() => navigate("/payments")}
                        className="cursor-pointer bg-white p-6 rounded-xl shadow hover:shadow-xl border border-gray-200 transition"
                    >
                        <div className="flex items-center gap-4 mb-4">
                            <Banknote className="w-10 h-10 text-green-600" />
                            <h2 className="text-xl font-semibold text-gray-900">Pagamenti</h2>
                        </div>
                        <p className="text-gray-600">
                            Gestisci i pagamenti.
                        </p>
                    </div>

                    <div
                        onClick={() => navigate("/subscriptions")}
                        className="cursor-pointer bg-white p-6 rounded-xl shadow hover:shadow-xl border border-gray-200 transition"
                    >
                        <div className="flex items-center gap-4 mb-4">
                            <AlarmClockCheck className="w-10 h-10 text-purple-600" />
                            <h2 className="text-xl font-semibold text-gray-900">Sottoscrizioni</h2>
                        </div>
                        <p className="text-gray-600">
                            Gestisci le sottoscrizioni degli utenti.
                        </p>
                    </div>

                </div>
            </div>
        </MainLayout>
    );
}
