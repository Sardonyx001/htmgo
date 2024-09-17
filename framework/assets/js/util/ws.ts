type WsOpts = {
    url: string;
    reconnectInterval?: number;
    onOpen?: () => void;
    onMessage: (message: string) => void;
    onError?: (error: Event) => void;
    onClose?: () => void;
}

export function createWebSocketClient(opts: WsOpts) {
    let socket: WebSocket | null = null;
    const connect = (tries: number) => {
        socket = new WebSocket(opts.url);
        // Handle connection open
        socket.onopen = () => {
        };
        // Handle incoming messages
        socket.onmessage = (event) => {
            opts.onMessage(event.data)
        };
        // Handle connection errors
        socket.onerror = (error) => {
        };
        // Handle connection close and attempt reconnection
        socket.onclose = () => {
            console.log('WebSocket connection closed. Attempting to reconnect...');
            let interval = tries * (opts.reconnectInterval || 50);
            setTimeout(() => connect(tries + 1), interval);
        };
    };
    connect(1);
    const sendMessage = (message: string) => {
        if (socket && socket.readyState === WebSocket.OPEN) {
            socket.send(message);
        } else {
           setTimeout(() => sendMessage(message), 100);
        }
    };
    return { sendMessage };
}