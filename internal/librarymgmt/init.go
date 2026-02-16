package librarymgmt

import (
	"fmt"
	"opal/internal/config"
	"os"
	"strings"
)

var CacheDir string

func Init() {
	CacheDir = config.FetchValue("/server.cfg", "cache_dir", true)

	libraryConfigList := config.FindNode("/libraries", &config.RootConfigNode)
	if libraryConfigList == nil {
		fmt.Println("Error: unable to find /libraries in config directory")
		os.Exit(1)
	}

	for _, libConfig := range libraryConfigList.Children {
		fmt.Printf("Parsing %s\n", libConfig.Name)

		newLib := &LibraryRecord{
			DisplayName: config.FetchValue(libConfig.Name, "display_name", false),
			Path:        config.FetchValue(libConfig.Name, "path", false),
		}

		if strings.HasPrefix(newLib.DisplayName, "ERROR") || strings.HasPrefix(newLib.Path, "ERROR") {
			fmt.Printf("Error in library config, skipping")
			continue
		}

		AllLibraries = append(AllLibraries, newLib)
	}

	initMetadata()
}
