import PlayingCardImg from "@/components/card/PlayingCardImg";
import {Box} from "@chakra-ui/react";
import {PlayingCard} from "@/app/game/models";

export type DeckProps = {
    deck: PlayingCard[];
}

export default function Deck({deck}: DeckProps) {
    return (
        <Box
            position="relative"
            top="46%"
            left="50%"
            transform="translate(-50%, -50%)"
            width="80px"    // Add width
            height="120px"  // Add height
        >
            {deck.map((card, index) => (
                <Box
                    key={index}
                    position="absolute"
                    top={0}
                    left={0}
                    style={{
                        transform: `translate(${index * 0.5}px, ${index * 0.5}px)`,
                        zIndex: index,
                    }}
                >
                    <PlayingCardImg card={card} isHidden />
                </Box>
            ))}
        </Box>


    )
}