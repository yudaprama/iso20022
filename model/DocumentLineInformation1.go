package model

// Provides document line information.
//
type DocumentLineInformation1 struct {

	// Provides identification of the document line.
	Identification []*DocumentLineIdentification1 `xml:"Id"`

	// Description associated with the document line.
	Description *Max2048Text `xml:"Desc,omitempty"`

	// Provides details on the amounts of the document line.
	Amount *RemittanceAmount3 `xml:"Amt,omitempty"`
}

func (d *DocumentLineInformation1) AddIdentification() *DocumentLineIdentification1 {
	newValue := new(DocumentLineIdentification1)
	d.Identification = append(d.Identification, newValue)
	return newValue
}

func (d *DocumentLineInformation1) SetDescription(value string) {
	d.Description = (*Max2048Text)(&value)
}

func (d *DocumentLineInformation1) AddAmount() *RemittanceAmount3 {
	d.Amount = new(RemittanceAmount3)
	return d.Amount
}
