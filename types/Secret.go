package types

type SecretJSONConfig struct {
	FirebaseConfig FirebaseConfig `json:"firebaseConfig"`
	Mailjet        MailjetConfig  `json:"mailjet"`
	Database       DatabaseConfig `json:"database"`
	OpenAI         OpenAIConfig   `json:"openai"`
	AESGCM         string         `json:"aesgcm"`
	Razorpay       RazorPayConfig `json:"razorpay"`
	JWTConfig      JWTConfig      `json:"jwtConfig"`
}

type FirebaseConfig struct {
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	AuthURI                 string `json:"auth_uri"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
	PrivateKey              string `json:"private_key"`
	PrivateKeyID            string `json:"private_key_id"`
	ProjectID               string `json:"project_id"`
	TokenURI                string `json:"token_uri"`
	Type                    string `json:"type"`
	UniverseDomain          string `json:"universe_domain"`
}

type MailjetConfig struct {
	PrivateApiKey string `json:"privateApiKey"`
	PublicApiKey  string `json:"publicApiKey"`
}

type RazorPayConfig struct {
	KeyId     string `json:"keyId"`
	KeySecret string `json:"keySecret"`
}

type DatabaseConfig struct {
	URL string `json:"url"`
}

type OpenAIConfig struct {
	Secret string `json:"secret"`
}

type JWTConfig struct {
	JWTSecret string `json:"jwtSecret"`
}
