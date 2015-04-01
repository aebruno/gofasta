// Copyright 2015 Andrew E. Bruno. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/aebruno/gofasta"
	"fmt"
	"os"
	"log"
	"flag"
)

var fastaFile = flag.String("fasta", "", "path to FASTA file")
var count = flag.Bool("count", false, "count sequences")
var idOnly = flag.Bool("id", false, "output ids")
var seqOnly = flag.Bool("seq", false, "output sequences")

func main() {
    flag.Parse()

    if *fastaFile == "" {
        log.Fatal("Please provide a FASTA file. See fastcat --help")
    }

    fastaFile, err := os.Open(*fastaFile)
    if err != nil {
        panic(fmt.Errorf("Failed to open fasta file: %s", err))
    }
    defer fastaFile.Close()

    i := 0
    for rec := range gofasta.SimpleParser(fastaFile) {
        if *idOnly {
            fmt.Printf("%s\n", rec.Id)
        } else if *seqOnly {
            fmt.Printf("%s\n", rec.Seq)
        } else if ! *count {
            fmt.Printf("%s\t%s\n", rec.Id, rec.Seq)
        }
        i++
    }

    if *count {
        fmt.Printf("%d\n", i)
    }
}
