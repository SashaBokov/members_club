package transport

import (
	"log"
	"net/http"
)

func Run(port string) {
	http.HandleFunc("/status", status)
	http.HandleFunc("/", index)
	http.HandleFunc("/new_member", newMember)
	// TODO: New handler
	http.HandleFunc("/remove_member", removeMember)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalln("[FATAL] ListenAndServe: ", err)
	}
}
