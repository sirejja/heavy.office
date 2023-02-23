package client_wrapper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

func Post[Req any, Res any](ctx context.Context, url string, request Req) (Res, error) {
	var response Res

	rawJSON, err := json.Marshal(request)
	if err != nil {
		return response, errors.WithMessage(err, "marshaling json")
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(rawJSON))
	if err != nil {
		return response, errors.WithMessage(err, "creating http request")
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return response, errors.WithMessage(err, "calling http")
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return response, fmt.Errorf("wrong status code: %d", httpResponse.StatusCode)
	}

	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return response, errors.WithMessage(err, "decoding json")
	}
	return response, nil
}
