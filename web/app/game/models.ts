export type Player = {
    Id: string;
    Name: string;
    Hand: PlayingCard[];
}

export type PlayingCard = {
    Suit: number
    Rank: number
}

export type GameState = {
    Id: string
}