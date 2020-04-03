package core

type Client struct {
	Name     string `hsk:"size(50)"`
	Secret string
	Domain string `hsk:"size(100)"`
	Public bool //Public must sign in using Login-Page
	AllowedScopes []string
}

func (c Client) Valid() (bool, error) {
	return true, nil
}

// GetID client id
func (c Client) GetID() string {
	return c.Name
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
	return "No USER"
}
