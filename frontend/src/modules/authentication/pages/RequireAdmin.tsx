import {Navigate} from "react-router-dom";
import type {JSX} from "react";

export function RequireAdmin({ children }: Readonly<{ children: JSX.Element }>) {
    const user = JSON.parse(localStorage.getItem("user") || "null");

    if (user?.role !== "admin") {
        return <Navigate to="/login?target=admin" replace />;
    }

    return children;
}
