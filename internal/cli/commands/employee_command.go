package cli_command

import (
	internal_cli_usecases "github.com/milfan/go-boilerplate/internal/cli/usecases"
	"github.com/urfave/cli/v2"
)

// Command to check employee
func CheckEmployee(c *cli.Context, usecase internal_cli_usecases.IEmployeeUsecase) error {

	err := usecase.FindEmployee(c.Context)
	if err != nil {
		return err
	}

	return nil
}
