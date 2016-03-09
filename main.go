package main

import (
	"log"
	"net/http"

	"github.com/dtoebe/slideshow-webserver/paths"
	"github.com/dtoebe/slideshow-webserver/utils"
	"github.com/dtoebe/slideshow-webserver/writer"

	"github.com/vaughan0/go-ini"
)

func parseSetting() (core, images map[string]string) {
	file, err := ini.LoadFile("settings.ini")
	if err != nil {
		log.Fatalf("[ERR] Cannot load settings file: %v\n", err)
	}
	return file["core"], file["images"]
}

func main() {
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

	log.Printf("[INF] Server started on port: %s\n", core["web_port"])
	log.Fatal(http.ListenAndServe(":"+core["web_port"], nil))
}
