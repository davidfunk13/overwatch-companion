package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type ruleset string

const (
	rulesetQuickplay   ruleset = "QUICKPLAY"
	rulesetCompetitive ruleset = "COMPETITIVE"
)


// GetStats takes a battletag and a set of rules
func GetStats(b string) []byte { 

	// dont forget about platform too. this will need to be passed in as well!
	url := "https://playoverwatch.com/en-us/career/pc/" + b

	fmt.Printf("HTML code of %s ...\n", url)

	resp, err := http.Get(url)

	// handle the error if there is one
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// show the HTML code as a string %s
	fmt.Printf("%s\n", html)
	return html
}
