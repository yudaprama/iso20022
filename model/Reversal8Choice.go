package model

// Choice between reversal by reference or by reversal details.
type Reversal8Choice struct {

	// Reference of the transfer confirmation to be reversed.
	References []*TransferReference10 `xml:"Refs"`

	// Details of the transfer out confirmation to be reversed.
	TransferOutConfirmationDetails *TransferOut16 `xml:"TrfOutConfDtls"`
}

func (r *Reversal8Choice) AddReferences() *TransferReference10 {
	newValue := new(TransferReference10)
	r.References = append(r.References, newValue)
	return newValue
}

func (r *Reversal8Choice) AddTransferOutConfirmationDetails() *TransferOut16 {
	r.TransferOutConfirmationDetails = new(TransferOut16)
	return r.TransferOutConfirmationDetails
}
