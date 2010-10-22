package main

import (
    "os"
    "fmt"
    "flag"
    "wc"
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
    ch := make(chan int64)
    che:=make(chan os.Error)
    go wc.WC(fd, ch, che)
    err = <-che
    if err != nil {
        fmt.Fprint(os.Stderr, "Error occured: "+err.String())
    }
    if *l {
        fmt.Printf("\t%d", <-ch)
    } else {
        _ = <-ch
    }
    if *w {
        fmt.Printf("\t%d", <-ch)
    } else {
        _ = <-ch
    }
    if *c {
        fmt.Printf("\t%d", <-ch)
    } else {
        _ = <-ch
    }
    fmt.Printf("\t%s\n", fname)
    os.Exit(0)
}
