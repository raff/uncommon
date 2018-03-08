# uncommon
Find entries (lines) that are not in common between two files

Given two files (a master file and a file to check) this tools checks that all the entries in the second file are contained
in the master file).

The tool uses a bloom filter to minimize the amount of space used to keep the master data, and it's tunable via command line flags:

## Install
    go get github.com/raff/uncommon
    
## Usage
    Usage of ncommon:
      -fp float
    	     false positive rate (default 0.01)
      -max uint
    	     maximum number of items to compare (lines in first file) (default 100000)
      -verbose
    	     verbose mode
