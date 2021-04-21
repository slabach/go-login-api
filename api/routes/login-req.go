package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"schnelllegal.com/api/models"
)

// define incoming request object
type loginRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	TimeStamp int    `json:"timestamp"`
}

// handler for login request endpoint
func LoginReq(user *models.UserList) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get all current users
		current_users := user.GetUsers()

		// get time for token comparison
		current_time := time.Now()
		valid_token := fmt.Sprintf("%d%02d", current_time.Hour(), current_time.Minute())
		token_int, tk_err := strconv.Atoi(valid_token)
		if tk_err != nil {
			handle_invalid_error("An error has occurred", c)
		}

		// bind incoming request to loginRequest object
		requestBody := loginRequest{}
		bind_err := c.Bind(&requestBody)

		// make sure object bind was successful
		if bind_err != nil {
			handle_invalid_error("An error has occurred", c)
		} else {
			// make sure user actually exists
			if len(current_users) == 0 {
				handle_invalid_error("No users exist. Please add user first.", c)
			} else {
				// validate token
				if requestBody.TimeStamp != token_int {
					handle_invalid_error(fmt.Sprintf("Invalid token. Please try again later. %v", token_int), c)
				} else {
					// iterate through user list. look for user in user list that matches the incoming user
					for _, v := range current_users {
						// compare stored hashed password to incoming password
						pw_match := compare_pwd(v.Password, requestBody.Password)
						if v.Username == requestBody.Username && pw_match {
							// if match is found, respond with success message
							c.JSON(http.StatusOK, map[string]bool{
								"success": true,
							})
						} else {
							handle_invalid_error("User not found.", c)
						}
					}
				}
			}
		}
	}
}

// function for handling all error responses
// would probably move this to a common library
// would also handle unique http status responses instead of just returning a 400
func handle_invalid_error(message string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, map[string]string{
		"message": message,
	})
}

// function for comparing stored hashed password to password from incoming request
func compare_pwd(hashed []byte, sent string) bool {
	err := bcrypt.CompareHashAndPassword(hashed, []byte(sent))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
