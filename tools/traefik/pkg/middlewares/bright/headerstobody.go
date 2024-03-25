package bright

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/opentracing/opentracing-go/ext"
	"github.com/traefik/traefik/v2/pkg/config/dynamic"
	"github.com/traefik/traefik/v2/pkg/log"
	"github.com/traefik/traefik/v2/pkg/middlewares"
	"github.com/traefik/traefik/v2/pkg/tracing"

	"github.com/google/uuid"
)

const (
	basicTypeName = "HeadersToBody"
)

type headersToBody struct {
	next        http.Handler
	name        string
	headerNames []string
}

// NewBasic creates a headersToBody middleware.
func NewHeadersToBody(ctx context.Context, next http.Handler, config dynamic.HeadersToBody, name string) (http.Handler, error) {
	log.FromContext(middlewares.GetLoggerCtx(ctx, name, basicTypeName)).Debug("Creating middleware")

	ctb := &headersToBody{
		name:        name,
		next:        next,
		headerNames: config.HeaderNames,
	}

	return ctb, nil
}

func (ctb *headersToBody) GetTracingInformation() (string, ext.SpanKindEnum) {
	return ctb.name, tracing.SpanKindNoneEnum
}

func (ctb *headersToBody) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	logger := log.FromContext(middlewares.GetLoggerCtx(req.Context(), ctb.name, basicTypeName))
	logger.Warnf("Read body: %v %v", ctb.name, basicTypeName)

	myBody, err := io.ReadAll(req.Body)
	if err != nil {
		logger.Warnf("Read body failed: %v", err)
		tracing.SetErrorWithEvent(req, "Read body failed")
		rw.WriteHeader(http.StatusForbidden)
		return
	}

	bodyMap := map[string]interface{}{}
	if len(myBody) > 0 {
		err = json.Unmarshal(myBody, &bodyMap)
		if err != nil {
			logger.Warnf("Unmarshal body failed: %v", err)
			tracing.SetErrorWithEvent(req, "Unmarshal body failed")
			rw.WriteHeader(http.StatusForbidden)
			return
		}
	}

	var infoMap map[string]interface{}
	infoExist := false

	if info, ok := bodyMap["Info"]; ok && info != nil {
		infoMap, infoExist = info.(map[string]interface{})
	}

	for _, name := range ctb.headerNames {
		header := req.Header.Get(name)
		if header == "" {
			continue
		}

		switch name {
		case authHeaderUser:
			if _, err := uuid.Parse(header); err != nil {
				logger.Warnf("invalid header value")
				continue
			}
		}

		bodyName := ""
		switch name {
		case authHeaderUser:
			bodyName = "UserID"
		case authHeaderRole:
			bodyName = "Token"
		}

		if bodyName == "" {
			logger.Warnf("unexpected header to body")
			continue
		}

		bodyMap[bodyName] = header

		switch name {
		case authHeaderUser:
			if infoMap != nil && infoExist {
				infoMap[bodyName] = header
			}
		}
	}

	bodyMap["Info"] = infoMap

	myBody, err = json.Marshal(&bodyMap)
	if err != nil {
		logger.Warnf("Marshal body failed: %v", err)
		tracing.SetErrorWithEvent(req, "Marshal body failed")
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	req.Body = io.NopCloser(strings.NewReader(string(myBody)))
	req.ContentLength = int64(len(myBody))

	ctb.next.ServeHTTP(rw, req)
}
