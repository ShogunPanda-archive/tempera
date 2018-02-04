/*
 * This file is part of tempera. Copyright (C) 2018 and above Shogun <shogun@cowtech.it>.
 * Licensed under the MIT license, which can be found at https://choosealicense.com/licenses/mit.
 */

package tempera

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var ansiMatcher = regexp.MustCompile("(?i)^(?:bg)?ansi:(\\d{0,3})(?:[,;](\\d)[,;](\\d))?")
var rgbMatcher = regexp.MustCompile("(?i)^(?:bg)?rgb:(\\d{0,3})[,;](\\d{0,3})[,;](\\d{0,3})")
var hexMatcher = regexp.MustCompile("(?i)^(?:bg)?hex:(?:#?)([0-9a-f]{2})([0-9a-f]{2})([0-9a-f]{2})")

func escapeANSI(codes ...int) string {
	return "\x1b[" + strings.Trim(strings.Replace(fmt.Sprint(codes), " ", ";", -1), "[]") + "m"
}

func specToANSICodes(spec string) ANSICode {
	if strings.HasPrefix(spec, "bg") {
		return ANSIBackgroundCodes
	}

	return ANSIForegroundCodes
}

func parseColorComponent(raw string, base int, min, max int64) (dest int64, err error) {
	dest, err = strconv.ParseInt(raw, base, 0)

	if dest < min || dest > max {
		err = errors.New("Invalid color component")
	}

	return
}

func parseColor(raw string, base int, min, max int64) (r, g, b int64, err error) {
	components := strings.Split(raw, ",")

	if r, err = parseColorComponent(components[0], base, min, max); err != nil {
		return
	}

	if g, err = parseColorComponent(components[1], base, min, max); err != nil {
		return
	}

	if b, err = parseColorComponent(components[2], base, min, max); err != nil {
		return
	}

	return
}

func convertANSIColor(spec ...string) (open, close string) {
	codes := specToANSICodes(spec[0])

	if spec[2] == "" { // Simple color code, range 16 to 255
		color, err := parseColorComponent(spec[1], 10, 16, 255)

		if err == nil {
			open = escapeANSI(codes.Open, 5, int(color))
		}
	} else {
		r, g, b, err := parseColor(strings.Join(spec[1:4], ","), 10, 0, 5)

		if err == nil {
			open = escapeANSI(codes.Open, 5, int(16+36*r+6*g+b))
		}
	}

	close = escapeANSI(codes.Close)
	return
}

func convertRGBColor(base int, spec ...string) (open, close string) {
	codes := specToANSICodes(spec[0])

	// Validate RGB
	r, g, b, err := parseColor(strings.Join(spec[1:4], ","), base, 0, 255)

	if err == nil {
		open = escapeANSI(codes.Open, 2, int(r), int(g), int(b))
		close = escapeANSI(codes.Close)
	}

	return
}

func styleToANSI(style string) (open, close string) {
	if spec := ansiMatcher.FindStringSubmatch(style); len(spec) > 1 { // ANSI 256 colors, either as color number or RGB 0-15 spec
		open, close = convertANSIColor(spec...)
	} else if spec := rgbMatcher.FindStringSubmatch(style); len(spec) > 1 { // ANSI RGB 16m colors
		open, close = convertRGBColor(10, spec...)
	} else if spec := hexMatcher.FindStringSubmatch(style); len(spec) > 1 { // ANSI RGB 16 colors in HEX form
		open, close = convertRGBColor(16, spec...)
	} else if codes, exists := ANSICodes[style]; exists { // Styles tag
		open = escapeANSI(codes.Open)
		close = escapeANSI(codes.Close)
	}

	return
}

// Colorize add colors to a string using a specific list of styles, in order
func Colorize(content string, styles ...string) string {
	var header, footer string

	for _, style := range resolveCustomStyles(styles) {
		open, close := styleToANSI(style)

		if open != "" && close != "" {
			header += open
			footer = close + footer
		}
	}

	return header + content + footer
}
