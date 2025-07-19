import {PlayingCard} from "@/app/game/models";
import Hand from "@/components/card/Hand";
import {Box, Heading} from "@chakra-ui/react";

export type MyHandProps = {
    hand: PlayingCard[];
    name: string
}

export default function MyHand({hand, name}: MyHandProps) {
    return (
        <Box position="absolute" bottom="20%" left="40%" width="100%">
            <Heading>{name}</Heading>
    <Hand hand={hand} canBeClicked/>
        </Box>
    )
}