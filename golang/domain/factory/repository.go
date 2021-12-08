package factory

import "github.com/ViniciusContini/imersao5-gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
