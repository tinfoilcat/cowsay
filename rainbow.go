package cowsay

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
)

const (
	red = iota + 31
	green
	yellow
	blue
	magenta
	cyan
)

func (cow *Cow) makeRainbow(mow string) string {
	attribute := ""
	if cow.Bold {
		attribute = ";1"
	}

	rainbow := []int{magenta, red, yellow, green, cyan, blue}
	b := bytes.NewBuffer(make([]byte, 0, len(mow)))
	i := 0
	for _, char := range mow {
		if char == '\n' {
			i = 0
			b.WriteRune(char)
			continue
		}
		b.WriteString(fmt.Sprintf("\x1b[%d%sm%c\x1b[0m", rainbow[i%6], attribute, char))
		i++
	}

	return b.String()
}

func (cow *Cow) makeAurora(mow string) string {
	attribute := ""
	if cow.Bold {
		attribute = ";1"
	}

	i := rand.Intn(256)
	freq := 0.01
	buf := bytes.NewBuffer(make([]byte, 0, len(mow)))
	for _, char := range mow {
		if char == '\n' {
			buf.WriteRune(char)
			continue
		}

		buf.WriteString(fmt.Sprintf("\033[38;5;%d%sm%c\033[0m", rgb(freq, float64(i)), attribute, char))
		i++
	}
	return buf.String()
}

func rgb(freq, i float64) int {
	red := int(6*((math.Sin(freq*i+0)*127+128)/256)) * 36
	green := int(6*((math.Sin(freq*i+2*math.Pi/3)*127+128)/256)) * 6
	blue := int(6*((math.Sin(freq*i+4*math.Pi/3)*127+128)/256)) * 1

	return 16 + red + green + blue
}
