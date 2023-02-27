package client_wrapper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func Post[Req any, Res any](ctx context.Context, url string, request Req) (Res, error) {
	op := "Post"

	var response Res

	rawJSON, err := json.Marshal(request)
	if err != nil {
		return response, fmt.Errorf("%s: %w", op, err)
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(rawJSON))
	if err != nil {
		return response, fmt.Errorf("%s: %w", op, err)
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return response, fmt.Errorf("%s: %w", op, err)
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return response, fmt.Errorf("%s: %w", op, err)

	}

	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return response, fmt.Errorf("%s: %w", op, err)
	}
	return response, nil
}
