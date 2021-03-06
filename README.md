# Goreact

This repo combines frontend react and Go backend by using Go's embed functionality. You get a single binary that has your assets in it.

## Build & Run
```sh
# you need npm install in frontend first
go generate # builds react
go build
./goreact
```