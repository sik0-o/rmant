package rmant

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type markAgent struct {
	redis *redis.Client
	conf  Conf
}

// Key generate key by using prefix from config of agent
func (m *markAgent) Key(subject string) RedisKey {
	return NewRedisKey(m.conf.Prefix, subject)
}

func MarkAgent(redis *redis.Client, conf Conf) Service {
	return &markAgent{
		redis: redis,
		conf:  conf,
	}
}

// Mark implements MarkAgent.
func (m *markAgent) Mark(key RedisKey) error {
	return m.redis.Set(context.Background(), key.String(), m.conf.MarkValue, m.conf.MarkTTL).Err()
}

// Check implements MarkAgent.
func (m *markAgent) Check(key RedisKey) (map[string]uint64, error) {
	var keys []string
	ctx := context.Background()

	iter := m.redis.Scan(ctx, 0, key.Any(), 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}

	points := map[string]uint64{}
	for _, k := range keys {
		subj, ok := key.Parse(k)["entity"]
		if !ok {
			continue
		}

		v, ok := points[subj]
		if !ok {
			v = 0
		}
		points[subj] = v + 1
	}

	return points, nil
}
