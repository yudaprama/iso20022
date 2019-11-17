package model

// Specifies the security option of a corporate event.
type SecuritiesOption56 struct {

	// Maximum quantity of financial instrument that may be instructed.
	MaximumQuantityToInstruct *FinancialInstrumentQuantity21Choice `xml:"MaxQtyToInst,omitempty"`

	// Minimum quantity of financial instrument that may be instructed.
	MinimumQuantityToInstruct *FinancialInstrumentQuantity21Choice `xml:"MinQtyToInst,omitempty"`

	// Minimum multiple quantity of financial  instrument that may be instructed.
	MinimumMultipleQuantityToInstruct *FinancialInstrumentQuantity22Choice `xml:"MinMltplQtyToInst,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity22Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity22Choice `xml:"NewDnmtnQty,omitempty"`

	// Specifies that if an order is prorated holders of odd lots who tender their full position will not have tendered position prorated but rather accepted in full.
	FrontEndOddLotQuantity *FinancialInstrumentQuantity22Choice `xml:"FrntEndOddLotQty,omitempty"`

	// Represents the presence of a back end odd lot provision and the quantity of equity required after proration to be eligible for this privilege.
	BackEndOddLotQuantity *FinancialInstrumentQuantity22Choice `xml:"BckEndOddLotQty,omitempty"`
}

func (s *SecuritiesOption56) AddMaximumQuantityToInstruct() *FinancialInstrumentQuantity21Choice {
	s.MaximumQuantityToInstruct = new(FinancialInstrumentQuantity21Choice)
	return s.MaximumQuantityToInstruct
}

func (s *SecuritiesOption56) AddMinimumQuantityToInstruct() *FinancialInstrumentQuantity21Choice {
	s.MinimumQuantityToInstruct = new(FinancialInstrumentQuantity21Choice)
	return s.MinimumQuantityToInstruct
}

func (s *SecuritiesOption56) AddMinimumMultipleQuantityToInstruct() *FinancialInstrumentQuantity22Choice {
	s.MinimumMultipleQuantityToInstruct = new(FinancialInstrumentQuantity22Choice)
	return s.MinimumMultipleQuantityToInstruct
}

func (s *SecuritiesOption56) AddNewBoardLotQuantity() *FinancialInstrumentQuantity22Choice {
	s.NewBoardLotQuantity = new(FinancialInstrumentQuantity22Choice)
	return s.NewBoardLotQuantity
}

func (s *SecuritiesOption56) AddNewDenominationQuantity() *FinancialInstrumentQuantity22Choice {
	s.NewDenominationQuantity = new(FinancialInstrumentQuantity22Choice)
	return s.NewDenominationQuantity
}

func (s *SecuritiesOption56) AddFrontEndOddLotQuantity() *FinancialInstrumentQuantity22Choice {
	s.FrontEndOddLotQuantity = new(FinancialInstrumentQuantity22Choice)
	return s.FrontEndOddLotQuantity
}

func (s *SecuritiesOption56) AddBackEndOddLotQuantity() *FinancialInstrumentQuantity22Choice {
	s.BackEndOddLotQuantity = new(FinancialInstrumentQuantity22Choice)
	return s.BackEndOddLotQuantity
}
