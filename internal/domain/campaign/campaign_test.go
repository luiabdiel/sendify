package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com", "email3@e.com"}

	// Act
	campaign := NewCampaign(name, content, contacts)

	//Assert
	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}
