package internal

import (
	"fmt"

	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/pkg/utils"
	"github.com/spf13/cobra"
)

var utilityVersionCmd = &cobra.Command{
	Use:     "mkpass",
	Aliases: []string{"password", "pass"},
	Short:   "Generate random passwords",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Generate random passwords.",
			"If count is greater than 1, then each password will be new-line separated.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl utility mkpass",
	}),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		count, _ := cmd.Flags().GetInt("count")
		length, _ := cmd.Flags().GetInt("length")
		noSymbols, _ := cmd.Flags().GetBool("no-symbols")
		noNumbers, _ := cmd.Flags().GetBool("no-numbers")
		noLowercase, _ := cmd.Flags().GetBool("no-lowercase")
		noUppercase, _ := cmd.Flags().GetBool("no-uppercase")
		if noSymbols && noNumbers && noLowercase && noUppercase {
			return fmt.Errorf("alphabet is empty")
		}
		if count < 1 {
			return fmt.Errorf("count must be greater than 0")
		}
		if length < 1 {
			return fmt.Errorf("length must be greater than 0")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		count, _ := cmd.Flags().GetInt("count")
		length, _ := cmd.Flags().GetInt("length")
		noSymbols, _ := cmd.Flags().GetBool("no-symbols")
		noNumbers, _ := cmd.Flags().GetBool("no-numbers")
		noLowercase, _ := cmd.Flags().GetBool("no-lowercase")
		noUppercase, _ := cmd.Flags().GetBool("no-uppercase")
		alphabet := ""
		if !noSymbols {
			alphabet += utils.AlphabetSymbols
		}
		if !noNumbers {
			alphabet += utils.AlphabetNumbers
		}
		if !noLowercase {
			alphabet += utils.AlphabetLowerCase
		}
		if !noUppercase {
			alphabet += utils.AlphabetUpperCase
		}
		for i := 0; i < count; i++ {
			fmt.Println(utils.GeneratePassword(alphabet, length))
		}
	},
}

func init() {
	utilityCmd.AddCommand(utilityVersionCmd)
	utilityVersionCmd.Flags().SortFlags = true
	utilityVersionCmd.Flags().IntP("count", "c", 1, "number of passwords to generate")
	utilityVersionCmd.Flags().IntP("length", "l", 32, "length of password")
	utilityVersionCmd.Flags().BoolP("no-symbols", "S", false, "do not include symbols in password")
	utilityVersionCmd.Flags().BoolP("no-numbers", "N", false, "do not include numbers in password")
	utilityVersionCmd.Flags().BoolP("no-lowercase", "L", false, "do not include lowercase chars in password")
	utilityVersionCmd.Flags().BoolP("no-uppercase", "U", false, "do not include uppercase chars in password")
}
