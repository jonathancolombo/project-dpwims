import { useNavigate } from "react-router-dom";
import MainLayout from "../layout/MainLayout";
import { user_authorization } from "../hooks/user_authorization";

export default function IndexPage() {
    const navigate = useNavigate();
    const { isLoggedIn, role } = user_authorization();

    if (isLoggedIn && role === "customer") {
        navigate("/user");
        return null;
    }

    if (isLoggedIn && role === "admin") {
        navigate("/admin");
        return null;
    }

    return (
        <MainLayout>
            <div className="p-10 max-w-5xl mx-auto space-y-12">
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
            </div>
        </MainLayout>
    );
}