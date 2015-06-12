library carm;

import 'dart:html';
part 'src/app.dart';


class Network {
  connect() {
	var ws = new WebSocket('ws://127.0.0.1:9876/ws');
	ws.send("{}");
  }
}

// {"command": "lock"} or {"command": "unlock"}
