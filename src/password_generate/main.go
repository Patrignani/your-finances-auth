package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	valid := verifyTwoFactorCode("FQ5J6TV4FV4BOEPU7TNRXVAATFTRZQNJ", "693105")
	if valid {
		fmt.Println("Código 2FA válido!")
	} else {
		fmt.Println("Código 2FA inválido!")
	}
}

func generator2fa() {
	v, _ := generateTwoFactorSecret("65b952697a0b1286f03b4a22")
	saveQRCode(v, "Patrignani", "E:\\git\\your-finances-auth\\src\\password_generate\\qrcode.png")

	println(v)
}

func verifyTwoFactorCode(secret, userInputCode string) bool {
	// Decodificar o segredo do usuário
	key, err := totp.Generate(totp.GenerateOpts{
		Secret:      []byte(secret),
		Issuer:      "your-finances-auth", // Nome do emissor (sua aplicação)
		AccountName: "Patrignani",
	})
	if err != nil {
		return false
	}

	// Verificar o código 2FA
	return totp.Validate(userInputCode, key.Secret())
}

func generatePassword() {
	seed := "213ca3d614b04698a94068afa45672cc"
	password := "admin" + seed

	value, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err == nil {
		println(string(value))
	}
}

func generateTwoFactorSecret(accountName string) (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "your-finances-auth", // Nome do emissor (sua aplicação)
		AccountName: accountName,
	})
	if err != nil {
		return "", err
	}

	return key.Secret(), nil
}

func saveQRCode(secret, accountName, filePath string) error {
	// Criar um objeto TOTP com o segredo do usuário
	key, err := totp.Generate(totp.GenerateOpts{
		Secret:      []byte(secret),
		Issuer:      "your-finances-auth",
		AccountName: accountName,
	})
	if err != nil {
		return err
	}

	// Criar a URL para o código 2FA
	url := key.URL()

	// Criar o QR code
	qrCode, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		return err
	}

	// Criar um arquivo para salvar o QR code
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Salvar o QR code no arquivo
	err = png.Encode(file, qrCode.Image(256))
	if err != nil {
		return err
	}

	fmt.Printf("QR code salvo em %s\n", filePath)

	return nil
}
