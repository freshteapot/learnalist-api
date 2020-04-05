package hugo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/freshteapot/learnalist-api/server/api/alist"
)

type HugoAListUserWriter struct {
	dataDirectory    string
	contentDirectory string
	publishDirectory string
}

func NewHugoAListByUserWriter(contentDirectory string, dataDirectory string, publishDirectory string) HugoAListUserWriter {
	return HugoAListUserWriter{
		dataDirectory:    dataDirectory,
		contentDirectory: contentDirectory,
		publishDirectory: publishDirectory,
	}
}

func (w HugoAListUserWriter) Content(userUUID string) {
	data := struct {
		UUID string
	}{
		UUID: userUUID,
	}

	base := template.Must(template.New("").Parse(`
---
title: "My Lists"
type: "alist"
layout: "user"
Uuid: {{.UUID}}
js_include: ["main"]
---
`))
	var tpl bytes.Buffer
	base.Execute(&tpl, data)

	content := strings.TrimSpace(tpl.String())
	path := fmt.Sprintf("%s/%s.md", w.contentDirectory, userUUID)
	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		fmt.Println(err)
	}

}

func (w HugoAListUserWriter) Data(userUUID string, lists []alist.ShortInfo) {
	content, _ := json.Marshal(lists)
	path := fmt.Sprintf("%s/%s.json", w.dataDirectory, userUUID)
	err := ioutil.WriteFile(path, content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func (w HugoAListUserWriter) GetFilesToPublish() []string {
	var (
		files     []string
		toPublish []string
	)

	err := filepath.Walk(w.contentDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		fmt.Print("Something has gone wrong in getFilesToPublish, when looking for files to process")
		fmt.Println(len(files))
		fmt.Println(err)
	}

	for _, path := range files {
		if strings.HasSuffix(path, ".md") {
			toPublish = append(toPublish, path)
		}
	}
	return toPublish
}

func (w HugoAListUserWriter) GetFilesToClean() []string {
	toPublish := w.GetFilesToPublish()
	var toClean []string

	for _, path := range toPublish {
		filename := strings.TrimPrefix(path, w.contentDirectory+"/")

		if strings.HasSuffix(filename, ".md") {
			uuid := strings.TrimSuffix(filename, ".md")

			filesToClean := []string{
				fmt.Sprintf("%s/%s.md", w.contentDirectory, uuid),
				fmt.Sprintf("%s/%s.json", w.dataDirectory, uuid),
				// TODO this might not be needed
				fmt.Sprintf("%s/alistsbyuser/%s.json", w.publishDirectory, uuid),
				fmt.Sprintf("%s/alistsbyuser/%s.html", w.publishDirectory, uuid),
			}
			toClean = append(toClean, filesToClean...)
		}
	}

	return toClean
}