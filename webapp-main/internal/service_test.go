package internal

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type MockTravellerStorage struct {
	GetFn    func(ctx context.Context, id uuid.UUID) (Traveller, error)
	CreateFn func(ctx context.Context, params CreateTravellerPayload) (uuid.UUID, error)
}

func (m *MockTravellerStorage) Get(ctx context.Context, id uuid.UUID) (Traveller, error) {
	return m.GetFn(ctx, id)
}

func (m *MockTravellerStorage) Create(ctx context.Context, params CreateTravellerPayload) (uuid.UUID, error) {
	return m.CreateFn(ctx, params)
}

func TestTravellers_GetTraveller(t *testing.T) {
	validID := uuid.New()
	mockTraveller := Traveller{
		ID:        validID,
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	tests := []struct {
		name      string
		id        uuid.UUID
		mockSetup func(mock *MockTravellerStorage)
		want      Traveller
		wantErr   bool
	}{
		{
			name: "Success flow",
			id:   validID,
			mockSetup: func(mock *MockTravellerStorage) {
				mock.GetFn = func(ctx context.Context, id uuid.UUID) (Traveller, error) {
					return mockTraveller, nil
				}
			},
			want:    mockTraveller,
			wantErr: false,
		},
		{
			name: "Fail flow: traveller not found",
			id:   validID,
			mockSetup: func(mock *MockTravellerStorage) {
				mock.GetFn = func(ctx context.Context, id uuid.UUID) (Traveller, error) {
					return Traveller{}, ErrNoResource
				}
			},
			want:    Traveller{},
			wantErr: true,
		},
		{
			name: "Edge case: empty nil UUID",
			id:   uuid.Nil,
			mockSetup: func(mock *MockTravellerStorage) {
				// Get() shouldn't be called for Nil UUID
			},
			want:    Traveller{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := &MockTravellerStorage{}
			if tt.mockSetup != nil {
				tt.mockSetup(mockStorage)
			}

			svc := NewTravellers(mockStorage)
			got, err := svc.GetTraveller(context.Background(), tt.id)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func TestTravellers_CreateTraveller(t *testing.T) {
	t.Parallel()
	createdID := uuid.New()

	tests := []struct {
		name      string
		payload   CreateTravellerPayload
		mockSetup func(mock *MockTravellerStorage)
		want      uuid.UUID
		wantErr   bool
	}{
		{
			name: "Success flow",
			payload: CreateTravellerPayload{
				FirstName: "Jane",
				LastName:  "Doe",
				Age:       25,
			},
			mockSetup: func(mock *MockTravellerStorage) {
				mock.CreateFn = func(ctx context.Context, params CreateTravellerPayload) (uuid.UUID, error) {
					return createdID, nil
				}
			},
			want:    createdID,
			wantErr: false,
		},
		{
			name: "Fail flow: storage error",
			payload: CreateTravellerPayload{
				FirstName: "Jane",
				LastName:  "Doe",
				Age:       25,
			},
			mockSetup: func(mock *MockTravellerStorage) {
				mock.CreateFn = func(ctx context.Context, params CreateTravellerPayload) (uuid.UUID, error) {
					return uuid.Nil, errors.New("db error")
				}
			},
			want:    uuid.Nil,
			wantErr: true,
		},
		{
			name: "Edge case: empty first name",
			payload: CreateTravellerPayload{
				FirstName: "",
				LastName:  "Doe",
				Age:       25,
			},
			mockSetup: func(mock *MockTravellerStorage) {
				// Storage shouldn't be called if validation fails
			},
			want:    uuid.Nil,
			wantErr: true,
		},
		{
			name: "Edge case: empty last name",
			payload: CreateTravellerPayload{
				FirstName: "Jane",
				LastName:  "",
				Age:       25,
			},
			mockSetup: func(mock *MockTravellerStorage) {
				// Storage shouldn't be called if validation fails
			},
			want:    uuid.Nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockStorage := &MockTravellerStorage{}
			if tt.mockSetup != nil {
				tt.mockSetup(mockStorage)
			}

			svc := NewTravellers(mockStorage)
			got, err := svc.CreateTraveller(context.Background(), tt.payload)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}
