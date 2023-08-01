package gfit

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const tokenFile = "token.json"

// CloseFile close file and log error
func CloseFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Println("unable to close file", f.Name(), err)
	}
}

// OpenOrCreateFile open file if it exists, create it if it doesn't.
// Will also create directories as necessary
func OpenOrCreateFile(fileName string) (*os.File, error) {
	f, err := os.OpenFile(fileName, os.O_RDWR, 0)
	if err == nil {
		return f, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	if err = os.MkdirAll(filepath.Dir(fileName), 0); err != nil {
		return nil, err
	}
	return os.Create(fileName)
}

// GetClient Retrieve a token, saves the token, then returns the generated client.
// The file token.json stores the user's access and refresh tokens, and is
// created automatically when the authorization flow completes for the first
// time.
func GetClient(ctx context.Context, config *oauth2.Config) (*http.Client, error) {
	tok, err := tokenFromFile(tokenFile)
	if err != nil {
		tok, err = getTokenFromWeb(ctx, config)
		if err != nil {
			return nil, fmt.Errorf("unable to start client: %w", err)
		}
		if err = saveToken(tokenFile, tok); err != nil {
			log.Println("saving token file failed:", err)
		}
	}
	return config.Client(ctx, tok), nil
}

func randomBytesInHex(count int) (string, error) {
	buf := make([]byte, count)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		return "", fmt.Errorf("Could not generate %d random bytes: %v", count, err)
	}

	return hex.EncodeToString(buf), nil
}

type AuthURL struct {
	URL          string
	State        string
	CodeVerifier string
}

func (u *AuthURL) String() string {
	return u.URL
}

func AuthorizationURL(config *oauth2.Config) (*AuthURL, error) {
	codeVerifier, verifierErr := randomBytesInHex(32) // 64 character string here
	if verifierErr != nil {
		return nil, fmt.Errorf("Could not create a code verifier: %v", verifierErr)
	}
	sha2 := sha256.New()
	io.WriteString(sha2, codeVerifier)
	codeChallenge := base64.RawURLEncoding.EncodeToString(sha2.Sum(nil))

	state, stateErr := randomBytesInHex(24)
	if stateErr != nil {
		return nil, fmt.Errorf("Could not generate random state: %v", stateErr)
	}

	authUrl := config.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
	)

	return &AuthURL{
		URL:          authUrl,
		State:        state,
		CodeVerifier: codeVerifier,
	}, nil
}

func stringScopes(scopes []AuthScope) []string {
	ss := make([]string, len(scopes))
	for i, s := range scopes {
		ss[i] = string(s)
	}
	return ss
}

// LoadConfig loads config from environment variables
func LoadConfig(scopes ...AuthScope) (*oauth2.Config, error) {
	var err error
	if err = godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  AuthorizationUrl,
			TokenURL: fmt.Sprintf("%s%s", Host, TokenUrl),
		},
		Scopes: stringScopes(scopes),
	}
	return &cfg, nil
}

// getTokenFromWeb Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(ctx context.Context, config *oauth2.Config) (*oauth2.Token, error) {
	//csrf := uuid.New().String()
	authURL, err := AuthorizationURL(config)
	if err != nil {
		return nil, fmt.Errorf("failed to read token: %w", err)
	}
	//authURL := config.AuthCodeURL(csrf, oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, fmt.Errorf("failed to read token: %w", err)
	}

	return config.Exchange(
		ctx,
		authCode,
		oauth2.SetAuthURLParam("code_verifier", authURL.CodeVerifier),
	)
}

// tokenFromFile Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer CloseFile(f)
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// saveToken Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) error {
	log.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("unable to cache oauth token, opening file failed: %w", err)
	}
	defer CloseFile(f)
	err = json.NewEncoder(f).Encode(token)
	if err != nil {
		return fmt.Errorf("unable to cache oauth token, json encoding failed: %w", err)
	}
	return nil
}

func makeRequest(client *http.Client, url string) ([]byte, error) {
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read all data from request
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return contents, nil
}
