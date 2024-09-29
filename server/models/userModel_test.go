package models_test

import (
	"task-inator3000/models"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		firstName string
		lastName  string
		password  string
		wantErr   bool
	}{
		{"Normal", "tester@testing.com", "Tester", "Doe", "Password", false},
		{"InvalidEmail", "NotEvenAnEmail", "Tester", "Doe", "Password", true},
		{"BlankEmail", "", "Tester", "Doe", "Password", true},
		{"BlankFirstName", "tester@testing.com", "", "Doe", "Password", true},
		{"BlankLastName", "tester@testing.com", "Tester", "", "Password", false},
		{"BlankPassword", "tester@testing.com", "Tester", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var user = models.User{
				Email:     tt.email,
				FirstName: tt.firstName,
				LastName:  tt.lastName,
				Password:  tt.password,
			}

			err := user.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Validate() returned error: %v", err)
			}
		})
	}
}
