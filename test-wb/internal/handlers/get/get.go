package get

import (
	"test-go/internal/storage"

	"github.com/labstack/echo"
	"github.com/patrickmn/go-cache"
)

type Getter interface {
	GetOrdes(int) (*cache.Cache, error)
	GetOrderById(string) (storage.Order, error)
}

func GetById(ctx echo.Context, getter Getter) (storage.Order, error) {
	idStr := ctx.Param("id")

	res, err := getter.GetOrderById(idStr)

	if err != nil {
		return storage.Order{}, err
	}

	return res, nil
}
