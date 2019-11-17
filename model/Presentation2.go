package model

// Information for the presentation of documents.
type Presentation2 struct {

	// Party, other than beneficiary, forwarding the documents.
	Presenter *PartyIdentification43 `xml:"Presntr,omitempty"`

	// Date on which the beneficiary presented the demand.
	BeneficiaryPresentationDate *ISODate `xml:"BnfcryPresntnDt,omitempty"`
}

func (p *Presentation2) AddPresenter() *PartyIdentification43 {
	p.Presenter = new(PartyIdentification43)
	return p.Presenter
}

func (p *Presentation2) SetBeneficiaryPresentationDate(value string) {
	p.BeneficiaryPresentationDate = (*ISODate)(&value)
}
