# Corinda

Corinda is a concurrency and entropy based hash cracker.
It takes advantage of the concurrency possibilities of [Go Language](https://golang.org/) 
to learn new password model entropy patterns, as well as to make guesses against a target hash.

Related publications on the subject:

 - Passfault: an Open Source Tool for Measuring Password Complexity and Strength: https://goo.gl/uYJsVr
 
## Installing

1. Clone to GOPATH:
```
go get github.com/bernardoaraujor/corinda
```

2. notes to self:
TODO:
 - doc dependencies
    - (jnigi check /usr/lib/jvm/..., sym link include/linux/jni_md.h to include/jni_md.h)
    - http://www.gnuplot.info/
    - https://github.com/sbinet/gsl
    - scipy, powerlaw
    - libboost-all-dev
 - future work:
    - CLI (cobra?)
    - go tests
    - salting
    - l337
    - misspelling
    - toggle case
    - model recognition
    - maximum likelihood for entropy estimation
    - entropy guess sorting

3. Dictionaries:

JohnTheRipper:
This list has been compiled by Solar Designer of Openwall Project,
http://www.openwall.com/wordlists/

This list is based on passwords most commonly seen on a set of Unix
systems in mid-1990's, sorted for decreasing number of occurrences
(that is, more common passwords are listed first).  It has been
revised to also include common website passwords from public lists
of "top N passwords" from major community website compromises that
occurred in 2006 through 2010.

Last update: 2011/11/20 (3546 entries)

spanishNames:
Copyright Rhett Butler of Mongay.com

indiaNames:
This OWASP Passfault word list is licensed under a Creative Commons Attribution 4.0 International Licence.
This list was compiled by Brandon Lyew, Georgina Matias, Kevin Sealy, Michael Glassman, and Scott Sands as part of their Capstone/Winter-code-sprint project.
The information was collected from public voting records.

usFirstNames:
???

usLastNames:
???