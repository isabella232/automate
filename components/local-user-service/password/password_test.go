package password_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chef/automate/components/local-user-service/password"
)

func TestValidate(t *testing.T) {
	validator, err := password.NewValidator()
	if err != nil {
		t.Fatal(err)
	}

	// These are the cases return simple error instances
	fails := map[string][]string{
		"password is empty or all whitespace":                              {"", "        "},
		"password does not contain enough distinct characters (minimum 3)": {"aaaaaaaaaa", "ababababababa"},
		"password is too short (must be at least 8 characters)":            {"foo", "short"},
	}
	passes := []string{
		"password",
		"applepie",
		"p@ssw0rd!",
		"axxleyie",
		"batteryhorsestaple",
		"ababababababac", // it's three different characters
		"telephone",
		"d4f108896d22656e692c39478d95a8dc", // Randomly generated bytes encoded as hex
	}

	t.Run("passing validation", func(t *testing.T) {
		for _, pass := range passes {
			t.Run(pass, func(t *testing.T) {
				assert.NoError(t, validator.Validate(pass))
			})
		}
	})

	t.Run("failing validation", func(t *testing.T) {
		for expErr, cases := range fails {
			for _, fail := range cases {
				t.Run(fail, func(t *testing.T) {
					err := validator.Validate(fail)
					require.Error(t, err)
					assert.Equal(t, expErr, err.Error())
				})
			}
		}
	})
}
