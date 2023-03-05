package resto

import (
	"context"
	"reflect"
	"resto-app/internal/model"
	"resto-app/internal/model/constant"
	"resto-app/internal/repository/menu"
	"resto-app/internal/repository/order"
	"resto-app/internal/repository/user"
	"testing"
)

func Test_restoUsecase_GetOrderData(t *testing.T) {
	type args struct {
		ctx     context.Context
		request model.GetOrderDataRequest
	}
	tests := []struct {
		name    string
		s       *restoUsecase
		args    args
		want    model.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetOrderData(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("restoUsecase.GetOrderData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoUsecase.GetOrderData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_restoUsecase_GetMenuList(t *testing.T) {
	type args struct {
		ctx      context.Context
		menuType string
	}
	tests := []struct {
		name    string
		r       *restoUsecase
		args    args
		want    []model.MenuItem
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success get menu",
			args: args{
				ctx:      context.Background(),
				menuType: string(constant.MenuTypeFood),
			},
			want: func() []model.MenuItem {
				// ctrl := gomock.NewController(t)
				// mock := mocks.NewMockMenuRepository(ctrl)
				MenuItem := []model.MenuItem{
					{
						Name:      "Nasi Uduk",
						OrderCode: "nasi_uduk",
						Price:     20000,
						Type:      constant.MenuTypeFood,
					},
				}
				// mock.EXPECT().GetMenuList(context.Background(), string(constant.MenuTypeFood)).
				// 	Times(1).
				// 	Return(MenuItem, nil)
				return MenuItem
			}(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetMenuList(tt.args.ctx, tt.args.menuType)
			if (err != nil) != tt.wantErr {
				t.Errorf("restoUsecase.GetMenuList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoUsecase.GetMenuList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUsecase(t *testing.T) {
	type args struct {
		menuRepo  menu.Repository
		orderRepo order.Repository
		userRepo  user.Repository
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUsecase(tt.args.menuRepo, tt.args.orderRepo, tt.args.userRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_restoUsecase_GetMenu(t *testing.T) {
	type args struct {
		ctx       context.Context
		orderCode string
	}
	tests := []struct {
		name    string
		s       *restoUsecase
		args    args
		want    model.MenuItem
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetMenu(tt.args.ctx, tt.args.orderCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("restoUsecase.GetMenu() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoUsecase.GetMenu() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_restoUsecase_Order(t *testing.T) {
	type args struct {
		ctx     context.Context
		request model.OrderMenuRequest
	}
	tests := []struct {
		name    string
		s       *restoUsecase
		args    args
		want    model.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Order(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("restoUsecase.Order() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoUsecase.Order() = %v, want %v", got, tt.want)
			}
		})
	}
}
