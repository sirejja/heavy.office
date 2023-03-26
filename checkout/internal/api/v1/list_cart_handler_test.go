package v1

import (
	"context"
	"errors"
	"fmt"
	"route256/checkout/internal/clients/grpc/products"
	productsClientMocks "route256/checkout/internal/clients/grpc/products/mocks"
	"route256/checkout/internal/models"
	"route256/checkout/internal/repositories/carts_products_repo"
	cartProductsRepoMocks "route256/checkout/internal/repositories/carts_products_repo/mocks"
	"route256/checkout/internal/repositories/carts_repo"
	cartRepoMocks "route256/checkout/internal/repositories/carts_repo/mocks"
	"route256/checkout/internal/services/cart"
	desc "route256/checkout/pkg/v1/api"
	"route256/libs/transactor"
	txMock "route256/libs/transactor/mocks"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"golang.org/x/time/rate"
)

func TestListCartHandler(t *testing.T) {
	type cartsRepoMockFunc func(mc *minimock.Controller) carts_repo.ICartsRepo
	type cartsProductsRepoMockFunc func(mc *minimock.Controller) carts_products_repo.ICartsProductsRepo
	type productsClientMockFunc func(mc *minimock.Controller) products.IProductServiceClient
	type txManagerMockFunc func(mc *minimock.Controller) transactor.ITransactor
	type queryEngineMockFunc func(mc *minimock.Controller) transactor.IQueryEngine

	type args struct {
		ctx context.Context
		req *desc.ListCartRequest
	}

	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()
		n   = 2

		getCartsProductsRepoRes []models.Item
		expectedRes             = &desc.ListCartResponse{}
		limiter                 = rate.NewLimiter(rate.Every(time.Second/100), 10)
		userID                  = gofakeit.Int64()
		Err                     = errors.New("awaited error")
	)
	t.Cleanup(mc.Finish)

	for i := 0; i < n; i++ {
		sku := gofakeit.Uint32()
		count := gofakeit.Uint32()
		price := gofakeit.Uint32()

		getCartsProductsRepoRes = append(getCartsProductsRepoRes, models.Item{SKU: sku, Count: count})

		expectedRes.Items = append(expectedRes.Items, &desc.CartItem{
			Sku:   sku,
			Count: count,
			Name:  gofakeit.Name(),
			Price: price,
		})
		expectedRes.TotalPrice += price * count
	}

	tests := []struct {
		name                  string
		args                  args
		want                  *desc.ListCartResponse
		err                   error
		cartsRepoMock         cartsRepoMockFunc
		cartsProductsRepoMock cartsProductsRepoMockFunc
		productClientMock     productsClientMockFunc

		txManagerMock   txManagerMockFunc
		queryEngineMock queryEngineMockFunc
	}{
		{
			name: "positive case",
			args: args{
				ctx: ctx,
				req: &desc.ListCartRequest{User: userID},
			},
			want: expectedRes,
			err:  nil,
			queryEngineMock: func(mc *minimock.Controller) transactor.IQueryEngine {
				mock := txMock.NewIQueryEngineMock(mc)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) transactor.ITransactor {
				mock := txMock.NewITransactorMock(mc)
				return mock
			},
			cartsProductsRepoMock: func(mc *minimock.Controller) carts_products_repo.ICartsProductsRepo {
				mock := cartProductsRepoMocks.NewICartsProductsRepoMock(mc)
				mock.GetCartsProductsMock.Expect(ctx, userID).Return(getCartsProductsRepoRes, nil)
				return mock
			},
			productClientMock: func(mc *minimock.Controller) products.IProductServiceClient {
				mock := productsClientMocks.NewIProductServiceClientMock(mc)
				for _, item := range expectedRes.Items {
					mock.GetProductMock.
						When(ctx, item.Sku).
						Then(
							&models.ProductAttrs{
								Name:  item.Name,
								Price: item.Price,
							}, nil)
				}
				return mock
			},
			cartsRepoMock: func(mc *minimock.Controller) carts_repo.ICartsRepo {
				mock := cartRepoMocks.NewICartsRepoMock(mc)
				return mock
			},
		},
		{
			name: "negative case - products service error - GetProduct",
			args: args{
				ctx: ctx,
				req: &desc.ListCartRequest{User: userID},
			},
			want: nil,
			err:  Err,
			queryEngineMock: func(mc *minimock.Controller) transactor.IQueryEngine {
				mock := txMock.NewIQueryEngineMock(mc)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) transactor.ITransactor {
				mock := txMock.NewITransactorMock(mc)
				return mock
			},
			cartsProductsRepoMock: func(mc *minimock.Controller) carts_products_repo.ICartsProductsRepo {
				mock := cartProductsRepoMocks.NewICartsProductsRepoMock(mc)
				mock.GetCartsProductsMock.Expect(ctx, userID).Return(getCartsProductsRepoRes, nil)
				return mock
			},
			productClientMock: func(mc *minimock.Controller) products.IProductServiceClient {
				mock := productsClientMocks.NewIProductServiceClientMock(mc)
				for _, item := range expectedRes.Items {
					mock.GetProductMock.
						When(ctx, item.Sku).
						Then(nil, Err)
				}
				return mock
			},
			cartsRepoMock: func(mc *minimock.Controller) carts_repo.ICartsRepo {
				mock := cartRepoMocks.NewICartsRepoMock(mc)
				return mock
			},
		},
		{
			name: "negative case - repo error - GetCartsProducts",
			args: args{
				ctx: ctx,
				req: &desc.ListCartRequest{User: userID},
			},
			want: nil,
			err:  Err,
			queryEngineMock: func(mc *minimock.Controller) transactor.IQueryEngine {
				mock := txMock.NewIQueryEngineMock(mc)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) transactor.ITransactor {
				mock := txMock.NewITransactorMock(mc)
				return mock
			},
			cartsProductsRepoMock: func(mc *minimock.Controller) carts_products_repo.ICartsProductsRepo {
				mock := cartProductsRepoMocks.NewICartsProductsRepoMock(mc)
				mock.GetCartsProductsMock.Expect(ctx, userID).Return(nil, Err)
				return mock
			},
			productClientMock: func(mc *minimock.Controller) products.IProductServiceClient {
				mock := productsClientMocks.NewIProductServiceClientMock(mc)
				return mock
			},
			cartsRepoMock: func(mc *minimock.Controller) carts_repo.ICartsRepo {
				mock := cartRepoMocks.NewICartsRepoMock(mc)
				return mock
			},
		},
		//{
		//	name: "negative case - repository error",
		//	args: args{
		//		ctx: ctx,
		//		req: &emptypb.Empty{},
		//	},
		//	want: nil,
		//	err:  repoErr,
		//	noteRepositoryMock: func(mc *minimock.Controller) noteRepository.Repository {
		//		mock := noteRepoMocks.NewRepositoryMock(mc)
		//		mock.GetListMock.Expect(ctx).Return(nil, repoErr)
		//		return mock
		//	},
		//},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// дожидаемся завершения горутин в worker pool
			defer mc.Wait(time.Second * 2)

			t.Parallel()
			api := New(cart.NewMockService(
				tt.cartsRepoMock(mc),
				tt.cartsProductsRepoMock(mc),
				tt.productClientMock(mc),
				transactor.New(tt.queryEngineMock(mc)),
				limiter,
			))

			res, err := api.ListCart(tt.args.ctx, tt.args.req)

			if tt.want != nil {
				require.Equal(t, tt.want.TotalPrice, res.TotalPrice)
				require.ElementsMatch(t, tt.want.Items, res.Items)
			}

			if tt.err != nil {
				fmt.Println(err)
				require.ErrorContains(t, err, tt.err.Error())
			} else {
				require.Equal(t, tt.err, err)
			}

		})
	}
}

