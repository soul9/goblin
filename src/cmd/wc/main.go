package main

import (
    "os"
    "fmt"
    "flag"
)

func usage() {
	fmt.Fprint(os.Stderr, "usage: wc [-l|-c|-w] file\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
    c = flag.Bool("c", true, "Count characters in a file")
    w = flag.Bool("w", true, "Count words in a file")
    l = flag.Bool("l", true, "Count lines in a file")
    h = flag.Bool("h", false, "Display usage")
    err os.Error
    fd *os.File
)

func main() {
    flag.Usage = usage
    flag.Parse()
    fname := flag.Arg(0)
    if *h {
        usage()
        os.Exit(2)
    }
    if fname == "" {
        fname = ""
        fd = os.Stdin
    } else {
        fd, err = os.Open(fname, os.O_RDONLY, 0)
    }
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error opening input file: "+err.String())
        os.Exit(1)
    }
    defer fd.Close()
    var ccount, wcount, lcount int64
    ch := make(chan int64)
    che:=make(chan os.Error)
    go Count(fd, ch, che)
    ccount=<-ch
    wcount=<-ch
    lcount=<-ch
    err = <-che
    if *l {
        fmt.Printf("\t%d", lcount)
    }
    if *w {
        fmt.Printf("\t%d", wcount)
    }
    if *c {
        fmt.Printf("\t%d", ccount)
    }
    fmt.Printf("\t%s\n", fname)
    os.Exit(0)
}
