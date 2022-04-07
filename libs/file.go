package libs

import (
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
)

func MultipleFileHandler(form *multipart.Form, postId string) ([2][]string, error) {

	files := form.File["images"]

	var metadata [2][]string

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
			metadata[0] = append(metadata[0], filename)
			metadata[1] = append(metadata[1], url)
		}
	}

	return metadata, nil
}
