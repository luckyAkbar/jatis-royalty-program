package console

import (
	"context"
	"fmt"
	"time"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/jatis-royalty-program/internal/db"
	"github.com/luckyAkbar/jatis-royalty-program/internal/model"
	"github.com/luckyAkbar/jatis-royalty-program/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var createUserCMD = &cobra.Command{
	Use: "create-user",
	Short: "create a new user",
	Run: createUser,
}

func init() {
	RootCmd.AddCommand(createUserCMD)
}

func createUser(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		logrus.Fatal("args at pos 1 is required as user password")
	}

	password := args[0]
	if len(password) < 4 {
		logrus.Fatal("password must be 4 or more characters")
	}

	user := &model.User{
		ID: utils.GenerateID(),
		Password: password,
		CreatedAt: time.Now(),
	}

	if err := user.Encrypt(); err != nil {
		logrus.Fatal("failed to encrypt user password: ", err)
	}

	db.InitializePostgresConn()
	userRepo := repository.NewUserRepository(db.PostgresDB)

	if err := userRepo.Create(context.TODO(), user); err != nil {
		logrus.Fatal("failed to create user: ", err)
	}

	fmt.Println("user ID: ", user.ID)
}