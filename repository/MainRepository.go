package repository

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
	"uid/models"

	"github.com/prevoty/pdo"
)

var TimeZone, _ = time.LoadLocation("America/New_York")

func CreateConnection() *pdo.MySQL {
	connection, err := pdo.NewMySQL("gouser:gopass@tcp(mysql-go-db:3306)/godb")
	if err != nil {
		log.Fatal(err)
	}
	return connection
}

func CurrentTimeStamp() string {
	return time.Now().In(TimeZone).Format("2006-01-02 15:04:05")
}

func sqlErrorChecker(err error) error {
	switch err {
		case sql.ErrNoRows:
			log.Println("Not found the row")
			return err
		case nil:
			log.Println("found a some row")
			return nil
		default:
			log.Fatal(err)
			return err
	}
}

func getUIDByUserID(userID int64, connection *pdo.MySQL) (*models.Uid, error) {
	uid := new(models.Uid)
	err := connection.Find(uid, "WHERE `id` = ?", userID)
	return uid, sqlErrorChecker(err)
}

func getRequestByID(requestID int64, connection *pdo.MySQL) (*models.Request, error) {
	request := new(models.Request)
	err := connection.Find(request, "WHERE `id` = ?", requestID)
	return request, sqlErrorChecker(err)
}

func MakeRequest(requestAPIRequest *models.RequestAPIRequest, connection *pdo.MySQL) (*models.Request, error) {
	uid, err := getUIDByUserID(requestAPIRequest.UserID, connection)
	
	if (err != nil) {
		uid_id, _ := connection.Create(&models.Uid{
			CreatedAt: CurrentTimeStamp(),
			UserId: requestAPIRequest.UserID,
			Access: requestAPIRequest.Access,
			Hidden: requestAPIRequest.Hidden,
		})
		uid, _ = getUIDByUserID(uid_id, connection)
	}


	hash := fmt.Sprintf("%x", sha1.Sum(
		[]byte(strconv.Itoa(int(uid.UserId)) + strconv.Itoa(int(time.Now().UnixMicro()))),
	))

	request_id, err := connection.Create(&models.Request{
		UId: uid.Id,
		CreatedAt: CurrentTimeStamp(),
		Hash: hash,
		ServiceFamily: requestAPIRequest.ServiceFamily,
		Data: requestAPIRequest.Data,
	})

	if err != nil {
		return nil, err
	}

	request, err := getRequestByID(request_id, connection)

	if err != nil {
		return nil, err
	}
	return request, nil
}
