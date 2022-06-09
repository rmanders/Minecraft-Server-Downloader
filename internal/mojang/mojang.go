package mojang

type VersionInfoLatest struct {
	Release  string
	Snapshot string
}

type VersionInfo struct {
	Id          string
	Type        string
	Url         string
	Time        string
	ReleaseTime string
}

type Versions struct {
	Latest   VersionInfoLatest
	Versions []VersionInfo
}

type PackageDownloadInfo struct {
	Sha1 string
	Size int
	Url  string
}

type PackageDownloads struct {
	Server PackageDownloadInfo
}

type PackageMetadata struct {
	Downloads PackageDownloads
}
