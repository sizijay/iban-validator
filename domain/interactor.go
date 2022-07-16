package domain

import "context"

type Service interface {
	Do(ctx context.Context, arg []interface{}) ([]interface{}, error)
}