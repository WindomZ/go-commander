package commander

type Version struct {
	version string
}

func (v *Version) Set(ver string) *Version {
	v.version = ver
	return v
}

func (v Version) Get() string {
	if len(v.version) != 0 {
		return v.version
	}
	return "v0.0.0"
}

func (v Version) Valid() bool {
	return len(v.Get()) != 0
}
