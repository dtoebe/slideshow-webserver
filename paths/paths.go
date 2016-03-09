package paths

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"webserver/utils"
)

//Paths from settings.ini
type Paths struct {
	Full   string
	Thumb  string
	Upload string
}

//StaticImageHandler serve the different static images
func (p *Paths) StaticImageHandler(w http.ResponseWriter, r *http.Request) {
	pth := r.URL.Path
	log.Println(pth)
	pth = strings.TrimSpace(pth)
	trimmedPath := strings.Split(pth, "/")
	log.Println(trimmedPath)

	length := len(trimmedPath[1]) + 1

	var localPath string
	switch trimmedPath[1] {
	case "full":
		localPath = p.Full
	case "thumb":
		localPath = p.Thumb
	}

	img := r.URL.Path[length:]
	log.Println(img)

	http.ServeFile(w, r, path.Join(localPath, img))
}

//UploadHandler allows new images to be uploaded
func (p *Paths) UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		log.Printf("[ERR] Did not receive image: %v\n", err)
		utils.SendMsg(w, "error", "Did not receive image")
		return
	}
	defer file.Close()

	out, err := os.Create(path.Join(p.Upload, header.Filename))
	if err != nil {
		log.Printf("[ERR], Cannot create file to write to: %v\n", err)
		utils.SendMsg(w, "err", "Cannot create file to write to")
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		log.Printf("[ERR] Error writing to file: %v\n", err)
		utils.SendMsg(w, "err", "Error writing to file")
		return
	}

	log.Printf("[INF] File uploaded successfully: %s\n", header.Filename)
	utils.SendMsg(w, "info", "File uploaded successfully")
}
