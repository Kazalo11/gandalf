export default async function GamePage({
  params,
}: Readonly<{
  params: { id: string };
}>) {
  const { id } = await params;
    return (
        <div>
        <h1>Game Page</h1>
        <p>Game ID: {id}</p>
        <p>WebSocket connection will be established here.</p>
        </div>
    );


}