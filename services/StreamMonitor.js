var twitch = require('./Twitch');

function StreamMonitor() {

  var doWatch;

  function watch(streamName, _interval, online, offline) {
    doWatch = true;

    function checkStatus() {
      twitch.streams.getStream(streamName, function(e, res){
        if (e) {
          console.error(e);
        } else {
          if (res.body.stream !== null) {
            online();
          } else {
            offline();
          }
        }

        if (doWatch) {
          setTimeout(checkStatus, _interval);
        }
      });
    }

    checkStatus();
  }

  function stop() {
    doWatch = false;
  }

  return {
    watch: watch
  };
}

module.exports = new StreamMonitor();
