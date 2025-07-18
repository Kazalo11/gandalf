export type Player = {
    Id: string;
    Name: string;
    Hand: Card[];
}

type Card = {
    Suit: number
    Rank: number
}