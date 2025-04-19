package rmant

import (
	"strings"

	"github.com/google/uuid"
)

type RedisKey interface {
	String() string
	Entity(subject string) string
	Any() string
	Parse(key string) map[string]string
}

type redisKey struct {
	prefix string
	entity string
	uniq   string
}

func (k *redisKey) String() string {
	return strings.Join([]string{k.prefix, k.entity, k.uniq}, ":")
}

func (k *redisKey) Any() string {
	return strings.Join([]string{k.prefix, k.entity, "*"}, ":")
}

func (k *redisKey) Entity(subject string) string {
	return strings.Join([]string{k.prefix, subject, k.uniq}, ":")
}

func (k *redisKey) Parse(key string) map[string]string {
	m := map[string]string{}
	if !strings.HasPrefix(key, k.prefix) {
		return m
	}
	// извлекаем префикс
	m["prefix"] = k.prefix
	// разбиваем оставшийся ключ по `:`
	ka := strings.Split(key[len(k.prefix):], ":")
	// последний элемент - это uniq
	m["uniq"] = ka[len(ka)-1]
	// остальное entity
	ka = ka[:len(ka)-1]
	m["entity"] = strings.Join(ka, ":")

	return m
}

func NewRedisKey(prefix string, subject string) RedisKey {
	return &redisKey{
		prefix: prefix,
		entity: subject,
		uniq:   uuid.NewString(),
	}
}
