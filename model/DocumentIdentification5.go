package model

// Identifies a document by a unique identification and its issuer.
type DocumentIdentification5 struct {

	// Identification of a set of data.
	Identification *Max35Text `xml:"Id"`

	// Uniquely identifies the financial institution which has issued the identification of the set of data by using a BIC.
	IdentificationIssuer *BICIdentification1 `xml:"IdIssr"`
}

func (d *DocumentIdentification5) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification5) AddIdentificationIssuer() *BICIdentification1 {
	d.IdentificationIssuer = new(BICIdentification1)
	return d.IdentificationIssuer
}
