package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body"
	contacts = []string{"email1@e.com", "email2@e.com", "email3@e.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	campaign := NewCampaign(name, content, contacts)

	//Assert
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNill(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	campaign := NewCampaign(name, content, contacts)

	//Assert
	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreateOnMustBeNow(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	// Act
	campaign := NewCampaign(name, content, contacts)

	//Assert
	assert.Greater(campaign.CreatedOn, now)
}
