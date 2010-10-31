package goblin

import (
  "os"
  "bufio"
)

func WC(f *os.File, ch chan <-int64, che chan <-os.Error) {
    var wcount, ccount, lcount int64
    //BUG: initializing wcount to 1 because first word isn't counted otherwise
    wcount, ccount, lcount = 1, 0, 0
    buf := bufio.NewReader(f)
    var err os.Error
    for c, _, err := buf.ReadRune(); err == nil; c, _, err = buf.ReadRune() {
        switch c {
            case '\t':
                next, _ := buf.Peek(1)
                if next[0] != ' ' && next[0] != '\t' && next[0] != '\n' {
                    wcount++
                }
                ccount++
            case ' ':
                next, _ := buf.Peek(1)
                if next[0] != ' ' && next[0] != '\t' && next[0] != '\n' {
                    wcount++
                }
                ccount++
            case '\n':
                next, err := buf.Peek(1)
                if err == nil && next[0] != ' ' && next[0] != '\t' && next[0] != '\n' {
                    wcount++
                }
                lcount++
                ccount++
            default:
                ccount++
        }
    }
    che<- err
    ch<- lcount
    ch<- wcount
    ch<-ccount
}
