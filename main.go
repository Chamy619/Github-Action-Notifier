package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	configFilenamePtr := flag.String("config", "config.yaml", "config file name")
	hostPtr := flag.String("host", "localhost", "host")
	portPtr := flag.Int64("port", 8080, "port")

	flag.Parse()
	
	config, err := NewConfig(*configFilenamePtr)
	if err != nil {
		log.Fatal(err)
	}

	configHandler := func(writer http.ResponseWriter, request *http.Request) {
		configStr := fmt.Sprintf("%v", config)
		writer.Write([]byte(configStr))
	}

	http.HandleFunc("/config", configHandler)

	http.ListenAndServe(*hostPtr + ":" + strconv.FormatInt(*portPtr, 10), nil)
}