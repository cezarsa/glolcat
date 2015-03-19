package main

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"regexp"
	"time"
)

var stripAnsiStart = regexp.MustCompile("^\033" + `\[(\d+)(;\d+)?(;\d+)?[m|K]`)

type LolWriter struct {
	base   io.Writer
	os     int
	li     int
	spread float64
	freq   float64
}

var tabSpaces = []byte("        ")

func (w *LolWriter) Write(data []byte) (int, error) {
	for i := 0; i < len(data); i++ {
		c := data[i]
		if c == '\n' {
			w.li = 0
			w.os++
			w.base.Write([]byte{'\n'})
		} else if c == '\t' {
			w.li += len(tabSpaces)
			w.base.Write(tabSpaces)
		} else {
			matchPos := stripAnsiStart.FindIndex(data[i:])
			if matchPos != nil {
				i += matchPos[1] - 1
				continue
			}
			r, g, b := rainbow(w.freq, float64(w.os)+(float64(w.li)/w.spread))
			fmt.Fprint(w.base, colored(string(c), r, g, b))
			w.li++
		}
	}
	return len(data), nil
}

func (w *LolWriter) Close() error {
	return nil
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
	cat(&writer, os.Stdin)
}

func cat(writer io.WriteCloser, reader io.Reader) {
	io.Copy(writer, reader)
	writer.Close()
}

func rainbow(freq, i float64) (int, int, int) {
	red := int(math.Sin(freq*i+0)*127 + 128)
	green := int(math.Sin(freq*i+2*math.Pi/3)*127 + 128)
	blue := int(math.Sin(freq*i+4*math.Pi/3)*127 + 128)
	return red, green, blue
}

func colored(str string, r, g, b int) string {
	return fmt.Sprintf("\033[38%sm%s\033[0m", rgb(float64(r), float64(g), float64(b)), str)
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
