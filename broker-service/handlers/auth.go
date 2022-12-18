package handlers

import (
	"broker/repositories"
	"broker/utils"
	"fmt"
	"os"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

const defaultSecretKey = "sercrethatmaycontainch@r$32chars"

func getSecretKey() string {
	secret := os.Getenv(utils.AppName + "_SECRET")
	if secret == "" {
		return defaultSecretKey
	}

	return secret
}

// UserClaims represents the user token claims.
type UserClaims struct {
	UserID string `json:"user_id"`
}

// Validate implements the custom struct claims validator,
// this is totally optionally and maybe unnecessary but good to know how.
func (u *UserClaims) Validate() error {
	if u.UserID == "" {
		return fmt.Errorf("%w: %s", jwt.ErrMissingKey, "user_id")
	}

	return nil
}

// Verify allows only authorized clients.
func Verify() iris.Handler {
	secret := getSecretKey()
	verifier := jwt.NewVerifier(jwt.HS256, []byte(secret), jwt.Expected{Issuer: utils.AppName})
	verifier.WithDefaultBlocklist()
	verifier.Extractors = []jwt.TokenExtractor{jwt.FromHeader} // extract token only from Authorization: Bearer $token
	return verifier.Verify(func() interface{} {
		return new(UserClaims)
	})
}

func SignIn(repo repositories.UserRepo) iris.Handler {
	secret := getSecretKey()
	signer := jwt.NewSigner(jwt.HS256, []byte(secret), 15*time.Minute)

	return func(ctx iris.Context) {
		var request SingUpRequest
		err := ctx.ReadJSON(&request)
		if err != nil {
			//todo create validation
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}

		user, ok := repo.GetByUsernameAndPassword(request.Mail, request.Password)
		if !ok {
			ctx.StopWithText(iris.StatusBadRequest, "wrong username or password")
			return
		}

		claims := UserClaims{
			UserID: user.ID,
		}

		// Optionally, generate a JWT ID.
		jti, err := utils.GenerateUUID()
		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, err)
			return
		}

		token, err := signer.Sign(claims, jwt.Claims{
			ID:     jti,
			Issuer: utils.AppName,
		})
		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, err)
			return
		}

		response := SingInResponse{
			Mail:  user.Username,
			Token: string(token),
		}

		ctx.JSON(response)
	}
}

func SignUp(repo repositories.UserRepo) iris.Handler {
	secret := getSecretKey()
	signer := jwt.NewSigner(jwt.HS256, []byte(secret), 15*time.Minute)

	return func(ctx iris.Context) {

		var request SingUpRequest
		err := ctx.ReadJSON(&request)
		if err != nil {
			//todo create validation
			//if errs, ok := err.(validator.ValidationErrors); ok {
			//	// Wrap the errors with JSON format, the underline library returns the errors as interface.
			//	validationErrors := wrapValidationErrors(errs)
			//
			//	// Fire an application/json+problem response and stop the handlers chain.
			//	ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			//		Title("Validation error").
			//		Detail("One or more fields failed to be validated").
			//		Type("/user/validation-errors").
			//		Key("errors", validationErrors))
			//
			//	return
			//}
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}

		user, err := repo.Create(request.Mail, request.Password)
		if err != nil {
			ctx.StopWithText(iris.StatusBadRequest, "wrong email or password")
			return
		}

		claims := UserClaims{
			UserID: user.ID,
		}

		jti, err := utils.GenerateUUID()
		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, err)
			return
		}

		token, err := signer.Sign(claims, jwt.Claims{
			ID:     jti,
			Issuer: utils.AppName,
		})

		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, err)
			return
		}
		response := SingInResponse{
			Mail:  user.Username,
			Token: string(token),
		}

		ctx.JSON(response)
	}
}

// SignOut invalidates a user from server-side using the jwt Blocklist.
func SignOut(ctx iris.Context) {
	err := ctx.Logout()

	if err != nil {
		ctx.WriteString(err.Error())
	} else {
		ctx.Writef("token invalidated, a new token is required to access the protected API")
	}
}

// GetClaims returns the current authorized client claims.
func GetClaims(ctx iris.Context) *UserClaims {
	claims := jwt.Get(ctx).(*UserClaims)
	return claims
}

// GetUserID returns the current authorized client's user id extracted from claims.
func GetUserID(ctx iris.Context) string {
	return GetClaims(ctx).UserID
}

type SingUpRequest struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type SingInResponse struct {
	Mail  string `json:"mail"`
	Token string `json:"token"`
}
