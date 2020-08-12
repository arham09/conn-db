package middleware

import (
	"context"
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
		var counter int
		userID := c.Request().Header.Get("x-user-id")
		key := "request:user:" + userID
		count, ok := m.cache.GetItem(context.Background(), key)

		if ok {
			i, err := strconv.Atoi(count)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"success": false,
					"message": err,
				})
			}

			counter = i

			if counter == 5 {
				return c.JSON(http.StatusTooManyRequests, echo.Map{
					"success": false,
					"message": "Too Many Requests on " + c.Request().URL.String() + " you can try again in a minute",
				})
			}
		}

		err := m.cache.SetItem(context.Background(), key, counter+1, 1*time.Minute)

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
