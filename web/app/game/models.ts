export type Player = {
    Id: string;
    Name: string;
    Hand: PlayingCard[];
}

export type PlayingCard = {
    Suit: number
    Rank: number
}

export type Round = {
    turn: number
    isGandalf: boolean
}

export type GameState = {
    id: string
    players: Record<string, Player>
    deck: PlayingCard[] | null
    discard: PlayingCard[] | null
    rounds: Round[] | null

}