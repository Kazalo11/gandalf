import {PlayingCard} from "@/app/game/models";
import Hand from "@/components/card/Hand";

export type OtherPlayerHandProps = {
    hand: PlayingCard[];
}

export default function OtherPlayerHand({hand}: OtherPlayerHandProps) {
    return <Hand hand={hand} canBeClicked={false} />
}