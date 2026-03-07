import type { ReactNode } from "react";
import { Link } from "react-router-dom";

export default function MainLayout({ children }: Readonly<{ children: ReactNode }>) {
    return (
        <div className="flex h-screen bg-gray-100">

            {/* SIDEBAR */}
            <aside className="w-64 bg-white shadow-lg border-r border-gray-200 flex flex-col">
                <div className="px-6 py-5 border-b border-gray-200">
                    <h1 className="text-xl font-bold text-blue-700">Railway Admin</h1>
                    <p className="text-xs text-gray-500">Control Panel</p>
                </div>

                <nav className="flex-1 px-4 py-6 space-y-2">
                    <Link
                        to="/"
                        className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium"
                    >
                        🏠 Home
                    </Link>

                    <Link
                        to="/trains"
                        className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium"
                    >
                        🚆 Treni
                    </Link>

                    <Link
                        to="/tickets"
                        className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium"
                    >
                        🎫 Biglietti
                    </Link>

                    <Link
                        to="/users"
                        className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium"
                    >
                        👤 Utenti
                    </Link>

                    <Link
                        to="/notifications"
                        className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium"
                    >
                        🔔 Notifiche
                    </Link>

                    <Link
                        to="/settings"
                        className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium"
                    >
                        ⚙️ Impostazioni
                    </Link>

                </nav>


                <div className="px-6 py-4 border-t border-gray-200 text-sm text-gray-500">
                    Admin logged in
                </div>
            </aside>

            {/* MAIN CONTENT */}
            <div className="flex-1 flex flex-col">

                {/* TOP BAR */}
                <header className="h-16 bg-white shadow-sm border-b border-gray-200 flex items-center justify-between px-6">
                    <h2 className="text-lg font-semibold text-gray-800">Dashboard</h2>

                    <div className="flex items-center gap-4">
                        <button className="text-gray-600 hover:text-gray-800">🔍</button>
                        <button className="text-gray-600 hover:text-gray-800">⚙️</button>
                        <button className="text-gray-600 hover:text-gray-800">🚪 Logout</button>
                    </div>
                </header>

                {/* PAGE CONTENT */}
                <main className="flex-1 overflow-y-auto p-6">
                    {children}
                </main>
            </div>
        </div>
    );
}
