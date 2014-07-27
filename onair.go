package main

import(
    "encoding/json"
    "log"
    "io/ioutil"
    "github.com/stevegood/onair/pi"
    "github.com/stevegood/onair/twitch"
    "time"
)

var (
    _twitch *twitch.Twitch
    config *Config
    pins []*pi.Pin
)

type Config struct {
    Username string
    Frequency int
    ClientId string `json:"client_id"`
    Pins []int
}

func GetStream(username string) {
    stream :=  _twitch.GetStream(username)
    if stream.Channel.DisplayName == "" {
        log.Print(username+" is OFFLINE\n")
        SetAllPinsTo("1")
    } else {
        log.Print(stream.Channel.DisplayName+" is ONLINE\n")
        SetAllPinsTo("0")
    }

    defer GetStream(username)
    if config.Frequency > 0 {
        delay := time.Duration(config.Frequency) * time.Second
        time.Sleep(delay)
    }

}

func SetAllPinsTo(Value string) {
    for _, pin := range pins {
        if Value == "0" {
            pin.SetLow()
        } else {
            pin.SetHigh()
        }
    }
}

func LoadConfig() {
    file, err := ioutil.ReadFile("Config.json")
    if err != nil {
        log.Print("Unable to load Config.json, using defaults instead")
        config = &Config{Username: "ThatArdothGuy", Frequency: 60, Pins: []int{14,15,18,23}}
    } else {
        err = json.Unmarshal(file, &config)
        if err != nil {
            panic(err)
        }
    }
}

func InitPins() {
    for _, pin := range config.Pins {
        log.Printf("Initializing pin %d", pin)
        p := pi.NewPin(pin, "out", 1)
        pins = append(pins, p)
    }
}

func main() {
    log.Print("Hello On Air!\n")
    LoadConfig()
    _twitch = twitch.NewTwitch()

    if config.ClientId == "" {
        _twitch.ClientId = config.ClientId
    }
    InitPins()
    GetStream(config.Username)
}
