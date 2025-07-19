import {PlayingCard} from "@/app/game/models";
import {mapCardToImagePath} from "@/app/game/mapper/mapCardToImagePath";
import Image from 'next/image'

export type PlayingCardProps = {
    card: PlayingCard
    isHidden?: boolean
}

export default function PlayingCardImg({card, isHidden}: PlayingCardProps) {
    const imgSrc = mapCardToImagePath(card);
    const handWidth = 80;
    const handHeight = 120;

    return isHidden ?  (
        <Image
            src={"/cards/back.svg"}
            alt={"Back of card"}
            width={handWidth}
            height={handHeight}
        />
    ): (
        <Image
            src={imgSrc}
            alt={`${card.Rank} of ${card.Suit}`}
            width={handWidth}
            height={handHeight}
        />
    )

}