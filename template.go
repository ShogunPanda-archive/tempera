/*
 * This file is part of tempera. Copyright (C) 2018 and above Shogun <shogun@cowtech.it>.
 * Licensed under the MIT license, which can be found at https://choosealicense.com/licenses/mit.
 */

package tempera

import (
	"regexp"
	"sort"
	"strings"
)

var parser = regexp.MustCompile("(?i)\\{(?:\\{?)[^\\{\\}]+?\\}")
var splitter = regexp.MustCompile("[\\|\\s]+")

func addStyles(replacement *string, tagStyles *[]string, id string) bool {
	open, _ := styleToANSI(id)

	if open == "" { // Validate the style
		return false
	}

	// Push to the the list of the styles of this tag
	*tagStyles = append(*tagStyles, id)

	// Set the style by applying the opening ANSI code
	*replacement += open

	return true
}

func removeStyles(replacement *string, appliedStyles *[][]string) {
	// Pop a the style from the list of the applied ones
	totalStyles := len(*appliedStyles)
	stylesToRemove := (*appliedStyles)[totalStyles-1]
	*appliedStyles = (*appliedStyles)[:totalStyles-1]
	totalStyles--

	for _, id := range stylesToRemove { // Unset the style by applying the closing ANSI codes
		_, close := styleToANSI(id)

		*replacement += close
	}

	if len(*appliedStyles) > 0 { // If the applied styles stack is still non-empty, it means we have to restore the previous style
		for _, id := range (*appliedStyles)[totalStyles-1] {
			open, _ := styleToANSI(id)
			*replacement += open
		}
	}
}

// ColorizeTemplate add colors to a template string
func ColorizeTemplate(template string) string {
	// Create a new styles stack
	appliedStyles := make([][]string, 0)

	// For each tag in the string
	stylesInserted := false
	modified := parser.ReplaceAllStringFunc(template, func(match string) string {
		var replacement string

		if strings.HasPrefix(match, "{{") {
			return match[1:]
		}

		ids := splitter.Split(strings.Trim(match, "{}"), -1)
		tagStyles := make([]string, 0) // Maintain a list of all styles applied in this tag

	TagStyles:
		for _, id := range resolveCustomStyles(ids) { // For each id in this tag
			id = strings.TrimSpace(id)
			switch id {
			case "-": // We are removing a style from the text
				if len(appliedStyles) > 0 { // If there is a previously applied tag
					removeStyles(&replacement, &appliedStyles)
				}

				break TagStyles // Do no process further styles in this tag
			case "reset": // Reset, it means forget all the applied styles so far
				appliedStyles = make([][]string, 0)
				break TagStyles // Do no process further styles in this tag
			default: // We are adding a new style
				if addStyles(&replacement, &tagStyles, id) {
					stylesInserted = true
				}
			}
		}

		// If we added any style in this tag, add to the applied styles stack (in reverse order in order to guarantee proper closing)
		if len(tagStyles) > 0 {
			sort.Sort(sort.Reverse(sort.StringSlice(tagStyles)))
			appliedStyles = append(appliedStyles, tagStyles)
		}

		return replacement
	})

	if stylesInserted { // Always apply the global reset at the end if we apply any style. This is to leave a clean state
		modified += escapeANSI(ANSICodes["reset"].Close)
	}

	return modified
}

// CleanTemplate removes all style tag from a template
func CleanTemplate(template string) string {
	return parser.ReplaceAllStringFunc(template, func(match string) string {
		if strings.HasPrefix(match, "{{") {
			return match[1:]
		}

		return ""
	})
}
