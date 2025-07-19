import Link from "next/link";

export default function NoGameData() {
    return (
        <div>
            <h1>Error</h1>
            <p>No game data available. Please check the Game ID and try again.</p>
            <p>
                Go back to <Link href="/">Home</Link>.
            </p>
        </div>
    );
}