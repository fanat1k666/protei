package ems

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"protei/internal/entity"
)

func (c *Client) HandlerFindUserId(u entity.UserMailIn) (*entity.FindUserIdResponseOut, error) {
	client := http.Client{}
	body, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("can't marshal FindUserId requset: %w", err)
	}
	b := bytes.NewBuffer(body)
	req, err := http.NewRequest(http.MethodPost, c.host+employeesPath, b)
	if err != nil {
		return nil, fmt.Errorf("cant't create request: %w", err)
	}
	req.Header.Set("Authorization", c.getBasicAuth())
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cant't do request: %w", err)
	}
	f := entity.FindUserIdResponseOut{}
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
