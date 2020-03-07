package peripherals

import (
    "io"
    "log"
)

type Buzzer struct {
    OutputWriter
}

func (b *Buzzer) Write(message string) {
    b.OutputWriter.Write(message)
}

func (b *Buzzer) Key() {
    log.Print("key beep")
    b.Write("440:50")
}

func (b *Buzzer) Cancel() {
    log.Print("cancel beep")
    b.Write("600:500")
}

func (b *Buzzer) Error() {
    log.Print("error beep")
    b.Write("100:500")
}

func (b *Buzzer) Success() {
    log.Print("success beep")
    b.Write("400:400")
}

func (b *Buzzer) Read() {
    log.Print("read beep")
    b.Write("440:200")
}

func NewBuzzer(writer io.Writer) Buzzer {
    return Buzzer{
        OutputWriter: NewOutputWriter(writer),
    }
}
