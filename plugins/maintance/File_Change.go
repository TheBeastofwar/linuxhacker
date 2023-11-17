package maintance

import (
	"fmt"
	"os"
	"strings"
)

func File_Change(exist_change string) {
	existchange := strings.Split(exist_change, " ")
	existingfile := existchange[0]
	changingfile := existchange[1]
	fileInfo, err := os.Stat(existingfile)
	if err != nil {
		fmt.Println("Unable to obtain file information", err)
		os.Exit(1)
	}
	atime := fileInfo.ModTime()
	mtime := fileInfo.ModTime()

	err = os.Chtimes(changingfile, atime, mtime)
	if err != nil {
		fmt.Println("Unable to change file information", err)
		os.Exit(1)
	}
	fmt.Println("success")
}
