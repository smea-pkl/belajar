package main

func filterUser(usr User) (bool, error) {

	blacklists := []string{
		"sutarman",
	}

	for _, item := range blacklists {
		if item == usr.Username {
			return false, nil
		}
	}

	if usr.Username == "fitri" && usr.Gender == "p" {
		return false, nil
	}

	return true, nil
}
