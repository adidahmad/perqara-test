package service

import (
	"reflect"
	"testing"
	"time"

	"github.com/adidahmad/perqara-test/core/users/domain"
	"github.com/adidahmad/perqara-test/core/users/entity"
	usersPort "github.com/adidahmad/perqara-test/core/users/port"
	mockRepository "github.com/adidahmad/perqara-test/core/users/port/mocks"
	"github.com/golang/mock/gomock"
)

func TestNewUsersService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepository.NewMockIUsersRepository(ctrl)

	type args struct {
		usersRepo usersPort.IUsersRepository
	}
	tests := []struct {
		name string
		args args
		want usersPort.IUsersService
	}{
		{
			name: "Success",
			args: args{
				usersRepo: mockRepo,
			},
			want: UsersService{
				UsersRepository: mockRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUsersService(tt.args.usersRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsersService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersService_GetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepository.NewMockIUsersRepository(ctrl)

	// Mock data
	mockData := []*entity.Users{
		{ID: 1, Email: "test1@example.com", Password: "password1", IsActive: true, CreatedAt: time.Now()},
		{ID: 2, Email: "test2@example.com", Password: "password2", IsActive: true, CreatedAt: time.Now()},
	}

	mockRepo.EXPECT().FindAll().Return(mockData, nil)

	type fields struct {
		UsersRepository usersPort.IUsersRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.Users
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				UsersRepository: mockRepo,
			},
			want: []domain.Users{
				{ID: 1, Email: "test1@example.com", Password: "password1", IsActive: true, CreatedAt: mockData[0].CreatedAt},
				{ID: 2, Email: "test2@example.com", Password: "password2", IsActive: true, CreatedAt: mockData[1].CreatedAt},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UsersService{
				UsersRepository: tt.fields.UsersRepository,
			}
			got, err := s.GetList()
			if (err != nil) != tt.wantErr {
				t.Errorf("UsersService.GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersService.GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepository.NewMockIUsersRepository(ctrl)

	// Mock data
	mockData := entity.Users{
		ID: 1, Email: "test@example.com", Password: "password", IsActive: true, CreatedAt: time.Now(),
	}

	mockRepo.EXPECT().FindById("1").Return(mockData, nil)

	type fields struct {
		UsersRepository usersPort.IUsersRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Users
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				UsersRepository: mockRepo,
			},
			args: args{
				id: "1",
			},
			want: domain.Users{
				ID:        mockData.ID,
				Email:     mockData.Email,
				Password:  mockData.Password,
				IsActive:  mockData.IsActive,
				CreatedAt: mockData.CreatedAt,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UsersService{
				UsersRepository: tt.fields.UsersRepository,
			}
			got, err := s.GetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsersService.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersService.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepository.NewMockIUsersRepository(ctrl)

	// Mock data
	mockData := entity.Users{
		ID: 1, Email: "test@example.com", Password: "password", IsActive: true, CreatedAt: time.Now(),
	}

	mockRepo.EXPECT().Insert(gomock.Any()).Return(mockData, nil)

	type fields struct {
		UsersRepository usersPort.IUsersRepository
	}
	type args struct {
		data domain.CreateUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Users
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				UsersRepository: mockRepo,
			},
			args: args{
				data: domain.CreateUserRequest{
					Email:    "test@example.com",
					Password: "password",
					IsActive: true,
				},
			},
			want: domain.Users{
				ID:        mockData.ID,
				Email:     mockData.Email,
				Password:  mockData.Password,
				IsActive:  mockData.IsActive,
				CreatedAt: mockData.CreatedAt,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UsersService{
				UsersRepository: tt.fields.UsersRepository,
			}
			got, err := s.Create(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsersService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepository.NewMockIUsersRepository(ctrl)

	// Mock data
	mockData := entity.Users{
		ID: 1, Email: "test@example.com", Password: "password", IsActive: true, CreatedAt: time.Now(),
	}

	mockRepo.EXPECT().Update("1", gomock.Any()).Return(mockData, nil)

	type fields struct {
		UsersRepository usersPort.IUsersRepository
	}
	type args struct {
		id   string
		data domain.UpdateUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Users
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				UsersRepository: mockRepo,
			},
			args: args{
				id: "1",
				data: domain.UpdateUserRequest{
					Email:    "test@example.com",
					Password: "password",
					IsActive: true,
				},
			},
			want: domain.Users{
				ID:        mockData.ID,
				Email:     mockData.Email,
				Password:  mockData.Password,
				IsActive:  mockData.IsActive,
				CreatedAt: mockData.CreatedAt,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UsersService{
				UsersRepository: tt.fields.UsersRepository,
			}
			got, err := s.Update(tt.args.id, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsersService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersService.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersService_DeleteById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepository.NewMockIUsersRepository(ctrl)

	mockRepo.EXPECT().DeleteById("1").Return(nil)

	type fields struct {
		UsersRepository usersPort.IUsersRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				UsersRepository: mockRepo,
			},
			args:    args{id: "1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UsersService{
				UsersRepository: tt.fields.UsersRepository,
			}
			if err := s.DeleteById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UsersService.DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
