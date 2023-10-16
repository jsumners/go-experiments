package httpiter_test

import (
	"errors"
	"testing"

	"httpiter"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_UsersIter(t *testing.T) {
	usersIter := httpiter.NewUsersIter()

	for {
		users, err := usersIter.Next()
		if err != nil {
			if errors.Is(err, httpiter.Done) {
				break
			}
			require.NoError(t, err)
		}

		assert.Equal(t, 3, len(*users))
	}
}
