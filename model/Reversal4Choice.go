package model

// Choice between reversal by reference or by reversal details.
type Reversal4Choice struct {

	// Reference of the transfer confirmation to be reversed.
	Reference *TransferReference6 `xml:"Ref"`

	// Details of the transfer in confirmation to be reversed.
	TransferInConfirmationDetails *TransferIn9 `xml:"TrfInConfDtls"`
}

func (r *Reversal4Choice) AddReference() *TransferReference6 {
	r.Reference = new(TransferReference6)
	return r.Reference
}

func (r *Reversal4Choice) AddTransferInConfirmationDetails() *TransferIn9 {
	r.TransferInConfirmationDetails = new(TransferIn9)
	return r.TransferInConfirmationDetails
}
