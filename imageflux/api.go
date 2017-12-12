package imageflux

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

// APIParam is parameter of ImageFlux API
type APIParam struct {
	Token string
	URL   string
}

// APIResponse is response of ImageFlux API
type APIResponse struct {
	OK   bool     `json:"ok"`
	URLs []string `json:"urls"`
}

// RawAPIResponse is raw response of ImageFlux API
type RawAPIResponse struct {
	StatusCode   int
	ResponseBody []byte
}

const (
	// EpImageFluxAPI is API endpoint of ImageFlux API
	EpImageFluxAPI = "https://console.imageflux.jp"
	// DefaultAPITimeout is default timeout of ImageFlux API
	DefaultAPITimeout = 5
)

var (
	// URLRe is regexp of URL scheme
	URLRe *regexp.Regexp
	// APIClient is http client of ImageFlux API
	APIClient *http.Client
)

func init() {
	URLRe = regexp.MustCompile("^http(s)?://")
	APIClient = &http.Client{
		Timeout: DefaultAPITimeout * time.Second,
	}
}

func validate(param *APIParam) error {
	if param.URL == "" {
		return fmt.Errorf("cache: key is empty")
	}

	if !URLRe.MatchString(param.URL) {
		return fmt.Errorf("cache: key is not URL")
	}

	if _, err := url.Parse(param.URL); err != nil {
		return err
	}

	if param.Token == "" {
		return fmt.Errorf("cache: token is empty")
	}

	return nil
}

func issue(command string, param *APIParam) (*RawAPIResponse, error) {
	URL := fmt.Sprintf("%s/api/%s", EpImageFluxAPI, command)
	reqBody := url.Values{
		"token": {param.Token},
		"url":   {param.URL},
	}

	// url may include character to urlencode required.
	// e.g. /c!/w=360,f=webp/
	resp, err := APIClient.Post(URL,
		"application/x-www-form-urlencoded",
		strings.NewReader(reqBody.Encode()),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &RawAPIResponse{
		StatusCode:   resp.StatusCode,
		ResponseBody: body,
	}, nil

}

func buildAPIRespinse(rawResponse *RawAPIResponse) (*APIResponse, error) {
	resp := &APIResponse{}
	if err := json.Unmarshal(rawResponse.ResponseBody, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Run issues request to ImageFlux API endpoint.
func Run(command string) error {
	f := flag.NewFlagSet(os.Args[0]+" "+command, flag.ExitOnError)
	url := f.String("k", "", "cache key")
	verbose := f.Bool("v", false, "verbose output mode")
	f.Parse(os.Args[2:])

	token, err := loadToken(imagefluxConfPath)
	if err != nil {
		return err
	}

	param := &APIParam{
		Token: token,
		URL:   *url,
	}

	if err := validate(param); err != nil {
		return err
	}

	rawResp, err := issue(command, param)
	if err != nil {
		return err
	}

	if *verbose {
		resp, err := buildAPIRespinse(rawResp)
		if err != nil {
			return err
		}
		if !resp.OK {
			fmt.Fprintln(os.Stderr, "API call to ImageFlux was failed.")
		}
		fmt.Fprintf(os.Stderr, "HTTP status code: %d\n", rawResp.StatusCode)
	}

	fmt.Println(string(rawResp.ResponseBody))

	return nil
}
