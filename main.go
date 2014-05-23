package main

import "code.google.com/p/go.net/html"
import "fmt"
import "os"
import "path/filepath"
import "log"
import "io/ioutil"

func main() {
	fmt.Printf("")
	var inplace = false
	var argsI = 1
	var found = 0
	var tempfile *os.File
	defer func() {
		if tempfile != nil {
			if _, err := tempfile.Stat(); err == nil {
				tempfile.Close()
				os.Remove(tempfile.Name())
			}
		}
	}()
	if os.Args[1] == "-i" {
		inplace = true
		argsI = 2
	}
	htmlfile := os.Args[argsI]
	dir := filepath.Dir(htmlfile)
	hf, err := os.Open(htmlfile)
	if err != nil {
		log.Fatal(err)
	}
	n, e := html.Parse(hf)
	if e != nil {
		log.Fatal(e)
	}
	hf.Close()
	defer hf.Close()
	for _, i := range os.Args[argsI+1:] {
		f, ok := Find(n, i)
		if ok {
			if f.Parent != nil {
				f.Parent.RemoveChild(f)
				os.Stderr.WriteString(fmt.Sprintf("FOUND: %s\n", i))
				found++
			}
		}
	}
	if !inplace {
		html.Render(os.Stdout, n)
	} else if found > 0 {
		tempfile, err := ioutil.TempFile(dir, "htmlcleaner")
		if err != nil {
			log.Fatal(err)
		}
		err = html.Render(tempfile, n)
		if err != nil {
			log.Fatal(err)
		}
		tempfile.Close()
		err = os.Rename(tempfile.Name(), htmlfile)
		if err != nil {
			log.Fatal(err)
		}
	}
}
