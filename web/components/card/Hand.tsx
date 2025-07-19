import {PlayingCard} from "@/app/game/models";
import {HStack} from "@chakra-ui/react";
import PlayingCardImg from "@/components/card/PlayingCardImg";

export type HandProps = {
    hand: PlayingCard[];
    canBeClicked?: boolean;
}

export default function Hand({hand, canBeClicked}: HandProps) {
    return (
        <HStack>
            {
                hand.map((card, key) => {
                    return (
                        <PlayingCardImg key={key} card={card} isHidden/>
                    )
                })
            }
        </HStack>
    )

}