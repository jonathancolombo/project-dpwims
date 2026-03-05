import {useEffect, useState} from "react";
import MainLayout from "../../../core/layout/MainLayout";
import {getTickets, type Ticket} from "../api/ticketsApi";

export default function TicketsPage() {
    const [tickets, setTickets] = useState<Ticket[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        getTickets().then(res => {
            setTickets(res.data);
            setLoading(false);
        });
    }, []);

    if (loading) return <MainLayout>Caricamento...</MainLayout>;

    return (
        <MainLayout>
            <div className="p-6 space-y-6">
                <h1 className="text-3xl font-bold">Biglietti</h1>

                <div className="space-y-4">
                    {tickets.map(ticket => (
                        <div key={ticket.uuid} className="p-4 bg-white shadow rounded-lg border">
                            <div className="flex justify-between">
                                <h2 className="text-lg font-semibold">
                                    Biglietto #{ticket.uuid}
                                </h2>

                                <span className={`px-2 py-1 rounded text-white ${
                                    ticket.status === "booked" ? "bg-green-600" :
                                        ticket.status === "used" ? "bg-gray-500" :
                                            "bg-red-600"
                                }`}>
                                    {ticket.status}
                                </span>
                            </div>

                            <div className="mt-2 text-gray-700 space-y-1">
                                <div><strong>User ID:</strong> {ticket.user_id}</div>
                                <div><strong>Train UUID:</strong> {ticket.train_id}</div>
                                <div><strong>Schedule ID:</strong> {ticket.schedule_id}</div>
                                <div><strong>Posto:</strong> {ticket.seat_number}</div>
                                <div><strong>Prezzo:</strong> € {ticket.price}</div>
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </MainLayout>
    );
}
