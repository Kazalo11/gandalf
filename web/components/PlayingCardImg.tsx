import {PlayingCard} from "@/app/game/models";
import {mapCardToImagePath} from "@/app/game/mapper/mapCardToImagePath";
import Image from 'next/image'

export type PlayingCardProps = {
    card: PlayingCard
    isHidden?: boolean
}

export default function PlayingCardImg({card, isHidden}: PlayingCardProps) {
    const imgSrc = mapCardToImagePath(card);

    return isHidden ?  (
        <Image
            src={"/cards/back.svg"}
            alt={"Back of card"}
            width={100}
            height={140}
            className="playing-card"
            draggable="true"
        />
    ): (
        <Image
            src={imgSrc}
            alt={`${card.Rank} of ${card.Suit}`}
            width={100}
            height={140}
            className="playing-card"
            draggable="true"
        />
    )

}