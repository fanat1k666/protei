package server

import (
	"protei/internal/contracts"
	proteigrpc "protei/internal/user"
)

type Server struct {
	proteigrpc.UnimplementedUserSearchServer
	c contracts.Employees
}

func NewServer(employees contracts.Employees) *Server {
	return &Server{c: employees}
}
