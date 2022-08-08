package dto

type ApplicationCredit struct {
	Description string `json:"description,omitempty"`
	ID          int64  `json:"id,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Test        bool   `json:"test"`
}

type ApplicationCreditCollection struct {
	ApplicationCredit ApplicationCredit `json:"application_credit"`
}

type GetApplicationCreditRequest struct {
	Fields string `json:"fields"`
}

type GetApplicationCreditResponse struct {
	ApplicationCredits []ApplicationCredit `json:"application_credits"`
}
