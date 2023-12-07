package entity

type UserPersonId struct {
	PersonIds int `json:"PersonId"`
}

type FindUserReasonIdResponse struct {
	Status string `json:"status"`
	Data   []struct {
		Id         int    `json:"id"`
		PersonId   int    `json:"PersonId"`
		CreateDate string `json:"createDate"`
		DateFrom   string `json:"dateFrom"`
		DateTo     string `json:"dateTo"`
		ReasonId   int    `json:"reasonId"`
	}
}
