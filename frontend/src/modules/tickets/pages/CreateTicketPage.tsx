import { useState, useEffect } from "react";
import MainLayout from "../../../core/layout/MainLayout";
import { createTicket } from "../api/tickets_api";
import { getSchedules } from "../../trains/api/schedules_api";
import { getUsers } from "../../users/api/users_api";
import { useNavigate } from "react-router-dom";
import type {User} from "../../users/types/user.ts";
import type {Schedule} from "../../trains/types/schedule.ts";

export default function CreateTicketPage() {
    const navigate = useNavigate();
    const [users, setUsers] = useState<User[]>([]);
    const [schedules, setSchedules] = useState<Schedule[]>([]);
    const [userId, setUserId] = useState(0);
    const [scheduleId, setScheduleId] = useState(0);
    const [seatNumber, setSeatNumber] = useState("");
    const [price, setPrice] = useState(0);
    const [message, setMessage] = useState("");
    const [trainId, setTrainId] = useState("");


    useEffect(() => {
        getUsers().then(response => setUsers(response.data));
        getSchedules().then(response => setSchedules(response.data));
    }, []);

    const handleCreate = async () => {
        try {
            await createTicket({
                user_id: userId,
                schedule_id: scheduleId,
                train_id: trainId,
                seat_number: seatNumber,
                price,
                status: "booked"
            });

            navigate("/tickets");
        } catch {
            setMessage("Errore durante la creazione del biglietto.");
        }
    };

    return (
        <MainLayout>
            <div className="p-6 max-w-xl mx-auto space-y-6">
                <h1 className="text-3xl font-bold">Nuovo Biglietto</h1>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">
                        {message}
                    </div>
                )}

                <div className="space-y-4 bg-white p-6 rounded-xl shadow border">

                    <div>
                        <label className="block text-sm font-medium">Utente</label>
                        <select
                            className="w-full border p-2 rounded"
                            value={userId}
                            onChange={(element) => setUserId(Number(element.target.value))}
                        >
                            <option value="">Seleziona utente</option>
                            {users.map(user => (
                                <option key={user.id} value={user.id}>{user.email}</option>
                            ))}
                        </select>
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Itinerario</label>
                        <select
                            className="w-full border p-2 rounded"
                            value={scheduleId}
                            onChange={(element) => {
                                const id = Number(element.target.value);
                                setScheduleId(id);

                                const selected = schedules.find(schedule => schedule.id === id);
                                if (selected) {
                                    setTrainId(selected.train_id);
                                    setPrice(selected.price);
                                }
                            }}
                        >

                        <option value={0}>Seleziona itinerario</option>
                            {schedules.map(schedule => (
                                <option key={schedule.id} value={schedule.id}>
                                    {schedule.departure} → {schedule.arrival} (Treno {schedule.train_id})
                                </option>
                            ))}
                        </select>
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Treno</label>
                        <input
                            type="text"
                            value={trainId}
                            disabled
                            className="w-full mt-1 p-2 border rounded bg-gray-100"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Stato</label>
                        <input
                            type="text"
                            value="Prenotato"
                            disabled
                            className="w-full mt-1 p-2 border rounded bg-gray-100"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Posto</label>
                        <input
                            className="w-full border p-2 rounded"
                            value={seatNumber}
                            onChange={(element) => setSeatNumber(element.target.value)}
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Prezzo</label>
                        <input
                            type="number"
                            className="w-full border p-2 rounded"
                            value={price}
                            onChange={(element) => setPrice(Number(element.target.value))}
                        />
                    </div>

                </div>

                <button
                    onClick={handleCreate}
                    className="w-full bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700"
                >
                    Crea Biglietto
                </button>
            </div>
        </MainLayout>
    );
}
