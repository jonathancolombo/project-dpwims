import {useEffect, useState} from "react";
import {getStations} from "../api/stations_api.ts";

interface StationSelectProps {
    value: number;
    onChange: (value: number) => void;
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

    return (
        <select
            className="w-full border p-2 rounded"
            value={value}
            onChange={(element) => onChange(Number(element.target.value))}
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


