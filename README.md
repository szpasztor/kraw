Kraw is a tiny concurrent crawler skeleton, used as a demo or starting point for more complex projects. 

What it does is start from the url given, search for links in its response that contain a predefined keyword, and repeat until the visiting limit has been reached.

## Build
`cd src/kraw`<br>
`go build kraw.go`

## Run 
`Usage: kraw start_url keyword visiting_limit [async|sync]"`
`Example usage: ./kraw http://wikipedia.org wiki 50 async`

## Benchmarks
The app has a switch to run in procedural mode, ie. visiting urls one after another.
A quick and unscientific result of starting the app from a github page and limiting to 100 pages:
`==> async: 7.6 s
==> sync: 70.1 s`

Please feel free to submit PRs for improvements or fixes. 
