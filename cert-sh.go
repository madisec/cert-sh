package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/common-nighthawk/go-figure"
	// "github.com/dustinkirkland/gofmt/sprinf"
)

func main() {
	domain_flag := flag.String("d", "", "Domain for finding subdomains")
	silent_flag := flag.Bool("silent", false, "Silent mode for show minimal output")
	flag.Parse()
	if !*silent_flag {
		tool_figlet()
		send_req(*domain_flag)
	} else {
		send_req(*domain_flag)
	}

}

func tool_figlet() {
	figlet := figure.NewFigure("Cert-sh", "cyberlarge", true)
	fmt.Println(figlet)
	fmt.Println("    SSL certificate search tool - Version: 1.0.0")
	fmt.Println("    		Powered by MadiSec				 ")
	fmt.Println("")
}

type Certificate_parse struct {
	CommonName string `json:"common_name"`
}

func send_req(get_dom string) {
	dom := get_dom
	s1 := "https://crt.sh/?q="
	s2 := "&output=json"
	make_query := s1 + dom + s2
	resp, err := http.Get(make_query)
	if err != nil {
		fmt.Print("Can't call api, please try again", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Can't read body from call API.", err)
		return
	}
	final_call_api := string(body)
	var cc []Certificate_parse
	jdata := json.Unmarshal([]byte(final_call_api), &cc)
	if jdata != nil {
		fmt.Println("Error unmarshalling JSON data:  ", jdata)
		return
	}
	for _, ccp := range cc {
		common_name := ccp.CommonName
		fmt.Println(common_name)
	}
	// commonName := cc[0].CommonName
	// fmt.Println(commonName)
}
