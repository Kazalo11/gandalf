import {PlayingCard} from "@/app/game/models";
import Hand from "@/components/card/Hand";
import {Box, Heading} from "@chakra-ui/react";

export type OtherPlayerHandProps = {
    hand: PlayingCard[];
    name: string
    position?: "left" | "right" | "top";
}

export default function OtherPlayerHand({hand, name, position}: OtherPlayerHandProps) {
    let style = {};
    switch (position) {
        case "left":
            style = {position: "absolute", left: "10%", top: "50%", transform: "translateY(-50%)"};
            break;
        case "right":
            style = {position: "absolute", right: "10%", top: "50%", transform: "translateY(-50%)"};
            break;
        case "top":
            style = {position: "absolute", top: "10%", left: "50%", transform: "translateX(-50%)"};
            break;
        default:
            style = {position: "absolute", bottom: "20%", left: "40%"};
    }
    return (
        <Box {...style}>
            <Heading>{name}</Heading>
            <Hand hand={hand} canBeClicked={false} />
        </Box>
    )
}