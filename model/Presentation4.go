package model

// Information for the presentation of documents.
type Presentation4 struct {

	// Medium through which the presentation can be submitted such as paper, electronic or both.
	Medium *PresentationMedium1Choice `xml:"Mdm,omitempty"`

	// Document required to be presented.
	Document []*Document11 `xml:"Doc,omitempty"`

	// Additional information related to the presentation.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (p *Presentation4) AddMedium() *PresentationMedium1Choice {
	p.Medium = new(PresentationMedium1Choice)
	return p.Medium
}

func (p *Presentation4) AddDocument() *Document11 {
	newValue := new(Document11)
	p.Document = append(p.Document, newValue)
	return newValue
}

func (p *Presentation4) AddAdditionalInformation(value string) {
	p.AdditionalInformation = append(p.AdditionalInformation, (*Max2000Text)(&value))
}
