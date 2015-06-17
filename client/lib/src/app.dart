part of carm;

class ComputerAlarm {
  Network network;
  ComputerAlarm() {
    window.addEventListener('blur', handleBrowserBlur);
    document.addEventListener('keypress', handleKeypress);
    network = new Network();
  }

  handleBrowserBlur(event) {
    // DO SHENANIGANS HERE
    network.lock();
  }

  handleKeypress(KeyboardEvent event) {
    if (event.charCode == 11) { // ctrl +  K
      network.disarm();
      showDisarmPage();
    }
  }

  showDisarmPage() {
    DivElement container = querySelector(".container");

    container.innerHtml = "Disarmed";
  }
}