import { useEffect } from "react";
import mqtt from "mqtt";

export function useMqttNotifications(userID: number, onMessage: (payload: string) => void) {
    useEffect(() => {
        if (!userID) return;

        const client = mqtt.connect("ws://localhost:9001");

        client.on("connect", () => {
            client.subscribe(`notifications/user/${userID}`);
        });

        client.on("message", (_topic, message) => {
            onMessage(message.toString());
        });

        return () => {
            client.end();
        };
    }, [userID]);
}