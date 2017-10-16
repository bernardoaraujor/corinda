
# Corinda

Corinda is a Bayesian hash cracker and password strength estimator.

Theoretical foundation relies on Set Theory, First-order Model Theory, and Statistical Inference. Corinda expands upon the notions already established in Sahin et al. [1]. Previous results by Rodrigues et al. [2] suggest that human-biased models generate passwords with distinct statistical patterns, raising the following questions:

 - Can a Password's Strength be modeled as a Statistical Inference problem? 
 - Can the process of Password Cracking be fully optmized by Machine Learning?
 
 [1] https://goo.gl/uYJsVr
 
 [2] https://goo.gl/mUVlw2

## installation

1. Clone recursively:
```git clone --recursive https://github.com/bernardoaraujor/corinda
cd corinda
```

2. If corinda/training/input is empty, fetch from lfs:
```git lfs fetch
```

3. Build Passfault:
```cd model_recognition/passfault/commandLine
../gradlew installDist
```
