import { apiAuth } from "../../../core/api/client";

export function login(email: string, password: string) {
    return apiAuth.post("/auth/login", {
        email,
        password
    });
}