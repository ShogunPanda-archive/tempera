/*
 * This file is part of tempera. Copyright (C) 2018 and above Shogun <shogun@cowtech.it>.
 * Licensed under the MIT license, which can be found at https://choosealicense.com/licenses/mit.
 */

package tempera

// ANSICode represents a ANSI code with opening and closing (resetting) code
type ANSICode struct {
	Open  int
	Close int
}

// ANSIForegroundCodes holds general code for foreground color management
var ANSIForegroundCodes = ANSICode{38, 39}

// ANSIBackgroundCodes holds general code for foreground color management
var ANSIBackgroundCodes = ANSICode{48, 49}

// ANSICodes is the list of ANSI Escape codes, taken from https://github.com/chalk/ansi-styles
var ANSICodes = map[string]ANSICode{
	"reset": {0, 0},
	// Style
	"bold":          {1, 22}, // 21 isn't widely supported and 22 does the same thing
	"dim":           {2, 22},
	"italic":        {3, 23},
	"underline":     {4, 24},
	"inverse":       {7, 27},
	"hidden":        {8, 28},
	"strikethrough": {9, 29},
	// Foreground colors
	"black":   {30, 39},
	"red":     {31, 39},
	"green":   {32, 39},
	"yellow":  {33, 39},
	"blue":    {34, 39},
	"magenta": {35, 39},
	"cyan":    {36, 39},
	"white":   {37, 39},
	"gray":    {90, 39},
	// Background colors
	"bgBlack":   {40, 49},
	"bgRed":     {41, 49},
	"bgGreen":   {42, 49},
	"bgYellow":  {43, 49},
	"bgBlue":    {44, 49},
	"bgMagenta": {45, 49},
	"bgCyan":    {46, 49},
	"bgWhite":   {47, 49},
	// Bright foreground colors
	"redBright":     {91, 39},
	"greenBright":   {92, 39},
	"yellowBright":  {93, 39},
	"blueBright":    {94, 39},
	"magentaBright": {95, 39},
	"cyanBright":    {96, 39},
	"whiteBright":   {97, 39},
	// Bright background colors
	"bgBlackBright":   {100, 49},
	"bgRedBright":     {101, 49},
	"bgGreenBright":   {102, 49},
	"bgYellowBright":  {103, 49},
	"bgBlueBright":    {104, 49},
	"bgMagentaBright": {105, 49},
	"bgCyanBright":    {106, 49},
	"bgWhiteBright":   {107, 49},
}
