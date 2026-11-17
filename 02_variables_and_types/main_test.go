// Author: Ravi Ranjan Singh
// License: MIT

package main

import "testing"

func TestRoleString(t *testing.T) {
	tests := []struct {
		role     Role
		expected string
	}{
		{RoleGuest, "Guest"},
		{RoleUser, "User"},
		{RoleModerator, "Moderator"},
		{RoleAdministrator, "Administrator"},
		{Role(99), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.role.String(); got != tt.expected {
				t.Errorf("Role.String() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
