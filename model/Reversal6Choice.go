package model

// Choice between reversal by reference or by reversal details.
type Reversal6Choice struct {

	// Reference of the transfer confirmation to be reversed.
	Reference *TransferReference6 `xml:"Ref"`

	// Details of the transfer in confirmation to be reversed.
	TransferInConfirmationDetails *TransferIn12 `xml:"TrfInConfDtls"`
}

func (r *Reversal6Choice) AddReference() *TransferReference6 {
	r.Reference = new(TransferReference6)
	return r.Reference
}

func (r *Reversal6Choice) AddTransferInConfirmationDetails() *TransferIn12 {
	r.TransferInConfirmationDetails = new(TransferIn12)
	return r.TransferInConfirmationDetails
}
