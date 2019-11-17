package model

// Information about a database.
type DataBaseCheck1 struct {

	// Indicates whether the individual or organisation is listed in an on-line global Know Your Customer (KYC) database.
	DatabaseCheck *YesNoIndicator `xml:"DBChck"`

	// Identification of the database.
	Identification *Max35Text `xml:"Id"`
}

func (d *DataBaseCheck1) SetDatabaseCheck(value string) {
	d.DatabaseCheck = (*YesNoIndicator)(&value)
}

func (d *DataBaseCheck1) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}
