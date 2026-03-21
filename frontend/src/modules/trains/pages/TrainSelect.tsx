import {useEffect, useState} from "react";
import {getTrains} from "../api/trains_api.ts";
import type {Train} from "../types/train.ts";

interface TrainSelectProps {
    value: string;
    onChange: (value: string) => void;
}

export default function TrainSelect({ value, onChange }: Readonly<TrainSelectProps>) {
    const [trains, setTrains] = useState<Train[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        getTrains()
            .then((response) => setTrains(response.data))
            .finally(() => setLoading(false));
    }, []);

    if (loading) {
        return <div className="text-gray-500">Caricamento treni...</div>;
    }

    return (
        <select
            className="w-full border p-2 rounded"
            value={value}
            onChange={(element) => onChange(element.target.value)}
        >
            <option value="">Seleziona un treno</option>
            {trains.map((train) => (
                <option key={train.uuid} value={train.uuid}>
                    {train.train_number}
                </option>
            ))}
        </select>
    );
}
