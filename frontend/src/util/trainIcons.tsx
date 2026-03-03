import type {ReactNode} from "react";
import {Train, TrainFront, TrainTrack} from "lucide-react";

// oppure qualsiasi set di icone preferisci

export function getTrainIcon(type: string): ReactNode {
    switch (type) {
        case "regional":
            return <Train className="w-6 h-6 text-blue-600" />;
        case "intercity":
            return <TrainFront className="w-6 h-6 text-green-600" />;
        case "highspeed":
            return <TrainTrack className="w-6 h-6 text-red-600" />;
        default:
            return <Train className="w-6 h-6 text-gray-500" />;
    }
}
