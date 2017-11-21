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