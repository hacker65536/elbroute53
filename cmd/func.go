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
	Bright                        = "\x1b[0001m"
	BlackText                     = "\x1b[0030m"
	RedText                       = "\x1b[0031m"
	GreenText                     = "\x1b[0032m"
	YellowText                    = "\x1b[0033m"
	BlueText                      = "\x1b[0034m"
	MagentaText                   = "\x1b[0035m"
	CyanText                      = "\x1b[0036m"
	WhiteText                     = "\x1b[0037m"
	DefaultText                   = "\x1b[0039m"
	BrightRedText                 = "\x1b[1;31m"
	BrightGreenText               = "\x1b[1;32m"
	BrightYellowText              = "\x1b[1;33m"
	BrightBlueText                = "\x1b[1;34m"
	BrightMagentaText             = "\x1b[1;35m"
	BrightCyanText                = "\x1b[1;36m"
	BrightWhiteText               = "\x1b[1;37m"
	BlackBackground               = "\x1b[0040m"
	RedBackground                 = "\x1b[0041m"
	GreenBackground               = "\x1b[0042m"
	YellowBackground              = "\x1b[0043m"
	BlueBackground                = "\x1b[0044m"
	MagentaBackground             = "\x1b[0045m"
	CyanBackground                = "\x1b[0046m"
	WhiteBackground               = "\x1b[0047m"
	BrightBlackBackground         = "\x1b[0100m"
	BrightRedBackground           = "\x1b[0101m"
	BrightGreenBackground         = "\x1b[0102m"
	BrightYellowBackground        = "\x1b[0103m"
	BrightBlueBackground          = "\x1b[0104m"
	BrightMagentaBackground       = "\x1b[0105m"
	BrightCyanBackground          = "\x1b[0106m"
	BrightWhiteBackground         = "\x1b[0107m"
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
