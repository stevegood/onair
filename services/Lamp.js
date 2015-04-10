var monitor = require('./StreamMonitor');

function Lamp() {

  var online = false;

  function onlineHandler() {
    // turn the lamp on
    if (!online) {
      online = true;
      console.log('Turning the lamp ON');
    }
  }

  function offlineHandler() {
    // turn the lamp off
    if (online) {
      online = false;
      console.log('Turning the lamp OFF');
    }
  }

  return {
    init: function() {
      var streamName = process.env.STREAM_NAME,
          _interval = process.env.INTERVAL || 30;

      console.log('Starting up the lamp!');
      console.log('Watching for status changes on ' + streamName + ' every ' + _interval + ' seconds');;

      if (streamName && _interval) {
        monitor.watch(streamName, (_interval * 1000), onlineHandler, offlineHandler);
      }
    }
  };
}

module.exports = new Lamp();
