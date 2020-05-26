package controller

import "github.com/sendgrid/rest"

type RestClient interface {
	Send(request rest.Request) (*rest.Response, error)
}

func NewRestClient() RestClient{
	return &RestClientImpl{}
}

type RestClientImpl struct {
}

func (*RestClientImpl) Send(request rest.Request) (*rest.Response, error) {
	return rest.Send(request)
}