package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os/user"
	"strings"
)

type resJSON struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

//ParsePath chacks valididity of the path sent
func ParsePath(path string) string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("[ERR] Unable to get current user: %v\n", err)
	}
	if string(path[0]) == "~" {
		return strings.Replace(path, "~", usr.HomeDir, 1)
	}
	if string(path[0]) != "/" {
		log.Fatalln("[ERR] Malformed settings.ini: path must be absolute")
	}
	return path
}

//SendMsg to send message back to client
func SendMsg(w http.ResponseWriter, title, msg string) {
	j := resJSON{Title: title, Message: msg}

	js, err := json.Marshal(j)
	if err != nil {
		log.Printf("[ERR] Cannot marshal JSON message: %s - %v\n", msg, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
