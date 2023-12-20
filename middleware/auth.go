package middleware

/*
	func Protected() func(ctx *fiber.Ctx) error {
		return func(ctx *fiber.Ctx) error {
			// Get token from header
			token := ctx.Get("Authorization")
			// Verify token
			if _, err := VerifyToken(token); err != nil {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Unauthorized",
				})
			}
			// If everything is ok, call next middleware
			return ctx.Next()
		}
	}
*/
