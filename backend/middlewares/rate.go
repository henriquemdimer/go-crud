package middlewares

import (
	"net/http"
	"time"
)

func RateLimit(next http.Handler) http.Handler {
	type ClientLimit struct {
		Expires      int
		Requests     int
		BlockedUntil int
	}

	limits := make(map[string]ClientLimit)

	RATE := 10000
	REQ_LIMIT := 20
	BLOCK_TIME := 60000

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Real-IP")
		if len(ip) < 1 {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		limit, ok := limits[ip]
		now := int(time.Now().UnixMilli())

		if ok {
			if limit.BlockedUntil > 0 && limit.BlockedUntil > now {
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}

			if limit.Expires < int(time.Now().UnixMilli()) {
				limit.Requests = 0
				limit.Expires = now + RATE
				limit.BlockedUntil = 0
			} else if limit.Requests >= REQ_LIMIT {
				limit.BlockedUntil = now + BLOCK_TIME
				limits[ip] = limit
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}

			limit.Requests += 1
			limits[ip] = limit
		} else {
			limits[ip] = ClientLimit{Expires: now + RATE, Requests: 1}
		}

		next.ServeHTTP(w, r)
	})
}
