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

    $ go get install github.com/aebruno/gofasta
    $ go get install github.com/aebruno/gofasta/faidx
    $ go get install github.com/aebruno/gofasta/fastcat

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
License
-------------------------------------------------------------------------------

gofasta was written by Andrew E. Bruno and released under the GNU General
Public License ("GPL") Version 3.0.  See the LICENSE file.

-------------------------------------------------------------------------------
References
-------------------------------------------------------------------------------

[1] http://en.wikipedia.org/wiki/FASTA_format 
[2] http://samtools.sourceforge.net/samtools.shtml
