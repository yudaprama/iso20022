package model

// Provides information about the corporate action event.
type CorporateAction34 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate49 `xml:"DtDtls,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage *CorporateActionEventStageFormat14Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat4Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction34) AddDateDetails() *CorporateActionDate49 {
	c.DateDetails = new(CorporateActionDate49)
	return c.DateDetails
}

func (c *CorporateAction34) AddEventStage() *CorporateActionEventStageFormat14Choice {
	c.EventStage = new(CorporateActionEventStageFormat14Choice)
	return c.EventStage
}

func (c *CorporateAction34) AddLotteryType() *LotteryTypeFormat4Choice {
	c.LotteryType = new(LotteryTypeFormat4Choice)
	return c.LotteryType
}
