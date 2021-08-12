<h1 align=center>Croissant</h1>

<div align=center>
  <img src="https://img.shields.io/github/languages/top/classPythonAddike/croissant">
  <img src="https://pkg.go.dev/badge/pkg.go.dev/github.com/classPythonAddike/croissant.svg">
  <!--[![codecov](https://codecov.io/gh/classPythonAddike/croissant/branch/master/graph/badge.svg)](https://codecov.io/gh/classPythonAddike/croissant)-->
  <img src="https://goreportcard.com/badge/github.com/classPythonAddike/croissant">
  <img src="https://sourcegraph.com/github.com/classPythonAddike/croissant/-/badge.svg">
  <img src="https://www.codetriage.com/classpythonaddike/croissant/badges/users.svg">
  <img src="https://img.shields.io/github/license/classPythonAddike/croissant?style=flat">
</div>

<br>

A fast, easy to use Rest API Framework implemented in Golang

Croissant was inspired by [FastAPI](https://fastapi.tiangolo.com/). It aims to implement FastAPI features such as Request Body Parsing, Built in Documentation Handling, and lastly, (but definitely not the least) less code required.

## Goals for the Project

I am aiming for > 90% test coverage of Croissant's source code. In addition to this, I also want Croissant to be completely type safe.

## Performance

Croissant's performance is on par with API's created with Golang's standard lib - `net/http`. This is because, as of now, it is just a simple wrapper around the library. It is likely to decrease in speed once more features are added, and I will definitely try to reduce the impact.

You can see the programs used for comparison in the [benchmarks directory](/benchmarks).

## Code Length

My primary goal for Croissant is to reduce the code needed for API's written in Golang.

*As of now, Croissant API's require more code than net/http. But remember, these values are taken from the benchmark programs, which only involve 1 route. In addition to this, Croissant doesn't provide much more than net/http at the moment. So these values are not suitable to compare with, and are more than likely to change in the future.*

1. Framework - The framework used while performing tests
2. Code Length - Lines of code written, expressed with respect to Croissant (smaller is better)
3. Time - Milliseconds taken to serve 1k requests, at 10 concurrent requests (Baton is used to test)

| Framework | Code Length  | Time |
|-----------|--------------|------|
| Croissant | 1	           | 40   |
| net/http  | 0.8          | 35   |
