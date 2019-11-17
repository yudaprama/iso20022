package model

// Specifies the security option of a corporate event.
type SecuritiesOption51 struct {

	// Maximum quantity of financial instrument that may be instructed.
	MaximumQuantityToInstruct *FinancialInstrumentQuantity19Choice `xml:"MaxQtyToInst,omitempty"`

	// Minimum quantity of financial instrument that may be instructed.
	MinimumQuantityToInstruct *FinancialInstrumentQuantity19Choice `xml:"MinQtyToInst,omitempty"`

	// Minimum multiple quantity of financial  instrument that may be instructed.
	MinimumMultipleQuantityToInstruct *FinancialInstrumentQuantity20Choice `xml:"MinMltplQtyToInst,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity20Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity20Choice `xml:"NewDnmtnQty,omitempty"`

	// Specifies that if an order is prorated holders of odd lots who tender their full position will not have tendered position prorated but rather accepted in full.
	FrontEndOddLotQuantity *FinancialInstrumentQuantity20Choice `xml:"FrntEndOddLotQty,omitempty"`

	// Represents the presence of a back end odd lot provision and the quantity of equity required after proration to be eligible for this privilege.
	BackEndOddLotQuantity *FinancialInstrumentQuantity20Choice `xml:"BckEndOddLotQty,omitempty"`
}

func (s *SecuritiesOption51) AddMaximumQuantityToInstruct() *FinancialInstrumentQuantity19Choice {
	s.MaximumQuantityToInstruct = new(FinancialInstrumentQuantity19Choice)
	return s.MaximumQuantityToInstruct
}

func (s *SecuritiesOption51) AddMinimumQuantityToInstruct() *FinancialInstrumentQuantity19Choice {
	s.MinimumQuantityToInstruct = new(FinancialInstrumentQuantity19Choice)
	return s.MinimumQuantityToInstruct
}

func (s *SecuritiesOption51) AddMinimumMultipleQuantityToInstruct() *FinancialInstrumentQuantity20Choice {
	s.MinimumMultipleQuantityToInstruct = new(FinancialInstrumentQuantity20Choice)
	return s.MinimumMultipleQuantityToInstruct
}

func (s *SecuritiesOption51) AddNewBoardLotQuantity() *FinancialInstrumentQuantity20Choice {
	s.NewBoardLotQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.NewBoardLotQuantity
}

func (s *SecuritiesOption51) AddNewDenominationQuantity() *FinancialInstrumentQuantity20Choice {
	s.NewDenominationQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.NewDenominationQuantity
}

func (s *SecuritiesOption51) AddFrontEndOddLotQuantity() *FinancialInstrumentQuantity20Choice {
	s.FrontEndOddLotQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.FrontEndOddLotQuantity
}

func (s *SecuritiesOption51) AddBackEndOddLotQuantity() *FinancialInstrumentQuantity20Choice {
	s.BackEndOddLotQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.BackEndOddLotQuantity
}
