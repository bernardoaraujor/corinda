# Corinda

Corinda is a concurrency and entropy based hash cracker.
It takes advantage of the concurrency possibilities of [Go Language](https://golang.org/) 
to learn new password model entropy patterns, as well as to make guesses against a target hash.

Related publications on the subject:

 - Passfault: an Open Source Tool for Measuring Password Complexity and Strength: [https://goo.gl/uYJsVr](https://goo.gl/uYJsVr)
 
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

This work is part of my MSc. in Computer Engineering @ UFG, Brazil. The Dissertation can be found at `/ufg`, however only a Brazilian Portuguese version is available.


## Dictionaries

- JohnTheRipper:

> This list has been compiled by Solar Designer of Openwall Project, http://www.openwall.com/wordlists/ .
> This list is based on passwords most commonly seen on a set of Unix
systems in mid-1990's, sorted for decreasing number of occurrences
(that is, more common passwords are listed first).  It has been
revised to also include common website passwords from public lists
of "top N passwords" from major community website compromises that
occurred in 2006 through 2010. Last update: 2011/11/20 (3546 entries).

- spanishNames:
> Copyright Rhett Butler of Mongay.com

- indiaNames:
> This OWASP Passfault word list is licensed under a Creative Commons Attribution 4.0 International Licence. This list was compiled by Brandon Lyew, Georgina Matias, Kevin Sealy, Michael Glassman, and Scott Sands as part of their Capstone/Winter-code-sprint project.
The information was collected from public voting records.

- usFirstNames:
TODO

- usLastNames:
TODO

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
