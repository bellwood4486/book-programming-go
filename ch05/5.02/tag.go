package tag

import (
	"golang.org/x/net/html"
	"io"
)

func tagfreq(rd io.Reader) (map[string]int, error) {
	doc, err := html.Parse(rd)
	if err != nil {
		return nil, err
	}
	freq := make(map[string]int)
	visit(freq, doc)
	return freq, nil
}

func visit(freq map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		freq[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(freq, c)
	}
}
