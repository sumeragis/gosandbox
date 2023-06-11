// FAIL
// https://github.com/golang/mock/issues/621

package usecase

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/google/go-cmp/cmp"
// 	"github.com/sumeragis/sandbox/backend/domain/entity"
// 	mock_repository "github.com/sumeragis/sandbox/backend/domain/repository/mock"
// 	errorx "github.com/sumeragis/sandbox/backend/errors"
// )

// func TestUserUseCase_Get(t *testing.T) {

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	type args struct {
// 		ctx context.Context
// 		id int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		mockRepository func(r mock_repository.MockUserRepository)
// 		want *entity.User
// 		wantErr error
// 	}{
// 		{
// 			name: "success",
// 			args: args{ctx: context.Background(), id: 1},
// 			mockRepository: func(m mock_repository.MockUserRepository) {
// 				m.EXPECT().FindByID(gomock.Any(), 1).Return(&entity.User{ID: 1, Name: "name1"}, nil)
// 			},
// 			want: &entity.User{ID: 1, Name: "name1"},
// 		},
// 		{
// 			name: "error: not_found",
// 			args: args{ctx: context.Background(), id: 1},
// 			mockRepository: func(m mock_repository.MockUserRepository) {
// 				m.EXPECT().FindByID(gomock.Any(), 1).Return(nil, errors.New("not found entity!"))
// 			},
// 			want: &entity.User{ID: 1, Name: "name1"},
// 			wantErr: errors.New("not found entity!"),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mr := mock_repository.NewMockUserRepository(ctrl)
// 			tt.mockRepository(*mr)

// 			u := NewUserUseCase(mr)
// 			res, err := u.Get(tt.args.ctx, tt.args.id)
// 			if err != nil {
// 				if tt.wantErr != nil {
// 					if diff := cmp.Diff(tt.wantErr.Error(), err.Error()); diff != "" {
// 						t.Errorf("diff = %v, want %v", err, tt.wantErr)
// 					}
// 					return
// 				} else {
// 					t.Errorf("Error: failed to Get err=%s", err.Error())
// 					return
// 				}
// 			}

// 			if diff := cmp.Diff(tt.want, res); diff != "" {
// 				t.Errorf("got = %v, want %v", res, tt.want)
// 				return
// 			}

// 		})
// 	}
// }

// func TestUserUseCase_Create(t *testing.T) {

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	type args struct {
// 		ctx context.Context
// 		entity *entity.User
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		mockRepository func(r mock_repository.MockUserRepository)
// 		want *entity.User
// 		wantErr error
// 	}{
// 		{
// 			name: "success",
// 			args: args{ctx: context.Background(), entity: &entity.User{ID: 1, Name: "name1"}},
// 			mockRepository: func(m mock_repository.MockUserRepository) {
// 				m.EXPECT().Save(gomock.Any(), &entity.User{ID: 1, Name: "name1"}).Return(nil)
// 			},
// 			want: nil,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mr := mock_repository.NewMockUserRepository(ctrl)
// 			tt.mockRepository(*mr)

// 			u := NewUserUseCase(mr)
// 			err := u.Create(tt.args.ctx, tt.args.entity)
// 			if err != nil {
// 				if tt.wantErr != nil {
// 					if diff := cmp.Diff(tt.wantErr.Error(), err.Error()); diff != "" {
// 						t.Errorf("diff = %v, want %v", err, tt.wantErr)
// 					}
// 					return
// 				} else {
// 					t.Errorf("Error: failed to Get err=%s", err.Error())
// 					return
// 				}
// 			}
// 		})
// 	}
// }

// func TestUserUseCase_Update(t *testing.T) {

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	type args struct {
// 		ctx context.Context
// 		entity *entity.User
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		mockRepository func(r mock_repository.MockUserRepository)
// 		want *entity.User
// 		wantErr error
// 	}{
// 		{
// 			name: "success",
// 			args: args{ctx: context.Background(), entity: &entity.User{ID: 1, Name: "name2"}},
// 			mockRepository: func(m mock_repository.MockUserRepository) {
// 				m.EXPECT().FindByID(gomock.Any(), 1).Return(&entity.User{ID: 1, Name: "name1"}, nil)
// 				m.EXPECT().Update(gomock.Any(), &entity.User{ID: 1, Name: "name2"}).Return(nil)
// 			},
// 			want: nil,
// 		},
// 		{
// 			name: "success",
// 			args: args{ctx: context.Background(), entity: &entity.User{ID: 1, Name: "name2"}},
// 			mockRepository: func(m mock_repository.MockUserRepository) {
// 				m.EXPECT().FindByID(gomock.Any(), 1).Return(nil, nil)
// 			},
// 			want: nil,
// 			wantErr: errorx.ERR_NOT_FOUND,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mr := mock_repository.NewMockUserRepository(ctrl)
// 			tt.mockRepository(*mr)

// 			u := NewUserUseCase(mr)
// 			err := u.Update(tt.args.ctx, tt.args.entity)
// 			if err != nil {
// 				if tt.wantErr != nil {
// 					if diff := cmp.Diff(tt.wantErr.Error(), err.Error()); diff != "" {
// 						t.Errorf("diff = %v, want %v", err, tt.wantErr)
// 					}
// 					return
// 				} else {
// 					t.Errorf("Error: failed to Get err=%s", err.Error())
// 					return
// 				}
// 			}
// 		})
// 	}
// }

// func TestUserUseCase_Delete(t *testing.T) {

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	type args struct {
// 		ctx context.Context
// 		id int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		mockRepository func(r mock_repository.MockUserRepository)
// 		want *entity.User
// 		wantErr error
// 	}{
// 		{
// 			name: "success",
// 			args: args{ctx: context.Background(), id: 1},
// 			mockRepository: func(m mock_repository.MockUserRepository) {
// 				m.EXPECT().Delete(gomock.Any(), 1).Return(nil)
// 			},
// 			want: nil,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mr := mock_repository.NewMockUserRepository(ctrl)
// 			tt.mockRepository(*mr)

// 			u := NewUserUseCase(mr)
// 			err := u.Delete(tt.args.ctx, tt.args.id)
// 			if err != nil {
// 				if tt.wantErr != nil {
// 					if diff := cmp.Diff(tt.wantErr.Error(), err.Error()); diff != "" {
// 						t.Errorf("diff = %v, want %v", err, tt.wantErr)
// 					}
// 					return
// 				} else {
// 					t.Errorf("Error: failed to Get err=%s", err.Error())
// 					return
// 				}
// 			}
// 		})
// 	}
// }