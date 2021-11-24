package cmd

import (
	"fmt"
	"io"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Color string

const (
	Reset                   Color = "\x1b[0000m"
	Bright                  Color = "\x1b[0001m"
	BlackText               Color = "\x1b[0030m"
	RedText                 Color = "\x1b[0031m"
	GreenText               Color = "\x1b[0032m"
	YellowText              Color = "\x1b[0033m"
	BlueText                Color = "\x1b[0034m"
	MagentaText             Color = "\x1b[0035m"
	CyanText                Color = "\x1b[0036m"
	WhiteText               Color = "\x1b[0037m"
	DefaultText             Color = "\x1b[0039m"
	BrightRedText           Color = "\x1b[1;31m"
	BrightGreenText         Color = "\x1b[1;32m"
	BrightYellowText        Color = "\x1b[1;33m"
	BrightBlueText          Color = "\x1b[1;34m"
	BrightMagentaText       Color = "\x1b[1;35m"
	BrightCyanText          Color = "\x1b[1;36m"
	BrightWhiteText         Color = "\x1b[1;37m"
	BlackBackground         Color = "\x1b[0040m"
	RedBackground           Color = "\x1b[0041m"
	GreenBackground         Color = "\x1b[0042m"
	YellowBackground        Color = "\x1b[0043m"
	BlueBackground          Color = "\x1b[0044m"
	MagentaBackground       Color = "\x1b[0045m"
	CyanBackground          Color = "\x1b[0046m"
	WhiteBackground         Color = "\x1b[0047m"
	BrightBlackBackground   Color = "\x1b[0100m"
	BrightRedBackground     Color = "\x1b[0101m"
	BrightGreenBackground   Color = "\x1b[0102m"
	BrightYellowBackground  Color = "\x1b[0103m"
	BrightBlueBackground    Color = "\x1b[0104m"
	BrightMagentaBackground Color = "\x1b[0105m"
	BrightCyanBackground    Color = "\x1b[0106m"
	BrightWhiteBackground   Color = "\x1b[0107m"
)

var (
	timeFormat = "2006-01-02 15:04:05.000 -0700 MST"
)

func timeConvert(t *time.Time) time.Time {
	if viper.Get("tz") != nil {
		tz, err := time.LoadLocation((viper.Get("tz").(string)))
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Fatal("failed Can't time zone")
		}
		return t.In(tz)
	}
	return *t
}

func Paint(color Color, value string) string {
	return fmt.Sprintf("%v%v%v", color, value, Reset)
}
func PaintRow(colors []Color, row []string) []string {
	paintedRow := make([]string, len(row))
	for i, v := range row {
		paintedRow[i] = Paint(colors[i], v)
	}
	return paintedRow
}
func PrintRow(writer io.Writer, line []string) {
	fmt.Fprintln(writer, strings.Join(line, "\t"))
}
