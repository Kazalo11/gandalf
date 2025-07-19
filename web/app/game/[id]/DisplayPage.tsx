import {GameState, Player, PlayingCard} from "@/app/game/models";
import {Box, BoxProps} from "@chakra-ui/react";
import Deck from "@/components/deck/Deck";
import MyHand from "@/components/card/MyHand";
import OtherPlayerHand from "@/components/card/OtherPlayerHand";

export type DisplayPageProps = {
    game: GameState
    currentPlayer: Player
}
const boxStyles: BoxProps = {
    w: "900px",
    h: "800px",
    bgImage: "url(/deck/deck.svg)",
    bgSize: "contain",
    bgRepeat: "no-repeat",
    display: "flex",
}
const positions = ["left", "right", "top"] as const;
export default function DisplayPage({game, currentPlayer}: DisplayPageProps) {
    const otherPlayers: Player[] = Object.values(game.players).filter(
        (player) => player.Id !== currentPlayer.Id
    );

    return (
        <Box {...boxStyles}>
            <Deck deck={game.deck} />
            <MyHand hand={currentPlayer.Hand} name={currentPlayer.Name} />
            {otherPlayers.map((player, index) => {
                const position = positions[index % positions.length];
                return (
                    <OtherPlayerHand hand={player.Hand} name={player.Name} position={position} key={player.Id} />
                );
            })}


        </Box>
    )
}