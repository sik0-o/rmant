package rmant

type Service interface {
	Mark(key RedisKey) error
	Check(key RedisKey) (map[string]uint64, error)
	Key(subject string) RedisKey
	Del(key RedisKey) error
}
