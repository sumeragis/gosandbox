package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/sumeragis/sandbox/backend/domain/entity"
	mock_repository "github.com/sumeragis/sandbox/backend/domain/repository/mock"
)


func TestUserUseCase_Create(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		id int
	}
	tests := []struct {
		name string
		args args
		mockRepository func(r mock_repository.MockUserRepository)
		want *entity.User
		wantErr error
	}{
		{
			name: "success",
			args: args{ctx: context.Background(), id: 1},
			mockRepository: func(m mock_repository.MockUserRepository) {
				m.EXPECT().FindByID(gomock.Any(), 1).Return(&entity.User{ID: 1, Name: "name1"}, nil)
			},
			want: &entity.User{ID: 1, Name: "name1"},
		},
		{
			name: "error: not_found",
			args: args{ctx: context.Background(), id: 1},
			mockRepository: func(m mock_repository.MockUserRepository) {
				m.EXPECT().FindByID(gomock.Any(), 1).Return(nil, errors.New("not found entity!"))
			},
			want: &entity.User{ID: 1, Name: "name1"},
			wantErr: errors.New("not found entity!"),
		},
	}

	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := mock_repository.NewMockUserRepository(ctrl)
			tt.mockRepository(*mr)
			
			u := NewUserUseCase(mr)
			res, err := u.Get(tt.args.ctx, tt.args.id)
			if err != nil {
				if tt.wantErr != nil {
					if diff := cmp.Diff(tt.wantErr.Error(), err.Error()); diff != "" {
						t.Errorf("diff = %v, want %v", err, tt.wantErr)
					}
					return
				} else {
					t.Errorf("Error: failed to Get err=%s", err.Error())	
					return
				}
			}

			if diff := cmp.Diff(tt.want, res); diff != "" {
				t.Errorf("got = %v, want %v", res, tt.want)
				return
			}

		})
	}
}