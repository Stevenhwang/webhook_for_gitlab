package main

import (
	"log"
	"net/http"
	"os/exec"
)

func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println("Get webhook!")
	} else if r.Method == "POST" {
		log.Println("Post webhook!")
		if r.Header.Get("X-Gitlab-Token") == "jWPkyFjPPIJ3pPx0oLM7YSvO3sSelACZ" {
			cmd := exec.Command("/bin/sh", "/mnt/www/webhook.sh")
			err := cmd.Start()
			if err != nil {
				log.Println(err)
			}
		} else {
			log.Println("Wrong token!")
		}
	} else {
		log.Println("Not allowed method!")
	}
}

func main() {
	http.HandleFunc("/webhook", webhook)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
