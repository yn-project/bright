package version

import (
	"fmt"

	"bytes"
	"runtime"
	"text/template"
	"time"

	"yun.tea/block/bright/proto/bright"

	"golang.org/x/xerrors"
	logger "yun.tea/block/bright/common/logger"
)

func Version() (*bright.VersionResponse, error) {
	info, err := GetVersion()
	if err != nil {
		logger.Sugar().Errorf("get service version error: %+w", err)
		return nil, fmt.Errorf("get service version error: %w", err)
	}
	return &bright.VersionResponse{
		Info: info,
	}, nil
}

var versionTemplate = `
	Version:      {{.Version}}
	Go version:   {{.GoVersion}}
	Built:        {{.BuildTime}}
	OS/Arch:      {{.Os}}/{{.Arch}}
	BranchCommit: {{.Branch}}-{{.Commit}}`

var (
	// Version holds the current version of traefik.
	gitVersion = "0.0.1"
	// BuildDate holds the build date of traefik.
	buildDate = "I don't remember exactly"
	// StartDate holds the start date of traefik.
	StartDate = time.Now()
	// Branch holds the compiled branch
	gitBranch = "master"
	// Commit hold the commit hash of compiled code
	gitCommit = "N/A"
)

func GetVersion() (string, error) {
	tmpl, err := template.New("").Parse(versionTemplate)
	if err != nil {
		return "", xerrors.Errorf("fail to parse version template: %v", err)
	}

	v := struct {
		Version   string
		GoVersion string
		BuildTime string
		Os        string
		Arch      string
		Branch    string
		Commit    string
	}{
		Version:   gitVersion,
		GoVersion: runtime.Version(),
		BuildTime: buildDate,
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		Branch:    gitBranch,
		Commit:    gitCommit,
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, v)
	if err != nil {
		return "", xerrors.Errorf("fail to parse version")
	}

	return buf.String(), nil
}

type VersionType struct {
	GitVersion string `json:"git_version"`
	BuildDate  string `json:"build_date"`
	GitBranch  string `json:"git_branch"`
	GitCommit  string `json:"git_commit"`
}

func (v VersionType) Equal(ver VersionType) bool {
	return v.GitVersion == ver.GitVersion &&
		v.GitBranch == ver.GitBranch &&
		v.GitCommit == ver.GitCommit
}

func MyVersion() VersionType {
	return VersionType{
		GitVersion: gitVersion,
		BuildDate:  buildDate,
		GitBranch:  gitBranch,
		GitCommit:  gitCommit,
	}
}
