package main

import (
	"code.google.com/p/go.net/html"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var splitClassPtn = regexp.MustCompile(`[,\s]+`)
var splitAttrPtn = regexp.MustCompile(`\s+`)

func Find(root *html.Node, path string) (node *html.Node, ok bool) {
	tokens := strings.Split(path, "/")
	node, ok = findHelper(tokens, root)
	return
}

func findHelper(tokens []string, parent *html.Node) (*html.Node, bool) {
	index := -1
	element := tokens[0]
	class := ""
	attrs := make([]*html.Attribute, 0, 0)

	if s := strings.Split(element, "*"); len(s) == 2 {
		var err error
		index, err = strconv.Atoi(s[0])
		if err != nil {
			index = -1
		}
		element = s[1]
	}
	if e, a := attrParser(element); len(a) > 0 {
		element = e
		attrs = a
	}
	if s := strings.Split(element, "."); len(s) == 2 {
		element = s[0]
		class = s[1]
	}

	matchElementCount := 0
	fmt.Printf("")
	//	fmt.Println(element)
	//	fmt.Println(class)
	//	fmt.Println(attrs)
	for c := parent.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			continue
		}

		if c.Data == element || element == "" {
			attrMatch := false
			classMatch := false
			attrMatchedCount := 0
			if len(attrs) == 0 {
				attrMatch = true
			}
			if class == "" {
				classMatch = true
			}
			for _, attr := range c.Attr {
				if classMatch && attrMatch {
					break
				}
				if !classMatch && attr.Key == "class" {
					classes := splitClassPtn.Split(attr.Val, -1)
					for _, oneClass := range classes {
						if oneClass == class {
							classMatch = true
							break
						}
					}
				}
				if !attrMatch && attr.Key != "class" {
					for _, a := range attrs {
						if a.Key == attr.Key {
							ptnVals := splitAttrPtn.Split(a.Val, -1)
							htmVals := splitAttrPtn.Split(attr.Val, -1)
							if len(ptnVals) > len(htmVals) {
								continue
							}
							attrsMatchCheck := make([]int, len(ptnVals))
							for idx, i := range ptnVals {
								for _, j := range htmVals {
									if i == j {
										attrsMatchCheck[idx] = 1
										if sum(attrsMatchCheck) == len(ptnVals) {
											break
										}
									}
								}
							}
							if sum(attrsMatchCheck) == len(ptnVals) {
								attrMatchedCount++
							}
						}
					}
				}
			}
			if !attrMatch && attrMatchedCount >= len(attrs) {
				attrMatch = true
			}
			if attrMatch && classMatch {
				matchElementCount++
				if index != -1 && index != matchElementCount {
					continue
				}
				if len(tokens) == 1 {
					return c, true
				} else {
					node, ok := findHelper(tokens[1:len(tokens)], c)
					if ok {
						return node, true
					}
				}
			}
		}
	}
	return nil, false
}

func sum(a []int) int {
	var total = 0
	for _, i := range a {
		total += i
	}
	return total
}
