package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Pattern struct {
	pattern *regexp.Regexp
	color   string
}

func main() {
	if len(os.Args) < 3 || len(os.Args)%2 == 0 {
		fmt.Fprint(os.Stderr, "Even number of arguments expected")
		fmt.Fprintf(os.Stderr, "Usage: %s pattern1 color1 [pattern2 color2 ...]\n", os.Args[0])
		os.Exit(1)
	}

	patterns, err := preparePatterns()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		// read line from stdin using newline as separator
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		for _, p := range patterns {
			line = p.pattern.ReplaceAllString(line, fmt.Sprintf("\033[%sm$1\033[0m", p.color))
		}
		fmt.Print(line)
	}
}

func preparePatterns() ([]Pattern, error) {
	patterns := make([]Pattern, (len(os.Args)-1)/2)
	for i := 1; i < len(os.Args); i += 2 {
		r, err := regexp.Compile(fmt.Sprintf("(%s)", os.Args[i]))
		if err != nil {
			return nil, err
		}
		patterns[(i-1)/2] = Pattern{pattern: r, color: os.Args[i+1]}
	}
	return patterns, nil
}
