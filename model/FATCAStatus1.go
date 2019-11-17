package model

// Foreign Account Tax Compliance Act (FATCA) status information.
type FATCAStatus1 struct {

	// Foreign Account Tax Compliance Act (FATCA) status.
	Type *FATCAStatus1Choice `xml:"Tp"`

	// Source of the Foreign Account Tax Compliance Act (FATCA) status.
	Source *FATCASource1Choice `xml:"Src,omitempty"`
}

func (f *FATCAStatus1) AddType() *FATCAStatus1Choice {
	f.Type = new(FATCAStatus1Choice)
	return f.Type
}

func (f *FATCAStatus1) AddSource() *FATCASource1Choice {
	f.Source = new(FATCASource1Choice)
	return f.Source
}
