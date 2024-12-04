package types

import "time"

type ResponseSistersVersion struct {
	CodeName            string    `json:"codeName"`
	Version             string    `json:"version"`
	GolangVersion       string    `json:"golangVersion"`
	DockerEngineVersion string    `json:"dockerEngineVersion"`
	SisterId            string    `json:"sisterId"`
	Name                string    `json:"name"`
	Address             string    `json:"address"`
	BuiltGitCommitHash  string    `json:"builtGitCommitHash"`
	BuiltDate           time.Time `json:"builtDate"`
	Os                  string    `json:"os"`
	Arch                string    `json:"arch"`
}

type ResponseLastOrderVersion struct {
	CodeName            string    `json:"codeName"`
	Version             string    `json:"version"`
	GolangVersion       string    `json:"golangVersion"`
	DockerEngineVersion string    `json:"dockerEngineVersion"`
	BuiltGitCommitHash  string    `json:"builtgitCommitHash"`
	BuiltDate           time.Time `json:"builtDate"`
	Os                  string    `json:"os"`
	Arch                string    `json:"arch"`
}
