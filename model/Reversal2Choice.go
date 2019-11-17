package model

// Choice between reversal by reference or by reversal details.
type Reversal2Choice struct {

	// Reference of the transfer confirmation to be reversed.
	Reference *TransferReference2 `xml:"Ref"`

	// Details of the transfer in confirmation to be reversed.
	TransferInConfirmationDetails *TransferIn8 `xml:"TrfInConfDtls"`
}

func (r *Reversal2Choice) AddReference() *TransferReference2 {
	r.Reference = new(TransferReference2)
	return r.Reference
}

func (r *Reversal2Choice) AddTransferInConfirmationDetails() *TransferIn8 {
	r.TransferInConfirmationDetails = new(TransferIn8)
	return r.TransferInConfirmationDetails
}
