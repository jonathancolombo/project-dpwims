import {apiUsers} from "../../../core/api/client.ts";
import type {UpdateUserRequest, User} from "../types/user.ts";

export const getUsers = () => apiUsers.get<User[]>("/users");

export const deleteUser = (id: number) => apiUsers.delete(`/users/${id}`);
export const getUserById = (id: number) => apiUsers.get<User>(`/users/${id}`);

export const patchUser = (id: number, data: UpdateUserRequest) =>
    apiUsers.patch(`/users/${id}`, data);


export const createUser = (data: {
    username: string;
    email: string;
    telephone: string;
    fiscal_code: string;
    role: string;
    password: string
}) => apiUsers.post("/users", data);
