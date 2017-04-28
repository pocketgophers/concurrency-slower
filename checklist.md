# Things to check for each section

- check style with vale
- check spelling, grammar with Marked

## Introduction
- go get import path
- machine specs

## The Posted Code
- go run must not produce enough lotto numbers

## Extract to Function
- go run must not produce enough lotto numbers

## Testing with go test
- test of original version must fail
- go run must produce enough lotto numbers
- test of other version must fail

## A Serial Version
- go run must produce enought lotto numbers
- go test must pass

## Benchmarks
- benchmark must work

## A Parallel Version
- minimal unit must reflect serial version
- minimal scalable unit must reflect serial version
- concurrency-safe msu must reflect msu
- go test must pass, no races
- benchmark must work
- benchstat must show parallel to be slower

## Reducing Blocking Operations
- test with blockprofile must pass
- pprof listing must work
- benchmark must work
- benchstat must be faster than serial

## Reducing Allocations
- build must work
- slice literal must reflect previous step
- second build must work
- remaining escapes must reflect build output
- benchmark must work
- benchstat must be faster, with less memory and fewer allocs

## Success
- benchstat must be faster, with less memory and fewer allocs

# Modifications to generated site
- sanitize directory names