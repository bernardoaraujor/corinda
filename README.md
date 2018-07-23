<h1 align="center">
  <br>
  Corinda <br>
  <img src="https://raw.githubusercontent.com/bernardoaraujor/corinda/master/corinda.jpg">
  <br>
</h1>
<h3 align="center">
A password hash cracker written in <a href="https://golang.org" target="_blank">Golang</a>.
</h3>

[https://github.com/bernardoaraujor/corinda](
https://github.com/bernardoaraujor/corinda)

Corinda uses concurrent heuristics based on model entropy and relative frequency from sample sets. Currently, Corinda has the following sample sets:

 - RockYou
 - LinkedIn
 - AntiPublic
 
The user feeds Corinda a password hash (SHA1, 2, or 3), chooses a sample set, and waits while concurrent goroutines try to crack it with computational load balancing for each password model.

Corinda only supports CPU cracking.

Corinda uses a modified version of [OWASP's Passfault](http://www.passfault.com/) to train password models.

Related publications on the subject:

 - General Framework for Evaluating Password Complexity and Strength: [https://goo.gl/uYJsVr](https://goo.gl/uYJsVr)
 - Passfault: an Open Source Tool for Measuring Password Complexity and Strength: [https://goo.gl/mUVlw2](https://goo.gl/mUVlw2)

Corinda is released under [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html). Corinda is distributed for research purposes only. We believe that people should understand the dangers of simple passwords, and Corinda is an effort to encourage people to protect their privacy with high entropy passwords. And remember, with great powers, come great responsabilitiy!

Contact:
bernardoar@protonmail.com
 
## Installing

1. Install Go tools: https://golang.org/doc/install
2. Set GOPATH: https://golang.org/doc/code.html#GOPATH
3. Clone to GOPATH:
```
cd $GOPATH
mkdir -p src/github.com/bernardoaraujor/
cd src/github.com/bernardoaraujor
git clone https://github.com/bernardoaraujor/corinda
cd corinda
git submodule update --init --recursive
```

## Training

1. Have input file correctly formatted as Comma Separated Values. Each line must be composed of `[frequency,password]`. The .csv must be compressed into .csv.gz, and placed inside the `/csv` directory.
2. Run train command:
```
corinda train <input>
```

## Cracking

1. Make sure you have the trained maps (elementary and composite) in `/maps`.
2. Make sure you have there is only one file for target list in `/targets` (run `merger.sh` script if necessary)
3. Run crack command:
```
crack <trained list> <target list> <sha1 or sha256>
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
    - go tests
    - support salting
    - support l337 (Passfault)
    - misspelling (Passfault)
    - support toggle case (Passfault)
    - entropy guess sorting

## Tony Corinda
[https://en.wikipedia.org/wiki/Tony_Corinda](https://en.wikipedia.org/wiki/Tony_Corinda)
