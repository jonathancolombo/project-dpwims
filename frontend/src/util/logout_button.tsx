import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { logout } from "../core/hooks/user_logout.ts";

export function LogoutButton() {
    const navigate = useNavigate();
    const [open, setOpen] = useState(false);
    const [loading, setLoading] = useState(false);

    const handleConfirm = async () => {
        setLoading(true);
        try {
            await logout(navigate);
        } finally {
            setLoading(false);
            setOpen(false);
        }
    };

    return (
        <>
            <button
                onClick={() => setOpen(true)}
                className="px-3 py-1 rounded bg-red-600 text-white hover:bg-red-700"
                aria-haspopup="dialog"
                aria-expanded={open}
            >
                Logout
            </button>

            {open && (
                <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/40">
                    <div className="bg-white rounded-lg shadow-lg w-full max-w-sm p-6">
                        <h3 className="text-lg font-semibold mb-2">Confermi il logout?</h3>
                        <p className="text-sm text-gray-600 mb-4">
                            Verrai disconnesso e tornerai alla homepage.
                        </p>

                        <div className="flex justify-end gap-3">
                            <button
                                onClick={() => setOpen(false)}
                                className="px-4 py-2 rounded bg-gray-100 hover:bg-gray-200"
                                disabled={loading}
                            >
                                Annulla
                            </button>

                            <button
                                onClick={handleConfirm}
                                className="px-4 py-2 rounded bg-red-600 text-white hover:bg-red-700 flex items-center gap-2"
                                disabled={loading}
                            >
                                {loading ? (
                                    <svg className="w-4 h-4 animate-spin" viewBox="0 0 24 24">
                                        <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" fill="none"/>
                                        <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v4a4 4 0 00-4 4H4z"/>
                                    </svg>
                                ) : null}
                                Esci
                            </button>
                        </div>
                    </div>
                </div>
            )}
        </>
    );
}
