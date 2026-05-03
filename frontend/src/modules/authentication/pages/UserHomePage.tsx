import { useNavigate } from "react-router-dom";
import { Train, Ticket, AlarmClockCheck } from "lucide-react";
import { user_authorization } from "../../../core/hooks/user_authorization";
import MainLayout from "../../../core/layout/MainLayout";

export default function UserHomePage() {
    const navigate = useNavigate();
    const { user } = user_authorization();

    const features = [
        {
            title: "Treni",
            description: "Consulta orari, disponibilità e dettagli dei treni.",
            icon: <Train className="w-10 h-10 text-blue-600" />,
            action: () => navigate("/trains"),
        },
        {
            title: "I miei biglietti",
            description: "Visualizza e gestisci i biglietti acquistati.",
            icon: <Ticket className="w-10 h-10 text-green-600" />,
            action: () => navigate("/my-tickets"),
        },
        {
            title: "Sottoscrizioni",
            description: "Ricevi notifiche sui treni che ti interessano.",
            icon: <AlarmClockCheck className="w-10 h-10 text-purple-600" />,
            action: () => navigate("/subscriptions"),
        },
    ];

    return (
        <MainLayout>
            <div className="p-10 max-w-5xl mx-auto space-y-12">

                {/* HEADER */}
                <div className="text-center space-y-3">
                    <h1 className="text-4xl font-extrabold text-gray-900 tracking-tight">
                        Ciao {user?.userID || "Cliente"} 👋
                    </h1>
                    <p className="text-gray-600 text-lg">
                        Gestisci i tuoi viaggi, i tuoi biglietti e le tue notifiche.
                    </p>
                </div>

                {/* GRID */}
                <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
                    {features.map((item) => (
                        <div
                            key={item.title}
                            onClick={item.action}
                            className="cursor-pointer bg-white p-6 rounded-2xl shadow-md border border-gray-200
                                       hover:shadow-xl hover:-translate-y-1 transition-all duration-200"
                        >
                            <div className="flex items-center gap-4 mb-4">
                                {item.icon}
                                <h2 className="text-xl font-semibold text-gray-900">{item.title}</h2>
                            </div>
                            <p className="text-gray-600">{item.description}</p>
                        </div>
                    ))}
                </div>
            </div>
        </MainLayout>
    );
}
