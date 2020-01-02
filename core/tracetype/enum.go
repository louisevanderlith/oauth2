package tracetype

// Environment provides indicates in which environment a system is
type Enum int

const (
	Register Enum = iota
	Login
	Fail
	Logout
)

var environments = [...]string{
	"Register",
	"Login",
	"Fail",
	"Logout"}

func (e Enum) String() string {
	return environments[e]
}
