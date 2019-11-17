package model

// Choice of identification of a party. The party can be identified by providing a BIC or a MIC.
type PartyIdentification24Choice struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICOrBEI *AnyBICIdentifier `xml:"BICOrBEI"`

	// Market Identifier Code. Identification of a financial market, as stipulated in the norm ISO 10383 "Codes for exchanges and market identifications".
	MIC *MICIdentifier `xml:"MIC"`
}

func (p *PartyIdentification24Choice) SetBICOrBEI(value string) {
	p.BICOrBEI = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification24Choice) SetMIC(value string) {
	p.MIC = (*MICIdentifier)(&value)
}
