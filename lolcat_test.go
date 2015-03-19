package main

import (
	"bytes"
	"testing"
)

var bigData = []byte(`
[0;34;10m2015-03-18 16:33:02 -0300 [app][0b4bd652bc85]:[0m  * Running on http://0.0.0.0:8888/ (Press CTRL+C to quit)
[0;34;10m2015-03-18 16:33:02 -0300 [app][a29e3bcbdfc3]:[0m  * Running on http://0.0.0.0:8888/ (Press CTRL+C to quit)
[0;34;10m2015-03-18 16:33:04 -0300 [app][2ee9e0fd53f2]:[0m  * Running on http://0.0.0.0:8888/ (Press CTRL+C to quit)
[0;34;10m2015-03-18 16:33:04 -0300 [app][486ba06b3c50]:[0m  * Running on http://0.0.0.0:8888/ (Press CTRL+C to quit)
[0;34;10m2015-03-18 16:33:04 -0300 [app][0cf9527e1764]:[0m  * Running on http://0.0.0.0:8888/ (Press CTRL+C to quit)
[0;34;10m2015-03-18 16:33:06 -0300 [app][0b4bd652bc85]:[0m 10.2.177.14 - - [18/Mar/2015 16:28:41] "GET / HTTP/1.1" 200 -
[0;34;10m2015-03-18 16:33:06 -0300 [app][2ee9e0fd53f2]:[0m 10.2.177.14 - - [18/Mar/2015 16:28:41] "GET / HTTP/1.1" 200 -
[0;34;10m2015-03-18 16:33:06 -0300 [app][486ba06b3c50]:[0m 10.2.177.14 - - [18/Mar/2015 16:28:41] "GET / HTTP/1.1" 200 -
[0;34;10m2015-03-18 16:33:06 -0300 [app][a29e3bcbdfc3]:[0m 10.2.177.14 - - [18/Mar/2015 16:28:41] "GET / HTTP/1.1" 200 -
[0;34;10m2015-03-18 16:33:06 -0300 [app][0cf9527e1764]:[0m 10.2.177.14 - - [18/Mar/2015 16:28:41] "GET / HTTP/1.1" 200 -
`)

