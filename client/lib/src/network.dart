part of carm;

class Network {
  WebSocket ws;
  connect() {
    ws = new WebSocket('ws://127.0.0.1:9876/ws');
  }

  lock() {
    ws.send({"command":"lock"});
  }

  unlock() {
    ws.send({"command": "unlock"});
  }
}

