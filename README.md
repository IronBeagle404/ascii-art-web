# ASCII-Art Web

A website that allow users to turn text into ASCII representation.

## Project structure

<pre>
ascii-art-web/
├── main.go -- Main backend file
├── go.mod -- Golang module for the backend
├── README.md -- Documentation (this file)
├── asciiart/
│ ├── asciiart.go -- Main ASCII rendering file 
│ └── RenderASCII.go -- Secondary ASCII rendering file
├── ascii-banners
│ ├── shadow.txt -- Shadow banner file 
│ ├── standard.txt -- Standard banner file
│ └── thinkertoy.txt -- Thinkertoy banner file
├── templates/
│ └── index.html -- HTML template for the site's main page
├── static/ 
│ └── style.css -- CSS file for the site's main page
│  └── img/
│   ├── favicon.ico -- Website's browser icon
│   └── logo.png -- Website's logo
│   └── maxwell.jpeg -- A picture of a loyal & wise companion
</pre>



## Design

We made a [basic yet efficient wireframe](https://wireframe.cc/aO2vU7) to visualize the design of our website.
<br>

Later on we also made a [graphical mockup](https://app.moqups.com/Kgrz1bMDGttmJ1nH7QROiUKWRF4ARnUZ/view/page/ad64222d5)

## Usage

To launch the server, you can simply run the main backend file with `go run main.go`<br>
Once it's ready, you will be prompted with an indication of the port hosting the server.

Since it is harcoded to use the port `8080`, you can connect to the website with the adress http://localhost:8080

To stop the server, you can just `CTRL + C` in the console *(the terminal window where you ran the previous command)*

<br>

### Acknowledgements

This is originally a study project part of Zone01's core curriculum. 

I made this project in collaboration with [Nathan Allain](https://github.com/nathuc69)