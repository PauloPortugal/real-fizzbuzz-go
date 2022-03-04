[![Build Status](https://api.travis-ci.com/PauloPortugal/real-fizzbuzz-go.svg?branch=main)](https://api.travis-ci.com/PauloPortugal/real-fizzbuzz-go.svg?branch=main)

# The Real FizzBuzz in Go

Prints out the following for a contiguous range of numbers:
* the number
* `fizz` for numbers that are multiples of `3`
* `buzz` for numbers that are multiples of `5`
* `fizzbuzz` for numbers that are multiples of `15`
* if the number contains a `three` output the text `lucky`, overriding any existing behaviour
* report at the end of the program showing how many times each `word` was substituted


There is also [Java version](https://github.com/PauloPortugal/real-fizzbuzz-java)
and a [Clojure version](https://github.com/PauloPortugal/real-fizzbuzz-clojure)

## Example

Running the program over a range from 1-20 one should get the following output:

```
[1 2 lucky 4 buzz fizz 7 8 fizz buzz 11 fizz lucky 14 fizzbuzz 16 17 fizz 19 buzz]

{
  "buzz": 3,
  "fizz": 4,
  "fizzbuzz": 1,
  "integer": 10,
  "lucky": 2
}
```

## Run

```
go run main.go
```

## Test

```
go test
```