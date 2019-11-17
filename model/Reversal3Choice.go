package model

// Choice between reversal by reference or by reversal details.
type Reversal3Choice struct {

	// Reference of the transfer confirmation to be reversed.
	References []*TransferReference6 `xml:"Refs"`

	// Details of the transfer out confirmation to be reversed.
	TransferOutConfirmationDetails *TransferOut12 `xml:"TrfOutConfDtls"`
}

func (r *Reversal3Choice) AddReferences() *TransferReference6 {
	newValue := new(TransferReference6)
	r.References = append(r.References, newValue)
	return newValue
}

func (r *Reversal3Choice) AddTransferOutConfirmationDetails() *TransferOut12 {
	r.TransferOutConfirmationDetails = new(TransferOut12)
	return r.TransferOutConfirmationDetails
}
