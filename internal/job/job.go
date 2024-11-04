package job

import "context"

type Job interface {
	Run(ctx context.Context, payload any) error
	GetName() string
}

var Jobs map[string]Job = make(map[string]Job)
