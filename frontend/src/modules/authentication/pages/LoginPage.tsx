import {useLocation, useNavigate} from "react-router-dom";
import {useState} from "react";
import {login} from "../api/auth_api";

export default function LoginPage() {
    const location = useLocation();
    const navigate = useNavigate();

    const params = new URLSearchParams(location.search);
    const target : string = params.get("target") ?? "user";

    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");

    async function handleLogin() {
        try {
            const response = await login(email, password);
            const user = response.data;

            localStorage.setItem("token", user.token);
            localStorage.setItem("user", JSON.stringify(user));

            if (target === "admin" && user.role !== "admin") {
                setError("Non hai i permessi per accedere all'area admin.");
                return;
            }

            if (target === "user" && user.role !== "customer") {
                setError("Non hai i permessi per accedere all'area utente.");
                return;
            }

            // Redirect
            if (user.role === "admin") {
                navigate("/admin");
            } else {
                navigate("/user");
            }

        } catch (err) {
            setError("Credenziali non valide");
        }
    }

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-100">
            <div className="bg-white p-8 rounded-xl shadow-lg w-full max-w-md space-y-6 border">

                <h1 className="text-2xl font-bold text-center">
                    {target === "admin" ? "Login Admin" : "Login Utente"}
                </h1>

                {error && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">
                        {error}
                    </div>
                )}

                <div className="space-y-4">
                    <div>
                        <label className="block text-sm font-medium">Email</label>
                        <input
                            type="email"
                            className="w-full border p-2 rounded mt-1"
                            value={email}
                            onChange={element => setEmail(element.target.value)}
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Password</label>
                        <input
                            type="password"
                            className="w-full border p-2 rounded mt-1"
                            value={password}
                            onChange={element => setPassword(element.target.value)}
                        />
                    </div>
                </div>

                <button
                    onClick={handleLogin}
                    className="w-full bg-gray-800 text-white py-2 rounded-lg hover:bg-gray-900"
                >
                    Accedi
                </button>

                <button
                    onClick={() => navigate("/")}
                    className="w-full bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg"
                >
                    Torna alla home
                </button>
            </div>
        </div>
    );
}
