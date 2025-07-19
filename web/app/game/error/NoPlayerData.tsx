import Link from "next/link";

export default function NoPlayerData() {
    return (
        <div>
            <h1>Error</h1>
            <p>No player data available. Please ensure you have joined a game.</p>
            <p>
                Go back to <Link href="/">Home</Link>.
            </p>
        </div>
    );
}