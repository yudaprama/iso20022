package model

// Choice between reversal by reference or by reversal details.
type Reversal7Choice struct {

	// Reference of the transfer confirmation to be reversed.
	Reference *TransferReference10 `xml:"Ref"`

	// Details of the transfer in confirmation to be reversed.
	TransferInConfirmationDetails *TransferIn14 `xml:"TrfInConfDtls"`
}

func (r *Reversal7Choice) AddReference() *TransferReference10 {
	r.Reference = new(TransferReference10)
	return r.Reference
}

func (r *Reversal7Choice) AddTransferInConfirmationDetails() *TransferIn14 {
	r.TransferInConfirmationDetails = new(TransferIn14)
	return r.TransferInConfirmationDetails
}
