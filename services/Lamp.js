var monitor = require('./StreamMonitor'), gpio;

try {
  gpio = require('pi-gpio');
} catch (e) {
  console.error(e);
}

function Lamp() {

  var online = false,
      pins;

  function setPinsTo(value) {
    if (gpio) {
      for(var i=0; i < pins.length; i++) {
        var pin = parseInt(pins[i].toString());
        console.log(!isNaN(pin));
        gpio.close(pin, function(){
          gpio.open(pin, "output", function(err) {
            gpio.write(pin, value || 1, function() {
              gpio.close(pin);
            });
          });
        });
      }
    }
  }

  function onlineHandler() {
    // turn the lamp on
    if (!online) {
      online = true;
      console.log('Turning the lamp ON');
      setPinsTo(0);
    }
  }

  function offlineHandler() {
    // turn the lamp off
    if (online) {
      online = false;
      console.log('Turning the lamp OFF');
      setPinsTo(1);
    }
  }

  return {
    init: function() {
      var streamName = process.env.STREAM_NAME,
          _interval = process.env.INTERVAL || 30;

      pins = process.env.PINS.split(',');

      console.log('Starting up the lamp!');
      console.log('Watching for status changes on ' + streamName + ' every ' + _interval + ' seconds');;
      console.log('Will manage pins ' + pins);
      console.log('Setting pins high to start');

      setPinsTo(1);

      if (streamName && _interval) {
        monitor.watch(streamName, (_interval * 1000), onlineHandler, offlineHandler);
      }
    }
  };
}

module.exports = new Lamp();
