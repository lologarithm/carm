part of carm;

class ComputerAlarm {
  Network network;
  bool armed = true;

  ComputerAlarm() {
    window.addEventListener('blur', handleBrowserBlur);
    document.addEventListener('keypress', handleKeypress);
    network = new Network();
  }

  handleBrowserBlur(event) {
    if (armed) {
      // DO SHENANIGANS HERE
      network.lock();
    }
  }

  handleKeypress(KeyboardEvent event) {
    if (event.charCode == 107) {
      //  K
      armed = false;
      network.disarm();
      showDisarmPage();
    } else if (event.charCode == 114) {
      // R
      armed = true;
      network.sendPing();
      showArmedPage();
    }
  }

  showDisarmPage() {
    DivElement container = querySelector(".container");
    container.innerHtml = "Disarmed";
  }

  showArmedPage() {
    DivElement container = querySelector(".container");
    container.innerHtml = """
        <img src="./images/google.png" style="position: fixed; top: 20%; left: 50%;  transform: translate(-50%, -20%)" />
    """;
  }
}
