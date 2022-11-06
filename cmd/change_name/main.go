package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Aliath/studies_db_project/database"
	"github.com/Aliath/studies_db_project/models"
	"github.com/tcnksm/go-input"
)

func main() {
	db := database.GetDatabase()
	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	textUserId, _ := ui.Ask("user_id", &input.Options{
		Required: true,
		Loop:     true,
		Default:  "",
	})

	password, _ := ui.Ask("password_hash", &input.Options{
		Required: true,
		Loop:     true,
		Default:  "",
	})

	userId, _ := strconv.Atoi(textUserId)

	user := &models.User{}
	db.Where(&models.User{ID: userId, PasswordHash: password}).First(&user)

	userIdToUpdate := user.ID

	if userIdToUpdate == 0 {
		log.Fatal("User was not found")
	}

	if user.IsAdmin {
		textUserIdToUpdate, _ := ui.Ask("user_id to update", &input.Options{
			Required: true,
			Loop:     true,
			Default:  "",
		})

		userIdToUpdate, _ = strconv.Atoi(textUserIdToUpdate)

		foundUser := &models.User{}
		db.Where(&models.User{ID: userIdToUpdate}).First(foundUser)

		if foundUser.ID == 0 {
			log.Fatal("User was not found")
		}
	}

	newFullName, _ := ui.Ask("new full name", &input.Options{
		Required: true,
		Loop:     true,
		Default:  "",
	})

	db.Model(&models.User{ID: userIdToUpdate}).Update("FullName", newFullName)

	fmt.Println("Name was updated successfully")
}
