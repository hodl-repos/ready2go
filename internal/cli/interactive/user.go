package interactive

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hodl-repos/ready2go/internal/cli/config"
	"github.com/hodl-repos/ready2go/internal/cli/helper"
	"github.com/hodl-repos/ready2go/r2o"
	"github.com/urfave/cli/v2"
)

var userMenus = map[string]func(c *cli.Context) error{
	"get all":      runUserGetAll,
	"get by id":    runUserGetById,
	"user-sign-in": runUserSignIn,
}

func userMenu(c *cli.Context) error {
	return menuSelectAndRun("select function to run:", c, userMenus)
}

func runUserGetAll(c *cli.Context) error {
	pagination, err := requestPaginationData()

	if err != nil {
		msg := fmt.Sprintf("error while requesting pagintion-data from user: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.User.GetUsers(c.Context, pagination)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}

func runUserGetById(c *cli.Context) error {
	couponId, err := getNumberPrompt("enter user-id")
	if err != nil {
		helper.CheckTerminalInterrupt(err)
		return err
	}

	if err != nil {
		msg := fmt.Sprintf("error while requesting pagintion-data from user: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	data, err := clt.User.GetUser(c.Context, couponId)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}

func runUserSignIn(c *cli.Context) error {
	id, err := getNumberPrompt("enter user-id")
	if err != nil {
		helper.CheckTerminalInterrupt(err)
		return err
	}

	accToken := config.AccountToken()
	clt := r2o.NewClient(&accToken, nil)

	seconds := 60
	crmUserId := fmt.Sprint(*id)

	data, err := clt.User.GetSigningTokenWithGet(c.Context, id, &seconds, &crmUserId)

	if err != nil {
		msg := fmt.Sprintf("error while loading from r2o API: %v", err.Error())
		fmt.Println(msg)
		return errors.New(msg)
	}

	formattedBytes, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("RESULT:\n\n%v\n\n", string(formattedBytes))

	return nil
}
