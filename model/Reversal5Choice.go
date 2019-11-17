package model

// Choice between reversal by reference or by reversal details.
type Reversal5Choice struct {

	// Reference of the transfer confirmation to be reversed.
	References []*TransferReference6 `xml:"Refs"`

	// Details of the transfer out confirmation to be reversed.
	TransferOutConfirmationDetails *TransferOut14 `xml:"TrfOutConfDtls"`
}

func (r *Reversal5Choice) AddReferences() *TransferReference6 {
	newValue := new(TransferReference6)
	r.References = append(r.References, newValue)
	return newValue
}

func (r *Reversal5Choice) AddTransferOutConfirmationDetails() *TransferOut14 {
	r.TransferOutConfirmationDetails = new(TransferOut14)
	return r.TransferOutConfirmationDetails
}
