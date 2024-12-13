package client

import (
	"context"
	"errors"
	"math/rand"
	"pkg/validate"
	v1 "user/pkg/api/grpc/v1"

	"github.com/Markuysa/pkg/tracer"
)

func (s *service) SignUpRequest(
	ctx context.Context,
	request *v1.SignUpInitRequest,
) (signUpToken string, codeTTL int64, err error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	if !validate.Email(request.Email) {
		// errs
		return "", 0, errors.New("invalid email")
	}

	err = s.repos.SaveEmailConfirmationCode(ctx, request.Email, rand.Intn(90000)+10000)
	if err != nil {
		// errs
		return "", 0, err
	}

	return "", 0, nil
}
