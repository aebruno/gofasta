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
                continue
            }

            seqbuf.WriteString(strings.TrimSpace(line))
        }

        c <- &SeqRecord{Id: strings.TrimSpace(id), Seq: seqbuf.String()}
    }();

    return c
}

/*
func SimpleParser(file *os.File) chan *SeqRecord {
    c := make(chan *SeqRecord)

    go func() {
        scanner := bufio.NewScanner(file)
        line := ""

        for ;; {
            if scanner.Scan() {
                line = scanner.Text()
                if line == "" {
                    continue
                } else if line[0] == '>' {
                    break
                }
            } else {
                close(c)
                return
            }
        }

        for ;; {
            if line[0] != '>' {
                break
            }

            id := line[1:]

            if scanner.Scan() {
                line = scanner.Text()
            } else {
                break
            }

            seq := ""
            for ;; {
                if len(line) == 0 {
                    break
                }
                if line[0] == '>' {
                    break
                }

                seq = seq + strings.TrimSpace(line)
                if scanner.Scan() {
                    line = scanner.Text()
                } else {
                    break
                }
            }

            c <- &SeqRecord{Id: id, Seq: seq}
        }

        close(c)
    }();

    return c
}
*/
