package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type opts struct {
	lineNumber, filename, ignoreCase, invert, exactMatch bool
}

func Search(pattern string, flags, files []string) []string {
	results := []string{}

	opts := parseFlags(flags)

	if opts.exactMatch {
		pattern = "^" + pattern + "$"
	}
	if opts.ignoreCase {
		pattern = `(?i)` + pattern
	}

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			return nil
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		index := 0
		done := false
		for scanner.Scan() && !done {
			index++
			line := scanner.Text()

			rx := regexp.MustCompile(pattern)
			match := rx.Match([]byte(line))
			if opts.invert {
				match = !match
			}
			if match {
				fileIdentifier := ""
				if len(files) > 1 {
					fileIdentifier = filename + ":"
				}
				if opts.filename {
					results = append(results, file.Name())
					done = true
				} else if opts.lineNumber {
					results = append(results, fmt.Sprintf("%s%d:%s", fileIdentifier, index, line))
				} else {
					results = append(results, fmt.Sprintf("%s%s", fileIdentifier, line))
				}
			}
		}
	}

	return results
}

func parseFlags(flags []string) opts {
	opts := opts{}

	for _, flag := range flags {
		switch flag {
		case "-n":
			opts.lineNumber = true
		case "-l":
			opts.filename = true
		case "-i":
			opts.ignoreCase = true
		case "-x":
			opts.exactMatch = true
		case "-v":
			opts.invert = true
		}
	}

	return opts
}
