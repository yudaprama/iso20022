package model

// Choice between source of identification of a financial instrument.
type IdentificationSource3Choice struct {

	// Unique and unambiguous identification source, as assigned via a pre-determined code list.
	Code *ExternalFinancialInstrumentIdentificationType1Code `xml:"Cd"`

	// Unique and unambiguous identification source using a proprietary identification scheme.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (i *IdentificationSource3Choice) SetCode(value string) {
	i.Code = (*ExternalFinancialInstrumentIdentificationType1Code)(&value)
}

func (i *IdentificationSource3Choice) SetProprietary(value string) {
	i.Proprietary = (*Max35Text)(&value)
}
