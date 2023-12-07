package entity

type UserMailIn struct {
	Mail string `json:"mail"`
}

type FindUserIdResponseOut struct {
	Status string `json:"status"`
	Data   []struct {
		Id          int    `json:"id"`
		DisplayName string `json:"displayName"`
		Email       string `json:"email"`
		WorkPhone   string `json:"workPhone"`
	}
}
