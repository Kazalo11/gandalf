import {GameState, Player, PlayingCard} from "@/app/game/models";
import {Box, BoxProps} from "@chakra-ui/react";
import Deck from "@/components/deck/Deck";

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
export default function DisplayPage({game, currentPlayer}: DisplayPageProps) {
    const otherPlayers: Player[] = Object.values(game.players).filter(
        (player) => player.Id !== currentPlayer.Id
    );

    return (
        <Box {...boxStyles}>
            <Deck deck={game.deck} />
        </Box>
    )
}