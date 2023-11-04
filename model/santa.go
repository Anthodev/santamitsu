package model

type SantaParticipant struct {
	UserId string
}

type ExcludedPair struct {
	UserId1 string
	UserId2 string
}

type SantaSecret struct {
	ChannelID     string
	Title         string
	Description   string
	MaxPrice      string
	Currency      string
	State         int
	Participants  []SantaParticipant
	ExcludedPairs []ExcludedPair
}

func CreateSantaSecret(s SetupSettings) SantaSecret {
	return SantaSecret{
		ChannelID:     s.ChannelID,
		Title:         s.Title,
		Description:   s.Description,
		MaxPrice:      s.MaxPrice,
		Currency:      s.Currency,
		State:         s.State,
		Participants:  []SantaParticipant{},
		ExcludedPairs: []ExcludedPair{},
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
