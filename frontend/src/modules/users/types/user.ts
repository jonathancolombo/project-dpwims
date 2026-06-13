export type UserRole = "admin" | "customer";

export interface User {
    id: number;
    username: string;
    email: string;
    fiscal_code: string;
    telephone: string;
    role: UserRole;
}

export interface UpdateUserRequest {
    username?: string;
    email?: string;
    telephone?: string;
    fiscal_code?: string;
    role?: UserRole;
    password?: string;
}

export type CreateUserDTO = Omit<User, "id"> & {
    password: string;
};
