import {ShieldCheck, User} from "lucide-react";

export function getUserRoleIcon(role: string) {
    switch (role) {
        case "admin":
            return <ShieldCheck className="w-6 h-6 text-purple-600" />;
        case "customer":
            return <User className="w-6 h-6 text-blue-600" />;
        default:
            return <User className="w-6 h-6 text-gray-500" />;
    }
}
