# Corinda Docs

Corinda is a concurrency and entropy based hash cracker. It takes advantage of the concurrency possibilities of Go Language to learn new password model entropy patterns, as well as to make guesses against a target hash.

## Golang Primitives

### Module
### Struct
### Go Routine
### Map

## Models

Corinda abstracts password substrings as Models. Models can be either Elementary or Composite.

For example, in the password `corinda123`, the Token `corinda` belongs to the Elementary Model of words belonging to a dictionary (let's call it **[dict]**), and `123` belongs to the Elementary Model of three number sequences (let's call it **[3N]**). In this case, the `corinda123` belongs to the Composite Model **[dict 3N]**.

## Elementary Models

Elementary Models are abstractions for the atomic rule structures inside the password string. Each Elementary Model is has a Name, an Entropy value, and a map of Token frequencies. Tokens are just substring instances of the atomic rule structure, and Frequencies are integers that account for the number of occurances of a specific Token inside a Training list.

In terms of Golang's primitives, Elementary Models are implemented as structs:

```
type Model struct {
	Name       string
	Entropy    float64
	TokenFreqs map[string]int
}
```

## Composite Models

## Training Module

## Cracking Module
