import {Navigate} from "react-router-dom";
import type {JSX} from "react";

export function RequireUser({ children }: Readonly<{ children: JSX.Element }>) {
    const user = JSON.parse(localStorage.getItem("user") || "null");

    if (user?.role !== "customer") {
        return <Navigate to="/login?target=user" replace />;
    }

    return children;
}
