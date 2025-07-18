'use client';

import { Box, Button, Center, Flex, Heading, Input, Stack } from '@chakra-ui/react';
import {useState, useCallback} from 'react';
import { useRouter } from "next/navigation";


type CreateGameResponse = {
    gameId: string;
    playerName: string;
}

export default function Home() {
const [gameId, setGameId] = useState('');
const [playerName, setPlayerName] = useState('');
const router = useRouter();

const handleCreateGame = useCallback(() => {
  if (!playerName) {
    alert('Please enter your name');
    return;
  }
  const ws = new WebSocket(`ws://localhost:8080/ws/create?name=${encodeURIComponent(playerName)}`);
  ws.onopen = () => {
    console.log('WebSocket connection established');
  }
  ws.onmessage = (event) => {
    console.log('Message received:', event.data);
    const response: CreateGameResponse = JSON.parse(event.data);
    if (response.gameId) {
    router.push('/game/' + response.gameId);
    }

  };
  ws.onclose = () => {
    console.log('WebSocket connection closed');
  };
  return () => {
    ws.close();
  };
},[playerName]);

const handleJoinGame = useCallback(() => {
  if (!gameId || !playerName) {
    alert('Please enter both Game ID and your name');
    return;
  }
  const ws = new WebSocket(`ws://localhost:8080/ws/join/${encodeURIComponent(gameId)}'?name=${encodeURIComponent(playerName)}`);
  ws.onopen = () => {
    console.log('WebSocket connection established');
  }
  ws.onmessage = (event) => {
    console.log('Message received:', event.data);
  };
  ws.onclose = () => {
    console.log('WebSocket connection closed');
  };
  return () => {
    ws.close();
  };
}, [gameId, playerName]);


return (
  <Center minH="100vh" >
    <Box p={8} borderRadius="lg" boxShadow="lg"  minW="350px">
      <Stack align="center">
        <Heading size="2xl" mb={4}>Gandalf</Heading>
        <Input
          placeholder="Enter your name"
            value={playerName}
            onChange={e => setPlayerName(e.target.value)}
            mb={4}
          />
        <Button colorScheme="teal" size="lg" width="100%" onClick={handleCreateGame}>
          Create Game
        </Button>
        <Flex width="100%" gap={2}>
          <Input
            placeholder="Game ID"
            value={gameId}
            onChange={e => setGameId(e.target.value)}
            flex="1"
          />
          <Button colorScheme="blue" size="md" onClick={handleJoinGame}>
            Join Game
          </Button>
        </Flex>
      </Stack>
    </Box>
  </Center>
);
}