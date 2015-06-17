part of carm;

class Network {
  WebSocket ws;
  Timer pinger;

  Network() {
    ws = new WebSocket('ws://127.0.0.1:9876/ws');
    ws.onOpen.listen(handleOpened);
  }

  handleOpened(_) {
	  sendPing();
  }

  sendPing() {
	  ws.send(JSON.encode({"command":"ping"}));
	  pinger = new Timer(new Duration(milliseconds: 250), sendPing);
  }

  lock() {
    ws.send(JSON.encode({"command":"lock"}));
  }

  disarm() {
    ws.send(JSON.encode({"command": "disarm"}));
  }
}
