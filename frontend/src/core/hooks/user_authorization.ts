import {jwtDecode} from "jwt-decode";
export function useUserAuthorization() {
    const token: string | null = localStorage.getItem("token");
    const storedUserRaw = localStorage.getItem("user");
    let storedUser: { id?: number; role?: string } | null = null;

    if (storedUserRaw) {
        try {
            storedUser = JSON.parse(storedUserRaw);
        } catch {
            storedUser = null;
        }
    }

    if (!token && !storedUser) return { user: null, isLoggedIn: false, role: null };

    try {
        const decoded = token ? jwtDecode(token) as {
            sub?: string;
            role: string;
            exp: number;
        } : null;

        const userIDFromToken = decoded?.sub ? Number(decoded.sub) : Number.NaN;
        if (!Number.isNaN(userIDFromToken) && userIDFromToken > 0) {
            const user = {
                userID: userIDFromToken,
                role: decoded?.role ?? storedUser?.role ?? null,
                exp: decoded?.exp,
            };

            return {
                user,
                isLoggedIn: !!user,
                role: user.role,
            };
        }

        const userIDFromStorage = Number(storedUser?.id);
        if (!Number.isNaN(userIDFromStorage) && userIDFromStorage > 0) {
            const user = {
                userID: userIDFromStorage,
                role: storedUser?.role ?? null,
                exp: decoded?.exp,
            };

            return {
                user,
                isLoggedIn: !!user,
                role: user.role,
            };
        }

        return { user: null, isLoggedIn: false, role: null };
    } catch {
        if (storedUser !== null)
        {
            return storedUser?.id
                ? { user: { userID: Number(storedUser.id), role: storedUser.role ?? null, exp: null }, isLoggedIn: true, role: storedUser.role ?? null }
                : { user: null, isLoggedIn: false, role: null };
        }
        return { user: null, isLoggedIn: false, role: null };
    }

}

export const user_authorization = useUserAuthorization;

