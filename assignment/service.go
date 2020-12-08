package main

import (
	"context"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"protoimport/assignment"
)

// service implements the ApikeyServiceServer interface
type service struct {
	apikeyRepository ApiKeyRepository
	crypt            Crypt
}

const IdLength = 36
const UserIdLength = 36
const ExchangeLength = 36
const ApiKeyLength = 4096
const SecretLength = 4096

func (s service) AddApikey(ctx context.Context, request *assignment.AddApikeyRequest) (*assignment.AddApikeyResponse, error) {
	//TODO: input validation could be improved by some library
	if len(request.UserId) > UserIdLength {
		return nil, status.Error(codes.InvalidArgument, "UserId")
	}
	if len(request.Exchange) > ExchangeLength {
		return nil, status.Error(codes.InvalidArgument, "Exchange")
	}
	if len(request.Apikey) > ApiKeyLength {
		return nil, status.Error(codes.InvalidArgument, "ApiKey")
	}
	if len(request.Secret) > SecretLength {
		return nil, status.Error(codes.InvalidArgument, "Secret")
	}

	encryptedApiKey, err := s.crypt.encrypt(request.Apikey)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	encryptedSecret, err := s.crypt.encrypt(request.Secret)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	preview := request.Apikey[:Min(ApiKeyPreviewChars, len(request.Apikey))]
	apiKey := NewApiKey(request.UserId, request.Exchange, preview, encryptedApiKey, encryptedSecret)
	err = s.apikeyRepository.Save(ctx, &apiKey)

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &assignment.AddApikeyResponse{ApikeyId: apiKey.Id.String()}, nil
}

func (s service) ListApikeys(ctx context.Context, request *assignment.ListApikeysRequest) (*assignment.ListApikeysResponse, error) {
	keys, err := s.apikeyRepository.FindByUser(ctx, request.UserId)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	apiKeys := make([]*assignment.Apikey, len(keys))
	for i, key := range keys {
		apiKeys[i] = &assignment.Apikey{
			ApikeyId:      key.Id.String(),
			Exchange:      key.Exchange,
			ApikeyPreview: key.ApiKeyPreview}
	}
	return &assignment.ListApikeysResponse{Apikeys: apiKeys}, nil
}

func (s service) GetApikey(ctx context.Context, request *assignment.GetApikeyRequest) (*assignment.GetApikeyResponse, error) {
	if len(request.ApikeyId) > IdLength {
		return nil, status.Error(codes.InvalidArgument, "ApikeyId")
	}

	id, err := uuid.FromString(request.ApikeyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "ApikeyId has invalid format: "+err.Error())
	}

	key, err := s.apikeyRepository.Get(ctx, id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "ApikeyId not found: "+err.Error())
	}

	apiKey, err := s.crypt.decrypt(key.ApiKey)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	secret, err := s.crypt.decrypt(key.Secret)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &assignment.GetApikeyResponse{
		Apikey: &assignment.ApikeyExtended{
			ApikeyDetails: &assignment.Apikey{
				ApikeyId:      key.Id.String(),
				Exchange:      key.Exchange,
				ApikeyPreview: key.ApiKeyPreview},
			Apikey: apiKey,
			Secret: secret,
		}}, nil
}
