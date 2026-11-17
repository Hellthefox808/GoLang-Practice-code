// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"errors"
	"testing"
)

func TestAuthorizeRequest(t *testing.T) {
	tests := []struct {
		name        string
		account     Account
		targetRole  string
		expectedOk  bool
		expectedErr error
	}{
		{
			name:        "Active admin allowed anywhere",
			account:     Account{Username: "ravi", Role: "admin", IsActive: true, LoginRate: 10},
			targetRole:  "editor",
			expectedOk:  true,
			expectedErr: nil,
		},
		{
			name:        "Suspended account denied",
			account:     Account{Username: "bad", Role: "admin", IsActive: false, LoginRate: 10},
			targetRole:  "admin",
			expectedOk:  false,
			expectedErr: ErrAccountInactive,
		},
		{
			name:        "Rate limit exceeded",
			account:     Account{Username: "flooder", Role: "admin", IsActive: true, LoginRate: 150},
			targetRole:  "admin",
			expectedOk:  false,
			expectedErr: nil, // custom error with fmt.Errorf
		},
		{
			name:        "Insufficient role",
			account:     Account{Username: "guestUser", Role: "guest", IsActive: true, LoginRate: 5},
			targetRole:  "admin",
			expectedOk:  false,
			expectedErr: ErrAccessDenied,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := AuthorizeRequest(tt.account, tt.targetRole)
			if tt.expectedErr != nil && !errors.Is(err, tt.expectedErr) {
				t.Fatalf("expected error %v, got %v", tt.expectedErr, err)
			}
			if ok != tt.expectedOk {
				t.Errorf("got ok=%t, expected %t", ok, tt.expectedOk)
			}
		})
	}
}
