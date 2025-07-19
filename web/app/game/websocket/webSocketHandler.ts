import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";

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

export type KnownMessages = JoinGameMessage | GetGameStateMessage;

export type HandlerContext = {
    router: AppRouterInstance;
    listeners?: ((msg: any) => void)[];
};

export type WebSocketHandler<T extends ServerMessage = ServerMessage> = (
    message: T,
    context: HandlerContext
) => void;


export type WebSocketHandlerMap = {
    "GameMessage:JoinGame": WebSocketHandler<JoinGameMessage>;
    "GameMessage:GetGame": WebSocketHandler<GetGameStateMessage>;
};

export const webSocketHandlerMap: WebSocketHandlerMap = {
    "GameMessage:JoinGame": (message: JoinGameMessage, context: HandlerContext) => {
        const { router } = context;
        localStorage.setItem("playerId", message.playerId);
        localStorage.setItem("gameId", message.gameId);
        router.push("/game/" + message.gameId);
    },

    "GameMessage:GetGame": (message: GetGameStateMessage, context: HandlerContext) => {
        console.log("Get game state for gameId:", message.gameId);
    }
};
