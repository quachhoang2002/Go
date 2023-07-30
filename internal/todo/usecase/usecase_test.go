package usecase

import (
	context "context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gma-vietnam/tanca-event/internal/model"
	"gitlab.com/gma-vietnam/tanca-event/internal/todo/repository"
	"gitlab.com/gma-vietnam/tanca-event/pkg/log"
)

type mockDeps struct {
	repo *repository.MockRepository
}

func initUseCase(t *testing.T) (Usecase, mockDeps) {
	repo := repository.NewMockRepository(t)
	l := log.InitializeZapLogger(log.NewTestZapConfig())

	return New(l, repo), mockDeps{
		repo: repo,
	}
}

func TestTodoUsecase_Create(t *testing.T) {
	type mockRepo struct {
		expCall bool
		input   model.Todo
		err     error
	}

	tcs := map[string]struct {
		input    CreateInput
		mockRepo mockRepo
		wantErr  error
	}{
		"success": {
			input: CreateInput{
				Name:        "test",
				Description: "test",
			},
			mockRepo: mockRepo{
				expCall: true,
				input: model.Todo{
					Name:        "test",
					Description: "test",
				},
			},
		},
		"error": {
			input: CreateInput{
				Name:        "test",
				Description: "test",
			},
			mockRepo: mockRepo{
				expCall: true,
				input: model.Todo{
					Name:        "test",
					Description: "test",
				},
				err: errors.New("Some error"),
			},
			wantErr: errors.New("Some error"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// GIVEN
			ctx := context.Background()

			uc, deps := initUseCase(t)

			if tc.mockRepo.expCall {
				deps.repo.EXPECT().Create(ctx, tc.mockRepo.input).
					Return(tc.mockRepo.err).
					Once()
			}

			// WHEN
			err := uc.Create(ctx, tc.input)

			// THEN
			if tc.wantErr != nil {
				require.EqualError(t, tc.wantErr, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTodoUsecase_All(t *testing.T) {
	type mockRepo struct {
		expCall bool
		output  []model.Todo
		err     error
	}

	tcs := map[string]struct {
		mockRepo mockRepo
		wantRes  []model.Todo
		wantErr  error
	}{
		"success": {
			mockRepo: mockRepo{
				expCall: true,
				output: []model.Todo{
					{
						ID:          1,
						Name:        "test-1",
						Description: "test",
					},
					{
						ID:          2,
						Name:        "test-2",
						Description: "test",
					},
				},
			},
			wantRes: []model.Todo{
				{
					ID:          1,
					Name:        "test-1",
					Description: "test",
				},
				{
					ID:          2,
					Name:        "test-2",
					Description: "test",
				},
			},
		},
		"error": {
			mockRepo: mockRepo{
				expCall: true,
				err:     errors.New("Some error"),
			},
			wantErr: errors.New("Some error"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// GIVEN
			ctx := context.Background()

			uc, deps := initUseCase(t)

			if tc.mockRepo.expCall {
				deps.repo.EXPECT().All(ctx).
					Return(tc.mockRepo.output, tc.mockRepo.err).
					Once()
			}

			// WHEN
			res, err := uc.All(ctx)

			// THEN
			if tc.wantErr != nil {
				require.EqualError(t, tc.wantErr, err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.wantRes, res)
			}
		})
	}
}

func TestTodoUsecase_Delete(t *testing.T) {
	type mockRepo struct {
		expCall bool
		input   int
		err     error
	}

	tcs := map[string]struct {
		input    int
		mockRepo mockRepo
		wantErr  error
	}{
		"success": {
			mockRepo: mockRepo{
				expCall: true,
			},
		},
		"error not found": {
			mockRepo: mockRepo{
				expCall: true,
				err:     repository.ErrNotFound,
			},
			wantErr: ErrNotFound,
		},
		"internal error": {
			mockRepo: mockRepo{
				expCall: true,
				err:     errors.New("Some error"),
			},
			wantErr: errors.New("Some error"),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// GIVEN
			ctx := context.Background()

			uc, deps := initUseCase(t)

			if tc.mockRepo.expCall {
				deps.repo.EXPECT().Delete(ctx, tc.mockRepo.input).
					Return(tc.mockRepo.err).
					Once()
			}

			// WHEN
			err := uc.Delete(ctx, tc.input)

			// THEN
			if tc.wantErr != nil {
				require.EqualError(t, tc.wantErr, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
