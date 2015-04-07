// Copyright 2015 Andrew E. Bruno. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gofasta

import (
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

        id, err := reader.ReadBytes('\n')
        if err != nil {
            return
        }

        var seqbuf bytes.Buffer
        for ;; {
            line, err := reader.ReadBytes('\n')
            if err != nil || len(line) == 0 {
                break
            }
            if line[0] == '>' {
                c <- &SeqRecord{Id: string(bytes.TrimSpace(id)), Seq: seqbuf.String()}
                id = line[1:]
                seqbuf.Reset()
                continue
            }

            seqbuf.Write(line[:len(line)-1])
        }

        c <- &SeqRecord{Id: string(bytes.TrimSpace(id)), Seq: seqbuf.String()}
    }();

    return c
}
