import {Player} from "@/app/game/models";
import {Box, BoxProps, Heading, VStack} from "@chakra-ui/react";
import Deck from "@/components/deck/Deck";
import MyHand from "@/components/card/MyHand";
import {useWebSocket} from "@/app/game/websocket/WebSocketProvider";
import {GetGameStateMessage} from "@/app/game/websocket/webSocketHandler";
import {useEffect, useRef} from "react";

export type GameMainPageProps = {
    gameId: string;
    player: Player;
}

const boxStyles: BoxProps = {
    w: "900px",
    h: "800px",
    bgImage: "url(/deck/deck.svg)",
    bgSize: "contain",
    bgRepeat: "no-repeat",
    display: "flex",
}
export default function GameMainPage({gameId, player}: GameMainPageProps) {
    const {sendMessage} = useWebSocket();

    useEffect(() => {
        const getGameStateMessage: GetGameStateMessage = {
            gameId,
            type: "GameMessage",
            subType: "GetGame",
        };

        console.log("Sending GetGameStateMessage:", getGameStateMessage);
        sendMessage(JSON.stringify(getGameStateMessage));
    }, [gameId, sendMessage]);

    const playerHand = player.Hand;
    return (

        <VStack>
            <Heading size={"4xl"}>Gandalf</Heading>
            <Box {...boxStyles} >
                <Deck />
                <MyHand hand={playerHand}/>
            </Box>
        </VStack>
    );

}