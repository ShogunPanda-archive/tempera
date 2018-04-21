/*
 * This file is part of tempera. Copyright (C) 2018 and above Shogun <shogun@cowtech.it>.
 * Licensed under the MIT license, which can be found at https://choosealicense.com/licenses/mit.
 */

package tempera

import (
	"errors"
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)

	g.Describe("Colorize", func() {
		g.It("Applies known styles", func() {
			g.Assert(Colorize("ABC", "bgBlack", "red")).Equal("\x1b[40m\x1b[31mABC\x1b[39m\x1b[49m")
		})

		g.It("Ignores known styles", func() {
			g.Assert(Colorize("ABC", "whatever", "red")).Equal("\x1b[31mABC\x1b[39m")
		})

		g.It("Supports ANSI 256 colors, ignoring invalid colors", func() {
			g.Assert(Colorize("ABC", "ANSI:232")).Equal("\x1b[38;5;232mABC\x1b[39m")
			g.Assert(Colorize("ABC", "bgANSI:333")).Equal("ABC")
			g.Assert(Colorize("ABC", "bgansi:2,4,0")).Equal("\x1b[48;5;112mABC\x1b[49m")
			g.Assert(Colorize("ABC", "ANSI:2,4,6")).Equal("ABC")
		})

		g.It("Supports ANSI 16m RGB colors, ignoring invalid colors", func() {
			g.Assert(Colorize("ABC", "rgb:255,232,0")).Equal("\x1b[38;2;255;232;0mABC\x1b[39m")
			g.Assert(Colorize("ABC", "bgRGB:33,66,99")).Equal("\x1b[48;2;33;66;99mABC\x1b[49m")
			g.Assert(Colorize("ABC", "bgRGB:999,999,999")).Equal("ABC")
			g.Assert(Colorize("ABC", "bgRGB:1,999,999")).Equal("ABC")
			g.Assert(Colorize("ABC", "bgRGB:1,2,999")).Equal("ABC")
		})

		g.It("Supports ANSI 16m HEX colors, ignoring invalid colors", func() {
			g.Assert(Colorize("ABC", "hex:F0d030")).Equal("\x1b[38;2;240;208;48mABC\x1b[39m")
			g.Assert(Colorize("ABC", "bgHEX:0099FF")).Equal("\x1b[48;2;0;153;255mABC\x1b[49m")
			g.Assert(Colorize("ABC", "bgHEX:0099GG")).Equal("ABC")
		})
	})

	g.Describe("ColorizeTemplate", func() {
		g.It("Applies known styles and closes them in the right order", func() {
			g.Assert(ColorizeTemplate("{red}ABC{green}CDE{-}EFG{-}HIJ")).Equal("\x1b[31mABC\x1b[32mCDE\x1b[39m\x1b[31mEFG\x1b[39mHIJ\x1b[0m")
		})

		g.It("Ignores unknown styles", func() {
			g.Assert(ColorizeTemplate("{red}ABC{yolla}CDE{-}EFG{-}HIJ")).Equal("\x1b[31mABCCDE\x1b[39mEFGHIJ\x1b[0m")
		})

		g.It("Ignores unbalanced parenthesis", func() {
			g.Assert(ColorizeTemplate("{red}}ABC{-}")).Equal("\x1b[31m}ABC\x1b[39m\x1b[0m")
		})

		g.It("Ignores unbalanced tags", func() {
			g.Assert(ColorizeTemplate("{red}ABC")).Equal("\x1b[31mABC\x1b[0m")
		})

		g.It("Double curly braces are respected", func() {
			g.Assert(ColorizeTemplate("{{red}")).Equal("{red}")
		})

		g.It("Closing tag ignores further specs", func() {
			g.Assert(ColorizeTemplate("{red}ABC{green}CDE{- yellow}EFG{-}HIJ")).Equal("\x1b[31mABC\x1b[32mCDE\x1b[39m\x1b[31mEFG\x1b[39mHIJ\x1b[0m")
		})

		g.It("Reset tag cleans the stack", func() {
			g.Assert(ColorizeTemplate("{red}ABC{green}CDE{reset red}EFG{-}HIJ")).Equal("\x1b[31mABC\x1b[32mCDEEFGHIJ\x1b[0m")
		})

		g.It("Supports ANSI, RGB and HEX colors", func() {
			g.Assert(ColorizeTemplate("{ANSI:5,0,0}ABC{RGB:0,255,0}CDE{bgHEX:#0000FF}EFG")).
				Equal("\x1b[38;5;196mABC\x1b[38;2;0;255;0mCDE\x1b[48;2;0;0;255mEFG\x1b[0m")
		})
	})

	g.Describe("CleanTemplate", func() {
		g.It("Removes style tags from a template", func() {
			g.Assert(CleanTemplate("{red}ABC{green}CDE{-}EFG{-}HIJ")).Equal("ABCCDEEFGHIJ")
			g.Assert(CleanTemplate("{red}ABC{yolla}CDE{-}EFG{-}HIJ")).Equal("ABCCDEEFGHIJ")
			g.Assert(CleanTemplate("{red}}ABC{-}")).Equal("}ABC")
			g.Assert(CleanTemplate("{red}ABC")).Equal("ABC")
			g.Assert(CleanTemplate("{{red}")).Equal("{red}")
			g.Assert(CleanTemplate("{red}ABC{green}CDE{- yellow}EFG{-}HIJ")).Equal("ABCCDEEFGHIJ")
			g.Assert(CleanTemplate("{red}ABC{green}CDE{reset red}EFG{-}HIJ")).Equal("ABCCDEEFGHIJ")
			g.Assert(CleanTemplate("{ANSI:5,0,0}ABC{RGB:0,255,0}CDE{bgHEX:#0000FF}EFG")).Equal("ABCCDEEFG")
		})
	})

	g.Describe("AddCustomStyle / DeleteCustomStyles", func() {
		g.It("Allow to define custom styles, supported both by Colorize and ColorizeTemplate", func() {
			g.Assert(Colorize("ABC", "customRed@@")).Equal("ABC")
			g.Assert(ColorizeTemplate("{customRed@@ green}ABC{-}")).Equal("\x1b[32mABC\x1b[39m\x1b[0m")

			AddCustomStyle("customRed@@", "red", "underline")

			g.Assert(Colorize("ABC", "customRed@@")).Equal("\x1b[31m\x1b[4mABC\x1b[24m\x1b[39m")
			g.Assert(ColorizeTemplate("{customRed@@ green}ABC{-}")).Equal("\x1b[31m\x1b[4m\x1b[32mABC\x1b[24m\x1b[39m\x1b[39m\x1b[0m")

			DeleteCustomStyles("customRed@@")

			g.Assert(Colorize("ABC", "customRed@@")).Equal("ABC")
			g.Assert(ColorizeTemplate("{customRed@@ green}ABC{-}")).Equal("\x1b[32mABC\x1b[39m\x1b[0m")
		})

		g.It("Should reject custom styles name which contain spaces or curly brace", func() {
			for _, s := range []string{"{invalid", "invalid}", "no spaces"} {
				g.Assert(AddCustomStyle(s, "red")).Equal(errors.New("The custom style name should not contain spaces or curly braces"))
			}
		})
	})
}
