package rpcservice

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"github.com/sirupsen/logrus"
	"log"
)

func serverReceiveReqLogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		err := fn(ctx, req, rsp)
		if err != nil {
			log.Println(logrus.Fields{
				"service":     req.Service(),
				"endpoint":    req.Endpoint(),
				"contentType": req.ContentType(),
				"req_body":    req.Body(),
				"err":         err,
			}, "service receive request, process err")
		}

		return err
	}
}

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	logrus.WithFields(logrus.Fields{
		"service":     req.Service(),
		"endpoint":    req.Endpoint(),
		"contentType": req.ContentType(),
	})

	err := l.Client.Call(ctx, req, rsp)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     req.Service(),
			"endpoint":    req.Endpoint(),
			"contentType": req.ContentType(),
			"req_body":    req.Body(),
			"err":         err,
		})
	}
	return err
}

// Implements client.Wrapper as logWrapper
func clientRequestLogWrap(c client.Client) client.Client {
	return &logWrapper{c}
}
