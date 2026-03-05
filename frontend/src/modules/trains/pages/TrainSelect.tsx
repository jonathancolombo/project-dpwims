import {useEffect, useState} from "react";
import {getTrains} from "../api/trainsApi.ts";
import type {Train} from "../types/Train.ts";

interface TrainSelectProps {
    value: string;
    onChange: (value: string) => void;
}

export default function TrainSelect({ value, onChange }: Readonly<TrainSelectProps>) {
    const [trains, setTrains] = useState<Train[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        getTrains()
            .then((res) => setTrains(res.data))
            .finally(() => setLoading(false));
    }, []);

    if (loading) {
        return <div className="text-gray-500">Caricamento treni...</div>;
    }

    return (
        <select
            className="w-full border p-2 rounded"
            value={value}
            onChange={(e) => onChange(e.target.value)}
        >
            <option value="">Seleziona un treno</option>
            {trains.map((t) => (
                <option key={t.uuid} value={t.uuid}>
                    {t.train_number}
                </option>
            ))}
        </select>
    );
}
