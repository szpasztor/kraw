package main


import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"github.com/mvdan/xurls"
	"net/url"
	"strings"
	"strconv"
)




// Filter url for a predefined keyword
// Also, check that it has not yet been visited
func Filter(url *url.URL, visited *map[string]bool) bool {
	return strings.Contains(url.Hostname(), keyword) && (*visited)[url.String()] == false
}

//
func Visit(destination string, queue *chan string, visited *map[string]bool) {
	// Optional delay
	// time.Sleep(time.Millisecond * 100)

	resp, err := http.Get(destination)

	if err != nil {
		// Handle error
		// Skip urls that cause an error
		//fmt.Println(err)
	} else {
		defer resp.Body.Close()
		htmlData, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			// Carry on if body cannot be parsed
			//fmt.Println(err)
		} else {
			urls := xurls.Relaxed.FindAllString(string(htmlData), -1)
			for x := range urls {
				url, err := url.Parse(urls[x])
				if err != nil {
					// carry on if url cannot be parsed
					//fmt.Print(err)
				} else {
					if Filter(url, visited) {
						*queue <- urls[x]
					}

				}
			}


			// Mark visited
			(*visited)[destination] = true

			// Optional logs
			fmt.Println("Visited: ", destination)
			//fmt.Println("Queue length: ", len(*queue))
			//fmt.Println("Visited so far: ", len(*visited))
		}
	}
}


var keyword string

func main() {
	fmt.Println("Starting")

	args := os.Args
	if len(args) != 5 {
		fmt.Println("Usage: kraw start_url keyword visiting_limit [async|sync]")
		os.Exit(0)
	}
	// Read arguments
	start_url := args[1]
	keyword = args[2]
	visiting_limit, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("Malformed visiting_limit")
		os.Exit(2)
	}
	async := args[4] == "async"

	// A map that stores urls that have already been visited
	visited := make(map[string]bool)

	// This channel will store the urls to visit
	queue := make(chan string, 10000)
	queue <- start_url

	// Iterate on queue of unvisited urls until limit reached
	for len(visited) < visiting_limit {
		next := <- queue
		if async {
			// Start new goroutine and carry on with next elements from queue
			go Visit(next, &queue, &visited)
		} else {
			// Invoke method and continue executing once it's finished
			Visit(next, &queue, &visited)
		}
	}

	fmt.Println("Finished")

}
