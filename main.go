package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"strconv"
)

type GithubActionHTTPRequestBody struct {
	Payload string `json:"payload"`
}

func main() {
	// 플래그 설정
	configFilenamePtr := flag.String("config", "config.yaml", "config file name")
	hostPtr := flag.String("host", "localhost", "host")
	portPtr := flag.Int64("port", 8080, "port")

	flag.Parse()

	// 설정 파일 불러오기
	config, err := NewConfig(*configFilenamePtr)
	if err != nil {
		log.Fatal(err)
	}

	// 서비스 생성
	githubActionService := NewGithubActionService(config)

	http.HandleFunc("/github-action", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			writer.WriteHeader(http.StatusNotFound)
			writer.Write([]byte("invalid method"))
			return
		}

		body, err := io.ReadAll(request.Body)
		if err != nil {
			log.Printf("not readable body: %v", err)
		}

		var githubActionHTTPRequestBody GithubActionHTTPRequestBody
		err = json.Unmarshal(body, &githubActionHTTPRequestBody)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("request body unmarshaling fail"))
			return
		}
		if githubActionHTTPRequestBody.Payload == "" {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("invalid body"))
			return
		}

		err = githubActionService.Pushed(githubActionHTTPRequestBody)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("internal server error"))
			return
		}

		writer.Write([]byte("ok"))
	})

	log.Printf("starting Notifier %v:%v", *hostPtr, *portPtr)
	log.Fatal(http.ListenAndServe(*hostPtr+":"+strconv.FormatInt(*portPtr, 10), nil))
}
