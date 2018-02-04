# tempera

[![Release](https://img.shields.io/github/release/ShogunPanda/tempera.svg)](https://github.com/ShogunPanda/tempera/releases/latest)
[![GoDoc](https://godoc.org/github.com/ShogunPanda/tempera?status.svg)](https://godoc.org/github.com/ShogunPanda/tempera)
[![Go Report Card](https://goreportcard.com/badge/github.com/ShogunPanda/tempera)](https://goreportcard.com/report/github.com/ShogunPanda/tempera)
[![Build Status](https://img.shields.io/travis/rust-lang/rust/master.svg)](https://travis-ci.org/ShogunPanda/tempera)
[![License](https://img.shields.io/github/license/ShogunPanda/tempera.svg)](https://github.com/ShogunPanda/tempera/blob/master/LICENSE.md)

Template based terminal coloring made really easy.

http://sw.cowtech.it/tempera

## Usage

Tempera allows to add coloring to terminal in a really easy way.

### Colorize strings

To colorize strings, simply use the [Colorize](https://godoc.org/github.com/ShogunPanda/tempera#Colorize) function, passing a list of styles you want to apply.
The list of supported color names correspondes to the keys of [ANSICodes](https://godoc.org/github.com/ShogunPanda/tempera#pkg-variables) variable.
```go

import "github.com/ShogunPanda/tempera"

inRed := tempera.Colorize("Colorized", "red")
inRedWithBlueBackground := tempera.Colorize("Colorized", "red bgBlue")
```

### Colorize templates

To colorize a template using a tagged template syntax, simply use the [ColorizeTemplate](https://godoc.org/github.com/ShogunPanda/tempera#ColorizeTemplate) function.

```go

import "github.com/ShogunPanda/tempera"

colored := tempera.ColorizeTemplate("{red}This is in red, {green underline}this in green underlined{-}, this in red again.")
```

The template recognizes styles between curly braces (use a double opening brace to escape them) and the token `{-}` as universal closing tag (which also restores the previous style).

The closing tag at the end of the string can be omitted, since tempera will append the global reset style (`\x1b[0m`) if any style was set.

If you want to discard styles to be restored, use the `{reset}` token.

### Setting custom styles

If you want to define custom styles, use the [AddCustomStyle](https://godoc.org/github.com/ShogunPanda/tempera#AddCustomStyle) function.

```go
import "github.com/ShogunPanda/tempera"

tempera.AddCustomStyle("important", "red underline")
colored := tempera.ColorizeTemplate("This is in red, underlined.", "important")
```

### 256 and 16 millions colors support

tempera supports 256 ANSI codes and 16m RGB colors. Just give it a try:

```go
import "github.com/ShogunPanda/tempera"

fgAnsi256 := tempera.Colorize("color me", "ansi:100")
bgAnsi256 := tempera.Colorize("color me", "bgANSI:3,4,5")

fgRgb := tempera.Colorize("color me", "rgb:255,0,0")
bgRgb := tempera.Colorize("color me", "bgRGB:0,255,0")

fgHex := tempera.Colorize("color me", "hex:#FF0000")
bgHex := tempera.Colorize("color me", "bgHEX:00FF00")
```

ANSI, RGB, and HEX can be used in style definitions and templates as well.

## API Documentation

The API documentation can be found [here](https://godoc.org/github.com/ShogunPanda/tempera).

## Contributing to tempera

* Check out the latest master to make sure the feature hasn't been implemented or the bug hasn't been fixed yet.
* Check out the issue tracker to make sure someone already hasn't requested it and/or contributed it.
* Fork the project.
* Start a feature/bugfix branch.
* Commit and push until you are happy with your contribution.
* Make sure to add tests for it. This is important so I don't break it in a future version unintentionally.

## Copyright

Copyright (C) 2018 and above Shogun (shogun@cowtech.it).

Licensed under the MIT license, which can be found at https://choosealicense.com/licenses/mit.

