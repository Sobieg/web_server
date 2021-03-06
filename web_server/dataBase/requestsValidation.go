package dataBase

import (
	"regexp"
	"web_server/entities"
)

func (db *DataBase) validateUserInfo(userInfo *entities.Registration) error {
	r := regexp.MustCompile(nickReg)
	if !r.MatchString(userInfo.UserId) {
		return WrongSymbolsError
	}
	if len(userInfo.UserId) > 256 {
		return db.errorConstructLong(FieldTooLongError, "256", userIdField)
	}
	if len(userInfo.Email) > 256 {
		return db.errorConstructLong(FieldTooLongError, "256", emailField)
	}
	if len(userInfo.Password) > 256 {
		return db.errorConstructLong(FieldTooLongError, "256", passwordField)
	}
	if len(userInfo.FirstName) > 256 {
		return db.errorConstructLong(FieldTooLongError, "256", firstNameField)
	}
	if len(userInfo.LastName) > 256 {
		return db.errorConstructLong(FieldTooLongError, "256", lastNameField)
	}
	if userInfo.Gender != male && userInfo.Gender != female && userInfo.Gender != another {
		return db.errorConstructValue(WrongValueError, "gender", male, female, another)
	}
	if len(userInfo.Picture) > 512 {
		return db.errorConstructLong(FieldTooLongError, "512", pictureField)
	}
	if len(userInfo.BackgroundPicture) > 512 {
		return db.errorConstructLong(FieldTooLongError, "512", bgPictureField)
	}
	result, err := db.Connection.Exec("SELECT user_id FROM user_private WHERE user_id = $1", userInfo.UserId)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 1 {
		return NicknameUniqueError
	}
	result, err = db.Connection.Exec("SELECT email FROM user_private WHERE email = $1", userInfo.Email)
	if err != nil {
		return err
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil || rowsAffected == 1 {
		return EmailUniqueError
	}
	return nil
}

func (db *DataBase) validateTag(tag *entities.Tag) error {
	if len(tag.Name) > 50 {
		return db.errorConstructLong(FieldTooLongError, "50", "tag_name")
	}
	if len(tag.Description) > 256 {
		return db.errorConstructLong(FieldTooLongError, "256", "description")
	}
	result, err := db.Connection.Exec("SELECT tag_name FROM tags WHERE tag_name = $1", tag.Name)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 1 {
		return TagUniqueError
	}
	return nil
}

func (db *DataBase) validateTask(task *entities.Task) error {
	if len(task.Name) > 50 {
		return db.errorConstructLong(FieldTooLongError, "50", "task_name")
	}
	if len(task.Description) > 256 {
		return db.errorConstructLong(FieldTooLongError, "256", "description")
	}
	if len(task.Picture) > 512 {
		return db.errorConstructLong(FieldTooLongError, "512", "picture")
	}
	if len(task.BackgroundPicture) > 512 {
		return db.errorConstructLong(FieldTooLongError, "512", "backgroundPicture")
	}
	if len(task.RecommendedTime) > 256 {
		return db.errorConstructLong(FieldTooLongError, "256", "recommendedTime")
	}
	return nil
}
