import {apiUsers} from "../../../core/api/client.ts";
import type {CreateUserDTO, UpdateUserRequest, User} from "../types/user.ts";

export const getUsers = () => apiUsers.get<User[]>("/users");

export const deleteUser = (id: number) => apiUsers.delete(`/users/${id}`);

export const getUserById = (id: number) => apiUsers.get<User>(`/users/${id}`);

export const patchUser = (id: number, data: UpdateUserRequest) =>
    apiUsers.patch(`/users/${id}`, data);

export const createUser = (data: CreateUserDTO) =>
    apiUsers.post("/users", data);
