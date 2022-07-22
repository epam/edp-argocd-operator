package argoclient

import (
	"errors"
	"net/url"
	"os"

	openapiruntime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/epam/edp-argocd-operator/pkg/argoclient/client/application_service"
	"github.com/epam/edp-argocd-operator/pkg/argoclient/client/repository_service"
)

type ApplicationHTTPClient struct {
	Client      application_service.ClientService
	AuthOptions application_service.ClientOption
}

type RepositoryHTTPClient struct {
	Client      repository_service.ClientService
	AuthOptions repository_service.ClientOption
}

type argoAccessData struct {
	Host   string
	Schema string
	Token  string
}

func GetApplicationClient() (*ApplicationHTTPClient, error) {
	accessData, err := getArgoConfig()
	if err != nil {
		return nil, err
	}

	argoApplicationClient := application_service.New(getTransport(accessData), strfmt.Default)

	opts := application_service.ClientOption(func(op *openapiruntime.ClientOperation) {
		op.AuthInfo = httptransport.BearerToken(accessData.Token)
	})

	return &ApplicationHTTPClient{
		Client:      argoApplicationClient,
		AuthOptions: opts,
	}, nil
}

func GetRepositoryClient() (*RepositoryHTTPClient, error) {
	accessData, err := getArgoConfig()
	if err != nil {
		return nil, err
	}

	argoRepositoryClient := repository_service.New(getTransport(accessData), strfmt.Default)

	opts := repository_service.ClientOption(func(op *openapiruntime.ClientOperation) {
		op.AuthInfo = httptransport.BearerToken(accessData.Token)
	})

	return &RepositoryHTTPClient{
		Client:      argoRepositoryClient,
		AuthOptions: opts,
	}, nil
}

func getArgoConfig() (*argoAccessData, error) {
	argoUrl, ok := os.LookupEnv("ARGOCD_URL")
	if !ok {
		return nil, errors.New("ARGOCD_URL env should be provided")
	}

	token, ok := os.LookupEnv("ARGOCD_TOKEN")
	if !ok {
		return nil, errors.New("ARGOCD_TOKEN env should be provided")
	}

	u, err := url.Parse(argoUrl)
	if err != nil {
		return nil, err
	}

	return &argoAccessData{
		Host:   u.Host,
		Schema: u.Scheme,
		Token:  token,
	}, nil
}

func getTransport(accessData *argoAccessData) *httptransport.Runtime {
	transport := httptransport.New(accessData.Host, "", []string{accessData.Schema})

	return transport
}
