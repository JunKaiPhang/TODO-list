package service

import (
	"errors"
	"personal/TODO-list/model"
	"personal/TODO-list/pkg/util"
)

func CreateInsertJwtToken(user model.User, email, name string) (tokenString string, err error) {
	watId := GenerateWebAccessTokenID()

	tokenString, _ = util.CreateJWT(watId, email, name)
	if tokenString == "" {
		return "", errors.New("unable_to_generate_auth_token")
	}

	_, err = InsertToken(user, watId)
	if err != nil {
		return "", errors.New("unable_to_insert_token")
	}

	return tokenString, nil
}

/* Use; logout (delete token) */ //
/* On success; return nil */ //
/* On error; return "failed_to_logout" or "failed_to_get_user_data" */ //
func Logout(watId, logoutmeth, username string) error {
	if err := model.DeleteTokenByWatId(watId); err != nil {
		return errors.New("failed_to_logout")
	}

	return nil
}
