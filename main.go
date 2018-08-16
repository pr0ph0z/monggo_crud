package main

import (
	"fmt"
	"time"

	"barengin/config"
	"barengin/src/modules/profile/model"
	"barengin/src/modules/profile/repository"
)

func main() {
	fmt.Println("Run Successfully ...")

	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}

	profileRepository := repository.NewProfleRepositoryMongo(db, "profile")

	// saveProfile(profileRepository)
	// updateProfile(profileRepository)
	// deleteProfile(profileRepository)
	// getProfile("U1", profileRepository)
	getProfiles(profileRepository)
}

func saveProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U2"
	p.FirstName = "Spwhehe"
	p.LastName = "Gada"
	p.Email = "gada@spw.com"
	p.Password = "gada"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Save(&p)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Profile saved..")
}

func updateProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U1"
	p.FirstName = "Spw"
	p.LastName = "Gada"
	p.Email = "bener@gmail.com"
	p.Password = "gad1a1222"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Update("U1",&p)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Profile Updated..")
}

func deleteProfile(profileRepository repository.ProfileRepository) {
	err := profileRepository.Delete("U1")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Profile Deleted..")
}

func getProfile(id string, profileRepository repository.ProfileRepository) {
	profile, err := profileRepository.FindByID(id)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ID: "+profile.ID)
	fmt.Println("First Name: "+profile.FirstName)
	fmt.Println("Last Name: "+profile.LastName)
	fmt.Println("Email: "+profile.Email)
}

func getProfiles(profileRepository repository.ProfileRepository) {
	profiles, err := profileRepository.FindAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, profile := range profiles {
		fmt.Println("---------------------------------")
		fmt.Println("ID "+profile.ID)
		fmt.Println("First Name: "+profile.FirstName)
		fmt.Println("Last Name: "+profile.LastName)
		fmt.Println("Email: "+profile.Email)
	}
}