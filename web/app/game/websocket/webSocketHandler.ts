import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";
import {GameState} from "@/app/game/models";

export type ServerMessage = {
    type: string;
    subType: string;
};

export type GameServerMessage = ServerMessage & {
    type: "GameMessage";
}

export type JoinGameMessage = GameServerMessage & {
    subType: "JoinGame";
    gameId: string;
    playerId: string;
};

export type GetGameStateMessage = GameServerMessage & {
    subType: "GetGame";
    gameId: string;
};

export type GameStateMessage = GameServerMessage & {
    game: GameState

}

export type ReceivedMessages = JoinGameMessage | GameStateMessage;

export type HandlerContext = {
    router: AppRouterInstance;
    setGameState: (gameState: GameState) => void;
    gameState?: GameState;
    listeners?: ((msg: never) => void)[];
};

export type WebSocketHandler<T extends ServerMessage = ServerMessage> = (
    message: T,
    context: HandlerContext
) => void;


export type WebSocketHandlerMap = {
    "GameMessage:JoinGame": WebSocketHandler<JoinGameMessage>;
    "GameMessage:GameState": WebSocketHandler<GameStateMessage>;
};

export const webSocketHandlerMap: WebSocketHandlerMap = {
    "GameMessage:JoinGame": (message: JoinGameMessage, context: HandlerContext) => {
        const { router } = context;
        localStorage.setItem("playerId", message.playerId);
        localStorage.setItem("gameId", message.gameId);
        router.push("/game/" + message.gameId);
    },

    "GameMessage:GameState": (message: GameStateMessage, context: HandlerContext) => {
        console.log("GameStateMessage received:", message);
        const { setGameState } = context;
        if (!message.game) {
            console.error("GameStateMessage is missing game data");
            return;
        }
        if (message.game === context.gameState) {
            console.log("GameStateMessage is the same as current game state, skipping update");
            return;
        }
        setGameState(message.game)
    }
};
