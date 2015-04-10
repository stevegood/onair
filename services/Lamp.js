var monitor = require('./StreamMonitor'), gpio;

try {
  gpio = require('pi-gpio');
} catch (e) {
  console.error(e);
}

function Lamp() {

  var online = false,
      pins;

  function setPinsTo(value, callback) {
    if (gpio) {
      var i = 0;

      function setPins(pin) {
        console.log("Setting pin " + pin + ' to ' + value);
        gpio.open(pin, "output", function(err) {
          gpio.write(pin, value, function() {
            gpio.close(pin, function(){
              i++;
              if (i === pins.length) {
                if (callback) callback();
              } else {
                setPins(parseInt(pins[i].toString()));
              }
            });
          });
        });
      }

      setPins(parseInt(pins[i].toString()));
    } else {
      if (callback) callback();
    }
  }

  function onlineHandler() {
    // turn the lamp on
    if (!online) {
      online = true;
      console.log('Turning the lamp ON');
      setPinsTo(1);
    }
  }

  function offlineHandler() {
    // turn the lamp off
    if (online) {
      online = false;
      console.log('Turning the lamp OFF');
      setPinsTo(0);
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

      if (streamName && _interval) {
        monitor.watch(streamName, (_interval * 1000), onlineHandler, offlineHandler);
      }
    }
  };
}

module.exports = new Lamp();
