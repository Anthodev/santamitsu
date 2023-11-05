package service

import "anthodev/santamitsu/model"

func CheckUserIsExcluded(s model.SantaSecret, uid string) bool {
	for _, e := range s.ExcludedMembers {
		if e.UserId == uid {
			return true
		}
	}

	return false
}
