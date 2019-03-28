package main

/*
import (
	"fmt"
	"regexp"
)

func main() {
	data := `#  123.domain.tld #`
	pat := regexp.MustCompile(`(?P<num>(\d+)).(?P<WAT>(\w+)).\w+`)

	mch := pat.FindStringSubmatch(data)
	res := make(map[string]string)
	for i, name := range pat.SubexpNames() {
		if i != 0 && name != "" {
			res[name] = mch[i]
		}
	}
	fmt.Printf("\n%q\n", res)
	fmt.Printf("\n%s\n", res["num"])
}
*/
