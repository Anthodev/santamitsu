package model

type SantaParticipant struct {
	UserId string
}

type ExcludedPair struct {
	UserId1 string
	UserId2 string
}

type ExcludedMember struct {
	UserId string
}

type ModeratorRole struct {
	RoleId string
}

type SantaSecret struct {
	ChannelID       string
	Title           string
	Description     string
	MaxPrice        string
	Currency        string
	State           int
	Participants    []SantaParticipant
	ExcludedPairs   []ExcludedPair
	ExcludedMembers []ExcludedMember
	ModeratorRoles  []ModeratorRole
}

func CreateSantaSecret(s SetupSettings) SantaSecret {
	return SantaSecret{
		ChannelID:       s.ChannelID,
		Title:           s.Title,
		Description:     s.Description,
		MaxPrice:        s.MaxPrice,
		Currency:        s.Currency,
		State:           s.State,
		Participants:    []SantaParticipant{},
		ExcludedPairs:   []ExcludedPair{},
		ExcludedMembers: []ExcludedMember{},
		ModeratorRoles:  []ModeratorRole{},
	}
}

func AddParticipant(s SantaSecret, p SantaParticipant) SantaSecret {
	s.Participants = append(s.Participants, p)

	return s
}

func RemoveParticipant(s SantaSecret, p SantaParticipant) SantaSecret {
	for i, v := range s.Participants {
		if v.UserId == p.UserId {
			s.Participants = append(s.Participants[:i], s.Participants[i+1:]...)
			break
		}
	}

	return s
}

func AddExcludedMember(s SantaSecret, p ExcludedMember) SantaSecret {
	s.ExcludedMembers = append(s.ExcludedMembers, p)

	return s
}

func AddExcludedPair(s SantaSecret, p ExcludedPair) SantaSecret {
	s.ExcludedPairs = append(s.ExcludedPairs, p)

	return s
}

func RemoveExcludedMember(s SantaSecret, p ExcludedMember) SantaSecret {
	for i, v := range s.ExcludedMembers {
		if v.UserId == p.UserId {
			s.ExcludedMembers = append(s.ExcludedMembers[:i], s.ExcludedMembers[i+1:]...)
			break
		}
	}

	return s
}

func RemoveExcludedPair(s SantaSecret, p ExcludedPair) SantaSecret {
	for i, v := range s.ExcludedPairs {
		if v.UserId1 == p.UserId1 && v.UserId2 == p.UserId2 {
			s.ExcludedPairs = append(s.ExcludedPairs[:i], s.ExcludedPairs[i+1:]...)
			break
		}
	}

	return s
}

func AddModeratorRole(s SantaSecret, r ModeratorRole) SantaSecret {
	s.ModeratorRoles = append(s.ModeratorRoles, r)

	return s
}

func RemoveModeratorRole(s SantaSecret, r ModeratorRole) SantaSecret {
	for i, v := range s.ModeratorRoles {
		if v.RoleId == r.RoleId {
			s.ModeratorRoles = append(s.ModeratorRoles[:i], s.ModeratorRoles[i+1:]...)
			break
		}
	}

	return s
}
