package librarymgmt

import (
	"log/slog"
	"opal/internal/config"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
)

var CacheDir string

func Init() {
	CacheDir = config.FetchValue("/server.cfg", "cache_dir", true)

	libraryConfigList := config.FindNode("/libraries", &config.RootConfigNode)
	if libraryConfigList == nil {
		slog.Error("unable to find /libraries in config directory")
		os.Exit(1)
	}

	for _, libConfig := range libraryConfigList.Children {
		slog.Info("Loading library", "configPath", libConfig.Name)

		newLib := &LibraryRecord{
			DisplayName: config.FetchValue(libConfig.Name, "display_name", false),
			Path:        config.FetchValue(libConfig.Name, "path", false),
		}

		if strings.HasPrefix(newLib.DisplayName, "ERROR") || strings.HasPrefix(newLib.Path, "ERROR") {
			slog.Error("Error in library config, skipping", "libName", libConfig.Name)
			continue
		}

		AllLibraries = append(AllLibraries, newLib)
	}

	initNameCardRenderer()
	initMetadata()

	nameCardFont = nil
	runtime.GC()
	debug.FreeOSMemory() //Forces runtime to free up all the memory it allocated when rendering the name cards (which can often be +50MiB)
}
