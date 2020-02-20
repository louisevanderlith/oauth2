package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/ory/fosite"
)

// Client represents a client or an app.
type Client struct {
	ID     string `hsk:"size(50)"`
	Secret string
	Domain string `hsk:"size(100)"`
	UserID string `hsk:"null"`
	RedirectURIs []string
	GrantTypes []string
	ResponseTypes []string
	Scopes []string
	Public bool
	Audience []string
}

func (c Client) Valid() (bool, error) {
	return husk.ValidateStruct(&c)
}


func GetAllClients() (husk.Collection, error) {
	return authStore.Clients.Find(1, 50, husk.Everything())
}

// GetID returns the client ID.
func (c Client) GetID() string {
	return c.ID
}

// GetHashedSecret returns the hashed secret as it is stored in the store.
func (c Client) GetHashedSecret() []byte {
	return []byte(c.Secret)
}

// GetRedirectURIs returns the client's allowed redirect URIs.
func (c Client) GetRedirectURIs() []string {
	return c.RedirectURIs
}

// GetGrantTypes returns the client's allowed grant types.
func (c Client) GetGrantTypes() fosite.Arguments {
	return c.GrantTypes
}

// GetResponseTypes returns the client's allowed response types.
func (c Client) GetResponseTypes() fosite.Arguments {
	return c.ResponseTypes
}

// GetScopes returns the scopes this client is allowed to request.
func (c Client) GetScopes() fosite.Arguments {
	return c.Scopes
}

// IsPublic returns true, if this client is marked as public.
func (c Client) IsPublic() bool {
	return c.Public
}

// GetAudience returns the allowed audience(s) for this client.
func (c Client) GetAudience() fosite.Arguments{
	return c.Audience
}
