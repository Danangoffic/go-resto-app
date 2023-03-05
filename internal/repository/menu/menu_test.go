package menu

import (
	"context"
	"database/sql"
	"reflect"
	"regexp"
	"resto-app/internal/model"
	"resto-app/internal/model/constant"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_menuRepo_GetMenuList(t *testing.T) {
	type args struct {
		ctx      context.Context
		menuType string
	}
	tests := []struct {
		name     string
		r        *menuRepo
		args     args
		want     []model.MenuItem
		wantErr  bool
		initMock func() (*sql.DB, sqlmock.Sqlmock, error)
	}{
		{
			name: "success get menu list",
			args: args{ctx: context.Background(), menuType: string(constant.MenuTypeFood)},
			initMock: func() (*sql.DB, sqlmock.Sqlmock, error) {
				db, mock, err := sqlmock.New()
				mock.ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "menu_items"`),
				).WillReturnRows(sqlmock.NewRows([]string{
					"name",
					"order_code",
					"price",
					"type",
				}).AddRow("Nasi Uduk", "nasi_uduk", 20000, constant.MenuTypeFood))
				return db, mock, err
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetMenuList(tt.args.ctx, tt.args.menuType)
			if (err != nil) != tt.wantErr {
				t.Errorf("menuRepo.GetMenuList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("menuRepo.GetMenuList() = %v, want %v", got, tt.want)
			}
		})
	}
}
