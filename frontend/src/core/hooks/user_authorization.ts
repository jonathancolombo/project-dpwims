import { jwtDecode } from "jwt-decode";

interface DecodedToken {
    sub?: string;
    role: string;
    exp: number;
}

interface StoredUser {
    id?: number;
    role?: string;
}

interface AuthUser {
    userID: number;
    role: string | null;
    exp: number | null;
}

interface AuthResult {
    user: AuthUser | null;
    isLoggedIn: boolean;
    role: string | null;
}

const EMPTY_RESULT: AuthResult = { user: null, isLoggedIn: false, role: null };

function parseStoredUser(): StoredUser | null {
    const raw = localStorage.getItem("user");
    if (!raw) return null;
    try {
        return JSON.parse(raw);
    } catch {
        return null;
    }
}

function decodeToken(token: string | null): DecodedToken | null {
    if (!token) return null;
    try {
        return jwtDecode(token) as DecodedToken;
    } catch {
        return null;
    }
}

function buildResultFromID(userID: number, role: string | null, exp: number | null): AuthResult {
    if (Number.isNaN(userID) || userID <= 0) return EMPTY_RESULT;

    const user: AuthUser = { userID, role, exp };
    return { user, isLoggedIn: true, role: user.role };
}

export function useUserAuthorization(): AuthResult {
    const token = localStorage.getItem("token");
    const storedUser = parseStoredUser();

    if (!token && !storedUser) return EMPTY_RESULT;

    const decoded = decodeToken(token);

    if (decoded?.sub) {
        const userIDFromToken = Number(decoded.sub);
        const role = decoded.role ?? storedUser?.role ?? null;
        const result = buildResultFromID(userIDFromToken, role, decoded.exp ?? null);
        if (result.isLoggedIn) return result;
    }

    if (storedUser?.id) {
        const userIDFromStorage = Number(storedUser.id);
        const role = storedUser.role ?? null;
        return buildResultFromID(userIDFromStorage, role, decoded?.exp ?? null);
    }

    return EMPTY_RESULT;
}

export const user_authorization = useUserAuthorization;