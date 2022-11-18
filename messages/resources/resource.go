package resources

type Resource struct {
	ResourceName string `json:"-"`
	FullPath     string `json:"-"`
	RelativePath string `json:"Path"`
	Checksum     string `json:"Checksum"`
}
