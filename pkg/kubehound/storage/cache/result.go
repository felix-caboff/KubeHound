package cache

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrNoEntry     = errors.New("no matching cache entry")
	ErrInvalidType = errors.New("cache entry value cannot be converted to requested type")
)

// CacheResult provides syntactic sugar around retrieval and type casting of entries from the cache.
type CacheResult struct {
	Value any
	Err   error
}

// Text returns the result value as a string alongside any errors.
func (r *CacheResult) Text() (string, error) {
	if r.Err != nil {
		return "", r.Err
	}

	if r.Value == nil {
		return "", ErrNoEntry
	}

	s, ok := r.Value.(string)
	if !ok {
		return "", ErrInvalidType
	}

	return s, nil
}

// Int64 returns the result value as a int64 alongside any errors.
func (r *CacheResult) Int64() (int64, error) {
	if r.Err != nil {
		return -1, r.Err
	}

	if r.Value == nil {
		return -1, ErrNoEntry
	}

	i, ok := r.Value.(int64)
	if !ok {
		return -1, ErrInvalidType
	}

	return i, nil
}

// ObjectID returns the result value as a bson ObjectID alongside any errors.
func (r *CacheResult) ObjectID() (primitive.ObjectID, error) {
	if r.Err != nil {
		return primitive.NilObjectID, r.Err
	}

	if r.Value == nil {
		return primitive.NilObjectID, ErrNoEntry
	}

	raw, ok := r.Value.(string)
	if !ok {
		return primitive.NilObjectID, ErrInvalidType
	}

	oid, err := primitive.ObjectIDFromHex(raw)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return oid, nil
}
