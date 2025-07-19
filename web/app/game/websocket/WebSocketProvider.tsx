'use client';

import { createContext, useContext, useRef, useState, useCallback } from 'react';
import { useRouter } from 'next/navigation';
import {
    ReceivedMessages,
    WebSocketHandler,
    webSocketHandlerMap
} from "@/app/game/websocket/webSocketHandler";
import {GameState} from "@/app/game/models";



type WebSocketContextType = {
    createGame: (playerName: string) => void;
    joinGame: (gameId: string, playerName: string) => void;
    sendMessage: (message: string) => void;
    socket: WebSocket | null;
    reconnect: () => void;
    addMessageListener: (listener: (msg: never) => void) => void;
    removeMessageListener: (listener: (msg: never) => void) => void;
    gameState?: GameState

};

const WebSocketContext = createContext<WebSocketContextType | undefined>(undefined);

export const WebSocketProvider = ({ children }: { children: React.ReactNode }) => {
    const [socket, setSocket] = useState<WebSocket | null>(null);
    const socketRef = useRef<WebSocket | null>(null);
    const listenersRef = useRef<((msg: never) => void)[]>([]);
    const router = useRouter();
    const [gameState, setGameState] = useState<GameState | undefined>(undefined);

    const setupSocketHandlers = useCallback((ws: WebSocket) => {
        ws.onopen = () => {
            console.log('WebSocket connected');
        };

        ws.onmessage = (event) => {
            try {
            const message: ReceivedMessages = JSON.parse(event.data);
            if (!message || !message.type || !message.subType) {
                console.error('Invalid message format:', event.data);
                return;
            }
            const key = `${message.type}:${message.subType}`;
            console.log('Received WebSocket message:', key, message);

            // @ts-expect-error - TypeScript doesn't know about the dynamic keys in webSocketHandlerMap
            const handler: WebSocketHandler = webSocketHandlerMap[key];
            handler(message, { router, listeners: listenersRef.current, setGameState, gameState });

            } catch (err) {
                console.error('Failed to parse WebSocket message', err);
            }
        };

        ws.onclose = () => {
            console.log('WebSocket disconnected');
        };
    },[gameState, router]);

    const createGame = useCallback((playerName: string) => {
        const ws = new WebSocket(`ws://localhost:8080/ws/create?name=${encodeURIComponent(playerName)}`);
        setupSocketHandlers(ws);
        socketRef.current = ws;
        setSocket(ws);
    }, [setupSocketHandlers]);

    const joinGame = useCallback((gameId: string, playerName: string) => {
        const ws = new WebSocket(`ws://localhost:8080/ws/game/${encodeURIComponent(gameId)}/join?name=${encodeURIComponent(playerName)}`);
        setupSocketHandlers(ws);
        socketRef.current = ws;
        setSocket(ws);
    }, [setupSocketHandlers]);

    const sendMessage = useCallback((message: string) => {
        if ( socketRef && socketRef.current?.readyState === WebSocket.OPEN) {
            socketRef.current.send(message);
        } else {
            console.warn('WebSocket not connected');
        }
    }, []);


    const addMessageListener = useCallback((listener: (msg: never) => void) => {
        listenersRef.current.push(listener);
    }, []);

    const removeMessageListener = useCallback((listener: (msg: never) => void) => {
        listenersRef.current = listenersRef.current.filter(l => l !== listener);
    }, []);

    const reconnect = useCallback(() => {
        const playerId = sessionStorage.getItem('playerId');
        const gameId = sessionStorage.getItem('gameId');
        if (!playerId || !gameId) {
            console.warn('Missing playerName or gameId for reconnection');
            return;
        }

        const ws = new WebSocket(`ws://localhost:8080/ws/game/${encodeURIComponent(gameId)}/reconnect?playerId=${encodeURIComponent(playerId)}`);
        setupSocketHandlers(ws);
        socketRef.current = ws;
        setSocket(ws);
    }, [ setupSocketHandlers]);



    return (
        <WebSocketContext.Provider value={{ createGame, joinGame, sendMessage, socket, reconnect, addMessageListener, removeMessageListener, gameState }}>
            {children}
        </WebSocketContext.Provider>
    );
};

export const useWebSocket = (): WebSocketContextType => {
    const context = useContext(WebSocketContext);
    if (!context) {
        throw new Error('useWebSocket must be used within a WebSocketProvider');
    }
    return context;
};
