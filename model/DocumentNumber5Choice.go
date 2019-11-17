package model

// Choice between a short document number, a long document number or a proprietary document number.
type DocumentNumber5Choice struct {

	// Message type number of the document referenced.
	ShortNumber *Exact3NumericText `xml:"ShrtNb"`

	// MX Message identifier of the referenced document.
	LongNumber *ISO20022MessageIdentificationText `xml:"LngNb"`

	// Proprietary document identification.
	ProprietaryNumber *GenericIdentification36 `xml:"PrtryNb"`
}

func (d *DocumentNumber5Choice) SetShortNumber(value string) {
	d.ShortNumber = (*Exact3NumericText)(&value)
}

func (d *DocumentNumber5Choice) SetLongNumber(value string) {
	d.LongNumber = (*ISO20022MessageIdentificationText)(&value)
}

func (d *DocumentNumber5Choice) AddProprietaryNumber() *GenericIdentification36 {
	d.ProprietaryNumber = new(GenericIdentification36)
	return d.ProprietaryNumber
}
