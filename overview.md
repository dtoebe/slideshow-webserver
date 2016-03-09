# Web Server overview

### Web Socket Server
* Send "real-time" data to the chrome app on the status of the processes running in JSON
* __DONE__

### Local Socket Server
* receive communication from the Process watcher bin

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
