package requests

type PurInfoRecdPrcgCndnValidity struct {
	ConditionRecord              string  `json:"ConditionRecord"`
	ConditionValidityEndDate     *string `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate   *string `json:"ConditionValidityStartDate"`
	ConditionApplication         *string `json:"ConditionApplication"`
	ConditionType                *string `json:"ConditionType"`
	PurgDocOrderQuantityUnit     *string `json:"PurgDocOrderQuantityUnit"`
	PurchasingOrganization       *string `json:"PurchasingOrganization"`
	PurchasingInfoRecordCategory *string `json:"PurchasingInfoRecordCategory"`
	PurchasingInfoRecord         *string `json:"PurchasingInfoRecord"`
	Supplier                     *string `json:"Supplier"`
	MaterialGroup                *string `json:"MaterialGroup"`
	Material                     *string `json:"Material"`
	Plant                        *string `json:"Plant"`
	// ToPurInfoRecdPrcgCndn struct {
	// 	ToPurInfoRecdPrcgCndnResults []*PurInfoRecdPrcgCndn `json:"results"`
	// } `json:"to_PurInfoRecdPrcgCndn"`
}
