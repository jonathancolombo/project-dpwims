import {useEffect, useState} from "react";
import MainLayout from "../../../core/layout/MainLayout.tsx";
import {deleteUser, getUsers} from "../api/users_api.ts";
import type {User} from "../types/user.ts";
import {useNavigate} from "react-router-dom";
import {getUserRoleIcon} from "../../../util/user_icons.tsx";

export default function UsersPage() {
    const [users, setUsers] = useState<User[]>([]);
    const navigate = useNavigate();

    useEffect(() => {
        getUsers().then((response) => setUsers(response.data));
    }, []);

    const handleDelete = async (id: number) => {
        if (!globalThis.confirm("Vuoi davvero cancellare questo utente?")) return;

        await deleteUser(id);
        setUsers((previousUsers) => previousUsers.filter((user) => user.id !== id));
    };

    return (
        <MainLayout>
            <div className="p-6 space-y-8">

                {/* HEADER */}
                <div className="flex justify-between items-center">
                    <div>
                        <h1 className="text-3xl font-bold text-gray-900">Gestione Utenti</h1>
                        <p className="text-gray-600 mt-1">Amministrazione degli account del sistema</p>
                    </div>

                    <div className="flex gap-3">
                        <button
                            onClick={() => navigate(-1)}
                            className="bg-gray-200 hover:bg-gray-300 text-gray-700 px-4 py-2 rounded-lg font-medium transition"
                        >
                            ← Torna indietro
                        </button>
                        <button
                            onClick={() => navigate("/users/create")}
                            className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg font-medium shadow transition"
                        >
                            + Crea Utente
                        </button>
                    </div>
                </div>
                {/* LISTA UTENTI */}
                <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                    {users.map((user) => (
                        <div
                            key={user.id}
                            className="bg-white rounded-xl shadow-md border border-gray-200 p-6"
                        >
                            <div className="flex justify-between items-start mb-4">
                                <div className="flex items-center gap-3">
                                    {getUserRoleIcon(user.role)}
                                    <div>
                                        <h2 className="text-xl font-semibold text-gray-900">{user.username}</h2>
                                        <p className="text-sm text-gray-500">{user.email}</p>
                                    </div>
                                </div>

                            </div>

                            <div className="space-y-2 text-gray-700">
                                <p><span className="font-semibold">Ruolo: </span>
                                    {user.role === "admin" ? "Admin" : "Cliente"}
                                </p>

                                <p><span className="font-semibold">Telefono:</span> {user.telephone}</p>
                                <p><span className="font-semibold">Codice Fiscale:</span> {user.fiscal_code}</p>
                            </div>

                            <div className="mt-5 flex gap-3">
                                <button
                                    className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg font-medium transition"
                                    onClick={() => navigate(`/users/${user.id}`)}
                                >
                                    Modifica
                                </button>

                                <button
                                    className="flex-1 bg-red-600 hover:bg-red-700 text-white py-2 rounded-lg font-medium transition"
                                    onClick={() => handleDelete(user.id)}
                                >
                                    Cancella
                                </button>
                            </div>

                        </div>

                    ))}

                </div>

            </div>

        </MainLayout>
    );
}
