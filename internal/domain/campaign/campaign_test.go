package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com", "email3@e.com"}

	// Act
	campaign := NewCampaign(name, content, contacts)

	//Assert
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNill(t *testing.T) {
	assert := assert.New(t)

	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com", "email3@e.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreateOnIsNotNill(t *testing.T) {
	assert := assert.New(t)

	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com", "email3@e.com"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}
