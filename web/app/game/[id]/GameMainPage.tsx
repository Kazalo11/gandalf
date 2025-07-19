import {Player} from "@/app/game/models";
import {HStack} from "@chakra-ui/react";
import PlayingCardImg from "@/components/PlayingCardImg";

export type GameMainPageProps = {
    gameId: string;
    player: Player;
}
export default function GameMainPage({gameId, player}: GameMainPageProps) {
    return (
        <div>
            <h1>Game Page</h1>
            <p>Game ID: {gameId}</p>
            <p>Player ID: {player.Id}</p>
            <p>Player Name: {player.Name}</p>
            <p>
                Player Hand:{" "}
                <HStack>
                    {
                        player.Hand.map((card, key) => {
                            return (
                                <PlayingCardImg key={key} card={card} isHidden={true}/>
                            )
                        })
                    }
                </HStack>

            </p>
            <p>WebSocket connection will be established here.</p>
        </div>
    );

}