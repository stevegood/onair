package twitch

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)

const (
    ENDPOINT string = "https://api.twitch.tv/kraken"
)

type Twitch struct {
    ClientId string
}

type Preview struct {
    Small string
    Medium string
    Large string
    Template string
}

type ChannelLinks struct {
    Self string
    Follows string
    Commercial string
    StreamKey string `json:"stream_key"`
    Chat string
    Features string
    Subscriptions string
    Editors string
    Videos string
    Teams string
}

type TwitchChannel struct {
    Background string
    Banner string
    DisplayName string `json:"display_name"`
    Game string
    Logo string
    Mature bool
    Status string
    Url string
    Delay int
    Followers int
    Name string
    Views int
}

type StreamLinks struct {
    Self string
}

type TwitchStream struct {
    Links StreamLinks `json:"_links"`
    Game string
    Viewers int
    Channel TwitchChannel
    Preview Preview
}

type TwitchStreamResponse struct {
    Stream TwitchStream
}

func (t *Twitch) GetStream(StreamName string) *TwitchStream {
    resp, err := http.Get(t.GetUrl("/streams/"+StreamName))
    if err != nil {
        panic(err)
    }

    // do stuff with the data, like unmarshal it into stream
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    var StreamResponse TwitchStreamResponse
    err = json.Unmarshal(body, &StreamResponse)
    if err != nil {
        panic(err)
    }

    return &StreamResponse.Stream
}

func (t *Twitch) GetUrl(uri string) string {
    var url string = ENDPOINT+uri
    if t.ClientId != "" {
        url = url+"?client_id="+t.ClientId
    }
    return url
}

func NewTwitch() *Twitch {
    return &Twitch{ClientId: DefaultClientId()}
}

func DefaultClientId() string {
    return ""
}
