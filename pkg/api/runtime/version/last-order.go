package version

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/docker/docker/client"
	"github.com/radio-noise-project/last-order/pkg/api/runtime/types"
)

func GetLastOrderVersion() types.ResponseLastOrderVersion {
	return types.ResponseLastOrderVersion{
		CodeName:            getCodeName(),
		Version:             getVersion(),
		BuiltDate:           getBuiltDate(),
		DockerEngineVersion: getDockerEngineVersion(),
		BuiltGitCommitHash:  getBuiltCommitHash(),
		Os:                  getOs(),
		Arch:                getArch(),
	}
}

func getCodeName() string {
	codeName, _, _ := getVersionLastOrder()
	return codeName
}

func getVersion() string {
	_, version, _ := getVersionLastOrder()
	return version
}

func getBuiltDate() time.Time {
	_, _, builtDate := getVersionLastOrder()
	t, _ := time.Parse("2006-01-02", builtDate)
	return t
}

func getOs() string {
	return os.Getenv("OS")
}

func getArch() string {
	return os.Getenv("ARCH")
}

func getBuiltCommitHash() string {
	return os.Getenv("BUILT_COMMIT_HASH")
}

func getDockerEngineVersion() string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	versionInfo, err := cli.ServerVersion(ctx)
	if err != nil {
		panic(err)
	}
	cli.Close()

	return versionInfo.Version
}

func getVersionLastOrder() (string, string, string) {
	type config struct {
		CodeName  string `toml:"codeName"`
		Version   string `toml:"version"`
		BuiltDate string `toml:"builtDate"`
	}

	c := map[string]config{}

	if _, err := toml.DecodeFile("../../VERSION.toml", &c); err != nil {
		panic(err)
	}

	return c["version"].CodeName, c["version"].Version, c["version"].BuiltDate
}

func getGolangVersion() string {
	version := runtime.Version()
	return version
}

func getOsArchVersion() (string, string) {
	os := runtime.GOOS
	arch := runtime.GOARCH
	return os, arch
}

func getGitCommitHash() string {
	var hash string
	info, err := debug.ReadBuildInfo()
	if !err {
		fmt.Print("Nothing build information")
	}
	for _, s := range info.Settings {
		if s.Key == "vcs.revision" {
			hash = s.Value
		}
	}
	return hash
}
