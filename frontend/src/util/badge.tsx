export function StatusBadge({ status }: Readonly<{ status: "active" | "inactive" }>) {
    const color =
        status === "active"
            ? "bg-green-100 text-green-700 border-green-300"
            : "bg-red-100 text-red-700 border-red-300";

    return (
        <span className={`px-3 py-1 text-sm rounded-full border ${color}`}>
            {status === "active" ? "Attiva" : "Non attiva"}
        </span>
    );
}
