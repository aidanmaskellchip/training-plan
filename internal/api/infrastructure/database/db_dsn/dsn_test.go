package dbdsn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDSN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		env      string
		host     string
		port     string
		user     string
		password string
		dbName   string
		expected string
	}{
		{
			name:     "Local setup",
			env:      "local",
			host:     "test",
			port:     "5432",
			user:     "user",
			password: "pwd",
			dbName:   "wealth-score",
			expected: "postgres://user:pwd@test:5432/wealth-score?sslmode=disable",
		},
		{
			name:     "Prod setup",
			env:      "production",
			host:     "prod",
			port:     "5432",
			user:     "user",
			password: "pwd",
			dbName:   "wealth-score",
			expected: "postgres://user:pwd@prod:5432/wealth-score",
		},
	}

	for _, tc := range tests {
		testCase := tc // rebind into this lexical scope to avoid concurrent access

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			res := DSNGenerator(testCase.env, testCase.host, testCase.port, testCase.user, testCase.password, testCase.dbName)

			assert.Equal(t, testCase.expected, res)
		})
	}
}
