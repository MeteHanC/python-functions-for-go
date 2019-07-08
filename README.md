[![Codacy Badge](https://api.codacy.com/project/badge/Grade/2d3d22a86a4241fe825b6c445549afc7)](https://app.codacy.com/app/MeteHanC/python-functions-for-go?utm_source=github.com&utm_medium=referral&utm_content=MeteHanC/python-functions-for-go&utm_campaign=Badge_Grade_Dashboard)
[![Go Report Card](https://goreportcard.com/badge/github.com/metehanc/python-functions-for-go)](https://goreportcard.com/report/github.com/metehanc/python-functions-for-go)
[![codecov](https://codecov.io/gh/MeteHanC/python-functions-for-go/branch/master/graph/badge.svg)](https://codecov.io/gh/MeteHanC/python-functions-for-go)
[![CircleCI](https://circleci.com/gh/MeteHanC/python-functions-for-go.svg?style=svg)](https://circleci.com/gh/MeteHanC/python-functions-for-go)

# Python Functions For Go
Some useful Python functions for Golang. Currently there are only 3; reverse, min, max



## Reverse
It takes an array/slice of any type and reverses it and returns it as an []interface. Example usage;

```go
pygo.Reverse( []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "golang"} )
```
Will give you an interface slice;

```
[9 8 7 6 5 4 3 2 1]
```

## Min 
Min function returns the smallest value of the given array/slice of any type as an interface. Example usage;

```go
pygo.Min([]int{1,6,-2,5,100,-10})
```
Will return 
```
-10
```

## Max
Max function returns the largest value of the given array/slice of any type as an interface. Example usage;

```go
pygo.Max([]float64{1,6,-2,5,100,-10})
```
Will return
```
100
```
