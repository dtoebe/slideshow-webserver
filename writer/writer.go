package writer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"webserver/utils"
)

type structure struct {
	Core   core     `json:"core"`
	Images []images `json:"images"`
}

type core struct {
	Name  string `json:"name"`
	Trans string `json:"transition"`
}

type images struct {
	Image []string `json:"image"`
}

//ReceiveSettings gets the settings from the client
func ReceiveSettings(w http.ResponseWriter, r *http.Request, uploadPath string) {
	decoder := json.NewDecoder(r.Body)

	var s structure
	if err := decoder.Decode(&s); err != nil {
		utils.SendMsg(w, "error", "Unable to decode JSON")
		return
	}

	fmt.Println(s)
	utils.SendMsg(w, "info", "Json Received")
	writeSettings(w, s, uploadPath)
}

func writeSettings(w http.ResponseWriter, s structure, uploadPath string) {
	pth := path.Join(uploadPath, "settings.ini")
	file, err := os.OpenFile(pth, os.O_WRONLY|os.O_CREATE, 0655)
	if err != nil {
		log.Printf("[ERR] Cannot open/create settings.ini file: %v\n", err)
		utils.SendMsg(w, "error", "Cannot write create settings.ini file")
		return
	}
	defer file.Close()

	data := "[core]\nname=" + s.Core.Name + "\ntransition=" + s.Core.Trans + "\n\n"
	data += "[image]\n"

	for _, v := range s.Images {
		for i := 0; i < len(v.Image); i++ {
			switch i {
			case 0:
				data += v.Image[i] + "="
			case len(v.Image) - 1:
				data += v.Image[i]
			default:
				data += v.Image[i] + ","
			}
		}
		data += "\n"
	}

	b, err := file.Write([]byte(data))
	if err != nil {
		log.Printf("[ERR] Cannot write settings.ini file: %v\n", err)
		utils.SendMsg(w, "error", "Cannot write data to settings file")
	} else {
		log.Printf("[INF] Wrote settings.ini file with %d bytes\n", b)
		utils.SendMsg(w, "info", "Wrote settings file")
	}
}
