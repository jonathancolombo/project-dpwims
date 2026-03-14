import {useNavigate} from "react-router-dom";
import MainLayout from "../layout/MainLayout.tsx";
import {AlarmClockCheck, Banknote, BellElectric, StopCircle, Ticket, Train, UserIcon} from "lucide-react";

export default function IndexPage() {
    const navigate = useNavigate();

    return (
        <MainLayout>
            <div className="p-10 max-w-5xl mx-auto space-y-12">

                {/* HERO */}
                <div className="text-center space-y-4">
                    <h1 className="text-4xl font-extrabold text-gray-900">
                        Sistema di Gestione Ferroviaria
                    </h1>
                    <p className="text-gray-600 text-lg">
                        Benvenuto nel pannello di controllo. Seleziona una sezione per iniziare.
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

                    {/* CARD Itinerari */}

                    <div
                        onClick={() => navigate("/schedules")}
                        className="cursor-pointer p-6 bg-white shadow rounded-xl border hover:shadow-lg transition"
                    >
                        <div className="flex items-center gap-4 mb-4">
                            <StopCircle className="w-10 h-10 text-blue-600" />
                            <h2 className="text-xl font-semibold text-gray-900">Itinerari</h2>
                        </div>
                        <p className="text-gray-600 mt-2">Gestisci gli orari dei e le varie fermate dei treni relativi agli itinerari.</p>
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

                    {/* CARD Pagamenti */}
                    <div
                        onClick={() => navigate("/payments")}
                        className="cursor-pointer bg-white p-6 rounded-xl shadow hover:shadow-xl border border-gray-200 transition"
                    >
                        <div className="flex items-center gap-4 mb-4">
                            <Banknote className="w-10 h-10 text-green-600" />
                            <h2 className="text-xl font-semibold text-gray-900">Pagamenti</h2>
                        </div>
                        <p className="text-gray-600">
                            Visualizza, modifica e gestisci i pagamenti dei biglietti acquistati dai passeggeri.
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
            </div>
        </MainLayout>
    );
}
