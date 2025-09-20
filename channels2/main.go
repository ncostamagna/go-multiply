package main

import (
    "fmt"
    "regexp"
    "strings"
	"sync"
)

const content = 
  `<html>
    <head>
      <title>Test Title</title>
    </head>
    <body>
      <h1>Title 1</h1><p>Lorem ipsum 1</p>
      <h1>Title 2</h1>
      <p>Lorem ipsum 2</p>
      <div class="top">
          <div class="mid">
              <div class="bot">
                  <p>Lorem ipsum 3</p>
              </div>
          </div>
      </div>
    </body>
  </html>`

func main() {
    fmt.Println(FindTags(content))
}

func FindTags(s string) []string {

	wg := sync.WaitGroup{}
    tags := make(chan string)
    lines := strings.Split(s, "\n")

    for _, line := range lines {
		wg.Add(1)
        go findTags(line, tags, &wg)
    }

	go func() {
		wg.Wait()
		close(tags)
	}()

    var res []string
    for tag := range tags {
        res = append(res, tag)
    }

    return res
}

var tagName = regexp.MustCompile(`<([a-zA-Z0-9_\-]+)\s?.*>`)

func findTags(s string, tags chan string, wg *sync.WaitGroup) {
	defer wg.Done()
    matches := tagName.FindAllStringSubmatch(s, -1)
    for _, match := range matches {
        if len(match) > 0 {
            tags <- match[1]
        }
    }
}