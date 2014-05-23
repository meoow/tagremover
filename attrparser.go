package main

import "code.google.com/p/go.net/html"
import "regexp"
import "strings"

var squareBracElement = regexp.MustCompile(`^([^\[]*)\[(.+)\]\s*$`)
var stripQuote = regexp.MustCompile(`^['"](.*)['"]$`)

func attrParser(str string) (string, []*html.Attribute) {
	var result = squareBracElement.FindStringSubmatch(str)
	var class = ""
	attrs := make([]*html.Attribute, 0, 3)
	if len(result) == 3 {
		class = result[1]
		for _, a := range strings.Split(result[2], ",") {
			kv := strings.Split(a, "=")
			if len(kv) == 2 {
				key := strings.TrimSpace(kv[0])
				if key == "class" {
					continue
				}
				val := strings.TrimSpace(kv[1])
				if len(val) >= 2 {
					if (val[0] == '\'' || val[0] == '"') && val[len(val)-1] == val[0] {
						val = stripQuote.FindStringSubmatch(val)[1]
					}
				}
				attrs = append(attrs, &html.Attribute{"", key, val})
			}
		}
	}
	return strings.TrimSpace(class), attrs
}
