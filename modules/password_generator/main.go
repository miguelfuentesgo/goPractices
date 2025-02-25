package password_generator

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Caracteres permitidos
const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars  = "0123456789"
	specialChars = "@$!%*?&"
	allChars     = lowerChars + upperChars + numberChars + specialChars // Longitud de la contraseña
)

func Run() {
	fmt.Println("****** PASSWORD GENERATOR ******")
	fmt.Println()

	// Initilization of the calculator reader
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please, provide password length")

	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	length, err := strconv.Atoi(input1)

	if err != nil {
		fmt.Println("Invalid number:", err)
		return
	}

	fmt.Printf("Password length will be %d", length)
	fmt.Println()
	fmt.Println("Generating password...")

	// Conditions of the calculator
	// Length will be as input length of the calculator
	// At least one upper case letter is required
	// At least one lower case letter is required
	// At least one number is required
	// At least one special characters are required

	password := ""

	password = password + getRandomChar(lowerChars)
	password = password + getRandomChar(upperChars)
	password = password + getRandomChar(numberChars)
	password = password + getRandomChar(specialChars)

	//Complete password with rest
	for len(password) <= int(length) {
		password = password + getRandomChar(allChars)

	}

	fmt.Println("Validating password...")
	if !validatePassword(password) {
		fmt.Println("Internal error: Invalid password")
		return
	}

	fmt.Println("Password validated successfully")

	fmt.Println("Your password is: ", password)
}

func getRandomChar(alternatives string) string {
	return string(alternatives[randomInt(0, len(alternatives)-1)])
}

// RandomInt genera un número aleatorio entre min y max (incluyéndolos)
func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano()) // Inicializa la semilla con la hora actual
	return rand.Intn(max-min+1) + min
}

func validatePassword(password string) bool {
	// Regex pattern
	// Check for at least one lowercase letter
	hasLower, _ := regexp.MatchString(`[a-z]`, password)
	// Check for at least one uppercase letter
	hasUpper, _ := regexp.MatchString(`[A-Z]`, password)
	// Check for at least one digit
	hasDigit, _ := regexp.MatchString(`\d`, password)
	// Check for at least one special character
	hasSpecial, _ := regexp.MatchString(`[@$!%*?&]`, password)

	// Password is valid only if all conditions are met
	return hasLower && hasUpper && hasDigit && hasSpecial

}
