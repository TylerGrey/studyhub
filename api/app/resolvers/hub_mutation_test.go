package resolvers_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/TylerGrey/studyhub/api/app/resolvers"
	"github.com/TylerGrey/studyhub/api/app/resolvers/args"
	mysqlLib "github.com/TylerGrey/studyhub/internal/mysql"
	"github.com/TylerGrey/studyhub/internal/mysql/repo"
	"github.com/joho/godotenv"
)

var resolver *resolvers.Resolver

func init() {
	err := godotenv.Load("../../../configs/.env.local")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	mysqlMaster, mysqlReplica, err := mysqlLib.IntializeDatabase(os.Getenv("MYSQL_DB_NAME"))
	if err != nil {
		fmt.Println("db initialize error", err.Error())
		return
	}
	userRepo := repo.NewUserRepository(mysqlMaster, mysqlReplica)
	hubRepo := repo.NewHubRepository(mysqlMaster, mysqlReplica)
	hubIncorrectInfoRepo := repo.NewHubIncorrectInfoRepository(mysqlMaster, mysqlReplica)

	resolver = &resolvers.Resolver{
		UserRepo:             userRepo,
		HubRepo:              hubRepo,
		HubIncorrectInfoRepo: hubIncorrectInfoRepo,
	}
}

func TestCreateHub(t *testing.T) {
	_, err := resolver.CreateHub(args.CreateHubInput{
		Input: args.CreateHubArgs{
			Name: "TEST HUB",
			Type: "CAFE",
			CoverImage: args.ImageInput{
				File: "",
			},
			Address: args.AddressInput{
				Address: "서울 관악구 양녕로6길 22-6",
				Lat:     23.123456,
				Lng:     22.123456,
			},
		},
	})
	if err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateHub(t *testing.T) {
	updatedName := "Hub test name"
	_, err := resolver.UpdateHub(args.UpdateHubInput{
		Input: args.UpdateHubArgs{
			ID:   "25",
			Name: &updatedName,
		},
	})
	if err != nil {
		t.Error(err.Error())
	}
}

func TestDeleteHub(t *testing.T) {
	_, err := resolver.DeleteHub(struct{ ID string }{
		ID: "31",
	})
	if err != nil {
		t.Error(err.Error())
	}
}

func TestAddHubIncorrectInfo(t *testing.T) {
	_, err := resolver.AddHubIncorrectInfo(args.AddHubIncorrectInfoInput{
		Input: args.AddHubIncorrectInfoArgs{
			HubID:   "25",
			Message: "장소가 달라요.",
		},
	})
	if err != nil {
		t.Error(err.Error())
	}
}
