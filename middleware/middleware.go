package middleware

import (
	"github.com/labstack/echo"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

// // LOG for logger
// func (m *GoMiddleware) LOG(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		makeLogEntry(c).Info("Request")
// 		return next(c)
// 	}
// }

// func (m *GoMiddleware) ErrorHandler(err error, c echo.Context) {
// 	report, ok := err.(*echo.HTTPError)
// 	if ok {
// 		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
// 	} else {
// 		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	makeLogEntry(c).Error(report.Message)
// 	c.HTML(report.Code, report.Message.(string))
// }

// InitMiddleware intialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}

// func makeLogEntry(c echo.Context) *logrus.Entry {
// 	if c == nil {
// 		return logrus.WithFields(logrus.Fields{
// 			"at": time.Now().Format("2006-01-02 15:04:05"),
// 		})
// 	}

// 	return logrus.WithFields(logrus.Fields{
// 		"at":     time.Now().Format("2006-01-02 15:04:05"),
// 		"method": c.Request().Method,
// 		"uri":    c.Request().URL.String(),
// 		"ip":     c.Request().RemoteAddr,
// 	})
// }
