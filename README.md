On Air
======

The On Air application is a small utility application written in Go that monitors a [Twitch.tv](http://twitch.tv) broadcaster's status that runs on a Raspberry Pi.
When an ```online``` status is detected the RPi's GPIO pin 17 is set low, when the broadcaster is offline pin 17 is set high.  This can be used in conjunction with
a relay to switch on and off physical devices like an On Air lamp, a camera or other devices.

Setup
=====

Because this application is intended to run on a RPi you'll need to [compile a version of Go for the ARM architecture](http://www.maketecheasier.com/build-go-from-source-on-raspberry-pi/).

Dependencies
============

On Air leverages the [gpio](http://github.com/davecheney/gpio) library to communicate with the RPi's GPIO.
