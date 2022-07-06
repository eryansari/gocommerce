package helpers

import (
	"context"
	"fmt"
)

// GetUsername from Context
func GetUsername(c context.Context) string {
	username := fmt.Sprintf("%v", c.Value("username"))
	if username == "" || username == "<nil>" {
		username = ""
	}
	return username
}

// GetUserphone from Context
func GetUserphone(c context.Context) string {
	phone := fmt.Sprintf("%v", c.Value("phone"))
	if phone == "" || phone == "<nil>" {
		phone = ""
	}
	return phone
}

// GetUserEmail from Context
func GetUserEmail(c context.Context) string {
	email := fmt.Sprintf("%v", c.Value("email"))
	if email == "" || email == "<nil>" {
		email = ""
	}
	return email
}

// GetUserID from Context
func GetUserID(c context.Context) string {
	userID := fmt.Sprintf("%v", c.Value("user_id"))
	if userID == "" || userID == "<nil>" {
		userID = ""
	}
	return userID
}

// GetDeviceID from Context
func GetDeviceID(c context.Context) string {
	deviceID := fmt.Sprintf("%v", c.Value("device_id"))
	if deviceID == "" || deviceID == "<nil>" {
		deviceID = ""
	}
	return deviceID
}
