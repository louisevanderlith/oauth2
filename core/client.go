package core

type Client struct {
	ID     string `hsk:"size(50)"`
	Secret string
	Domain string `hsk:"size(100)"`
	UserID string `hsk:"null"`
}

func (c Client) Valid() (bool, error) {
	return true, nil
}

// GetID client id
func (c Client) GetID() string {
	return c.ID
}

// GetSecret client domain
func (c Client) GetSecret() string {
	return c.Secret
}

// GetDomain client domain
func (c Client) GetDomain() string {
	return c.Domain
}

// GetUserID user id
func (c Client) GetUserID() string {
	return c.UserID
}
