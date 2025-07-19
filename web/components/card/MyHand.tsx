import {PlayingCard} from "@/app/game/models";
import Hand from "@/components/card/Hand";

export type MyHandProps = {
    hand: PlayingCard[];
}

export default function MyHand({hand}: MyHandProps) {
    return <Hand hand={hand} canBeClicked/>
}