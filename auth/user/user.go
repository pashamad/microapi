package user

import (
	jwt "bitbucket.org/onlifedevelop/bcc-jwt/src/service/proto"
	"context"
	"github.com/google/uuid"
	"github.com/micro/micro/v3/service/context/metadata"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/pashamad/microapi/auth/errors"
	"google.golang.org/grpc"
	"strings"
)

//noinspection GoUnusedExportedFunction
func GetAppUserTest() (id uuid.UUID, err error) {
	return uuid.Parse("9f11846b-fbef-45d1-8f09-03318aca878c")
}

func GetAppUser(ctx context.Context) (id uuid.UUID, err error) {
	// Get authorization header
	m, ok := metadata.FromContext(ctx)
	if !ok {
		log.Errorf("failed to get context metadata")
		return id, errors.CtxMetadataNotRetrieved
	}
	h, ok := m.Get("Authorization")
	if !ok {
		//log.Error("Failed to get authorization header from context metadata")
		return id, errors.AuthHeaderNotFound
	}
	log.Infof("Loaded authorization header: %s", h)

	// Get token from header
	split := strings.Split(h, " ")
	if len(split) != 2 {
		log.Errorf("Invalid authorization header: %s", h)
		return id, errors.AuthHeaderInvalidFormat
	}
	token := split[1]
	log.Infof("Extracted token from context: %s", token)

	// @todo load jwt service address from configuration or service context
	// Call decode method on bcc-jwt service
	jwtConn, err := grpc.Dial("bcc-jwt.default.svc.cluster.local:5304", []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		log.Errorf("Connection to bcc-jwt failed: %v", err)
		return id, err
	}
	log.Info("Opened connection to bcc-jwt service")

	jc := jwt.NewJWTClient(jwtConn)
	r := jwt.Token{Token: token}
	res, err := jc.DecodeJWT(ctx, &r)
	if err != nil {
		log.Errorf("Failed to decode token: %s", err)
		return id, err
	}
	//noinspection GoUnhandledErrorResult
	defer jwtConn.Close()

	s, l, v := res.Body.Sub, res.Body.Role, res.Valid
	if !v {
		log.Error("Token is not valid: ", token)
	}
	log.Debugf("Resolved authenticated user: %s : %s", s, l)

	u, err := uuid.Parse(s)
	if err != nil {
		log.Errorf("Failed to parse user id; %s", err)
		return id, err
	}

	// @todo check the user against database

	return u, nil
}
