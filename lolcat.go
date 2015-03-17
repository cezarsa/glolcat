package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var stripAnsi = regexp.MustCompile(`\[(\d+)(;\d+)?(;\d+)?[m|K]`)

func rainbow(freq, i float64) (int, int, int) {
	red := int(math.Sin(freq*i+0)*127 + 128)
	green := int(math.Sin(freq*i+2*math.Pi/3)*127 + 128)
	blue := int(math.Sin(freq*i+4*math.Pi/3)*127 + 128)
	return red, green, blue
}

type LolWriter struct {
	base   io.Writer
	buf    bytes.Buffer
	os     int
	spread float64
	freq   float64
}

func (w *LolWriter) Write(data []byte) (int, error) {
	for _, c := range data {
		if c == '\n' {
			w.flush()
			w.base.Write([]byte{'\n'})
		} else {
			w.buf.WriteByte(c)
		}
	}
	return len(data), nil
}

func (w *LolWriter) flush() {
	line := w.buf.String()
	line = stripAnsi.ReplaceAllString(line, "")
	line = strings.Replace(line, "\t", "        ", -1)
	for i, chr := range []rune(line) {
		r, g, b := rainbow(w.freq, float64(w.os)+(float64(i)/w.spread))
		fmt.Fprint(w.base, colored(string(chr), r, g, b))
	}
	w.buf.Reset()
	w.os += 1
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	seed := int(rand.Int31n(256))
	writer := LolWriter{
		os:     seed,
		base:   os.Stdout,
		freq:   0.1,
		spread: 3.0,
	}
	io.Copy(&writer, os.Stdin)
	writer.flush()
}

func colored(str string, r, g, b int) string {
	return fmt.Sprintf("\033[38%s;m%s\033[0m", rgb(float64(r), float64(g), float64(b)), str)
}

func toBaseColor(color float64, mod int) int {
	return int(6*(color/256)) * mod
}

func rgb(red, green, blue float64) string {
	grayPossible := true
	sep := 42.5
	var gray bool
	for grayPossible {
		if red < sep || green < sep || blue < sep {
			gray = red < sep && green < sep && blue < sep
			grayPossible = false
		}
		sep += 42.5
	}
	if gray {
		return fmt.Sprintf(";5;%d", 232+int((red+green+blue)/33.0))
	} else {
		value := 16 + toBaseColor(red, 36) + toBaseColor(green, 6) + toBaseColor(blue, 1)
		return fmt.Sprintf(";5;%d", value)
	}
}
