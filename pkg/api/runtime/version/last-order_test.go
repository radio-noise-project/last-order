package version

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGetLastOrderVersion(t *testing.T) {
	apath, _ := filepath.Abs("../../")
	os.Chdir(apath)
	conf := GetLastOrderVersion()
	fmt.Printf("%v", conf)
}
