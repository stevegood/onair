On Air
======

The On Air application is a small utility application written in Go that monitors a [Twitch.tv](http://twitch.tv) broadcaster's status that runs on a Raspberry Pi.
When an ```online``` status is detected the RPi's GPIO pins specified in ```Config.json``` are set low, when the broadcaster is ```offline``` the pins are set high.  This can be used in conjunction with
a relay to switch on and off physical devices like an On Air lamp, a camera or other devices.

Setup
=====

Because this application is intended to run on a RPi you'll need to [compile a version of Go for the ARM architecture](http://www.maketecheasier.com/build-go-from-source-on-raspberry-pi/).

Dependencies
============

None!  Everything you need is in this repo.
