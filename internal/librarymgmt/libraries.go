package librarymgmt

const uuidNamespaceString = "9bbc49d0-f9b0-11f0-95e8-325096b39f47"

// Use library tree to fetch uuids and actual library info and metadata
type LibraryRecord struct {
	DisplayName string
	Path        string
	Uuid        string
}

// TODO: add mutex
var AllLibraries []*LibraryRecord
var AllLibrariesMap = make(map[string]*LibraryRecord)
