package model

// Deposited media put in the safe.
type ATMDepositedMedia3 struct {

	// Type of deposited media.
	MediaType *ATMMediaType2Code `xml:"MdiaTp"`

	// Category of deposited media items.
	MediaCategory *ATMMediaType3Code `xml:"MdiaCtgy,omitempty"`

	// Media item that are deposited.
	MediaItems []*ATMDepositedMedia2 `xml:"MdiaItms"`
}

func (a *ATMDepositedMedia3) SetMediaType(value string) {
	a.MediaType = (*ATMMediaType2Code)(&value)
}

func (a *ATMDepositedMedia3) SetMediaCategory(value string) {
	a.MediaCategory = (*ATMMediaType3Code)(&value)
}

func (a *ATMDepositedMedia3) AddMediaItems() *ATMDepositedMedia2 {
	newValue := new(ATMDepositedMedia2)
	a.MediaItems = append(a.MediaItems, newValue)
	return newValue
}
