import { useMemo } from "react";
import {jwtDecode} from "jwt-decode";
export function user_authorization() {
    const token: string | null = localStorage.getItem("token");
    const user = useMemo(() => {
        if (!token) return null;
        try {
            return jwtDecode(token) as {
                email: string;
                userID: number;
                role: string;
                exp: number;
            };
        } catch {
            return null;
        }
    }, [token]);

    return {
        user,
        isLoggedIn: !!user,
        role: user?.role,
    };
}
