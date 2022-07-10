package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func showInfo() {
	name := "Igor Oliveira"
	fmt.Println("Hello, ", name)
}

func showOptions() {
	fmt.Println("1- Start")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
}

func readOption() int {
	var option int
	fmt.Print("Select the option: ")
	fmt.Scan(&option)
	return option
}

func startMonitor() {
	fmt.Println("Start the monitor...")
	sites := openFile()
	/* primeiro retorno a chave do array e o segundo o valor, caso não use a chave, colocar um _ */
	for i := 0; i < 5; i++ {
		for _, site := range sites {
			testSite(site)
		}
		// espero 10s
		time.Sleep(10 * time.Second)
	}
}

func testSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("He have a error...")
	}

	if response.StatusCode == 200 {
		fmt.Println(site, " is online")
	} else {
		fmt.Println(site, " is not online | ", response.StatusCode)
	}
}

func openFile() []string {
	sites := []string{}
	openFile, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println(err)
	}

	line := bufio.NewReader(openFile)
	for {
		result, err := line.ReadString('\n')
        result = strings.TrimSpace(result)
        sites = append(sites, result)
		if err != io.EOF {
            break
		}
	}
    
    openFile.Close()
	return sites
}

func main() {
	showInfo()
	showOptions()
	option := readOption()

	switch option {
	case 1:
		startMonitor()
	case 2:
		fmt.Println("Show the logs...")
	case 0:
		fmt.Println("Exit the program")
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
	}
}
