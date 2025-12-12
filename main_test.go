package main

import (
	"ascii-art-web/asciiart"
	"bytes"
	"html/template"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func LoadBanner(str string) []string {
	banner := ""
	switch str {
	case "shadow":
		banner = "shadow.txt"
	case "standard":
		banner = "standard.txt"
	case "thinkertoy":
		banner = "thinkertoy.txt"
	default:
		banner = "standard.txt"
	}
	banner = "ascii-banners/" + banner
	content, err := os.ReadFile(banner)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n\n")
}

func Test_RENDER_Standard(t *testing.T) {
	input := "hello"
	input2 := "standard"
	banner := LoadBanner(input2)

	output := asciiart.RenderASCIIToString(input, banner)
	expected := ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               `
	if output != expected {
		t.Errorf("Expected:\n%q\n\nGot:\n%q", expected, output)
	}

}

func Test_RENDER_Shadow(t *testing.T) {
	input := "hello world"
	input2 := "shadow"
	banner := LoadBanner(input2)

	output := asciiart.RenderASCIIToString(input, banner)
	expected :=
		`                                                                                        
_|                _| _|                                                     _|       _| 
_|_|_|     _|_|   _| _|   _|_|         _|      _|      _|   _|_|   _|  _|_| _|   _|_|_| 
_|    _| _|_|_|_| _| _| _|    _|       _|      _|      _| _|    _| _|_|     _| _|    _| 
_|    _| _|       _| _| _|    _|         _|  _|  _|  _|   _|    _| _|       _| _|    _| 
_|    _|   _|_|_| _| _|   _|_|             _|      _|       _|_|   _|       _|   _|_|_| 
                                                                                        
                                                                                        `
	if output != expected {
		t.Errorf("Expected:\n%q\n\nGot:\n%q", expected, output)
	}

}
func Test_RENDER_Thinkertoy(t *testing.T) {
	input := "nice 2 meet you"
	input2 := "thinkertoy"
	banner := LoadBanner(input2)

	output := asciiart.RenderASCIIToString(input, banner)
	expected :=
		`                                                     
                       --                       o                      
     o                o  o                      |                      
o-o     o-o o-o         /        o-O-o o-o o-o -o-       o  o o-o o  o 
|  | | |    |-'        /         | | | |-' |-'  |        |  | | | |  | 
o  o |  o-o o-o       o--o       o o o o-o o-o  o        o--O o-o o--o 
                                                            |          
                                                         o--o          `

	if output != expected {
		t.Errorf("\nExpected:\n%q\n\nGot:\n%q", expected, output)
	}

}

func Test_GET_Home(t *testing.T) {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}
	var expected bytes.Buffer
	data := pageData{Result: "", Input: ""}
	err = template.Execute(&expected, data)
	if err != nil {
		panic(err)
	}
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	response := httptest.NewRecorder()
	home(response, request)
	got := response.Body.String()
	want := expected.String()
	if got != want {
		t.Errorf("\nExpected:\n%q\n\nGot:\n%q", want, got)
	}
}
