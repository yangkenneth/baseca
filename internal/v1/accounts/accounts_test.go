package accounts

import (
	mock_store "github.com/coinbase/baseca/db/mock"
	db "github.com/coinbase/baseca/db/sqlc"
	"github.com/coinbase/baseca/internal/config"
	"github.com/coinbase/baseca/internal/lib/util/validator"
)

func buildAccountsConfig(store *mock_store.MockStore) (*Service, error) {
	config, err := config.GetTestConfigurationPath()
	if err != nil {
		return nil, err
	}

	endpoints := db.DatabaseEndpoints{Writer: store, Reader: store}
	validator.SupportedConfig(config)
	validator.SupportedEnvironments(config)

	return &Service{
		store:       endpoints,
		acmConfig:   config.ACMPCA,
		environment: config.Environment,
	}, nil
}
