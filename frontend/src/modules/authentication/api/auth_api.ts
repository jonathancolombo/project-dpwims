import axios from "axios";

export async function login(email: string, password: string) {
    return axios.post("/auth/login", { email, password });
}
