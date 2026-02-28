package config

//TODO: change folder terminology to directory

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type ConfigNode struct {
	Name       string
	Children   []ConfigNode
	Properties []string
}

var ConfigDirectoryPath string = "./config"
var RootConfigNode ConfigNode

func Init() {
	//TODO: use flags module instead
	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg, "--config-dir=") {
			continue
		}

		dirStrSplit := strings.Split(arg, "=")
		ConfigDirectoryPath = dirStrSplit[len(dirStrSplit)-1]
	}

	slog.Info("parsing config directory:", "path", ConfigDirectoryPath)
	RootConfigNode = LoadConfig(ConfigDirectoryPath, "/")

	if slog.Default().Enabled(context.Background(), slog.LevelDebug) {
		PrintConfigTree(RootConfigNode)
	}
}

func LoadConfig(filePath string, name string) ConfigNode {
	var retNode ConfigNode
	retNode.Name = name

	fileStat, err := os.Lstat(filePath)
	if err != nil {
		slog.Error("failed to parse config", "file", filePath, "reason", err.Error())
		os.Exit(1)
	}

	fileMode := fileStat.Mode()

	if !fileMode.IsDir() && (filePath == ConfigDirectoryPath) {
		slog.Error("config directory path does not lead to a directory", "path", ConfigDirectoryPath)
		os.Exit(1)
	}

	if fileMode&os.ModeSymlink != 0 {
		slog.Error("symlinks are currently not supported in the config tree", "offendingFile", filePath)
		os.Exit(1)
	}

	if fileMode.IsDir() {
		retNode.Properties = append(retNode.Properties, "folder=1")
		subDirEntries, err := os.ReadDir(filePath)
		if err != nil {
			slog.Error("failed to open directory", "path", filePath, "reason", err.Error())
			os.Exit(1)
		}

		for _, subFile := range subDirEntries {
			var childName string
			if filePath == ConfigDirectoryPath {
				childName = "/" + subFile.Name()
			} else {
				childName = name + "/" + subFile.Name()
			}
			retNode.Children = append(retNode.Children, LoadConfig(filePath+"/"+subFile.Name(), childName))
		}
		return retNode
	}

	if !fileMode.IsRegular() {
		slog.Error("unable to determine filetype", "file", filePath)
		os.Exit(1)
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("failed to open config file", "file", filePath, "reason", err.Error())
	}

	fileString := string(fileBytes)
	retNode.Properties = strings.Split(fileString, "\n")

	retNode.Properties = slices.DeleteFunc(retNode.Properties, func(property string) bool {
		return strings.HasPrefix(property, "#")
	})

	return retNode
}

func PrintConfigTree(node ConfigNode) {
	if NodeIsFolder(&node) {
		slog.Debug("listing config directory", "dir", node.Name)
		for _, child := range node.Children {
			PrintConfigTree(child)
		}
		return
	}

	slog.Debug("listing config node properties", "file", node.Name)

	for i, property := range node.Properties {
		slog.Debug("property", "index", i, "value", property)
	}
}

func NodeIsFolder(node *ConfigNode) bool {
	return slices.Contains(node.Properties, "folder=1")
}

func FindNode(path string, startNode *ConfigNode) *ConfigNode {
	if startNode.Name == path {
		return startNode
	}

	if !NodeIsFolder(startNode) {
		return nil
	}

	for _, child := range startNode.Children {
		foundNode := FindNode(path, &child)
		if foundNode != nil {
			return foundNode
		}
	}

	return nil
}

//TODO: return errors not strings

func FetchValue(path string, key string, errorOnFail bool) string {
	pathNode := FindNode(path, &RootConfigNode)
	if pathNode != nil {
		slog.Debug("found config node", "path", path)
	} else {
		slog.Error("unable to find config node", "configPath", path)
		if errorOnFail {
			os.Exit(1)
		}
		return "ERROR_COULD_NOT_FIND_NODE"
	}

	for _, property := range pathNode.Properties {
		if strings.HasPrefix(property, key+"=") {
			val := strings.SplitN(property, "=", 2)[1]
			val = strings.TrimSpace(val)
			val = strings.Trim(val, `"'`)
			return val
		}
	}

	absPath := filepath.Join(ConfigDirectoryPath, path)
	slog.Error("config node doesnt not contain key", "path", absPath, "key", key)
	if errorOnFail {
		os.Exit(1)
	}
	return "ERROR_NODE_DOES_NOT_CONTAIN_KEY"
}
