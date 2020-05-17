const socket = new WebSocket('https://7b4qjmnofj.execute-api.us-east-2.amazonaws.com/development');

socket.onopen = (event) => {
    console.log('Websocket Connected.');
};

socket.onmessage = (event) => {
    console.log('Message Received');
    console.log(event);
    const messages = document.getElementById('messages');
    const messageData = JSON.parse(event.data);
    messages.innerHTML = messages.innerHTML + `<p><strong>${messageData.connectionId}</strong>: ${messageData.message}</p>`;
};

socket.onerror = (event) => {
    console.error(event);
};

function sendMessage() {
    const input = document.getElementById('message-input');
    const message = input.value.trim();
    const requestPayload = JSON.stringify({
        action: 'message',
        message
    });

    if (message) {
        console.log('Message Sent.');
        socket.send(requestPayload);
    }
};