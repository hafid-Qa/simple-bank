package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	mockdb "go_api/db/mock"
	db "go_api/db/sqlc"
	"go_api/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetAccount(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/accounts/%d", account.ID)
	req, err := http.NewRequest("GET", url, nil)
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, req)
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchAccount(t, recorder.Body, account)

}

func randomAccount() db.Account {
	return db.Account{
		ID:       utils.RandomInt(1, 1000),
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomAmount(),
		Currency: utils.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	var gotAccount db.Account
	err := json.Unmarshal(body.Bytes(), &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}
