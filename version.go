package main

import (
	"runtime/debug"

	"github.com/carlmjohnson/versioninfo"
)

func getVersion() string {
	if version == "" {
		version = versioninfo.Revision

		if versioninfo.DirtyBuild {
			version += "-dirty"
		}
	}

	buildinfo, ok := debug.ReadBuildInfo()

	if ok && (buildinfo != nil) && (buildinfo.Main.Version != "(devel)") {
		version = buildinfo.Main.Version
	}

	return version
}
