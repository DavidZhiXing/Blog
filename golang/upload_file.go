// create package to upload file
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

// set max file size to 10MB
const maxFileSize = 10 * 1024 * 1024

// report the upload progress

type Progress struct {
	TotalSize int64
	CurrentSize int64
}

func (p *Progress) Write(b []byte) (int, error) {
	n := len(b)
	p.CurrentSize += int64(n)

	fmt.Printf("\rUploaded %d/%d bytes", p.CurrentSize, p.TotalSize)
	return n, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "./upload_file.html")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxFileSize)
	if err := r.ParseMultipartForm(maxFileSize); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(maxFileSize); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]
	for _, fileHeader := range files {
		if fileHeader.Size > maxFileSize {
			http.Error(w, "File too big", http.StatusBadRequest)
			return
		}


		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fileType := http.DetectContentType(buff)
		if fileType != "image/jpeg" && fileType != "image/png" {
			http.Error(w, "Unsupported file type", http.StatusBadRequest)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}


		// create the upload file folder if not exists
		if _, err := os.Stat("upload"); os.IsNotExist(err) {
			os.Mkdir("upload", 0777)
		}

		// create a new file plus time in the upload folder
		newFile, err := os.Create(fmt.Sprintf("upload/%s_%s", fileHeader.Filename, time.Now().Format("20060102150405")))
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		defer newFile.Close()

		pr := &Progress{
			TotalSize: fileHeader.Size,
		}

		// copy the uploaded file to the file system
		_, err = io.Copy(newFile, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", fileHeader.Filename)

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}