package auth
//
//import (
//	"github.com/gin-gonic/gin"
//	"log"
//)
//
////
//type validator struct {
//	repository *repository
//}
//
//func NewValidator(repository *repository) *validator {
//	//validate , err := zValidator.GetValidate(locale)
//	//if err != nil {
//	//	log.Fatal(err.Error())
//	//}
//	return &validator{
//		repository: repository,
//	}
//}
//
//// Validate validate SingUpParam
//func (v *validator) validateSignUp (c *gin.Context, json *SignUpParam) Errs {
//
//	if err := c.ShouldBindJSON(&json); err != nil {
//		return err.Error()
//	}
//	var err error
//
//	// register validation for 'SingUpParam'
//	v.RegisterStructValidation(signUpParamStructLevelValidation, SignUpParam{})
//
//	err = v.validator.RegisterValidation("is_exists", func(fl goValidator.FieldLevel) bool {
//		user := v.repository.firstByUsername(fl.Field().String())
//		return user == nil
//	})
//	if err != nil {
//		return ErrInternalServerError
//	}
//
//	return zValidator.Validate(sup)
//}
////
//// signUpParamStructLevelValidation 自定义SignUpParam结构体成员校验
////
//// password_confirmation must be equal to Password
//// password_confirmation must be equal to password
//func signUpParamStructLevelValidation(sl goValidator.StructLevel)  {
//	sup := sl.Current().Interface().(SignUpParam)
//	if sup.Password != sup.PasswordConfirmation {
//		sl.ReportError(
//			sup.PasswordConfirmation,
//			"password_confirmation",
//			"PasswordConfirmation",
//			"eqfield",
//			"password",
//		)
//	}
//}
