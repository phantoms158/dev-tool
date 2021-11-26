package models

import (
	"devtools-backend/databases"
	"devtools-backend/forms"
	"devtools-backend/helpers"
	"time"
)

type User struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Role       uint      `json:"role"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Signup handles registering a user
func Signup(data forms.RegisterCommand) error {
	user := User{Email: data.Email, Username: data.Username, Password: helpers.GeneratePasswordHash([]byte(data.Password)), CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err := databases.DB.Create(&user).Error
	return err

}

func GetUserByEmail(email string) (user User) {
    databases.DB.Where("email = ?", email).First(&user)
    return user
}

// type AuthCustomClaims struct {
// 	Name string `json:"name"`
// 	User bool   `json:"user"`
// 	jwt.StandardClaims
// }

// type jwtServices struct {
// 	secretKey string
// 	issure    string
// }

// func GetSecretKey() string {
// 	secret := os.Getenv("SECRET")
// 	if secret == "" {
// 		secret = "secret"
// 	}
// 	return secret
// }

// func GenerateToken(email string, isUser bool) string {
// 	claims := &AuthCustomClaims{
// 		email,
// 		isUser,
// 		jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
// 			Issuer:    "Trongpq",
// 			IssuedAt:  time.Now().Unix(),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	//encoded string
// 	t, err := token.SignedString([]byte(GetSecretKey()))
// 	if err != nil {
// 		panic(err)
// 	}
// 	return t
// }

// func ValidateToken(encodedToken string) (*jwt.Token, error) {
// 	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
// 		_, isvalid := token.Method.(*jwt.SigningMethodHMAC)
// 		if !isvalid {
// 			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

// 		}
// 		fmt.Println(isvalid)
// 		return []byte(GetSecretKey()), nil
// 	})

// }
