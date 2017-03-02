package commander

type Usage struct {
	usage string
}

func (u *Usage) Set(usage string) *Usage {
	u.usage = usage
	return u
}

func (u Usage) Get() string {
	return u.usage
}

func (u Usage) Valid() bool {
	return len(u.usage) != 0
}
