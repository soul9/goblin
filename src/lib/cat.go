package goblin
import (
    "os"
    "bufio"
)

func Cat(f *os.File, ch chan <- []byte, che chan <- os.Error) {
    buf := bufio.NewReader(f)
    var err os.Error
    var line []byte
    for line, err = buf.ReadSlice('\n'); err == nil; line, err = buf.ReadSlice('\n') {
        che <- err
        ch <- line
    }
    che <- err
}
