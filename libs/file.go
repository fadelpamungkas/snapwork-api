package libs

import (
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
)

type Metadata struct {
	Filename string `json:"filename"`
	Url      string `json:"url"`
}

func MultipleFileHandler(form *multipart.Form, postId string) ([]Metadata, error) {
	files := form.File["images"]

	var metadata []Metadata

	for i, file := range files {
		var filename string
		var url string

		if file != nil {
			extenstionFile := filepath.Ext(file.Filename)
			filename = fmt.Sprintf("%s-%d%s", "pict", i, extenstionFile)

			out, err := file.Open()
			if err != nil {
				log.Println("err: ", err)
			}
			log.Println("File Header: ", file.Header)

			fileUrl, err := UploadFileInS3Bucket(out, fmt.Sprintf("5dzt/%s/%s", postId, filename))
			if err != nil {
				log.Println("err: ", err)
			}
			url = fileUrl
			log.Println("Saved to: ", fileUrl)

		} else {
			log.Println("Nothing file to uploading.")
		}

		if filename != "" && url != "" {
			metadata = append(metadata, Metadata{
				Filename: filename,
				Url:      url,
			})
		}
	}

	return metadata, nil
}
