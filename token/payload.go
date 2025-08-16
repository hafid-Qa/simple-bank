package token

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	expiredAt := now.Add(duration)

	payload := &Payload{
		Id:        id,
		Username:  username,
		IssuedAt:  now,
		ExpiredAt: expiredAt,
	}

	return payload, nil
}
