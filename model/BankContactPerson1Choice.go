package model

// Contact person details in a bank.
type BankContactPerson1Choice struct {

	// Person to be contacted in the buyer's bank.
	BuyerBankContactPerson []*ContactIdentification1 `xml:"BuyrBkCtctPrsn"`

	// Person to be contacted in the seller's bank.
	SellerBankContactPerson []*ContactIdentification1 `xml:"SellrBkCtctPrsn"`
}

func (b *BankContactPerson1Choice) AddBuyerBankContactPerson() *ContactIdentification1 {
	newValue := new(ContactIdentification1)
	b.BuyerBankContactPerson = append(b.BuyerBankContactPerson, newValue)
	return newValue
}

func (b *BankContactPerson1Choice) AddSellerBankContactPerson() *ContactIdentification1 {
	newValue := new(ContactIdentification1)
	b.SellerBankContactPerson = append(b.SellerBankContactPerson, newValue)
	return newValue
}
