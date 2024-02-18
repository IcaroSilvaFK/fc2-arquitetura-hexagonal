package cli

import (
	"fmt"

	"github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/application"
)

func Run(
	service application.ProductServiceInterface,
	action string,
	productId string,
	productName string,
	price float64,
) (string, error) {

	result := ""

	switch action {
	case "create":
		p, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}

		result := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f then status is %s", p.GetID(), p.GetName(), p.GetPrice(), p.GetStatus())

		return result, nil

	case "enable":
		p, err := service.Get(productId)

		if err != nil {
			return result, err
		}

		res, err := service.Enable(p)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been enabled", res.GetName())

		return result, nil

	case "disable":
		p, err := service.Get(productId)

		if err != nil {
			return result, err
		}

		res, err := service.Disable(p)

		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been disabled", res.GetName())

		return result, nil

	default:
		r, err := service.Get(productId)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf(
			"The Product Details \n Name: %s\n ID: %s\n Price: %f\n Status: %s",
			r.GetName(), r.GetID(), r.GetPrice(), r.GetStatus(),
		)

		return result, err
	}

}
