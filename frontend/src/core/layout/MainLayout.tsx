import type {ReactNode} from "react";

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
                    <a className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium cursor-pointer">
                        🚆 Treni
                    </a>
                    <a className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium cursor-pointer">
                        🎫 Biglietti
                    </a>
                    <a className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium cursor-pointer">
                        👤 Utenti
                    </a>
                    <a className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium cursor-pointer">
                        🔔 Notifiche
                    </a>
                    <a className="block px-4 py-2 rounded-lg hover:bg-blue-50 text-gray-700 font-medium cursor-pointer">
                        ⚙️ Impostazioni
                    </a>
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
