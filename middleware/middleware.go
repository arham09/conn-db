package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/arham09/conn-db/helpers/caching"
	"github.com/labstack/echo"

	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	cache caching.Caching
	// another stuff , may be needed by middleware
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

// UserLimiter for handle request per user
func (m *GoMiddleware) UserLimiter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Request().Header.Get("x-user"))

		key := "request:user:" + c.Request().Header.Get("x-user")
		counter := 0

		err := m.cache.SetItem(context.Background(), key, counter, 2000)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"success": false,
				"message": err,
			})
		}

		return next(c)
	}
}

func (m *GoMiddleware) IPRateLimiter(timeCount time.Duration) echo.MiddlewareFunc {
	rate := limiter.Rate{
		Period: timeCount * time.Second,
		Limit:  1,
	}

	store := memory.NewStore()
	ipRateLimiter := limiter.New(store, rate)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			ip := c.RealIP()

			limiterCtx, err := ipRateLimiter.Get(c.Request().Context(), ip)

			if err != nil {
				log.Printf("IPRateLimit - ipRateLimiter.Get - err: %v, %s on %s", err, ip, c.Request().URL)
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"success": false,
					"message": err,
				})
			}

			h := c.Response().Header()
			h.Set("X-RateLimit-Limit", strconv.FormatInt(limiterCtx.Limit, 10))
			h.Set("X-RateLimit-Remaining", strconv.FormatInt(limiterCtx.Remaining, 10))
			h.Set("X-RateLimit-Reset", strconv.FormatInt(limiterCtx.Reset, 10))

			if limiterCtx.Reached {
				log.Printf("Too Many Requests from %s on %s", ip, c.Request().URL)
				return c.JSON(http.StatusTooManyRequests, echo.Map{
					"success": false,
					"message": "Too Many Requests on " + c.Request().URL.String(),
				})
			}

			return next(c)
		}
	}
}

// InitMiddleware intialize the middleware
func InitMiddleware(c caching.Caching) *GoMiddleware {
	return &GoMiddleware{
		cache: c,
	}
}
