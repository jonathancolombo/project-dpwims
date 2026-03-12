import { useEffect, useState } from "react";
import { getStations } from "../api/stations_api.ts";
import * as React from "react";

interface StationSelectProps {
    value: number;
    onChange: (id: number, name: string) => void;
}

export default function StationSelect({ value, onChange }: Readonly<StationSelectProps>) {
    const [stations, setStations] = useState<{ id: number; name: string }[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        getStations()
            .then((response) => setStations(response.data))
            .finally(() => setLoading(false));
    }, []);

    if (loading) {
        return <div className="text-gray-500">Caricamento stazioni...</div>;
    }

    const handleChange = (element: React.ChangeEvent<HTMLSelectElement>) => {
        const id = Number(element.target.value);
        const station = stations.find((station) => station.id === id);

        if (station) {
            onChange(station.id, station.name);
        }
    };

    return (
        <select
            className="w-full border p-2 rounded"
            value={value}
            onChange={handleChange}
        >
            <option value={0}>Seleziona una stazione</option>
            {stations.map((station) => (
                <option key={station.id} value={station.id}>
                    {station.name}
                </option>
            ))}
        </select>
    );
}
