# Web Server overview

### Web Socket Server
* Send "real-time" data to the chrome app on the status of the processes running in JSON
* __DONE__

### Local Socket Server
* receive communication from the Process watcher bin
* __DONE__

### Execute Thumbnail Generator
* Since the Thumbnail generator is its own executable I need a func to execute it

### Thumbnail Generator
* created: https://github.com/dtoebe/SlideShowThumbnailCreator.git
* __DONE__

### Serve Static Thumbnails
* Static path "/thumb/"
* __DONE__

### Serve Full Size images
* Static path "/full/"
* __DONE__

### Image Uploader
* Simple http multipart-form uploader
* __DONE__

### Settings writer
* receives the settings as JSON from the Chrome App the writes the slideshows settings.ini file
* __DONE__

### Stettings Reader
* read the slideshow settings and send it to client in json format

### Comment Code
* Since I am the only one working on this I haven't taken the time to comment what I have done.. Minus what GoOracle yells at me to do

### Fix Logging
* Format "[Type of msg] (portion of server) message: err if needed"
    * _Type of msg_
        * ERR, INF
    * _Portion of server_
        * Socket server, Websocket, Upload...
