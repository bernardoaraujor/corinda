# Corinda

![alt text](https://github.com/bernardoaraujor/corinda/corinda.jpg "Corinda")


Corinda is a concurrency and entropy based hash cracker.
It takes advantage of the concurrency possibilities of [Go Language](https://golang.org/) 
to learn new password model entropy patterns, as well as to make guesses against a target hash.

Related publications on the subject:

 - General Framework for Evaluating Password Complexity and Strength: [https://goo.gl/uYJsVr](https://goo.gl/uYJsVr)
 - Passfault: an Open Source Tool for Measuring Password Complexity and Strength: [https://goo.gl/mUVlw2](https://goo.gl/mUVlw2)
 
## Installing

1. Install Go tools: https://golang.org/doc/install
2. Set GOPATH: https://golang.org/doc/code.html#GOPATH
3. Clone to GOPATH:
```
cd $GOPATH
go get github.com/bernardoaraujor/corinda
```

## Training

1. Have .csv file corretly formatted in `/csv`. In other words, each line must be composed of `[frequency,password]`
2. Run `main/mainTrain.go`:
```
cd $GOPATH/src/github.com/bernardoaraujor/corinda/
go run main/mainCrack.go <input list>
```

## Cracking

1. Make sure you have the trained maps (elementary and composite) in `/maps`.
2. Make sure you have there is only one file for target list in `/targets` (run `merger.sh` script if necessary)
3. Run `main/mainCrack.go`
```
cd $GOPATH/src/github.com/bernardoaraujor/corinda/
go run main/mainCrack.go <trained list> <target list> <sha1 or sha256>
```

## Passfault

This project uses a modified version of [OWASP's Passfault](http://www.passfault.com/) as a submodule.

## Masters

This work is part of my MSc. in Computer Engineering @ UFG, Brazil. The Dissertation will be available soon, however only a Brazilian Portuguese version.


## Passfault Wordlists

Found at `passfault_corinda/src/org/owasp/passfault/wordlists/`.

#### Language words (xxPopular and xxLongTail):
Generated with help of the Python module [wordfreq](https://pypi.python.org/pypi/wordfreq), maintained by [Luminoso Technologies, Inc.](https://luminoso.com/). The module gathers information about word usage on different topics at different levels of formality, using data collected from the following sources: LeedsIC, SUBTLEX, OpenSub, Twitter, Wikipedia, Reddit, and CCrawl.

xxPopular contain the 80% head of the Zipf Distribution of words, while xxLongTail contain the 20% long tail.

#### JohnTheRipper:

Downloaded from [https://wiki.skullsecurity.org/Passwords](https://wiki.skullsecurity.org/Passwords).

> This list has been compiled by Solar Designer of Openwall Project, http://www.openwall.com/wordlists/ .
> This list is based on passwords most commonly seen on a set of Unix
systems in mid-1990's, sorted for decreasing number of occurrences
(that is, more common passwords are listed first).  It has been
revised to also include common website passwords from public lists
of "top N passwords" from major community website compromises that
occurred in 2006 through 2010. Last update: 2011/11/20 (3546 entries).

#### cain-and-abel:

Downloaded from [https://wiki.skullsecurity.org/Passwords](https://wiki.skullsecurity.org/Passwords).

#### 500-worst-passwords:
Downloaded from [https://github.com/danielmiessler/SecLists/tree/master/Passwords](https://github.com/danielmiessler/SecLists/tree/master/Passwords).

#### 10k-worst-passwords:
Downloaded from [https://github.com/danielmiessler/SecLists/tree/master/Passwords](https://github.com/danielmiessler/SecLists/tree/master/Passwords).

#### spanishNames:
> Copyright Rhett Butler of Mongay.com

#### indiaNames:
> This OWASP Passfault word list is licensed under a Creative Commons Attribution 4.0 International Licence. This list was compiled by Brandon Lyew, Georgina Matias, Kevin Sealy, Michael Glassman, and Scott Sands as part of their Capstone/Winter-code-sprint project.
The information was collected from public voting records.

#### usFirstNames:
Downloaded from the US Social Security website: [https://www.ssa.gov/oact/babynames/limits.html](https://www.ssa.gov/oact/babynames/limits.html)

#### usLastNames:
Downloaded from the US Census of 2000: [https://www.census.gov/topics/population/genealogy/data/2000_surnames.html](https://www.census.gov/topics/population/genealogy/data/2000_surnames.html)

## TODO
 - doc dependencies
    - (jnigi check /usr/lib/jvm/..., sym link include/linux/jni_md.h to include/jni_md.h)
 - future work:
    - CLI (cobra?)
    - go tests
    - support salting
    - support l337 (Passfault)
    - misspelling (Passfault)
    - support toggle case (Passfault)
    - entropy guess sorting

## Tony Corinda
[https://en.wikipedia.org/wiki/Tony_Corinda](https://en.wikipedia.org/wiki/Tony_Corinda)
