package ems

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"protei/internal/entity"
)

func (c *Client) HandlerFindReasonId(u *entity.UserPersonId) (*entity.FindUserReasonIdResponse, error) {
	body, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("can't marshal: %w", err)
	}
	b2 := bytes.NewBuffer(body)
	client := http.Client{}
	req, _ := http.NewRequest(http.MethodPost, c.host+absencesPath, b2)
	req.Header.Set("Authorization", c.getBasicAuth())
	res, _ := client.Do(req)
	f := entity.FindUserReasonIdResponse{}
	str, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read str: %w", err)
	}
	err = json.Unmarshal(str, &f)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal: %w", err)
	}
	return &f, nil
}
