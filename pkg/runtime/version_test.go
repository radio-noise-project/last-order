package runtime

import (
	"fmt"
	"testing"
)

func TestDockerEngineVersion(t *testing.T) {
	infoVersion := GetDockerEngineVersion()
	fmt.Printf("Docker Engine Version: %s", infoVersion)
}
