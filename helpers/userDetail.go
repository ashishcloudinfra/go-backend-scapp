package helpers

import (
	"fmt"
	"log"
	"net/http"

	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/secrets"
	"server.simplifycontrol.com/types"
)

func InsertIntoUserDetailTable(w http.ResponseWriter, values types.UserDetailReq, iamId string) error {
	query := fmt.Sprintf(`INSERT INTO %s (firstName, lastName, email, phone, address, city, state, zip, country, dob, gender, metadata, iamId) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`, database.PERSONAL_DETAIL_TABLE_NAME)

	encryptFirstName, err := secrets.EncryptAESGCM(values.FirstName)
	if err != nil {
		return err
	}
	encryptLastName, err := secrets.EncryptAESGCM(values.LastName)
	if err != nil {
		return err
	}
	encryptEmail, err := secrets.EncryptAESGCM(values.Email)
	if err != nil {
		return err
	}
	encryptContact, err := secrets.EncryptAESGCM(values.Phone)
	if err != nil {
		return err
	}
	encryptAddress, err := secrets.EncryptAESGCM(values.Address)
	if err != nil {
		return err
	}
	encryptCity, err := secrets.EncryptAESGCM(values.City)
	if err != nil {
		return err
	}
	encryptState, err := secrets.EncryptAESGCM(values.State)
	if err != nil {
		return err
	}
	encryptZIP, err := secrets.EncryptAESGCM(values.Zip)
	if err != nil {
		return err
	}
	encryptCountry, err := secrets.EncryptAESGCM(values.Country)
	if err != nil {
		return err
	}
	encryptDob, err := secrets.EncryptAESGCM(values.Dob)
	if err != nil {
		return err
	}
	encryptGender, err := secrets.EncryptAESGCM(values.Gender)
	if err != nil {
		return err
	}
	encryptMetadata, err := secrets.EncryptAESGCM(values.Metadata)
	if err != nil {
		return err
	}
	_, err = database.Query(
		query,
		encryptFirstName,
		encryptLastName,
		encryptEmail,
		encryptContact,
		encryptAddress,
		encryptCity,
		encryptState,
		encryptZIP,
		encryptCountry,
		encryptDob,
		encryptGender,
		encryptMetadata,
		iamId,
	)
	if err != nil {
		log.Println("Error inserting into user detail table:", err.Error())
		return err
	}
	return nil
}

func UpdateUserDetailTable(w http.ResponseWriter, values types.UserDetailReq, iamId string) error {
	query := fmt.Sprintf(`UPDATE %s SET firstName = $1, lastName = $2, email = $3, phone = $4, address = $5, city = $6, state = $7, zip = $8, country = $9, dob = $10, gender = $11, metadata = $12 WHERE iamId = $13`, database.PERSONAL_DETAIL_TABLE_NAME)

	encryptFirstName, err := secrets.EncryptAESGCM(values.FirstName)
	if err != nil {
		return err
	}
	encryptLastName, err := secrets.EncryptAESGCM(values.LastName)
	if err != nil {
		return err
	}
	encryptEmail, err := secrets.EncryptAESGCM(values.Email)
	if err != nil {
		return err
	}
	encryptContact, err := secrets.EncryptAESGCM(values.Phone)
	if err != nil {
		return err
	}
	encryptAddress, err := secrets.EncryptAESGCM(values.Address)
	if err != nil {
		return err
	}
	encryptCity, err := secrets.EncryptAESGCM(values.City)
	if err != nil {
		return err
	}
	encryptState, err := secrets.EncryptAESGCM(values.State)
	if err != nil {
		return err
	}
	encryptZIP, err := secrets.EncryptAESGCM(values.Zip)
	if err != nil {
		return err
	}
	encryptCountry, err := secrets.EncryptAESGCM(values.Country)
	if err != nil {
		return err
	}
	encryptDob, err := secrets.EncryptAESGCM(values.Dob)
	if err != nil {
		return err
	}
	encryptGender, err := secrets.EncryptAESGCM(values.Gender)
	if err != nil {
		return err
	}
	encryptMetadata, err := secrets.EncryptAESGCM(values.Metadata)
	if err != nil {
		return err
	}
	_, err = database.Query(
		query,
		encryptFirstName,
		encryptLastName,
		encryptEmail,
		encryptContact,
		encryptAddress,
		encryptCity,
		encryptState,
		encryptZIP,
		encryptCountry,
		encryptDob,
		encryptGender,
		encryptMetadata,
		iamId,
	)
	if err != nil {
		log.Println("Error inserting into user detail table:", err.Error())
		return err
	}
	return nil
}
