package http

import (
	"strings"

	accountUseCase "github.com/quachhoang2002/Go/internal/account/usecase"
	"github.com/quachhoang2002/Go/internal/domain"
	pkgErrors "github.com/quachhoang2002/Go/pkg/errors"
)

type addRequest struct {
	Owner    string `json:"owner"`
	Balance  int32  `json:"balance"`
	Currency string `json:"currency"`
}

func (r addRequest) toInput() (accountUseCase.CreateInput, error) {
	vErrCollect := pkgErrors.NewValidationErrorCollector()

	if strings.TrimSpace(r.Owner) == "" {
		vErrCollect.Add(pkgErrors.NewValidationError("Onwer", errMsgNameIsRequired))
	}

	if vErrCollect.HasError() {
		return accountUseCase.CreateInput{}, vErrCollect
	}

	return accountUseCase.CreateInput{
		Owner:    r.Owner,
		Balance:  r.Balance,
		Currency: r.Currency,
	}, nil
}

type accountReponse struct {
	ID        int64  `json:"id"`
	Owner     string `json:"owner"`
	Balance   int32  `json:"balance"`
	Currency  string `json:"currency"`
	CreatedAt string `json:"created_at"`
}

func newListResponse(accounts []domain.Account) []accountReponse {
	data := []accountReponse{}
	for _, account := range accounts {
		data = append(data, accountReponse{
			ID:       account.ID,
			Owner:    account.Owner,
			Balance:  account.Balance,
			Currency: account.Currency,
		})
	}

	return data
}
