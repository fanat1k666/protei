package ems

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

const (
	absencesPath  = "/Portal/springApi/api/absences"
	employeesPath = "/Portal/springApi/api/employees"
)

type Client struct {
	c        *http.Client
	host     string
	login    string
	password string
}

func NewClient(c *http.Client, h string, l string, p string) *Client {
	return &Client{
		c:        c,
		host:     h,
		login:    l,
		password: p,
	}
}

func (c *Client) getBasicAuth() string {
	return "Basic" + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.login, c.password)))
}
