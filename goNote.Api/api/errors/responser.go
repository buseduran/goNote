package responser

import "github.com/gofiber/fiber/v2"

// Utils struct (placeholder if additional utilities are needed)
type Utils struct{}

// Unknown Error
func UnknownError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "An unknown error occurred.",
	})
}

// Not Empty
func NotEmpty(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Field cannot be empty.",
	})
}

// Not Found
func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Resource not found.",
	})
}

// Already Exists
func AlreadyExists(c *fiber.Ctx) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"error": "Resource already exists.",
	})
}

// Invalid
func Invalid(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Invalid input.",
	})
}

// Invalid Body
func InvalidBody(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Invalid request body.",
	})
}

// Fetch Failed (Internal Server Error)
func FetchFailed(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Failed to fetch data. Please try again later.",
	})
}

// Fetch Success
func FetchSuccess(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": "Fetched datas successfully.",
		"data":  data,
	})
}

// Bad Request
func BadRequest(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Bad request. Please check your input.",
	})
}

// Invalid Date Format
func InvalidDatetime(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Invalid date or time format.",
	})
}

// Invalid ID
func InvalidID(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Invalid ID format.",
	})
}

// Create Failed
func CreateFailed(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Failed to create resource.",
	})
}

// Create Success
func CreateSuccess(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Resource created successfully.",
	})
}

// Update Failed
func UpdateFailed(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Failed to update resource.",
	})
}

// Update Success
func UpdateSuccess(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Resource updated successfully.",
	})
}

// Delete Failed
func DeleteFailed(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Failed to delete resource.",
	})
}

// Delete Success
func DeleteSuccess(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Resource deleted successfully.",
	})
}

// Get Failed
func GetFailed(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Failed to retrieve data.",
	})
}

// Get Success
func GetSuccess(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data retrieved successfully.",
	})
}

// Unauthorized
func Unauthorized(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized access. Please log in.",
	})
}

// Successful Login
func SuccessfulLogin(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful.",
	})
}

// Successful Logout
func SuccessfulLogout(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout successful.",
	})
}

// Successful Signup
func SuccessfulSignup(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Signup successful.",
	})
}
