package main

import "strings"

func filterUser(usr User) (bool, error) {

	blacklists := []string{
		"sutarman",
	}

	for _, item := range blacklists {
		if item == strings.ToLower(usr.Username) {
			return false, nil
		}
	}

	if usr.Username == "fitri" && usr.Gender == "p" {
		return false, nil
	}

	// check gender
	switch usr.Gender {
	case "l":
	case "p":
	default:
		return false, nil
	}

	return true, nil
}
