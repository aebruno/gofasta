===============================================================================
gofasta - Process FASTA files in Go
===============================================================================

-------------------------------------------------------------------------------
About
-------------------------------------------------------------------------------

gofasta is library for processing FASTA [1] files and extracting sequences from
indexed FASTA files in faidx [2] format. 

-------------------------------------------------------------------------------
Install
-------------------------------------------------------------------------------

Fetch from github::

    $ go get github.com/aebruno/gofasta/...

-------------------------------------------------------------------------------
Usage
-------------------------------------------------------------------------------

Extract sequences from indexed fasta file::

    $ faidx /genomes/hg19.fa chr1:1000-100000

Cat FASTA files::

    $ fastcat -fasta test.fasta --count
    $ fastcat --help
    Usage of fastcat:
      -count=false: count sequences
      -fasta="": path to FASTA file
      -id=false: output ids
      -seq=false: output sequences

-------------------------------------------------------------------------------
References
-------------------------------------------------------------------------------

[1] http://en.wikipedia.org/wiki/FASTA_format 

[2] http://samtools.sourceforge.net/samtools.shtml
