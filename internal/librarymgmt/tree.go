package librarymgmt

import (
	"sync"
	"time"
)

/*
The library tree is a global tree that stores all library information alongside associated metadata,
Each Node on the tree is either a directory, a tvshow, or a movie.
The metadata subsystem handles the building of the library tree because:
	1. It already has to fully scan all libraries, may as well reuse the results of the scan
	2. It already monitors changes in libraries to automatically update metadata
	3. Inserting the metadata into the tree as each file and directory is being parsed is easier than inserting it after the fact
*/

type TreeNode struct {
	Uuid           string
	RootUuid       string
	Path           string //relative path, combine with library abs path obtained using root uuid to get the real path
	Name           string
	Type           string
	Children       []*TreeNode
	TvshowMetadata *tmdbTvshow
	MovieMetadata  *tmdbMovie

	ImdbId       string
	Mtime        time.Time
	ReleasedTime time.Time
}

// TODO: most functions dont respect the mutex but atleast they only read
var LibTreeMutex sync.RWMutex
var LibTree []*TreeNode

// key is uuid
var LibTreeMap = make(map[string]*TreeNode)

type MtimeSortEntry struct {
	Node  *TreeNode
	Mtime time.Time
}

var NewestMedia = make(map[string][]MtimeSortEntry)
