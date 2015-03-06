package gofasta

import (
    "os"
    "strings"
    "strconv"
    "fmt"
    "bufio"
    "io"
    "unicode"
)

type Faidx struct {
    fastaFile    *os.File
    IndexCache   map[string][]int64
}

func NewFaidx(fileName string) (*Faidx,error) {
    fastaFile, err := os.Open(fileName)
    if err != nil {
        return nil, fmt.Errorf("Failed to open fasta file: %s", err)
    }
    indexFile, err := os.Open(fileName+".fai")
    if err != nil {
        return nil, fmt.Errorf("Failed to open fasta index (.fai) file: %s", err)
    }
    defer indexFile.Close()

    cache := make(map[string][]int64)
    scanner := bufio.NewScanner(indexFile)
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Split(line, "\t")
        chrom := fields[0]
        data := make([]int64, 4)
        for i,v := range fields[1:] {
            data[i],_ = strconv.ParseInt(v, 10, 64)
        }
        cache[chrom] = data
    }

    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("Failed to read fasta index file: %s", err)
    }

    return &Faidx{fastaFile: fastaFile, IndexCache: cache}, nil
}

func (f *Faidx) Fetch(chr string, start, end int) ([]byte,error) {
    chrom,exists := f.IndexCache[chr]
    if !exists {
        return nil, fmt.Errorf("chrom not found in fasta file: %s", chr)
    }

    slen,offset,blen,bytelen := chrom[0],chrom[1],chrom[2],chrom[3]

    start = start-1
    if start < 0 {
        return nil, fmt.Errorf("start location out of bounds: %s %d-%d", chr,start+1, end)
    } else if start >= end {
        return nil, fmt.Errorf("Invalid start/end location: %s %d-%d", chr, start+1, end)
    } else if int64(end) > slen {
        return nil, fmt.Errorf("end location out of bounds: %s %d-%d", chr, start+1, end)
    }

    f.fastaFile.Seek(offset+int64(start)/blen*bytelen+int64(start)%blen, 0)

    seq := make([]byte, end-start)
    bufSize := 4096
    if end-start < bufSize {
        bufSize = end-start
    }

    buf := make([]byte, bufSize)
    index := 0
    for {
        n,err := f.fastaFile.Read(buf)
        if err != nil && err != io.EOF {
            return nil, fmt.Errorf("I/O error reading file: %s", err)
        }
        if n == 0 || index >= end-start {
            break
        }

        for _,b := range buf {
            if unicode.IsPrint(rune(b)) {
                seq[index] = b
                index++
                if index >= end-start {
                    break
                }
            }
        }
    }

    return seq, nil
}

func (f *Faidx) Close() (error) {
    return f.fastaFile.Close()
}
