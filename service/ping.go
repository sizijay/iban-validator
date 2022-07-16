package service

import (
	"context"
	"time"
)

type PingService struct{}

func (i *PingService) Do(ctx context.Context, arg []interface{}) ([]interface{}, error) {
	var pinging = make([]interface{}, 1)
	pinging[0] = "P o n g! @ " + time.Now().Format("2006-01-02 15:04:05") + " from : github.com/sizijay/iban-validator"

	return pinging, nil
}
