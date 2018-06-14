# Corinda Docs

Corinda is a concurrency and entropy based hash cracker. It takes advantage of the concurrency possibilities of the [Go Language](https://golang.org/) to learn new password model entropy patterns, as well as to make guesses against a target hash.

## Golang Primitives

### Module
### Struct
### Go Routine
### Map

## Models

Corinda abstracts password substrings as Models. Models can be either Elementary or Composite.

For example, in the password `corinda123`, the Token `corinda` belongs to the Elementary Model of words belonging to a dictionary (let's call it **[dict]**), and `123` belongs to the Elementary Model of three number sequences (let's call it **[3N]**). In this case, the `corinda123` belongs to the Composite Model **[dict 3N]**.

## Elementary Models

Elementary Models are abstractions for the atomic rule structures dictating the password string. Each Elementary Model is has a Name, an Entropy value, and a map of Token frequencies. Tokens are just substring instances of the atomic rule structure, and Frequencies are integers that account for the number of occurances of a specific Token inside a Training list.

Elementary Models are implemented in the [Elementary Module](https://github.com/bernardoaraujor/corinda/blob/master/elementary/elementary.go).

In terms of Golang's primitives, an Elementary Model is implemented as a `struct`:

```
type Model struct {
	Name       string
	Entropy    float64
	TokenFreqs map[string]int
}
```

Elementary Models have the following functions:

 - **UpdateEntropy**: TODO
 - **UpdateTokenFreq**: TODO
 - **SortedTokens**: TODO

## Composite Models

Composite Models are abstractions for the composition of atomic rule structures dictating the password string. Each Composite Model has a Name, a Frequency value, a Probability value, a total Entropy value, and an array of strings representing the Elementary Models that form the composed structure.

Composite Models are implemented in the [Composite Module](https://github.com/bernardoaraujor/corinda/blob/master/composite/composite.go).

A Composite Model is also implemented as a `struct`:

```
type Model struct{
	Name             string
	Freq    int
	Prob 	float64
	Entropy float64
	Models  []string
}
```

Composite Models have the following functions:

 - **UpdateProb**: TODO
 - **UpdateFreq**: TODO
 - **UpdateEntropy**: TODO
 - **recursive**:

## Training Module

## Cracking Module
