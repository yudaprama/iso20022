package model

// Container for proprietary information. Business content of this element is not specified.
type ProprietaryData4 struct {

	// Specifies the type of proprietary document
	Type *Max35Text `xml:"Tp"`

	// Proprietary data content.
	Data *ProprietaryData3 `xml:"Data"`
}

func (p *ProprietaryData4) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryData4) AddData() *ProprietaryData3 {
	p.Data = new(ProprietaryData3)
	return p.Data
}
