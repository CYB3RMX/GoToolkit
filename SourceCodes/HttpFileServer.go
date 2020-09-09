/* This program allows to you creating a HTTP file server */
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// init variables
	directory := [PUT YOUR HOME DIRECTORY OR ANYTHING YOU WANT]
	var bindAddr string

	// Lets ask the user
	fmt.Printf("[*] Enter Bind Address: ")
	fmt.Scanln(&bindAddr)

	// Parsing bind address
	bino := fmt.Sprintf("%s:9000", bindAddr)

	// Start
	fmt.Printf("\n[+] HTTP file server started on %s:9000\n", bindAddr)
	fmt.Printf("[+] Serving: %s\n", string(directory))

	// Handling server
	fileSystem := http.FileServer(http.Dir(directory))
	log.Fatal(http.ListenAndServe(bino, fileSystem))
}
