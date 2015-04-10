var superagent = require('superagent');

function Twitch() {

  var ENDPOINT = "https://api.twitch.tv/kraken";

  function Streams() {
    return {
      getStream: function(streamName, callback) {
        superagent.get(ENDPOINT + '/streams/' + streamName)
          .end(function(e, res){
            callback(e, res);
          });
      }
    };
  }

  return {
    streams: new Streams()
  };
}

module.exports = new Twitch();
