'use client';

import { use, useEffect, useState } from "react";
import Link from "next/link";
import {Player} from "@/app/game/models";
import PlayingCardImg from "@/components/PlayingCardImg";
import {HStack} from "@chakra-ui/react";


export default function GamePage({
                                     params: paramsPromise,
                                 }: {
    params: Promise<{ id: string }>;
}) {
    const { id } = use(paramsPromise);
    const [playerData, setPlayerData] = useState<Player | null>(null);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const playerId = localStorage.getItem("playerId");
        if (!playerId) {
            setError("no-player-id");
            setLoading(false);
            return;
        }

        fetch(`http://localhost:8080/game/${id}/player/${playerId}`)
            .then((res) => {
                if (!res.ok) throw new Error("fetch-failed");
                return res.json();
            })
            .then((data: Player) => {
                setPlayerData(data);
            })
            .catch(() => {
                setError("fetch-failed");
            })
            .finally(() => setLoading(false));
    }, [id]);

    if (loading) return <p>Loading...</p>;

    if (error === "no-player-id") {
        return (
            <div>
                <h1>Error</h1>
                <p>You must have a player ID to view this page.</p>
                <p>
                    Please create or join a game first by going to <Link href="/">Home</Link>.
                </p>
            </div>
        );
    }

    if (error === "fetch-failed") {
        return (
            <div>
                <h1>Error</h1>
                <p>Failed to fetch game data. Please check the Game ID and try again.</p>
                <p>
                    Go back to <Link href="/">Home</Link>.
                </p>
            </div>
        );
    }

    return playerData ? (
        <div>
            <h1>Game Page</h1>
            <p>Game ID: {id}</p>
            <p>Player ID: {playerData.Id}</p>
            <p>Player Name: {playerData.Name}</p>
            <p>
                Player Hand:{" "}
                <HStack>
                    {
                        playerData.Hand.map((card, key) => {
                            return (
                                <PlayingCardImg key={key} card={card} isHidden={true}/>
                            )
                        })
                    }
                </HStack>

            </p>
            <p>WebSocket connection will be established here.</p>
        </div>
    ): (
        <div>
            <h1>Error</h1>
            <p>No player data found for this game.</p>
            <p>
                Go back to <Link href="/">Home</Link>.
            </p>
        </div>
    );
}
