package services

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"technical-test-go/models/domain"
	"technical-test-go/models/web"
	"technical-test-go/repositories"
	"time"
)

type CartServiceImpl struct {
	CartRepository  repositories.CartRepository
	OrderRepository repositories.OrderRepository
	ProductService  ProductService
	UserService     UserService
	BankService     BankService
	DB              *sqlx.DB
}

func NewCartService(db *sqlx.DB, cartRepository repositories.CartRepository, orderRepository repositories.OrderRepository, productService ProductService, userService UserService, bankService BankService) CartService {
	return &CartServiceImpl{
		CartRepository:  cartRepository,
		OrderRepository: orderRepository,
		ProductService:  productService,
		UserService:     userService,
		BankService:     bankService,
		DB:              db,
	}
}

func (service *CartServiceImpl) Create(ctx context.Context, request web.CartCreateRequest) (web.CartResponse, error) {
	cart := domain.Cart{
		UserId:    request.UserId,
		ProductId: request.ProductId,
		Quantity:  request.Quantity,
	}

	_, err := service.CartRepository.Save(ctx, service.DB, cart)
	if err != nil {
		return web.CartResponse{}, fmt.Errorf("failed to create cart: %w", err)
	}

	return service.FindByUserId(ctx, request.UserId)
}

func (service *CartServiceImpl) Update(ctx context.Context, request web.CartUpdateRequest) (web.CartResponse, error) {

	cart := domain.Cart{
		Id:       request.Id,
		Quantity: request.Quantity,
	}

	userCart, err := service.CartRepository.Update(ctx, service.DB, cart)
	if err != nil {
		return web.CartResponse{}, fmt.Errorf("failed to update cart: %w", err)
	}

	return service.FindById(ctx, userCart.Id)
}

func (service *CartServiceImpl) Delete(ctx context.Context, cartId int) error {
	_, err := service.CartRepository.FindById(ctx, service.DB, cartId)
	if err != nil {
		return fmt.Errorf("failed to retrieve cart:: %w", err)
	}

	err = service.CartRepository.Delete(ctx, service.DB, cartId)
	if err != nil {
		return fmt.Errorf("failed to delete cart item: %w", err)
	}

	return nil
}

func (service *CartServiceImpl) FindAll(ctx context.Context) ([]web.CartResponse, error) {
	carts, err := service.CartRepository.FindAll(ctx, service.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve carts: %w", err)
	}

	return service.mapToCartResponses(ctx, carts)
}

func (service *CartServiceImpl) FindByUserId(ctx context.Context, userId int) (web.CartResponse, error) {
	carts, err := service.CartRepository.FindByUserId(ctx, service.DB, userId)
	if err != nil {
		return web.CartResponse{}, fmt.Errorf("failed to retrieve cart for user Id %d: %w", userId, err)
	}

	cartItems := []web.CartItem{}
	for _, cart := range carts {
		product, err := service.ProductService.FindById(ctx, cart.ProductId)
		if err != nil {
			return web.CartResponse{}, fmt.Errorf("failed to retrieve product details: %w", err)
		}

		cartItems = append(cartItems, web.CartItem{
			ProductResponse: product,
			Quantity:        cart.Quantity,
		})
	}

	user, err := service.UserService.FindById(ctx, userId)
	if err != nil {
		return web.CartResponse{}, fmt.Errorf("failed to retrieve user details: %w", err)
	}

	return web.CartResponse{
		User:      user,
		CartItems: cartItems,
	}, nil
}

func (service *CartServiceImpl) FindById(ctx context.Context, cartId int) (web.CartResponse, error) {
	cart, err := service.CartRepository.FindById(ctx, service.DB, cartId)
	if err != nil {
		return web.CartResponse{}, fmt.Errorf("failed to retrieve cart item: %w", err)
	}

	product, err := service.ProductService.FindById(ctx, cart.ProductId)
	if err != nil {
		return web.CartResponse{}, fmt.Errorf("failed to retrieve product details: %w", err)
	}

	user, err := service.UserService.FindById(ctx, cart.UserId)
	if err != nil {
		return web.CartResponse{}, fmt.Errorf("failed to retrieve user details: %w", err)
	}

	return web.CartResponse{
		User: user,
		CartItems: []web.CartItem{
			{
				ProductResponse: product,
				Quantity:        cart.Quantity,
			},
		},
	}, nil
}

func (service *CartServiceImpl) Checkout(ctx context.Context, userId int) (web.CartResponse, error) {
	cartItems, err := service.CartRepository.FindByUserId(ctx, service.DB, userId)
	if err != nil {
		return web.CartResponse{}, fmt.Errorf("failed to retrieve cart items: %w", err)
	}

	if len(cartItems) == 0 {
		return web.CartResponse{}, fmt.Errorf("cart is empty")
	}

	// Step 2: Create the order
	tx, err := service.DB.Beginx()
	if err != nil {
		return web.CartResponse{}, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	order, err := service.OrderRepository.Save(ctx, tx, domain.Order{
		UserID:    userId,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})
	if err != nil {
		return web.CartResponse{}, fmt.Errorf("failed to save order: %w", err)
	}

	for _, cartItem := range cartItems {
		detail := domain.OrderDetail{
			OrderID:   order.Id,
			ProductID: cartItem.ProductId,
			Quantity:  cartItem.Quantity,
		}

		product, err := service.ProductService.FindById(ctx, cartItem.ProductId)
		if err != nil {
			return web.CartResponse{}, fmt.Errorf("failed to retrieve product: %w", err)
		}

		detail.Price = int64(product.Price * cartItem.Quantity)

		_, err = service.OrderRepository.SaveDetail(ctx, tx, detail)
		if err != nil {
			return web.CartResponse{}, fmt.Errorf("failed to create order details: %w", err)
		}

		transfer := web.BankTransferRequest{
			FromAccountId: userId,
			ToAccountId:   product.User.Id,
			Amount:        detail.Price,
			CreatedAt:     time.Time{},
		}

		err = service.BankService.Transfer(ctx, transfer)
		if err != nil {
			return web.CartResponse{}, fmt.Errorf("failed to transfer money to seller: %w", err)
		}

		err = service.CartRepository.Delete(ctx, service.DB, cartItem.Id)
	}

	err = tx.Commit()
	if err != nil {
		return web.CartResponse{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	response, err := service.FindByUserId(ctx, userId)
	if err != nil {
		return web.CartResponse{}, fmt.Errorf("failed to find cart: %w", err)
	}

	return response, nil
}

func (service *CartServiceImpl) mapToCartResponses(ctx context.Context, carts []domain.Cart) ([]web.CartResponse, error) {
	cartResponses := []web.CartResponse{}
	userCartMap := make(map[int][]web.CartItem)

	for _, cart := range carts {
		product, err := service.ProductService.FindById(ctx, cart.ProductId)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve product details: %w", err)
		}

		cartItem := web.CartItem{
			ProductResponse: product,
			Quantity:        cart.Quantity,
		}
		userCartMap[cart.UserId] = append(userCartMap[cart.UserId], cartItem)
	}

	for userId, cartItems := range userCartMap {
		user, err := service.UserService.FindById(ctx, userId)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve user details: %w", err)
		}

		cartResponses = append(cartResponses, web.CartResponse{
			User: web.UserResponse{
				Id:    userId,
				Name:  user.Name,
				Email: user.Email,
			},
			CartItems: cartItems,
		})
	}

	return cartResponses, nil
}
