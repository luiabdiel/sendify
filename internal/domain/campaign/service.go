package campaign

import "sendify/internal/contract"

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) error {
	return nil
}
