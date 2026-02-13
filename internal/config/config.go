package config

//TODO: change folder terminology to directory

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type ConfigNode struct {
	Name       string
	Children   []ConfigNode
	Properties []string
}

var configDirectoryPath string = "./config"
var RootConfigNode ConfigNode

func Init() {
	//TODO: use flags module instead
	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg, "--config-dir=") {
			continue
		}

		dirStrSplit := strings.Split(arg, "=")
		configDirectoryPath = dirStrSplit[len(dirStrSplit)-1]
	}

	fmt.Println("Config directory:", configDirectoryPath)
	RootConfigNode = LoadConfig(configDirectoryPath, "/")
	PrintConfigTree(RootConfigNode)
}

func LoadConfig(filePath string, name string) ConfigNode {
	var retNode ConfigNode
	retNode.Name = name

	fileStat, err := os.Lstat(filePath)
	if err != nil {
		fmt.Printf("Config parsing error: %s\n", err.Error())
		os.Exit(1)
	}

	fileMode := fileStat.Mode()

	if !fileMode.IsDir() && (filePath == configDirectoryPath) {
		fmt.Printf("Config parsing error: %s is not a directory\n", configDirectoryPath)
		os.Exit(1)
	}

	if fileMode&os.ModeSymlink != 0 {
		fmt.Println("Config parsing error: symlinks currently not supported in config tree")
		os.Exit(1)
	}

	if fileMode.IsDir() {
		retNode.Properties = append(retNode.Properties, "folder=1")
		subDirEntries, err := os.ReadDir(filePath)
		if err != nil {
			fmt.Printf("Config parsing error: %s\n", err.Error())
			os.Exit(1)
		}

		for _, subFile := range subDirEntries {
			var childName string
			if filePath == configDirectoryPath {
				childName = "/" + subFile.Name()
			} else {
				childName = name + "/" + subFile.Name()
			}
			retNode.Children = append(retNode.Children, LoadConfig(filePath+"/"+subFile.Name(), childName))
		}
		return retNode
	}

	if !fileMode.IsRegular() {
		fmt.Printf("Config parsing error: unable to determine filetype of %s\n", filePath)
		os.Exit(1)
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Config parsing error: %s\n", err.Error())
	}

	fileString := string(fileBytes)
	retNode.Properties = strings.Split(fileString, "\n")

	retNode.Properties = slices.DeleteFunc(retNode.Properties, func(property string) bool {
		return strings.HasPrefix(property, "#")
	})

	return retNode
}

func PrintConfigTree(node ConfigNode) {
	fmt.Printf("%s\n", node.Name)

	if NodeIsFolder(&node) {
		for _, child := range node.Children {
			PrintConfigTree(child)
		}
		return
	}

	for i, property := range node.Properties {
		fmt.Printf("%v: %s\n", i, property)
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

func FetchValue(path string, key string, errorOnFail bool) string {
	pathNode := FindNode(path, &RootConfigNode)
	if pathNode != nil {
		//fmt.Printf("Found node %s @ %p\n", path, pathNode)
	} else {
		fmt.Printf("Config fetching error: unable to find %s in config directory\n", path)
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

	if errorOnFail {
		fmt.Printf("Config fetching error: unable to find %s in %s\n", key, configDirectoryPath+path)
		os.Exit(1)
	}
	return "ERROR_NODE_DOES_NOT_CONTAIN_KEY"
}
