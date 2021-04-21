package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"schnelllegal.com/api/models"
)

// define incoming request object
type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// handler for registration endpoint
func RegisterReq(user *models.UserList) gin.HandlerFunc {
	return func(c *gin.Context) {
		// bind incoming request to loginRequest object
		requestBody := registerRequest{}
		err := c.Bind(&requestBody)

		// make sure object bind was successful
		if err != nil {
			// this is where moving the error response to a common library would be smart, so this doesnt need repeated
			c.JSON(http.StatusBadRequest, map[string]string{
				"message": c.Errors.String(),
			})
		} else {
			// hash incoming password for storage
			hash, hash_err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.MinCost)
			// make sure hash worked successfully
			if hash_err != nil {
				c.JSON(http.StatusTeapot, map[string]string{
					"success": "false",
				})
			}
			// create User instance using incoming username + hashed password
			new_user := models.User{
				Username: requestBody.Username,
				Password: hash,
			}

			// add user instance to user list
			user.Add(new_user)

			// return OK status to requestor
			c.Status(http.StatusAccepted)
		}
	}
}
