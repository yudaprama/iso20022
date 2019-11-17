package model

// Provides information about the corporate action event.
type CorporateAction13 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate30 `xml:"DtDtls,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage *CorporateActionEventStageFormat6Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat1Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction13) AddDateDetails() *CorporateActionDate30 {
	c.DateDetails = new(CorporateActionDate30)
	return c.DateDetails
}

func (c *CorporateAction13) AddEventStage() *CorporateActionEventStageFormat6Choice {
	c.EventStage = new(CorporateActionEventStageFormat6Choice)
	return c.EventStage
}

func (c *CorporateAction13) AddLotteryType() *LotteryTypeFormat1Choice {
	c.LotteryType = new(LotteryTypeFormat1Choice)
	return c.LotteryType
}
