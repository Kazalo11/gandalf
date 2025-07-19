'use client';

import { Box, Button, Center, Flex, Heading, Input, Stack } from '@chakra-ui/react';
import { useState, useEffect, useCallback } from 'react';
import { useRouter } from 'next/navigation';
import {useWebSocket} from "@/app/game/WebSocketProvider";



export default function Home() {
  const [gameId, setGameId] = useState('');
  const [playerName, setPlayerName] = useState('');
  const router = useRouter();
  const { createGame, joinGame, addMessageListener, removeMessageListener } = useWebSocket();

  const handleCreateGame = useCallback(() => {
    if (!playerName) {
      alert('Please enter your name');
      return;
    }
    createGame(playerName);
  }, [playerName, createGame]);

  const handleJoinGame = useCallback(() => {
    if (!gameId || !playerName) {
      alert('Please enter both Game ID and your name');
      return;
    }
    joinGame(gameId, playerName);
  }, [gameId, playerName, joinGame]);

  useEffect(() => {
    const handleMessage = (data: any) => {
      if (data?.gameId && data?.playerId) {
        localStorage.setItem('playerId', data.playerId);
        router.push('/game/' + data.gameId);
      }
    };

    addMessageListener(handleMessage);
    return () => {
      removeMessageListener(handleMessage);
    };
  }, [addMessageListener, removeMessageListener, router]);

  return (
      <Center minH="100vh">
        <Box p={8} borderRadius="lg" boxShadow="lg" minW="350px">
          <Stack align="center">
            <Heading size="2xl" mb={4}>Gandalf</Heading>
            <Input
                placeholder="Enter your name"
                value={playerName}
                onChange={(e) => setPlayerName(e.target.value)}
                mb={4}
            />
            <Button colorScheme="teal" size="lg" width="100%" onClick={handleCreateGame}>
              Create Game
            </Button>
            <Flex width="100%" gap={2}>
              <Input
                  placeholder="Game ID"
                  value={gameId}
                  onChange={(e) => setGameId(e.target.value)}
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
