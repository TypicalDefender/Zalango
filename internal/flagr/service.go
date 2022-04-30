package flagr

import (
	"context"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/checkr/goflagr"
)

type Service interface {
	GetConfigByDescription(ctx context.Context, description string) (map[string]string, error)
	GetConfigByTag(ctx context.Context, tag string) map[string]string
}

type service struct {
	apiClient *goflagr.APIClient
}

func NewService(client *goflagr.APIClient) Service {
	return &service{apiClient: client}
}

func (s *service) GetConfigByDescription(ctx context.Context, serviceName string) (map[string]string, error) {

	opts := goflagr.FindFlagsOpts{
		Enabled:     optional.NewBool(true),
		Description: optional.NewString(serviceName),
		Preload:     optional.NewBool(true),
	}

	flags, _, err := s.apiClient.FlagApi.FindFlags(ctx, &opts)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch flags by service description from flagr %s", err.Error())
	}

	if len(flags) != 1 {
		fmt.Printf("%+v\n len: %v", flags, len(flags))
		return nil, errors.New(fmt.Sprintf("No config available for service : %s", serviceName))
	}

	confs := make(map[string]string)
	for _, flag := range flags[0].Variants {
		attach := (*flag.Attachment).(map[string]interface{})
		val, ok := attach["value"].(string)
		if !ok {
			return nil, errors.New("could not transform flagr response to key value")
		}
		confs[flag.Key] = val
	}

	return confs, nil
}
func (s *service) GetConfigByTag(ctx context.Context, tag string) map[string]string {
	return map[string]string{}
}
