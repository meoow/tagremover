package main

import "testing"

//import "code.google.com/p/go.net/html"

func Test_arrtParser_1(t *testing.T) {
	var sample = `div[id=tnav]`
	e, a := attrParser(sample)
	t.Logf("Element: %s\n", e)
	t.Logf("Attrs: %#V\n", a)
	if e == "div" {
		if a[0].Key == "id" && a[0].Val == "tnav" {
			return
		}
	}
	t.Error("Failed")
}

func Test_arrtParser_2(t *testing.T) {
	var sample = `p[name = tnav, code = "classone"]`
	e, a := attrParser(sample)
	t.Logf("Element: %s\n", e)
	t.Logf("Attrs: %#V\n", a)
	if e == "p" {
		if a[0].Key == "name" && a[0].Val == "tnav" {
			if a[1].Key == "code" && a[1].Val == "classone" {
				return
			}
		}
	}
	t.Error("Failed")
}
