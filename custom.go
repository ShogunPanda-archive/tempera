/*
 * This file is part of tempera. Copyright (C) 2018 and above Shogun <shogun@cowtech.it>.
 * Licensed under the MIT license, which can be found at https://choosealicense.com/licenses/mit.
 */

package tempera

import (
	"errors"
	"regexp"
)

var customStyles = make(map[string][]string)
var invalidStyleMatcher = regexp.MustCompile("(?i)[\\s\\{\\}]")

func resolveCustomStyles(styles []string) []string {
	var resolvedStyles = make([]string, 0)

	for _, style := range styles {
		cs, exists := customStyles[style]

		if !exists { // No custom styles, add itself
			resolvedStyles = append(resolvedStyles, style)
		} else {
			resolvedStyles = append(resolvedStyles, cs...)
		}
	}

	return resolvedStyles
}

// AddCustomStyle add a new custom style
func AddCustomStyle(name string, styles ...string) error {
	if invalidStyleMatcher.MatchString(name) {
		return errors.New("The custom style name should not contain spaces or curly braces")
	}

	customStyles[name] = styles

	return nil
}

// DeleteCustomStyles removes one or more custom styles
func DeleteCustomStyles(names ...string) {
	for _, n := range names {
		delete(customStyles, n)
	}
}
