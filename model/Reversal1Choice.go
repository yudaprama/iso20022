package model

// Choice between reversal by reference or by reversal details.
type Reversal1Choice struct {

	// Reference of the transfer confirmation to be reversed.
	Reference *TransferReference2 `xml:"Ref"`

	// Details of the transfer out confirmation to be reversed.
	TransferOutConfirmationDetails *TransferOut10 `xml:"TrfOutConfDtls"`
}

func (r *Reversal1Choice) AddReference() *TransferReference2 {
	r.Reference = new(TransferReference2)
	return r.Reference
}

func (r *Reversal1Choice) AddTransferOutConfirmationDetails() *TransferOut10 {
	r.TransferOutConfirmationDetails = new(TransferOut10)
	return r.TransferOutConfirmationDetails
}
