package usecase

import "github.com/virb30/goexpert-cleanarch/internal/entity"

type ListOrdersOrderDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersOutputDTO struct {
	Orders []ListOrdersOrderDTO `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (u *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {
	orders, err := u.OrderRepository.GetAll()
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}
	var ordersOutput []ListOrdersOrderDTO
	for _, order := range orders {
		ordersOutput = append(ordersOutput, ListOrdersOrderDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return ListOrdersOutputDTO{
		Orders: ordersOutput,
	}, nil
}
