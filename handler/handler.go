package handler

import (
	"github.com/bradfitz/gomemcache/memcache"
)

type Handler struct {
	Cache *memcache.Client
}
