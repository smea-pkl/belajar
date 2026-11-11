package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputUser(t *testing.T) {

	t.Run("testing gender", func(t *testing.T) {

		t.Run("test l kecil", func(t *testing.T) {
			usr := User{
				Name:     "Dono",
				Username: "dono",
				Gender:   "l",
			}
			check, err := filterUser(usr)
			assert.Nil(t, err)
			assert.True(t, check)
		})

		t.Run("test gender ngasal", func(t *testing.T) {
			usr := User{
				Name:     "Dono",
				Username: "dono",
				Gender:   "asdasdasd",
			}
			check, err := filterUser(usr)
			assert.Nil(t, err)
			assert.False(t, check)
		})

		t.Run("sutarman harus diblokir", func(t *testing.T) {
			usr := User{
				Name:     "Dono",
				Username: "sutarman",
				Gender:   "asdasdasd",
			}
			check, err := filterUser(usr)
			assert.Nil(t, err)
			assert.False(t, check)

			usr = User{
				Name:     "Dono",
				Username: "sutaRman",
				Gender:   "l",
			}
			check, err = filterUser(usr)
			assert.Nil(t, err)
			assert.False(t, check)
		})

	})

}
