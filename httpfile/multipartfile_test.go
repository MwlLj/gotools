package httpfile

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

type CUploadFileHandler struct {
}

func (*CUploadFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urls, err := DownloadFile(r, "file", "uploads", true)
	fmt.Println(urls, err)
	io.WriteString(w, "success")
}

// func TestDownloadFile(t *testing.T) {
// 	mux := http.NewServeMux()
// 	mux.Handle("/multi/file", &CUploadFileHandler{})
// 	http.ListenAndServe(":8080", mux)
// }

func TestUploadFile(t *testing.T) {
	var response string = ""
	UploadFile("pictureurl", "uploads/test.jpg", "http://127.0.0.1:59000/resource/picture", &response)
	fmt.Println(response)
}
