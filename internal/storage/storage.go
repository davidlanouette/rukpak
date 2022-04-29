package storage

import (
	"context"
	"io/fs"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Storage interface {
	Load(ctx context.Context, owner client.Object) (fs.FS, error)
	Store(ctx context.Context, owner client.Object, bundle fs.FS) error
	Delete(ctx context.Context, owner client.Object) error
}
