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

type MatchedPair struct {
	UserId1 string
	UserId2 string
}

type SecretSanta struct {
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
	MatchedPairs    []MatchedPair
}

func CreateSecretSanta(s SetupSettings) SecretSanta {
	return SecretSanta{
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
		MatchedPairs:    []MatchedPair{},
	}
}

func AddParticipant(s SecretSanta, p SantaParticipant) SecretSanta {
	s.Participants = append(s.Participants, p)

	return s
}

func RemoveParticipant(s SecretSanta, p SantaParticipant) SecretSanta {
	for i, v := range s.Participants {
		if v.UserId == p.UserId {
			s.Participants = append(s.Participants[:i], s.Participants[i+1:]...)
			break
		}
	}

	return s
}

func AddExcludedMember(s SecretSanta, p ExcludedMember) SecretSanta {
	s.ExcludedMembers = append(s.ExcludedMembers, p)

	return s
}

func AddExcludedPair(s SecretSanta, p ExcludedPair) SecretSanta {
	s.ExcludedPairs = append(s.ExcludedPairs, p)

	return s
}

func RemoveExcludedMember(s SecretSanta, p ExcludedMember) SecretSanta {
	for i, v := range s.ExcludedMembers {
		if v.UserId == p.UserId {
			s.ExcludedMembers = append(s.ExcludedMembers[:i], s.ExcludedMembers[i+1:]...)
			break
		}
	}

	return s
}

func RemoveExcludedPair(s SecretSanta, p ExcludedPair) SecretSanta {
	for i, v := range s.ExcludedPairs {
		if v.UserId1 == p.UserId1 && v.UserId2 == p.UserId2 {
			s.ExcludedPairs = append(s.ExcludedPairs[:i], s.ExcludedPairs[i+1:]...)
			break
		}
	}

	return s
}

func AddModeratorRole(s SecretSanta, r ModeratorRole) SecretSanta {
	s.ModeratorRoles = append(s.ModeratorRoles, r)

	return s
}

func RemoveModeratorRole(s SecretSanta, r ModeratorRole) SecretSanta {
	for i, v := range s.ModeratorRoles {
		if v.RoleId == r.RoleId {
			s.ModeratorRoles = append(s.ModeratorRoles[:i], s.ModeratorRoles[i+1:]...)
			break
		}
	}

	return s
}

func AddMatchedPair(s SecretSanta, p MatchedPair) SecretSanta {
	s.MatchedPairs = append(s.MatchedPairs, p)

	return s
}
