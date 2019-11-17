package model

// Choice between source of identification of a financial instrument.
type IdentificationSource4Choice struct {

	// Unique and unambiguous identification source, as assigned via a pre-determined code list.
	Code *ExternalFinancialInstrumentIdentificationType1Code `xml:"Cd"`

	// Unique and unambiguous identification source using a proprietary identification scheme.
	Proprietary *RestrictedFINExact2Text `xml:"Prtry"`
}

func (i *IdentificationSource4Choice) SetCode(value string) {
	i.Code = (*ExternalFinancialInstrumentIdentificationType1Code)(&value)
}

func (i *IdentificationSource4Choice) SetProprietary(value string) {
	i.Proprietary = (*RestrictedFINExact2Text)(&value)
}
