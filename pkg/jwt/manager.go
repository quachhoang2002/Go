package jwt

type Manager interface {
	ValidateToken(token string) (string, error)
}

type implManager struct {
	secretKey string
}

func NewManager(secretKey string) Manager {
	return &implManager{
		secretKey: secretKey,
	}
}

func (m *implManager) ValidateToken(token string) (string, error) {
	return "", nil
}
