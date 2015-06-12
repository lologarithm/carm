part of carm;

class ComputerAlarm {
  Network network;
  ComputerAlarm() {
    window.addEventListener('blur', handleBrowserBlur);
    network = new Network();
  }

  handleBrowserBlur(event) {
    // DO SHENANIGANS HERE
    network.lock();
  }
}