var expected = "\n\x1b[38;5;154m2\x1b[0m\x1b[38;5;154m0\x1b[0m\x1b[38;5;154m1\x1b[0m\x1b[38;5;154m5\x1b[0m\x1b[38;5;148m-\x1b[0m\x1b[38;5;184m0\x1b[0m\x1b[38;5;184m3\x1b[0m\x1b[38;5;184m-\x1b[0m\x1b[38;5;184m1\x1b[0m\x1b[38;5;184m8\x1b[0m\x1b[38;5;184m \x1b[0m\x1b[38;5;184m1\x1b[0m\x1b[38;5;184m6\x1b[0m\x1b[38;5;184m:\x1b[0m\x1b[38;5;184m3\x1b[0m\x1b[38;5;184m3\x1b[0m\x1b[38;5;178m:\x1b[0m\x1b[38;5;214m0\x1b[0m\x1b[38;5;214m2\x1b[0m\x1b[38;5;214m \x1b[0m\x1b[38;5;214m-\x1b[0m\x1b[38;5;214m0\x1b[0m\x1b[38;5;214m3\x1b[0m\x1b[38;5;214m0\x1b[0m\x1b[38;5;214m0\x1b[0m\x1b[38;5;214m \x1b[0m\x1b[38;5;208m[\x1b[0m\x1b[38;5;208ma\x1b[0m\x1b[38;5;208mp\x1b[0m\x1b[38;5;208mp\x1b[0m\x1b[38;5;208m]\x1b[0m\x1b[38;5;208m[\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208mb\x1b[0m\x1b[38;5;208m4\x1b[0m\x1b[38;5;209mb\x1b[0m\x1b[38;5;203md\x1b[0m\x1b[38;5;203m6\x1b[0m\x1b[38;5;203m5\x1b[0m\x1b[38;5;203m2\x1b[0m\x1b[38;5;203mb\x1b[0m\x1b[38;5;203mc\x1b[0m\x1b[38;5;203m8\x1b[0m\x1b[38;5;203m5\x1b[0m\x1b[38;5;203m]\x1b[0m\x1b[38;5;203m:\x1b[0m\x1b[38;5;203m \x1b[0m\x1b[38;5;204m \x1b[0m\x1b[38;5;198m*\x1b[0m\x1b[38;5;198m \x1b[0m\x1b[38;5;198mR\x1b[0m\x1b[38;5;198mu\x1b[0m\x1b[38;5;198mn\x1b[0m\x1b[38;5;198mn\x1b[0m\x1b[38;5;198mi\x1b[0m\x1b[38;5;198mn\x1b[0m\x1b[38;5;198mg\x1b[0m\x1b[38;5;199m \x1b[0m\x1b[38;5;199mo\x1b[0m\x1b[38;5;199mn\x1b[0m\x1b[38;5;199m \x1b[0m\x1b[38;5;199mh\x1b[0m\x1b[38;5;199mt\x1b[0m\x1b[38;5;199mt\x1b[0m\x1b[38;5;199mp\x1b[0m\x1b[38;5;199m:\x1b[0m\x1b[38;5;163m/\x1b[0m\x1b[38;5;163m/\x1b[0m\x1b[38;5;164m0\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;164m0\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;164m0\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;164m0\x1b[0m\x1b[38;5;164m:\x1b[0m\x1b[38;5;164m8\x1b[0m\x1b[38;5;164m8\x1b[0m\x1b[38;5;128m8\x1b[0m\x1b[38;5;128m8\x1b[0m\x1b[38;5;129m/\x1b[0m\x1b[38;5;129m \x1b[0m\x1b[38;5;129m(\x1b[0m\x1b[38;5;129mP\x1b[0m\x1b[38;5;129mr\x1b[0m\x1b[38;5;129me\x1b[0m\x1b[38;5;129ms\x1b[0m\x1b[38;5;129ms\x1b[0m\x1b[38;5;129m \x1b[0m\x1b[38;5;93mC\x1b[0m\x1b[38;5;93mT\x1b[0m\x1b[38;5;93mR\x1b[0m\x1b[38;5;93mL\x1b[0m\x1b[38;5;93m+\x1b[0m\x1b[38;5;93mC\x1b[0m\x1b[38;5;93m \x1b[0m\x1b[38;5;93mt\x1b[0m\x1b[38;5;93mo\x1b[0m\x1b[38;5;99m \x1b[0m\x1b[38;5;63mq\x1b[0m\x1b[38;5;63mu\x1b[0m\x1b[38;5;63mi\x1b[0m\x1b[38;5;63mt\x1b[0m\x1b[38;5;63m)\x1b[0m\n\x1b[38;5;154m2\x1b[0m\x1b[38;5;148m0\x1b[0m\x1b[38;5;184m1\x1b[0m\x1b[38;5;184m5\x1b[0m\x1b[38;5;184m-\x1b[0m\x1b[38;5;184m0\x1b[0m\x1b[38;5;184m3\x1b[0m\x1b[38;5;184m-\x1b[0m\x1b[38;5;184m1\x1b[0m\x1b[38;5;184m8\x1b[0m\x1b[38;5;184m \x1b[0m\x1b[38;5;184m1\x1b[0m\x1b[38;5;184m6\x1b[0m\x1b[38;5;178m:\x1b[0m\x1b[38;5;214m3\x1b[0m\x1b[38;5;214m3\x1b[0m\x1b[38;5;214m:\x1b[0m\x1b[38;5;214m0\x1b[0m\x1b[38;5;214m2\x1b[0m\x1b[38;5;214m \x1b[0m\x1b[38;5;214m-\x1b[0m\x1b[38;5;214m0\x1b[0m\x1b[38;5;214m3\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208m \x1b[0m\x1b[38;5;208m[\x1b[0m\x1b[38;5;208ma\x1b[0m\x1b[38;5;208mp\x1b[0m\x1b[38;5;208mp\x1b[0m\x1b[38;5;208m]\x1b[0m\x1b[38;5;208m[\x1b[0m\x1b[38;5;209ma\x1b[0m\x1b[38;5;203m2\x1b[0m\x1b[38;5;203m9\x1b[0m\x1b[38;5;203me\x1b[0m\x1b[38;5;203m3\x1b[0m\x1b[38;5;203mb\x1b[0m\x1b[38;5;203mc\x1b[0m\x1b[38;5;203mb\x1b[0m\x1b[38;5;203md\x1b[0m\x1b[38;5;203mf\x1b[0m\x1b[38;5;203mc\x1b[0m\x1b[38;5;203m3\x1b[0m\x1b[38;5;204m]\x1b[0m\x1b[38;5;198m:\x1b[0m\x1b[38;5;198m \x1b[0m\x1b[38;5;198m \x1b[0m\x1b[38;5;198m*\x1b[0m\x1b[38;5;198m \x1b[0m\x1b[38;5;198mR\x1b[0m\x1b[38;5;198mu\x1b[0m\x1b[38;5;198mn\x1b[0m\x1b[38;5;198mn\x1b[0m\x1b[38;5;199mi\x1b[0m\x1b[38;5;199mn\x1b[0m\x1b[38;5;199mg\x1b[0m\x1b[38;5;199m \x1b[0m\x1b[38;5;199mo\x1b[0m\x1b[38;5;199mn\x1b[0m\x1b[38;5;199m \x1b[0m\x1b[38;5;199mh\x1b[0m\x1b[38;5;199mt\x1b[0m\x1b[38;5;163mt\x1b[0m\x1b[38;5;163mp\x1b[0m\x1b[38;5;164m:\x1b[0m\x1b[38;5;164m/\x1b[0m\x1b[38;5;164m/\x1b[0m\x1b[38;5;164m0\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;164m0\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;164m0\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;164m0\x1b[0m\x1b[38;5;128m:\x1b[0m\x1b[38;5;128m8\x1b[0m\x1b[38;5;129m8\x1b[0m\x1b[38;5;129m8\x1b[0m\x1b[38;5;129m8\x1b[0m\x1b[38;5;129m/\x1b[0m\x1b[38;5;129m \x1b[0m\x1b[38;5;129m(\x1b[0m\x1b[38;5;129mP\x1b[0m\x1b[38;5;129mr\x1b[0m\x1b[38;5;129me\x1b[0m\x1b[38;5;93ms\x1b[0m\x1b[38;5;93ms\x1b[0m\x1b[38;5;93m \x1b[0m\x1b[38;5;93mC\x1b[0m\x1b[38;5;93mT\x1b[0m\x1b[38;5;93mR\x1b[0m\x1b[38;5;93mL\x1b[0m\x1b[38;5;93m+\x1b[0m\x1b[38;5;93mC\x1b[0m\x1b[38;5;99m \x1b[0m\x1b[38;5;63mt\x1b[0m\x1b[38;5;63mo\x1b[0m\x1b[38;5;63m \x1b[0m\x1b[38;5;63mq\x1b[0m\x1b[38;5;63mu\x1b[0m\x1b[38;5;63mi\x1b[0m\x1b[38;5;63mt\x1b[0m\x1b[38;5;63m)\x1b[0m\n\x1b[38;5;184m2\x1b[0m\x1b[38;5;184m0\x1b[0m\x1b[38;5;184m1\x1b[0m\x1b[38;5;184m5\x1b[0m\x1b[38;5;184m-\x1b[0m\x1b[38;5;184m0\x1b[0m\x1b[38;5;184m3\x1b[0m\x1b[38;5;184m-\x1b[0m\x1b[38;5;184m1\x1b[0m\x1b[38;5;184m8\x1b[0m\x1b[38;5;178m \x1b[0m\x1b[38;5;214m1\x1b[0m\x1b[38;5;214m6\x1b[0m\x1b[38;5;214m:\x1b[0m\x1b[38;5;214m3\x1b[0m\x1b[38;5;214m3\x1b[0m\x1b[38;5;214m:\x1b[0m\x1b[38;5;214m0\x1b[0m\x1b[38;5;214m4\x1b[0m\x1b[38;5;214m \x1b[0m\x1b[38;5;208m-\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208m3\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208m \x1b[0m\x1b[38;5;208m[\x1b[0m\x1b[38;5;208ma\x1b[0m\x1b[38;5;208mp\x1b[0m\x1b[38;5;209mp\x1b[0m\x1b[38;5;203m]\x1b[0m\x1b[38;5;203m[\x1b[0m\x1b[38;5;203m2\x1b[0m\x1b[38;5;203me\x1b[0m\x1b[38;5;203me\x1b[0m\x1b[38;5;203m9\x1b[0m\x1b[38;5;203me\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;203mf\x1b[0m\x1b[38;5;203md\x1b[0m\x1b[38;5;203m5\x1b[0m\x1b[38;5;204m3\x1b[0m\x1b[38;5;198mf\x1b[0m\x1b[38;5;198m2\x1b[0m\x1b[38;5;198m]\x1b[0m\x1b[38;5;198m:\x1b[0m\x1b[38;5;198m \x1b[0m\x1b[38;5;198m \x1b[0m\x1b[38;5;198m*\x1b[0m\x1b[38;5;198m \x1b[0m\x1b[38;5;198mR\x1b[0m\x1b[38;5;199mu\x1b[0m\x1b[38;5;199mn\x1b[0m\x1b[38;5;199mn\x1b[0m\x1b[38;5;199mi\x1b[0m\x1b[38;5;199mn\x1b[0m\x1b[38;5;199mg\x1b[0m\x1b[38;5;199m \x1b[0m\x1b[38;5;199mo\x1b[0m\x1b[38;5;199mn\x1b[0m\x1b[38;5;163m \x1b[0m\x1b[38;5;163mh\x1b[0m\x1b[38;5;164mt\x1b[0m\x1b[38;5;164mt\x1b[0m\x1b[38;5;164mp\x1b[0m\x1b[38;5;164m:\x1b[0m\x1b[38;5;164m/\x1b[0m\x1b[38;5;164m/\x1b[0m\x1b[38;5;164m0\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;164m0\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;128m0\x1b[0m\x1b[38;5;128m.\x1b[0m\x1b[38;5;129m0\x1b[0m\x1b[38;5;129m:\x1b[0m\x1b[38;5;129m8\x1b[0m\x1b[38;5;129m8\x1b[0m\x1b[38;5;129m8\x1b[0m\x1b[38;5;129m8\x1b[0m\x1b[38;5;129m/\x1b[0m\x1b[38;5;129m \x1b[0m\x1b[38;5;129m(\x1b[0m\x1b[38;5;93mP\x1b[0m\x1b[38;5;93mr\x1b[0m\x1b[38;5;93me\x1b[0m\x1b[38;5;93ms\x1b[0m\x1b[38;5;93ms\x1b[0m\x1b[38;5;93m \x1b[0m\x1b[38;5;93mC\x1b[0m\x1b[38;5;93mT\x1b[0m\x1b[38;5;93mR\x1b[0m\x1b[38;5;99mL\x1b[0m\x1b[38;5;63m+\x1b[0m\x1b[38;5;63mC\x1b[0m\x1b[38;5;63m \x1b[0m\x1b[38;5;63mt\x1b[0m\x1b[38;5;63mo\x1b[0m\x1b[38;5;63m \x1b[0m\x1b[38;5;63mq\x1b[0m\x1b[38;5;63mu\x1b[0m\x1b[38;5;63mi\x1b[0m\x1b[38;5;63mt\x1b[0m\x1b[38;5;63m)\x1b[0m\n\x1b[38;5;184m2\x1b[0m\x1b[38;5;184m0\x1b[0m\x1b[38;5;184m1\x1b[0m\x1b[38;5;184m5\x1b[0m\x1b[38;5;184m-\x1b[0m\x1b[38;5;184m0\x1b[0m\x1b[38;5;184m3\x1b[0m\x1b[38;5;178m-\x1b[0m\x1b[38;5;214m1\x1b[0m\x1b[38;5;214m8\x1b[0m\x1b[38;5;214m \x1b[0m\x1b[38;5;214m1\x1b[0m\x1b[38;5;214m6\x1b[0m\x1b[38;5;214m:\x1b[0m\x1b[38;5;214m3\x1b[0m\x1b[38;5;214m3\x1b[0m\x1b[38;5;214m:\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208m4\x1b[0m\x1b[38;5;208m \x1b[0m\x1b[38;5;208m-\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208m3\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208m \x1b[0m\x1b[38;5;209m[\x1b[0m\x1b[38;5;203ma\x1b[0m\x1b[38;5;203mp\x1b[0m\x1b[38;5;203mp\x1b[0m\x1b[38;5;203m]\x1b[0m\x1b[38;5;203m[\x1b[0m\x1b[38;5;203m4\x1b[0m\x1b[38;5;203m8\x1b[0m\x1b[38;5;203m6\x1b[0m\x1b[38;5;203mb\x1b[0m\x1b[38;5;203ma\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;204m6\x1b[0m\x1b[38;5;198mb\x1b[0m\x1b[38;5;198m3\x1b[0m\x1b[38;5;198mc\x1b[0m\x1b[38;5;198m5\x1b[0m\x1b[38;5;198m0\x1b[0m\x1b[38;5;198m]\x1b[0m\x1b[38;5;198m:\x1b[0m\x1b[38;5;198m \x1b[0m\x1b[38;5;198m \x1b[0m\x1b[38;5;199m*\x1b[0m\x1b[38;5;199m \x1b[0m\x1b[38;5;199mR\x1b[0m\x1b[38;5;199mu\x1b[0m\x1b[38;5;199mn\x1b[0m\x1b[38;5;199mn\x1b[0m\x1b[38;5;199mi\x1b[0m\x1b[38;5;199mn\x1b[0m\x1b[38;5;199mg\x1b[0m\x1b[38;5;163m \x1b[0m\x1b[38;5;163mo\x1b[0m\x1b[38;5;164mn\x1b[0m\x1b[38;5;164m \x1b[0m\x1b[38;5;164mh\x1b[0m\x1b[38;5;164mt\x1b[0m\x1b[38;5;164mt\x1b[0m\x1b[38;5;164mp\x1b[0m\x1b[38;5;164m:\x1b[0m\x1b[38;5;164m/\x1b[0m\x1b[38;5;164m/\x1b[0m\x1b[38;5;164m0\x1b[0m\x1b[38;5;128m.\x1b[0m\x1b[38;5;128m0\x1b[0m\x1b[38;5;129m.\x1b[0m\x1b[38;5;129m0\x1b[0m\x1b[38;5;129m.\x1b[0m\x1b[38;5;129m0\x1b[0m\x1b[38;5;129m:\x1b[0m\x1b[38;5;129m8\x1b[0m\x1b[38;5;129m8\x1b[0m\x1b[38;5;129m8\x1b[0m\x1b[38;5;129m8\x1b[0m\x1b[38;5;93m/\x1b[0m\x1b[38;5;93m \x1b[0m\x1b[38;5;93m(\x1b[0m\x1b[38;5;93mP\x1b[0m\x1b[38;5;93mr\x1b[0m\x1b[38;5;93me\x1b[0m\x1b[38;5;93ms\x1b[0m\x1b[38;5;93ms\x1b[0m\x1b[38;5;93m \x1b[0m\x1b[38;5;99mC\x1b[0m\x1b[38;5;63mT\x1b[0m\x1b[38;5;63mR\x1b[0m\x1b[38;5;63mL\x1b[0m\x1b[38;5;63m+\x1b[0m\x1b[38;5;63mC\x1b[0m\x1b[38;5;63m \x1b[0m\x1b[38;5;63mt\x1b[0m\x1b[38;5;63mo\x1b[0m\x1b[38;5;63m \x1b[0m\x1b[38;5;63mq\x1b[0m\x1b[38;5;63mu\x1b[0m\x1b[38;5;69mi\x1b[0m\x1b[38;5;33mt\x1b[0m\x1b[38;5;33m)\x1b[0m\n\x1b[38;5;184m2\x1b[0m\x1b[38;5;184m0\x1b[0m\x1b[38;5;184m1\x1b[0m\x1b[38;5;184m5\x1b[0m\x1b[38;5;178m-\x1b[0m\x1b[38;5;214m0\x1b[0m\x1b[38;5;214m3\x1b[0m\x1b[38;5;214m-\x1b[0m\x1b[38;5;214m1\x1b[0m\x1b[38;5;214m8\x1b[0m\x1b[38;5;214m \x1b[0m\x1b[38;5;214m1\x1b[0m\x1b[38;5;214m6\x1b[0m\x1b[38;5;214m:\x1b[0m\x1b[38;5;208m3\x1b[0m\x1b[38;5;208m3\x1b[0m\x1b[38;5;208m:\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208m4\x1b[0m\x1b[38;5;208m \x1b[0m\x1b[38;5;208m-\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208m3\x1b[0m\x1b[38;5;209m0\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;203m \x1b[0m\x1b[38;5;203m[\x1b[0m\x1b[38;5;203ma\x1b[0m\x1b[38;5;203mp\x1b[0m\x1b[38;5;203mp\x1b[0m\x1b[38;5;203m]\x1b[0m\x1b[38;5;203m[\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;203mc\x1b[0m\x1b[38;5;203mf\x1b[0m\x1b[38;5;204m9\x1b[0m\x1b[38;5;198m5\x1b[0m\x1b[38;5;198m2\x1b[0m\x1b[38;5;198m7\x1b[0m\x1b[38;5;198me\x1b[0m\x1b[38;5;198m1\x1b[0m\x1b[38;5;198m7\x1b[0m\x1b[38;5;198m6\x1b[0m\x1b[38;5;198m4\x1b[0m\x1b[38;5;198m]\x1b[0m\x1b[38;5;199m:\x1b[0m\x1b[38;5;199m \x1b[0m\x1b[38;5;199m \x1b[0m\x1b[38;5;199m*\x1b[0m\x1b[38;5;199m \x1b[0m\x1b[38;5;199mR\x1b[0m\x1b[38;5;199mu\x1b[0m\x1b[38;5;199mn\x1b[0m\x1b[38;5;199mn\x1b[0m\x1b[38;5;163mi\x1b[0m\x1b[38;5;163mn\x1b[0m\x1b[38;5;164mg\x1b[0m\x1b[38;5;164m \x1b[0m\x1b[38;5;164mo\x1b[0m\x1b[38;5;164mn\x1b[0m\x1b[38;5;164m \x1b[0m\x1b[38;5;164mh\x1b[0m\x1b[38;5;164mt\x1b[0m\x1b[38;5;164mt\x1b[0m\x1b[38;5;164mp\x1b[0m\x1b[38;5;164m:\x1b[0m\x1b[38;5;128m/\x1b[0m\x1b[38;5;128m/\x1b[0m\x1b[38;5;129m0\x1b[0m\x1b[38;5;129m.\x1b[0m\x1b[38;5;129m0\x1b[0m\x1b[38;5;129m.\x1b[0m\x1b[38;5;129m0\x1b[0m\x1b[38;5;129m.\x1b[0m\x1b[38;5;129m0\x1b[0m\x1b[38;5;129m:\x1b[0m\x1b[38;5;129m8\x1b[0m\x1b[38;5;93m8\x1b[0m\x1b[38;5;93m8\x1b[0m\x1b[38;5;93m8\x1b[0m\x1b[38;5;93m/\x1b[0m\x1b[38;5;93m \x1b[0m\x1b[38;5;93m(\x1b[0m\x1b[38;5;93mP\x1b[0m\x1b[38;5;93mr\x1b[0m\x1b[38;5;93me\x1b[0m\x1b[38;5;99ms\x1b[0m\x1b[38;5;63ms\x1b[0m\x1b[38;5;63m \x1b[0m\x1b[38;5;63mC\x1b[0m\x1b[38;5;63mT\x1b[0m\x1b[38;5;63mR\x1b[0m\x1b[38;5;63mL\x1b[0m\x1b[38;5;63m+\x1b[0m\x1b[38;5;63mC\x1b[0m\x1b[38;5;63m \x1b[0m\x1b[38;5;63mt\x1b[0m\x1b[38;5;63mo\x1b[0m\x1b[38;5;69m \x1b[0m\x1b[38;5;33mq\x1b[0m\x1b[38;5;33mu\x1b[0m\x1b[38;5;33mi\x1b[0m\x1b[38;5;33mt\x1b[0m\x1b[38;5;33m)\x1b[0m\n\x1b[38;5;184m2\x1b[0m\x1b[38;5;178m0\x1b[0m\x1b[38;5;214m1\x1b[0m\x1b[38;5;214m5\x1b[0m\x1b[38;5;214m-\x1b[0m\x1b[38;5;214m0\x1b[0m\x1b[38;5;214m3\x1b[0m\x1b[38;5;214m-\x1b[0m\x1b[38;5;214m1\x1b[0m\x1b[38;5;214m8\x1b[0m\x1b[38;5;214m \x1b[0m\x1b[38;5;208m1\x1b[0m\x1b[38;5;208m6\x1b[0m\x1b[38;5;208m:\x1b[0m\x1b[38;5;208m3\x1b[0m\x1b[38;5;208m3\x1b[0m\x1b[38;5;208m:\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208m6\x1b[0m\x1b[38;5;208m \x1b[0m\x1b[38;5;209m-\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;203m3\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;203m \x1b[0m\x1b[38;5;203m[\x1b[0m\x1b[38;5;203ma\x1b[0m\x1b[38;5;203mp\x1b[0m\x1b[38;5;203mp\x1b[0m\x1b[38;5;203m]\x1b[0m\x1b[38;5;203m[\x1b[0m\x1b[38;5;204m0\x1b[0m\x1b[38;5;198mb\x1b[0m\x1b[38;5;198m4\x1b[0m\x1b[38;5;198mb\x1b[0m\x1b[38;5;198md\x1b[0m\x1b[38;5;198m6\x1b[0m\x1b[38;5;198m5\x1b[0m\x1b[38;5;198m2\x1b[0m\x1b[38;5;198mb\x1b[0m\x1b[38;5;198mc\x1b[0m\x1b[38;5;199m8\x1b[0m\x1b[38;5;199m5\x1b[0m\x1b[38;5;199m]\x1b[0m\x1b[38;5;199m:\x1b[0m\x1b[38;5;199m \x1b[0m\x1b[38;5;199m1\x1b[0m\x1b[38;5;199m0\x1b[0m\x1b[38;5;199m.\x1b[0m\x1b[38;5;199m2\x1b[0m\x1b[38;5;163m.\x1b[0m\x1b[38;5;163m1\x1b[0m\x1b[38;5;164m7\x1b[0m\x1b[38;5;164m7\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;164m1\x1b[0m\x1b[38;5;164m4\x1b[0m\x1b[38;5;164m \x1b[0m\x1b[38;5;164m-\x1b[0m\x1b[38;5;164m \x1b[0m\x1b[38;5;164m-\x1b[0m\x1b[38;5;164m \x1b[0m\x1b[38;5;128m[\x1b[0m\x1b[38;5;128m1\x1b[0m\x1b[38;5;129m8\x1b[0m\x1b[38;5;129m/\x1b[0m\x1b[38;5;129mM\x1b[0m\x1b[38;5;129ma\x1b[0m\x1b[38;5;129mr\x1b[0m\x1b[38;5;129m/\x1b[0m\x1b[38;5;129m2\x1b[0m\x1b[38;5;129m0\x1b[0m\x1b[38;5;129m1\x1b[0m\x1b[38;5;93m5\x1b[0m\x1b[38;5;93m \x1b[0m\x1b[38;5;93m1\x1b[0m\x1b[38;5;93m6\x1b[0m\x1b[38;5;93m:\x1b[0m\x1b[38;5;93m2\x1b[0m\x1b[38;5;93m8\x1b[0m\x1b[38;5;93m:\x1b[0m\x1b[38;5;93m4\x1b[0m\x1b[38;5;99m1\x1b[0m\x1b[38;5;63m]\x1b[0m\x1b[38;5;63m \x1b[0m\x1b[38;5;63m\"\x1b[0m\x1b[38;5;63mG\x1b[0m\x1b[38;5;63mE\x1b[0m\x1b[38;5;63mT\x1b[0m\x1b[38;5;63m \x1b[0m\x1b[38;5;63m/\x1b[0m\x1b[38;5;63m \x1b[0m\x1b[38;5;63mH\x1b[0m\x1b[38;5;63mT\x1b[0m\x1b[38;5;69mT\x1b[0m\x1b[38;5;33mP\x1b[0m\x1b[38;5;33m/\x1b[0m\x1b[38;5;33m1\x1b[0m\x1b[38;5;33m.\x1b[0m\x1b[38;5;33m1\x1b[0m\x1b[38;5;33m\"\x1b[0m\x1b[38;5;33m \x1b[0m\x1b[38;5;33m2\x1b[0m\x1b[38;5;33m0\x1b[0m\x1b[38;5;39m0\x1b[0m\x1b[38;5;39m \x1b[0m\x1b[38;5;39m-\x1b[0m\n\x1b[38;5;214m2\x1b[0m\x1b[38;5;214m0\x1b[0m\x1b[38;5;214m1\x1b[0m\x1b[38;5;214m5\x1b[0m\x1b[38;5;214m-\x1b[0m\x1b[38;5;214m0\x1b[0m\x1b[38;5;214m3\x1b[0m\x1b[38;5;214m-\x1b[0m\x1b[38;5;208m1\x1b[0m\x1b[38;5;208m8\x1b[0m\x1b[38;5;208m \x1b[0m\x1b[38;5;208m1\x1b[0m\x1b[38;5;208m6\x1b[0m\x1b[38;5;208m:\x1b[0m\x1b[38;5;208m3\x1b[0m\x1b[38;5;208m3\x1b[0m\x1b[38;5;208m:\x1b[0m\x1b[38;5;209m0\x1b[0m\x1b[38;5;203m6\x1b[0m\x1b[38;5;203m \x1b[0m\x1b[38;5;203m-\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;203m3\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;203m \x1b[0m\x1b[38;5;203m[\x1b[0m\x1b[38;5;203ma\x1b[0m\x1b[38;5;203mp\x1b[0m\x1b[38;5;204mp\x1b[0m\x1b[38;5;198m]\x1b[0m\x1b[38;5;198m[\x1b[0m\x1b[38;5;198m2\x1b[0m\x1b[38;5;198me\x1b[0m\x1b[38;5;198me\x1b[0m\x1b[38;5;198m9\x1b[0m\x1b[38;5;198me\x1b[0m\x1b[38;5;198m0\x1b[0m\x1b[38;5;198mf\x1b[0m\x1b[38;5;199md\x1b[0m\x1b[38;5;199m5\x1b[0m\x1b[38;5;199m3\x1b[0m\x1b[38;5;199mf\x1b[0m\x1b[38;5;199m2\x1b[0m\x1b[38;5;199m]\x1b[0m\x1b[38;5;199m:\x1b[0m\x1b[38;5;199m \x1b[0m\x1b[38;5;199m1\x1b[0m\x1b[38;5;163m0\x1b[0m\x1b[38;5;163m.\x1b[0m\x1b[38;5;164m2\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;164m1\x1b[0m\x1b[38;5;164m7\x1b[0m\x1b[38;5;164m7\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;164m1\x1b[0m\x1b[38;5;164m4\x1b[0m\x1b[38;5;164m \x1b[0m\x1b[38;5;164m-\x1b[0m\x1b[38;5;128m \x1b[0m\x1b[38;5;128m-\x1b[0m\x1b[38;5;129m \x1b[0m\x1b[38;5;129m[\x1b[0m\x1b[38;5;129m1\x1b[0m\x1b[38;5;129m8\x1b[0m\x1b[38;5;129m/\x1b[0m\x1b[38;5;129mM\x1b[0m\x1b[38;5;129ma\x1b[0m\x1b[38;5;129mr\x1b[0m\x1b[38;5;129m/\x1b[0m\x1b[38;5;93m2\x1b[0m\x1b[38;5;93m0\x1b[0m\x1b[38;5;93m1\x1b[0m\x1b[38;5;93m5\x1b[0m\x1b[38;5;93m \x1b[0m\x1b[38;5;93m1\x1b[0m\x1b[38;5;93m6\x1b[0m\x1b[38;5;93m:\x1b[0m\x1b[38;5;93m2\x1b[0m\x1b[38;5;99m8\x1b[0m\x1b[38;5;63m:\x1b[0m\x1b[38;5;63m4\x1b[0m\x1b[38;5;63m1\x1b[0m\x1b[38;5;63m]\x1b[0m\x1b[38;5;63m \x1b[0m\x1b[38;5;63m\"\x1b[0m\x1b[38;5;63mG\x1b[0m\x1b[38;5;63mE\x1b[0m\x1b[38;5;63mT\x1b[0m\x1b[38;5;63m \x1b[0m\x1b[38;5;63m/\x1b[0m\x1b[38;5;69m \x1b[0m\x1b[38;5;33mH\x1b[0m\x1b[38;5;33mT\x1b[0m\x1b[38;5;33mT\x1b[0m\x1b[38;5;33mP\x1b[0m\x1b[38;5;33m/\x1b[0m\x1b[38;5;33m1\x1b[0m\x1b[38;5;33m.\x1b[0m\x1b[38;5;33m1\x1b[0m\x1b[38;5;33m\"\x1b[0m\x1b[38;5;39m \x1b[0m\x1b[38;5;39m2\x1b[0m\x1b[38;5;39m0\x1b[0m\x1b[38;5;39m0\x1b[0m\x1b[38;5;39m \x1b[0m\x1b[38;5;39m-\x1b[0m\n\x1b[38;5;214m2\x1b[0m\x1b[38;5;214m0\x1b[0m\x1b[38;5;214m1\x1b[0m\x1b[38;5;214m5\x1b[0m\x1b[38;5;214m-\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208m3\x1b[0m\x1b[38;5;208m-\x1b[0m\x1b[38;5;208m1\x1b[0m\x1b[38;5;208m8\x1b[0m\x1b[38;5;208m \x1b[0m\x1b[38;5;208m1\x1b[0m\x1b[38;5;208m6\x1b[0m\x1b[38;5;208m:\x1b[0m\x1b[38;5;209m3\x1b[0m\x1b[38;5;203m3\x1b[0m\x1b[38;5;203m:\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;203m6\x1b[0m\x1b[38;5;203m \x1b[0m\x1b[38;5;203m-\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;203m3\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;203m \x1b[0m\x1b[38;5;204m[\x1b[0m\x1b[38;5;198ma\x1b[0m\x1b[38;5;198mp\x1b[0m\x1b[38;5;198mp\x1b[0m\x1b[38;5;198m]\x1b[0m\x1b[38;5;198m[\x1b[0m\x1b[38;5;198m4\x1b[0m\x1b[38;5;198m8\x1b[0m\x1b[38;5;198m6\x1b[0m\x1b[38;5;198mb\x1b[0m\x1b[38;5;199ma\x1b[0m\x1b[38;5;199m0\x1b[0m\x1b[38;5;199m6\x1b[0m\x1b[38;5;199mb\x1b[0m\x1b[38;5;199m3\x1b[0m\x1b[38;5;199mc\x1b[0m\x1b[38;5;199m5\x1b[0m\x1b[38;5;199m0\x1b[0m\x1b[38;5;199m]\x1b[0m\x1b[38;5;163m:\x1b[0m\x1b[38;5;163m \x1b[0m\x1b[38;5;164m1\x1b[0m\x1b[38;5;164m0\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;164m2\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;164m1\x1b[0m\x1b[38;5;164m7\x1b[0m\x1b[38;5;164m7\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;164m1\x1b[0m\x1b[38;5;128m4\x1b[0m\x1b[38;5;128m \x1b[0m\x1b[38;5;129m-\x1b[0m\x1b[38;5;129m \x1b[0m\x1b[38;5;129m-\x1b[0m\x1b[38;5;129m \x1b[0m\x1b[38;5;129m[\x1b[0m\x1b[38;5;129m1\x1b[0m\x1b[38;5;129m8\x1b[0m\x1b[38;5;129m/\x1b[0m\x1b[38;5;129mM\x1b[0m\x1b[38;5;93ma\x1b[0m\x1b[38;5;93mr\x1b[0m\x1b[38;5;93m/\x1b[0m\x1b[38;5;93m2\x1b[0m\x1b[38;5;93m0\x1b[0m\x1b[38;5;93m1\x1b[0m\x1b[38;5;93m5\x1b[0m\x1b[38;5;93m \x1b[0m\x1b[38;5;93m1\x1b[0m\x1b[38;5;99m6\x1b[0m\x1b[38;5;63m:\x1b[0m\x1b[38;5;63m2\x1b[0m\x1b[38;5;63m8\x1b[0m\x1b[38;5;63m:\x1b[0m\x1b[38;5;63m4\x1b[0m\x1b[38;5;63m1\x1b[0m\x1b[38;5;63m]\x1b[0m\x1b[38;5;63m \x1b[0m\x1b[38;5;63m\"\x1b[0m\x1b[38;5;63mG\x1b[0m\x1b[38;5;63mE\x1b[0m\x1b[38;5;69mT\x1b[0m\x1b[38;5;33m \x1b[0m\x1b[38;5;33m/\x1b[0m\x1b[38;5;33m \x1b[0m\x1b[38;5;33mH\x1b[0m\x1b[38;5;33mT\x1b[0m\x1b[38;5;33mT\x1b[0m\x1b[38;5;33mP\x1b[0m\x1b[38;5;33m/\x1b[0m\x1b[38;5;33m1\x1b[0m\x1b[38;5;39m.\x1b[0m\x1b[38;5;39m1\x1b[0m\x1b[38;5;39m\"\x1b[0m\x1b[38;5;39m \x1b[0m\x1b[38;5;39m2\x1b[0m\x1b[38;5;39m0\x1b[0m\x1b[38;5;39m0\x1b[0m\x1b[38;5;39m \x1b[0m\x1b[38;5;39m-\x1b[0m\n\x1b[38;5;214m2\x1b[0m\x1b[38;5;214m0\x1b[0m\x1b[38;5;208m1\x1b[0m\x1b[38;5;208m5\x1b[0m\x1b[38;5;208m-\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208m3\x1b[0m\x1b[38;5;208m-\x1b[0m\x1b[38;5;208m1\x1b[0m\x1b[38;5;208m8\x1b[0m\x1b[38;5;208m \x1b[0m\x1b[38;5;209m1\x1b[0m\x1b[38;5;203m6\x1b[0m\x1b[38;5;203m:\x1b[0m\x1b[38;5;203m3\x1b[0m\x1b[38;5;203m3\x1b[0m\x1b[38;5;203m:\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;203m6\x1b[0m\x1b[38;5;203m \x1b[0m\x1b[38;5;203m-\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;203m3\x1b[0m\x1b[38;5;204m0\x1b[0m\x1b[38;5;198m0\x1b[0m\x1b[38;5;198m \x1b[0m\x1b[38;5;198m[\x1b[0m\x1b[38;5;198ma\x1b[0m\x1b[38;5;198mp\x1b[0m\x1b[38;5;198mp\x1b[0m\x1b[38;5;198m]\x1b[0m\x1b[38;5;198m[\x1b[0m\x1b[38;5;198ma\x1b[0m\x1b[38;5;199m2\x1b[0m\x1b[38;5;199m9\x1b[0m\x1b[38;5;199me\x1b[0m\x1b[38;5;199m3\x1b[0m\x1b[38;5;199mb\x1b[0m\x1b[38;5;199mc\x1b[0m\x1b[38;5;199mb\x1b[0m\x1b[38;5;199md\x1b[0m\x1b[38;5;199mf\x1b[0m\x1b[38;5;163mc\x1b[0m\x1b[38;5;163m3\x1b[0m\x1b[38;5;164m]\x1b[0m\x1b[38;5;164m:\x1b[0m\x1b[38;5;164m \x1b[0m\x1b[38;5;164m1\x1b[0m\x1b[38;5;164m0\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;164m2\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;164m1\x1b[0m\x1b[38;5;164m7\x1b[0m\x1b[38;5;128m7\x1b[0m\x1b[38;5;128m.\x1b[0m\x1b[38;5;129m1\x1b[0m\x1b[38;5;129m4\x1b[0m\x1b[38;5;129m \x1b[0m\x1b[38;5;129m-\x1b[0m\x1b[38;5;129m \x1b[0m\x1b[38;5;129m-\x1b[0m\x1b[38;5;129m \x1b[0m\x1b[38;5;129m[\x1b[0m\x1b[38;5;129m1\x1b[0m\x1b[38;5;93m8\x1b[0m\x1b[38;5;93m/\x1b[0m\x1b[38;5;93mM\x1b[0m\x1b[38;5;93ma\x1b[0m\x1b[38;5;93mr\x1b[0m\x1b[38;5;93m/\x1b[0m\x1b[38;5;93m2\x1b[0m\x1b[38;5;93m0\x1b[0m\x1b[38;5;93m1\x1b[0m\x1b[38;5;99m5\x1b[0m\x1b[38;5;63m \x1b[0m\x1b[38;5;63m1\x1b[0m\x1b[38;5;63m6\x1b[0m\x1b[38;5;63m:\x1b[0m\x1b[38;5;63m2\x1b[0m\x1b[38;5;63m8\x1b[0m\x1b[38;5;63m:\x1b[0m\x1b[38;5;63m4\x1b[0m\x1b[38;5;63m1\x1b[0m\x1b[38;5;63m]\x1b[0m\x1b[38;5;63m \x1b[0m\x1b[38;5;69m\"\x1b[0m\x1b[38;5;33mG\x1b[0m\x1b[38;5;33mE\x1b[0m\x1b[38;5;33mT\x1b[0m\x1b[38;5;33m \x1b[0m\x1b[38;5;33m/\x1b[0m\x1b[38;5;33m \x1b[0m\x1b[38;5;33mH\x1b[0m\x1b[38;5;33mT\x1b[0m\x1b[38;5;33mT\x1b[0m\x1b[38;5;39mP\x1b[0m\x1b[38;5;39m/\x1b[0m\x1b[38;5;39m1\x1b[0m\x1b[38;5;39m.\x1b[0m\x1b[38;5;39m1\x1b[0m\x1b[38;5;39m\"\x1b[0m\x1b[38;5;39m \x1b[0m\x1b[38;5;39m2\x1b[0m\x1b[38;5;39m0\x1b[0m\x1b[38;5;38m0\x1b[0m\x1b[38;5;38m \x1b[0m\x1b[38;5;44m-\x1b[0m\n\x1b[38;5;208m2\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208m1\x1b[0m\x1b[38;5;208m5\x1b[0m\x1b[38;5;208m-\x1b[0m\x1b[38;5;208m0\x1b[0m\x1b[38;5;208m3\x1b[0m\x1b[38;5;208m-\x1b[0m\x1b[38;5;209m1\x1b[0m\x1b[38;5;203m8\x1b[0m\x1b[38;5;203m \x1b[0m\x1b[38;5;203m1\x1b[0m\x1b[38;5;203m6\x1b[0m\x1b[38;5;203m:\x1b[0m\x1b[38;5;203m3\x1b[0m\x1b[38;5;203m3\x1b[0m\x1b[38;5;203m:\x1b[0m\x1b[38;5;203m0\x1b[0m\x1b[38;5;203m6\x1b[0m\x1b[38;5;203m \x1b[0m\x1b[38;5;204m-\x1b[0m\x1b[38;5;198m0\x1b[0m\x1b[38;5;198m3\x1b[0m\x1b[38;5;198m0\x1b[0m\x1b[38;5;198m0\x1b[0m\x1b[38;5;198m \x1b[0m\x1b[38;5;198m[\x1b[0m\x1b[38;5;198ma\x1b[0m\x1b[38;5;198mp\x1b[0m\x1b[38;5;198mp\x1b[0m\x1b[38;5;199m]\x1b[0m\x1b[38;5;199m[\x1b[0m\x1b[38;5;199m0\x1b[0m\x1b[38;5;199mc\x1b[0m\x1b[38;5;199mf\x1b[0m\x1b[38;5;199m9\x1b[0m\x1b[38;5;199m5\x1b[0m\x1b[38;5;199m2\x1b[0m\x1b[38;5;199m7\x1b[0m\x1b[38;5;163me\x1b[0m\x1b[38;5;163m1\x1b[0m\x1b[38;5;164m7\x1b[0m\x1b[38;5;164m6\x1b[0m\x1b[38;5;164m4\x1b[0m\x1b[38;5;164m]\x1b[0m\x1b[38;5;164m:\x1b[0m\x1b[38;5;164m \x1b[0m\x1b[38;5;164m1\x1b[0m\x1b[38;5;164m0\x1b[0m\x1b[38;5;164m.\x1b[0m\x1b[38;5;164m2\x1b[0m\x1b[38;5;128m.\x1b[0m\x1b[38;5;128m1\x1b[0m\x1b[38;5;129m7\x1b[0m\x1b[38;5;129m7\x1b[0m\x1b[38;5;129m.\x1b[0m\x1b[38;5;129m1\x1b[0m\x1b[38;5;129m4\x1b[0m\x1b[38;5;129m \x1b[0m\x1b[38;5;129m-\x1b[0m\x1b[38;5;129m \x1b[0m\x1b[38;5;129m-\x1b[0m\x1b[38;5;93m \x1b[0m\x1b[38;5;93m[\x1b[0m\x1b[38;5;93m1\x1b[0m\x1b[38;5;93m8\x1b[0m\x1b[38;5;93m/\x1b[0m\x1b[38;5;93mM\x1b[0m\x1b[38;5;93ma\x1b[0m\x1b[38;5;93mr\x1b[0m\x1b[38;5;93m/\x1b[0m\x1b[38;5;99m2\x1b[0m\x1b[38;5;63m0\x1b[0m\x1b[38;5;63m1\x1b[0m\x1b[38;5;63m5\x1b[0m\x1b[38;5;63m \x1b[0m\x1b[38;5;63m1\x1b[0m\x1b[38;5;63m6\x1b[0m\x1b[38;5;63m:\x1b[0m\x1b[38;5;63m2\x1b[0m\x1b[38;5;63m8\x1b[0m\x1b[38;5;63m:\x1b[0m\x1b[38;5;63m4\x1b[0m\x1b[38;5;69m1\x1b[0m\x1b[38;5;33m]\x1b[0m\x1b[38;5;33m \x1b[0m\x1b[38;5;33m\"\x1b[0m\x1b[38;5;33mG\x1b[0m\x1b[38;5;33mE\x1b[0m\x1b[38;5;33mT\x1b[0m\x1b[38;5;33m \x1b[0m\x1b[38;5;33m/\x1b[0m\x1b[38;5;33m \x1b[0m\x1b[38;5;39mH\x1b[0m\x1b[38;5;39mT\x1b[0m\x1b[38;5;39mT\x1b[0m\x1b[38;5;39mP\x1b[0m\x1b[38;5;39m/\x1b[0m\x1b[38;5;39m1\x1b[0m\x1b[38;5;39m.\x1b[0m\x1b[38;5;39m1\x1b[0m\x1b[38;5;39m\"\x1b[0m\x1b[38;5;38m \x1b[0m\x1b[38;5;38m2\x1b[0m\x1b[38;5;44m0\x1b[0m\x1b[38;5;44m0\x1b[0m\x1b[38;5;44m \x1b[0m\x1b[38;5;44m-\x1b[0m\n"

func BenchmarkLolRegular(b *testing.B) {
	reader := bytes.NewReader(bigData)
	buf := bytes.NewBuffer(nil)
	w := &LolWriter{
		os:     1,
		base:   buf,
		freq:   0.1,
		spread: 3.0,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cat(w, reader)
	}
	if buf.String() != expected {
		b.Fatalf("incorrect output, expected:\n%s\ngot:\n%s", expected, buf.String())
	}
}

func BenchmarkLolBuffered(b *testing.B) {
	reader := bytes.NewReader(bigData)
	buf := bytes.NewBuffer(nil)
	w := &LolBufferedWriter{
		os:     1,
		base:   buf,
		freq:   0.1,
		spread: 3.0,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cat(w, reader)
	}
	if buf.String() != expected {
		b.Fatalf("incorrect output, expected:\n%q\ngot:\n%q", expected, buf.String())
	}
}
