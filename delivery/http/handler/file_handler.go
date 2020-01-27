package handler

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

//UploadImage saves a Image in a directory
func UploadImage(w http.ResponseWriter, r *http.Request, file multipart.File, header *multipart.FileHeader) (string, error) {

	// r.ParseMultipartForm(10 << 20)
	// file, handler, err := r.FormFile("myImage")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	fmt.Println(file)

	defer file.Close()

	fname := header.Filename
	fmt.Println(fname)

	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := filepath.Join(wd, "assets", "images", fname)
	image, err := os.Create(path)
	if err != nil {

		return "", err
	}
	defer image.Close()

	io.Copy(image, file)

	// http.Redirect(w, r, "/done", http.StatusSeeOther)
	w.Header().Set("Content-Type", "application/json")

	// http.Error(w, http.StatusText(http.StatusCreated), http.StatusCreated)
	return fname, nil
}
