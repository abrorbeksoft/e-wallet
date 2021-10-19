package v1

import (
	"github.com/abrorbeksoft/e-wallet.git/storage"
)

type HandlerV1Options struct {
	Storage storage.StorageI
}

type handlerv1 struct {
	Storage storage.StorageI
}

func New(options *HandlerV1Options) *handlerv1 {
	return &handlerv1{
		Storage: options.Storage,
	}
}