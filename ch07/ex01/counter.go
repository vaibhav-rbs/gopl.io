package counter

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	var n int
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	words, n, err := count(scanner)
	*c += WordCounter(words)
	return n, err
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	lines, n, err := count(scanner)
	*c += LineCounter(lines)
	return n, err
}

func count(scanner *bufio.Scanner) (units, bytes int, err error) {
	for {
		ok := scanner.Scan()
		if !ok {
			return units, bytes, fmt.Errorf("error counting using scanner %v", scanner)
		}
		if ok {
			units++
			bytes += len(scanner.Bytes())
		} else {
			break
		}
	}
	return units, bytes, scanner.Err()
}
