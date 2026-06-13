import type {NavigateFunction} from "react-router-dom";
import { apiTrains, apiUsers, apiTickets, apiNotifications } from "../api/client.ts";

export async function logout(navigate?: NavigateFunction) {
    try {
        // chiamo il backend per invalidare il token
        const token : string | null = localStorage.getItem("token");
        if (token) {
            try {
                await apiUsers.post("/auth/logout", {}, { headers: { Authorization: `Bearer ${token}` } });
            } catch {

            }
        }
    } finally {
        // Rimuovo credenziali
        localStorage.removeItem("token");
        localStorage.removeItem("user");

        // Pulisco header Authorization dalle istanze axios
        try { delete apiTrains.defaults.headers.common["Authorization"]; } catch {}
        try { delete apiUsers.defaults.headers.common["Authorization"]; } catch {}
        try { delete apiTickets.defaults.headers.common["Authorization"]; } catch {}
        try { delete apiNotifications.defaults.headers.common["Authorization"]; } catch {}

        // Redirect alla login
        if (navigate) navigate("/login");
        else globalThis.location.href = "/login";
    }
}
