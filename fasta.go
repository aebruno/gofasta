package gofasta

import (
    "strings"
    "os"
    "bufio"
    "bytes"
)

type SeqRecord struct {
    Id  string
    Seq string
}

func SimpleParser(file *os.File) chan *SeqRecord {
    c := make(chan *SeqRecord)

    go func() {
        defer close(c)

        reader := bufio.NewReader(file)

        // skip bytes until the first record
        _, err := reader.ReadBytes('>')
        if err != nil {
            return
        }

        id, err := reader.ReadString('\n')
        if err != nil {
            return
        }

        var seqbuf bytes.Buffer
        for ;; {
            line, err := reader.ReadString('\n')
            if err != nil || line == "" {
                break
            }
            if line[0] == '>' {
                c <- &SeqRecord{Id: strings.TrimSpace(id), Seq: seqbuf.String()}
                id = line[1:]
                seqbuf.Reset()
                continue
            }

            seqbuf.WriteString(strings.TrimSpace(line))
        }

        c <- &SeqRecord{Id: strings.TrimSpace(id), Seq: seqbuf.String()}
    }();

    return c
}
