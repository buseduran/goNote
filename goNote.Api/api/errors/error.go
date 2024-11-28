package errorw

import "github.com/gofiber/fiber/v2"

type Utils struct{}

// Error
func UnknownError(c *fiber.Ctx) error {
	return c.Status(400).JSON(fiber.Map{"error": "Unknown error"})
}

// Not empty
func NotEmpty(c *fiber.Ctx) error {
	return c.Status(400).JSON(fiber.Map{"error": "Not empty"})
}

// Not found
func NotFound(c *fiber.Ctx) error {
	return c.Status(400).JSON(fiber.Map{"error": "Not empty"})
}

// Already exists
func AlreadyExists(c *fiber.Ctx) error {
	return c.Status(400).JSON(fiber.Map{"error": "Already exists"})
}

// Invalid
func Invalid(c *fiber.Ctx) error {
	return c.Status(400).JSON(fiber.Map{"error": "Invalid"})
}

// Bad request
func BadRequest(c *fiber.Ctx) error {
	return c.Status(400).JSON(fiber.Map{"error": "Bad request"})
}

//create failed
func CreateFailed(c *fiber.Ctx) error {
	return c.Status(400).JSON(fiber.Map{"error": "Create failed"})
}

//create success
func CreateSuccess(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"error": "Create success"})
}

// Update failed
func UpdateFailed(c *fiber.Ctx) error {
	return c.Status(400).JSON(fiber.Map{"error": "Update failed"})
}

// Update success
func UpdateSuccess(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"error": "Update success"})
}

// Delete failed
func DeleteFailed(c *fiber.Ctx) error {
	return c.Status(400).JSON(fiber.Map{"error": "Delete failed"})
}

// Delete success
func DeleteSuccess(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"error": "Delete success"})
}

// Get failed
func GetFailed(c *fiber.Ctx) error {
	return c.Status(400).JSON(fiber.Map{"error": "Get failed"})
}

// Get success
func GetSuccess(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"error": "Get success"})
}

// Unauthorized
func Unauthorized(c *fiber.Ctx) error {
	return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
}

// Successful login
func SuccessfulLogin(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"error": "Successful login"})
}

// Successful logout
func SuccessfulLogout(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"error": "Successful logout"})
}

// Successful signup
func SuccessfulSignup(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"error": "Successful Signup"})
}
