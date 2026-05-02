import { useNavigate } from "react-router-dom";
import MainLayout from "../layout/MainLayout";
import { useAuth } from "../hooks/useAuth";
import {
    AlarmClockCheck,
    Banknote,
    BellElectric,
    StopCircle,
    Ticket,
    Train,
    UserIcon
} from "lucide-react";

export default function IndexPage() {
    const navigate = useNavigate();
    const { isLoggedIn, role } = useAuth();

    // Se l'utente è loggato come USER → reindirizza alla sua dashboard
    if (isLoggedIn && role === "user") {
        navigate("/user/dashboard");
        return null;
    }

    return (
        <MainLayout>
            <div className="p-10 max-w-5xl mx-auto space-y-12">

                {/* SEZIONE DI SCELTA */}
                {!isLoggedIn && (
                    <div className="text-center space-y-6">
                        <h1 className="text-4xl font-extrabold text-gray-900">
                            Benvenuto
                        </h1>
                        <p className="text-gray-600 text-lg">
                            Seleziona l'area di accesso
                        </p>

                        <div className="flex justify-center gap-6">
                            <button
                                onClick={() => navigate("/login?target=user")}
                                className="px-6 py-3 rounded-lg bg-blue-600 text-white hover:bg-blue-700 transition"
                            >
                                Area Utente
                            </button>

                            <button
                                onClick={() => navigate("/login?target=admin")}
                                className="px-6 py-3 rounded-lg bg-gray-800 text-white hover:bg-gray-900 transition"
                            >
                                Area Admin
                            </button>
                        </div>
                    </div>
                )}

                {/* SEZIONE ADMIN (solo se loggato come admin) */}
                {isLoggedIn && role === "admin" && (
                    <>
                        <div className="text-center space-y-4">
                            <h2 className="text-3xl font-bold text-gray-900">
                                Pannello di Controllo Admin
                            </h2>
                            <p className="text-gray-600 text-lg">
                                Seleziona una sezione per iniziare.
                            </p>
                        </div>

                        {/* GRID DELLE SEZIONI */}
                        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">

                            {/* CARD TRENI */}
                            <div
                                onClick={() => navigate("/trains")}
                                className="cursor-pointer bg-white p-6 rounded-xl shadow hover:shadow-xl border border-gray-200 transition"
                            >
                                <div className="flex items-center gap-4 mb-4">
                                    <Train className="w-10 h-10 text-blue-600" />
                                    <h2 className="text-xl font-semibold text-gray-900">Treni</h2>
                                </div>
                                <p className="text-gray-600">
                                    Gestisci la flotta ferroviaria: crea, modifica e monitora i treni.
                                </p>
                            </div>

                            {/* CARD UTENTI */}
                            <div
                                onClick={() => navigate("/users")}
                                className="cursor-pointer bg-white p-6 rounded-xl shadow hover:shadow-xl border border-gray-200 transition"
                            >
                                <div className="flex items-center gap-4 mb-4">
                                    <UserIcon className="w-10 h-10 text-purple-600" />
                                    <h2 className="text-xl font-semibold text-gray-900">Utenti</h2>
                                </div>
                                <p className="text-gray-600">
                                    Gestisci gli account del sistema: ruoli, email, credenziali e informazioni personali.
                                </p>
                            </div>

                            {/* CARD STAZIONI */}
                            <div
                                onClick={() => navigate("/stations")}
                                className="cursor-pointer bg-white p-6 rounded-xl shadow hover:shadow-xl border border-gray-200 transition"
                            >
                                <div className="flex items-center gap-4 mb-4">
                                    <BellElectric className="w-10 h-10 text-blue-600" />
                                    <h2 className="text-xl font-semibold text-gray-900">Stazioni</h2>
                                </div>
                                <p className="text-gray-600">
                                    Gestisci le stazioni: crea, modifica e visualizza le stazioni disponibili.
                                </p>
                            </div>

                            {/* CARD ITINERARI */}
                            <div
                                onClick={() => navigate("/schedules")}
                                className="cursor-pointer p-6 bg-white shadow rounded-xl border hover:shadow-lg transition"
                            >
                                <div className="flex items-center gap-4 mb-4">
                                    <StopCircle className="w-10 h-10 text-blue-600" />
                                    <h2 className="text-xl font-semibold text-gray-900">Itinerari</h2>
                                </div>
                                <p className="text-gray-600 mt-2">
                                    Gestisci gli orari e le fermate dei treni.
                                </p>
                            </div>

                            {/* CARD BIGLIETTI */}
                            <div
                                onClick={() => navigate("/tickets")}
                                className="cursor-pointer bg-white p-6 rounded-xl shadow hover:shadow-xl border border-gray-200 transition"
                            >
                                <div className="flex items-center gap-4 mb-4">
                                    <Ticket className="w-10 h-10 text-green-600" />
                                    <h2 className="text-xl font-semibold text-gray-900">Biglietti</h2>
                                </div>
                                <p className="text-gray-600">
                                    Visualizza, modifica e gestisci i biglietti acquistati dai passeggeri.
                                </p>
                            </div>

                            {/* CARD PAGAMENTI */}
                            <div
                                onClick={() => navigate("/payments")}
                                className="cursor-pointer bg-white p-6 rounded-xl shadow hover:shadow-xl border border-gray-200 transition"
                            >
                                <div className="flex items-center gap-4 mb-4">
                                    <Banknote className="w-10 h-10 text-green-600" />
                                    <h2 className="text-xl font-semibold text-gray-900">Pagamenti</h2>
                                </div>
                                <p className="text-gray-600">
                                    Gestisci i pagamenti dei biglietti.
                                </p>
                            </div>

                            {/* CARD SOTTOSCRIZIONI */}
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
                    </>
                )}
            </div>
        </MainLayout>
    );
}
