package main

import (
	"log"
	"net/http"

	"github.com/dtoebe/slideshow-webserver/paths"
	"github.com/dtoebe/slideshow-webserver/socket"
	"github.com/dtoebe/slideshow-webserver/utils"
	"github.com/dtoebe/slideshow-webserver/writer"

	"github.com/vaughan0/go-ini"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

func parseSetting() (core, images map[string]string) {
	file, err := ini.LoadFile("settings.ini")
	if err != nil {
		log.Fatalf("[ERR] Cannot load settings file: %v\n", err)
	}
	return file["core"], file["images"]
}

func main() {
	// go socket.TestData()
	core, images := parseSetting()
	uploadPath := utils.ParsePath(images["upload_path"])
	fullPath := utils.ParsePath(images["static_path"])
	thumbPath := utils.ParsePath(images["thumb_path"])

	p := paths.Paths{Full: fullPath, Thumb: thumbPath, Upload: uploadPath}

	http.HandleFunc("/full/", p.StaticImageHandler)
	http.HandleFunc("/thumb/", p.StaticImageHandler)

	http.HandleFunc("/upload", p.UploadHandler)
	http.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		writer.ReceiveSettings(w, r, uploadPath)
	})

	http.Handle("/sock/", sockjs.NewHandler("/sock", sockjs.DefaultOptions, socket.Server))
	go socket.TCPServer("localhost", core["socket_port"])

	log.Printf("[INF] Web server started on port: %s\n", core["web_port"])
	log.Fatal(http.ListenAndServe(":"+core["web_port"], nil))
}
