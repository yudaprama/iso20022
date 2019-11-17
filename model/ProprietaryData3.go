package model

// Container for proprietary information. Business content of this element is not specified.
type ProprietaryData3 struct {

	// Proprietary content.
	Any *SkipProcessing `xml:"Any"`
}

func (p *ProprietaryData3) AddAny() *SkipProcessing {
	p.Any = new(SkipProcessing)
	return p.Any
}
