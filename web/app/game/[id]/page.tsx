'use client';

import {use, useEffect, useState} from "react";
import {Player} from "@/app/game/models";
import NoPlayerId from "@/app/game/error/NoPlayerId";
import NoGameData from "@/app/game/error/NoGameData";
import NoPlayerData from "@/app/game/error/NoPlayerData";
import GameMainPage from "@/app/game/[id]/GameMainPage";


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

    if (error === "no-player-id") return <NoPlayerId/>;
    if (error === "fetch-failed") return <NoGameData/>;

    return playerData ? <GameMainPage gameId={id} player={playerData}/> : <NoPlayerData/>;
}