//func TestPurchaseHandler(t *testing.T) {
//	type cartsRepoMockFunc func(mc *minimock.Controller) carts_repo.ICartsRepo
//	type cartsProductsRepoMockFunc func(mc *minimock.Controller) carts_products_repo.ICartsProductsRepo
//	type lomsClientMockFunc func(mc *minimock.Controller) loms.ILOMSClient
//	type txManagerMockFunc func(mc *minimock.Controller) transactor.ITransactor
//	type queryEngineMockFunc func(mc *minimock.Controller) transactor.IQueryEngine
//
//	type args struct {
//		ctx context.Context
//		req *desc.PurchaseRequest
//	}
//
//	var (
//		mc      = minimock.NewController(t)
//		ctx     = context.Background()
//		n       = 2
//		userID  = gofakeit.Int64()
//		orderID = gofakeit.Int64()
//
//		tx                      = txMock.NewTxMock(t)
//		ctxTx                   = context.WithValue(ctx, transactor.TXKey, tx)
//		getCartsProductsRepoRes []models.Item
//		expectedRes             = &desc.PurchaseResponse{OrderID: orderID}
//		Err                     = errors.New("awaited error")
//	)
//	t.Cleanup(mc.Finish)
//
//	for i := 0; i < n; i++ {
//		getCartsProductsRepoRes = append(getCartsProductsRepoRes, models.Item{SKU: gofakeit.Uint32(), Count: gofakeit.Uint32()})
//	}
//
//	tests := []struct {
//		name string
//		args args
//		want *desc.PurchaseResponse
//		err  error
//
//		txManagerMock         txManagerMockFunc
//		queryEngineMock       queryEngineMockFunc
//		cartsProductsRepoMock cartsProductsRepoMockFunc
//		lomsClientMock        lomsClientMockFunc
//		cartsRepoMock         cartsRepoMockFunc
//	}{
//		{
//			name: "positive case",
//			args: args{
//				ctx: ctx,
//				req: &desc.PurchaseRequest{User: userID},
//			},
//			want: expectedRes,
//			err:  nil,
//			queryEngineMock: func(mc *minimock.Controller) transactor.IQueryEngine {
//				mock := txMock.NewIQueryEngineMock(mc)
//				mock.BeginTxMock.Expect(ctx, pgx.TxOptions{
//					IsoLevel: pgx.RepeatableRead,
//				}).Return(tx, nil)
//				tx.CommitMock.Expect(ctxTx).Return(nil)
//				return mock
//			},
//			txManagerMock: func(mc *minimock.Controller) transactor.ITransactor {
//				mock := txMock.NewITransactorMock(mc)
//				return mock
//			},
//			cartsProductsRepoMock: func(mc *minimock.Controller) carts_products_repo.ICartsProductsRepo {
//				mock := cartProductsRepoMocks.NewICartsProductsRepoMock(mc)
//				mock.GetCartsProductsMock.Expect(ctxTx, userID).Return(getCartsProductsRepoRes, nil)
//				return mock
//			},
//			lomsClientMock: func(mc *minimock.Controller) loms.ILOMSClient {
//				mock := lomsClientMocks.NewILOMSClientMock(mc)
//				mock.CreateOrderMock.Expect(ctxTx, userID, getCartsProductsRepoRes).Return(orderID, nil)
//				return mock
//			},
//			cartsRepoMock: func(mc *minimock.Controller) carts_repo.ICartsRepo {
//				mock := cartRepoMocks.NewICartsRepoMock(mc)
//				mock.PurchaseCartMock.Expect(ctxTx, userID).Return(uint64(orderID), nil)
//				return mock
//			},
//		},
//		{
//			name: "negative cas - validation error",
//			args: args{
//				ctx: ctx,
//				req: &desc.PurchaseRequest{User: 0},
//			},
//			want: nil,
//			err:  models.ErrEmptyUser,
//			queryEngineMock: func(mc *minimock.Controller) transactor.IQueryEngine {
//				mock := txMock.NewIQueryEngineMock(mc)
//				return mock
//			},
//			txManagerMock: func(mc *minimock.Controller) transactor.ITransactor {
//				mock := txMock.NewITransactorMock(mc)
//				return mock
//			},
//			cartsProductsRepoMock: func(mc *minimock.Controller) carts_products_repo.ICartsProductsRepo {
//				mock := cartProductsRepoMocks.NewICartsProductsRepoMock(mc)
//				return mock
//			},
//			lomsClientMock: func(mc *minimock.Controller) loms.ILOMSClient {
//				mock := lomsClientMocks.NewILOMSClientMock(mc)
//				return mock
//			},
//			cartsRepoMock: func(mc *minimock.Controller) carts_repo.ICartsRepo {
//				mock := cartRepoMocks.NewICartsRepoMock(mc)
//				return mock
//			},
//		},
//		{
//			name: "negative case - repo error - GetCartsProducts",
//			args: args{
//				ctx: ctx,
//				req: &desc.PurchaseRequest{User: userID},
//			},
//			want: nil,
//			err:  Err,
//			queryEngineMock: func(mc *minimock.Controller) transactor.IQueryEngine {
//				mock := txMock.NewIQueryEngineMock(mc)
//				mock.BeginTxMock.Expect(ctx, pgx.TxOptions{
//					IsoLevel: pgx.RepeatableRead,
//				}).Return(tx, nil)
//				tx.RollbackMock.Expect(ctxTx).Return(Err)
//				return mock
//			},
//			txManagerMock: func(mc *minimock.Controller) transactor.ITransactor {
//				mock := txMock.NewITransactorMock(mc)
//				return mock
//			},
//			cartsProductsRepoMock: func(mc *minimock.Controller) carts_products_repo.ICartsProductsRepo {
//				mock := cartProductsRepoMocks.NewICartsProductsRepoMock(mc)
//				mock.GetCartsProductsMock.Expect(ctxTx, userID).Return(nil, Err)
//				return mock
//			},
//			lomsClientMock: func(mc *minimock.Controller) loms.ILOMSClient {
//				mock := lomsClientMocks.NewILOMSClientMock(mc)
//				return mock
//			},
//			cartsRepoMock: func(mc *minimock.Controller) carts_repo.ICartsRepo {
//				mock := cartRepoMocks.NewICartsRepoMock(mc)
//				return mock
//			},
//		},
//		{
//			name: "negative case - loms client error - CreateOrder",
//			args: args{
//				ctx: ctx,
//				req: &desc.PurchaseRequest{User: userID},
//			},
//			want: nil,
//			err:  Err,
//			queryEngineMock: func(mc *minimock.Controller) transactor.IQueryEngine {
//				mock := txMock.NewIQueryEngineMock(mc)
//				mock.BeginTxMock.Expect(ctx, pgx.TxOptions{
//					IsoLevel: pgx.RepeatableRead,
//				}).Return(tx, nil)
//				tx.RollbackMock.Expect(ctxTx).Return(Err)
//				return mock
//			},
//			txManagerMock: func(mc *minimock.Controller) transactor.ITransactor {
//				mock := txMock.NewITransactorMock(mc)
//				return mock
//			},
//			cartsProductsRepoMock: func(mc *minimock.Controller) carts_products_repo.ICartsProductsRepo {
//				mock := cartProductsRepoMocks.NewICartsProductsRepoMock(mc)
//				mock.GetCartsProductsMock.Expect(ctxTx, userID).Return(getCartsProductsRepoRes, nil)
//				return mock
//			},
//			lomsClientMock: func(mc *minimock.Controller) loms.ILOMSClient {
//				mock := lomsClientMocks.NewILOMSClientMock(mc)
//				mock.CreateOrderMock.Expect(ctxTx, userID, getCartsProductsRepoRes).Return(0, Err)
//				return mock
//			},
//			cartsRepoMock: func(mc *minimock.Controller) carts_repo.ICartsRepo {
//				mock := cartRepoMocks.NewICartsRepoMock(mc)
//				return mock
//			},
//		},
//		{
//			name: "negative case - repo error - PurchaseCart",
//			args: args{
//				ctx: ctx,
//				req: &desc.PurchaseRequest{User: userID},
//			},
//			want: nil,
//			err:  Err,
//			queryEngineMock: func(mc *minimock.Controller) transactor.IQueryEngine {
//				mock := txMock.NewIQueryEngineMock(mc)
//				mock.BeginTxMock.Expect(ctx, pgx.TxOptions{
//					IsoLevel: pgx.RepeatableRead,
//				}).Return(tx, nil)
//				tx.RollbackMock.Expect(ctxTx).Return(Err)
//				return mock
//			},
//			txManagerMock: func(mc *minimock.Controller) transactor.ITransactor {
//				mock := txMock.NewITransactorMock(mc)
//				return mock
//			},
//			cartsProductsRepoMock: func(mc *minimock.Controller) carts_products_repo.ICartsProductsRepo {
//				mock := cartProductsRepoMocks.NewICartsProductsRepoMock(mc)
//				mock.GetCartsProductsMock.Expect(ctxTx, userID).Return(getCartsProductsRepoRes, nil)
//				return mock
//			},
//			lomsClientMock: func(mc *minimock.Controller) loms.ILOMSClient {
//				mock := lomsClientMocks.NewILOMSClientMock(mc)
//				mock.CreateOrderMock.Expect(ctxTx, userID, getCartsProductsRepoRes).Return(orderID, nil)
//				return mock
//			},
//			cartsRepoMock: func(mc *minimock.Controller) carts_repo.ICartsRepo {
//				mock := cartRepoMocks.NewICartsRepoMock(mc)
//				mock.PurchaseCartMock.Expect(ctxTx, userID).Return(0, Err)
//				return mock
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		tt := tt
//		t.Run(tt.name, func(t *testing.T) {
//			t.Parallel()
//			api := New(cart.NewMockService(
//				tt.cartsRepoMock(mc),
//				tt.cartsProductsRepoMock(mc),
//				tt.lomsClientMock(mc),
//				transactor.New(tt.queryEngineMock(mc)),
//			))
//
//			res, err := api.Purchase(tt.args.ctx, tt.args.req)
//			require.Equal(t, tt.want, res)
//			if tt.err != nil {
//				require.ErrorContains(t, err, tt.err.Error())
//			} else {
//				require.Equal(t, tt.err, err)
//			}
//		})
//	}
//}
