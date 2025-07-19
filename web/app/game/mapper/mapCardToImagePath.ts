import {PlayingCard} from "@/app/game/models";

const suitMap: Record<number, string> = {
    0: "spades",
    1: "clubs",
    2: "diamonds",
    3: "hearts",
}

const faceMap: Record<number, string> = {
    1: "ace",
    11: "jack",
    12: "queen",
    13: "king",
}

export function mapCardToImagePath(card: PlayingCard): string {
    const suit = suitMap[card.Suit];
    const rank = faceMap[card.Rank] || card.Rank.toString();

    if (!suit || !rank) {
        throw new Error(`Invalid card: Suit ${card.Suit}, Rank ${card.Rank}`);
    }

    return `/cards/${suit}_${rank}.svg`;

}