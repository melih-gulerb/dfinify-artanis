package middlewares

import (
	"artanis/src/logging"
	"artanis/src/models/base"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"runtime/debug"
	"strings"
)

func PanicRecoveryMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				stackTrace := debug.Stack()

				errMsg := fmt.Sprintf("Panic recovered: %v", r)
				logging.Log(logging.ERROR, errMsg+"\n"+string(stackTrace))

				c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
				_ = c.Status(fiber.StatusInternalServerError).JSON(base.Response{
					Success: false,
					Message: "Internal Server Error: The server encountered an unexpected condition",
					Data:    nil,
				})
			}
		}()

		return c.Next()
	}
}

func CustomErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	switch e := err.(type) {
	case *fiber.Error:
		code = e.Code
		message = e.Message
	case *json.SyntaxError:
		code = fiber.StatusBadRequest
		message = "Invalid JSON format: " + e.Error()
	default:
		errMsg := err.Error()

		if strings.Contains(errMsg, "runtime error: invalid memory address") ||
			strings.Contains(errMsg, "nil pointer dereference") {
			message = "Server encountered a nil pointer error"
			logging.Log(logging.ERROR, "Nil pointer error: "+errMsg)
		} else if strings.Contains(errMsg, "sql: no rows") {
			code = fiber.StatusNotFound
			message = "Requested resource not found"
		} else if strings.Contains(errMsg, "context deadline exceeded") ||
			strings.Contains(errMsg, "context canceled") {
			code = fiber.StatusGatewayTimeout
			message = "Request timed out"
		} else {
			logging.Log(logging.ERROR, "Unhandled error: "+errMsg)
		}
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return c.Status(code).JSON(base.Response{
		Success: false,
		Message: message,
		Data:    nil,
	})
}
