package model

// Foreign Account Tax Compliance Act (FATCA) status information.
type FATCAStatus2 struct {

	// Foreign Account Tax Compliance Act (FATCA) status.
	Type *FATCAStatus2Choice `xml:"Tp"`

	// Source of the Foreign Account Tax Compliance Act (FATCA) status.
	Source *FATCASource1Choice `xml:"Src,omitempty"`
}

func (f *FATCAStatus2) AddType() *FATCAStatus2Choice {
	f.Type = new(FATCAStatus2Choice)
	return f.Type
}

func (f *FATCAStatus2) AddSource() *FATCASource1Choice {
	f.Source = new(FATCASource1Choice)
	return f.Source
}
