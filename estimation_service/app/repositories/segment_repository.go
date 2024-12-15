package repositories

import (
	"context"
	"time"
)

func (obj *UserSegmentPair) Store() (err error) {
	ctx := context.Background()

	// Key format: segment:user_id
	key := obj.Segment + ":" + obj.UserId
	// 2 weeks expire time
	duration := 14 * 24 * time.Hour

	err = Connection.Set(ctx, key, time.Now().Unix(), duration).Err()

	return
}

func (obj *UserSegmentPair) GetSegmentCount(segment string) (count int, err error) {
	ctx := context.Background()

	// Scan for keys matching the segment
	pattern := segment + ":*"

	iter := Connection.Scan(ctx, 0, pattern, 0).Iterator()

	err = iter.Err()
	if err != nil {
		return 0, err
	}

	count = 0
	for iter.Next(ctx) {
		count++
	}

	return
}
