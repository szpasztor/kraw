# kraw
Kraw is a tiny concurrent crawler skeleton, used as a demo or starting point for more complex projects. 

## Build
`cd src/kraw`
`go build kraw.go`

## Run 
`Usage: kraw start_url keyword visiting_limit [async|sync]"`
`Example usage: ./kraw http://wikipedia.org wiki 50 async`

## Benchmarks
The app has a switch to run in procedural mode, ie. visiting urls one after another.
A quick and unscientific result of starting the app from a github page and limiting to 100 pages:
`==> async: 7.6 s
==> sync: 70.1 s`
