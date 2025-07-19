import PlayingCardImg from "@/components/card/PlayingCardImg";
import {Box} from "@chakra-ui/react";



export default function Deck() {
    return (
        <Box
            position="absolute"
            top="46%"
            left="50%"
            transform="translate(-50%, -50%)"
        >
            <PlayingCardImg card={{
                Suit: 0, // Spades
                Rank: 1 // Ace
            }} isHidden/>
        </Box>

    )
}