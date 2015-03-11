package main

import (
    "github.com/aebruno/gofasta"
    "fmt"
    "log"
    "os"
    "regexp"
    "strconv"
)

var regionPattern = regexp.MustCompile(`(\w+):(\d+)-(\d+)`)

func main() {
    if len(os.Args) != 3 {
        log.Fatal("usage: faidx [fasta] [chrom:start-end]")
    }

    f, err := gofasta.NewFaidx(os.Args[1])
    if err != nil {
        log.Fatalf("[FATAL] Failed to open fastsa file: %s", err)
    }
    defer f.Close()


    str := os.Args[2]
    matches := regionPattern.FindStringSubmatch(str)
    if len(matches) != 4 {
        log.Fatalf("[FATAL] Invalid region format '%s', must be chr:start-end", str)
    }
    chrom := matches[1]
    start,err := strconv.Atoi(matches[2])
    if err != nil {
        log.Fatalf("Invalid region '%s', start is not an int", str)
    }
    end,err := strconv.Atoi(matches[3])
    if err != nil {
        log.Fatalf("Invalid region '%s', end is not an int", str)
    }

    base,err := f.Fetch(chrom, start, end)
    fmt.Printf(">%s\n%s\n", str, base)
}
