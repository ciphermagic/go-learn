package filelisting

import (
	"github.com/gpmgo/gopm/modules/log"
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	i, err := writer.Write(all)
	if err != nil {
		return err
	}
	log.Debug("write success %d bytes", i)
	return nil
}
