package server

import (
	"context"
	"fmt"
	"log"
	"protei/internal/entity"
	proteigrpc "protei/internal/user"
	"time"
)

var emoji = map[int]string{
	1:  "🏠",
	2:  "",
	3:  "✈",
	4:  "✈",
	5:  "🌡",
	6:  "🌡",
	7:  "",
	8:  "",
	9:  "🎓",
	10: "🏠",
	11: "☀️",
	12: "☀️",
	13: "☀️",
}

func (s *Server) FindUser(ctx context.Context, in *proteigrpc.UserRequest) (*proteigrpc.UserReply, error) {
	log.Printf("Received: %v", in.GetMail())
	u := entity.UserMailIn{Mail: in.GetMail()}
	f, err := s.c.HandlerFindUserId(u)
	if err != nil {
		return nil, fmt.Errorf("can't create FindUserId handler: %w", err)
	}
	if len(f.Data) == 0 {
		return nil, fmt.Errorf("user not found")
	}
	r, err := s.c.HandlerFindReasonId(&entity.UserPersonId{PersonIds: f.Data[0].Id})
	for i := 0; i < len(r.Data); i++ {
		dataFrom, err := time.Parse("2006-01-02T15:04:05", r.Data[i].DateFrom)
		if err != nil {
			return nil, fmt.Errorf("can't parse dataFrom: %w", err)
		}
		dataTo, err := time.Parse("2006-01-02T15:04:05", r.Data[i].DateTo)
		if err != nil {
			return nil, fmt.Errorf("can't parse dataFrom: %w", err)
		}
		if (dataFrom.Before(time.Now())) && (dataTo.After(time.Now())) {
			return &proteigrpc.UserReply{DisplayName: f.Data[0].DisplayName + " " + emoji[r.Data[i].ReasonId]}, nil
		}
	}
	return nil, nil
}
