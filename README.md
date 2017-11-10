# Corinda

Corinda is a concurrency and entropy based hash cracker.
It takes advantage of the concurrency possibilities of [Go Language](https://golang.org/) 
to learn new password model entropy patterns, as well as to make guesses against a target hash.

Related publications on the subject:

 - Passfault: an Open Source Tool for Measuring Password Complexity and Strength: https://goo.gl/uYJsVr
 - zxcvbn: realistic password strength estimation: https://goo.gl/9AfUhv
 
## Installing

1. Clone to ~/go/src/github/bernardoaraujor/:
```
mkdir -p ~/go/src/github/bernardoaraujor/
cd ~/go/src/github/bernardoaraujor/
git clone --recursive https://github.com/bernardoaraujor/corinda
```

2. If corinda/training/input is empty, fetch from lfs (might take a while):
```
cd corinda
git lfs fetch
