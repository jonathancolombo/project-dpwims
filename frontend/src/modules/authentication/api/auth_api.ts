import axios from "axios";

const AUTH_API = "http://localhost:8085";

export function login(email: string, password: string) {
    return axios.post(`${AUTH_API}/auth/login`, {
        email,
        password
    });
}
