package jfstructs

type ResponseSystemInfo struct {
	LocalAddress               string                                       `json:"LocalAddress"`
	ServerName                 string                                       `json:"ServerName"`
	Version                    string                                       `json:"Version"`
	ProductName                string                                       `json:"ProductName"`
	OperatingSystem            string                                       `json:"OperatingSystem"`
	ID                         string                                       `json:"Id"`
	StartupWizardCompleted     bool                                         `json:"StartupWizardCompleted"`
	OperatingSystemDisplayName string                                       `json:"OperatingSystemDisplayName"`
	PackageName                string                                       `json:"PackageName"`
	HasPendingRestart          bool                                         `json:"HasPendingRestart"`
	IsShuttingDown             bool                                         `json:"IsShuttingDown"`
	SupportsLibraryMonitor     bool                                         `json:"SupportsLibraryMonitor"`
	WebSocketPortNumber        int                                          `json:"WebSocketPortNumber"`
	CompletedInstallations     []ResponseSystemInfoCompletedInstallations   `json:"CompletedInstallations"`
	CanSelfRestart             bool                                         `json:"CanSelfRestart"`
	CanLaunchWebBrowser        bool                                         `json:"CanLaunchWebBrowser"`
	ProgramDataPath            string                                       `json:"ProgramDataPath"`
	WebPath                    string                                       `json:"WebPath"`
	ItemsByNamePath            string                                       `json:"ItemsByNamePath"`
	CachePath                  string                                       `json:"CachePath"`
	LogPath                    string                                       `json:"LogPath"`
	InternalMetadataPath       string                                       `json:"InternalMetadataPath"`
	TranscodingTempPath        string                                       `json:"TranscodingTempPath"`
	CastReceiverApplications   []ResponseSystemInfoCastReceiverApplications `json:"CastReceiverApplications"`
	HasUpdateAvailable         bool                                         `json:"HasUpdateAvailable"`
	EncoderLocation            string                                       `json:"EncoderLocation"`
	SystemArchitecture         string                                       `json:"SystemArchitecture"`
}

type ResponseSystemInfoCompletedInstallations struct {
	GUID        string                                              `json:"Guid"`
	Name        string                                              `json:"Name"`
	Version     string                                              `json:"Version"`
	Changelog   string                                              `json:"Changelog"`
	SourceURL   string                                              `json:"SourceUrl"`
	Checksum    string                                              `json:"Checksum"`
	PackageInfo ResponseSystemInfoCompletedInstallationsPackageInfo `json:"PackageInfo"`
}

type ResponseSystemInfoCompletedInstallationsPackageInfo struct {
	Name        string                                             `json:"name"`
	Description string                                             `json:"description"`
	Overview    string                                             `json:"overview"`
	Owner       string                                             `json:"owner"`
	Category    string                                             `json:"category"`
	GUID        string                                             `json:"guid"`
	Versions    []ResponseSystemInfoCompletedInstallationsVersions `json:"versions"`
	ImageURL    string                                             `json:"imageUrl"`
}

type ResponseSystemInfoCompletedInstallationsVersions struct {
	Version        string `json:"version"`
	VersionNumber  string `json:"VersionNumber"`
	Changelog      string `json:"changelog"`
	TargetAbi      string `json:"targetAbi"`
	SourceURL      string `json:"sourceUrl"`
	Checksum       string `json:"checksum"`
	Timestamp      string `json:"timestamp"`
	RepositoryName string `json:"repositoryName"`
	RepositoryURL  string `json:"repositoryUrl"`
}

type ResponseSystemInfoCastReceiverApplications struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
}
