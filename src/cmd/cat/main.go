package main

import (
    "os"
    "flag"
    "fmt"
    "cat"
)

func usage() {
	fmt.Fprint(os.Stderr, "usage: cat file1 [file2⋯]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
    h = flag.Bool("h", false, "Display usage")
    files = flag.Args()
    err os.Error
    fd *os.File
)

func main(){
    if len(files) == 0 {
        files = make([]string, 1)
        files[0] = "os.Stdin"
    }
    bch := make(chan []byte)
    errch := make(chan os.Error)
    for i, _ := range files {
        if files[i] == "os.Stdin" {
            fd = os.Stdin
        } else {
            fd, err = os.Open(files[i], os.O_RDONLY, 0)
            if err != nil {
                fmt.Fprintln(os.Stderr, "Error opening input file: " + err.String())
                os.Exit(1)
            }
        }
        go cat.Cat(fd, bch, errch)
        for err = <- errch; err == nil; err = <-errch {
            fmt.Printf("%s", <- bch)
        }
        if err != os.EOF {
            fmt.Fprintln(os.Stderr, "Unexpected End-Of-File: "+ err.String())
        }
        fd.Close()
    }
    os.Exit(0)
}