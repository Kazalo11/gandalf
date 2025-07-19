import Link from "next/link";

export default function NoPlayerId() {
    return (
        <div>
            <h1>Error</h1>
            <p>You must have a player ID to view this page.</p>
            <p>
                Please create or join a game first by going to <Link href="/">Home</Link>.
            </p>
        </div>
    );
}