part of carm;

class Network {
  WebSocket ws;
  Network() {
    ws = new WebSocket('ws://127.0.0.1:9876/ws');
  }

  lock() {
    ws.send(JSON.encode({"command":"lock"}));
  }

  unlock() {
    ws.send(JSON.encode({"command": "unlock"}));
  }
}

