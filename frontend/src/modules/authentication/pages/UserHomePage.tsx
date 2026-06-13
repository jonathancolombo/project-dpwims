import { Train, Ticket, AlarmClockCheck } from "lucide-react";
import { Link } from "react-router-dom";
import { user_authorization } from "../../../core/hooks/user_authorization";
import MainLayout from "../../../core/layout/MainLayout";

export default function UserHomePage() {
    user_authorization();

    const features = [
        {
            title: "Itinerari e Biglietti",
            description: "Consulta gli orari dei treni e acquista i tuoi biglietti.",
            icon: <Train className="w-10 h-10 text-blue-600" />,
            to: "/user/schedules",
        },
        {
            title: "I miei biglietti",
            description: "Visualizza e gestisci i biglietti acquistati.",
            icon: <Ticket className="w-10 h-10 text-green-600" />,
            to: "/user/tickets",
        },
        {
            title: "Sottoscrizioni",
            description: "Ricevi notifiche sui treni che ti interessano.",
            icon: <AlarmClockCheck className="w-10 h-10 text-purple-600" />,
            to: "/user/subscriptions",
        },
    ];

    return (
        <MainLayout>
            <div className="p-10 max-w-5xl mx-auto space-y-12">

                {/* HEADER */}
                <div className="text-center space-y-3">
                    <h1 className="text-4xl font-extrabold text-gray-900 tracking-tight">
                        Ciao {"Cliente"} 👋
                    </h1>
                    <p className="text-gray-600 text-lg">
                        Gestisci i tuoi viaggi, i tuoi biglietti e le tue notifiche.
                    </p>
                </div>

                {/* GRID */}
                <div className="grid grid-cols-1 sm:grid-cols-2 gap-8">
                    {features.map((item) => (
                        <Link
                            key={item.title}
                            to={item.to}
                            className="block bg-white p-6 rounded-2xl shadow-md border border-gray-200 transition-all duration-200"
                        >
                            <div className="flex items-center gap-4 mb-4">
                                {item.icon}
                                <h2 className="text-xl font-semibold text-gray-900">{item.title}</h2>
                            </div>
                            <p className="text-gray-600">{item.description}</p>
                        </Link>
                    ))}
                </div>
            </div>
        </MainLayout>
    );
}