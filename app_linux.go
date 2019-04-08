//+build !windows
//+build !amd64
//+build !386
//+build arm
package main

import (
         "errors"
	"github.com/visualfc/goqt/ui"
        "os"
        "os/exec"
        "path/filepath"
        "strings"
)
func GetCurrentPath() (string, error) {
        file, err := exec.LookPath(os.Args[0])
        if err != nil {
            return "",err
        }
        path, err := filepath.Abs(file)
        if err != nil {
            return "",err
        }
        index := strings.LastIndex(path, "/")
        if index < 0 {
           return "", errors.New(`Can't find "/" or "\".`)
        }
        return string(path[0:index+1]),nil

}

func AddLibraryPath(app *ui.QApplication) {
        bin_dir, _ := GetCurrentPath()
        plugins_dir := bin_dir + "../plugins"
        lib_dir := bin_dir + "../lib"
	app.AddLibraryPath(plugins_dir)
        app.AddLibraryPath(lib_dir)
}
