import { useEffect, useState } from "react";
import { getStations } from "../api/stations_api";

interface Station {
    id: number;
    name: string;
}

interface StationNameSelectProps {
    value: string;
    onChange: (value: string) => void;
}

export default function StationNameSelect({ value, onChange }: Readonly<StationNameSelectProps>) {
    const [stations, setStations] = useState<Station[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        getStations()
            .then((res) => setStations(res.data))
            .finally(() => setLoading(false));
    }, []);

    if (loading) {
        return (
            <select className="w-full border p-2 rounded mt-1" disabled>
                <option>Caricamento stazioni...</option>
            </select>
        );
    }

    return (
        <select
            className="w-full border p-2 rounded mt-1"
            value={value}
            onChange={(element) => onChange(element.target.value)}
        >
            <option value="">Seleziona una stazione</option>

            {stations.map((station) => (
                <option key={station.id} value={station.name}>
                    {station.name}
                </option>
            ))}
        </select>
    );
}
