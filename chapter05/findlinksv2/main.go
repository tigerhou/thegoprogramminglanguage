package main

import (
	"fmt"
	//"github.com/tigerhou/thegoprogramminglanguage/chapter05/findlinksv1"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findlinksv2(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "find links:%", err)
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func Vlist(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = Vlist(links, c)
	}
	return links
}

func findlinksv2(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s : %s ", url, resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML %s", url, err)
	}

	return Vlist(nil, doc), nil

}
