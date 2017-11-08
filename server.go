package fileserver

import (
	"net/http"
	"os"
)

/*
	http.Handle("/", http.FileServer(http.Dir("/tmp")))

	The original http.FileServer() will show directory listings, this one does not

	https://groups.google.com/forum/#!topic/golang-nuts/bStLPdIVM6w

	Example:

	func main() {
		fs := simplefileserver.OnlyFilesFilesystem{http.Dir("/tmp/")}
		http.ListenAndServe(":8080", http.FileServer(fs))
	}

*/

type OnlyFilesFilesystem struct {
	FS http.FileSystem
}

type neuteredReaddirFile struct {
	http.File
}

func (fs OnlyFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.FS.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}
