package struct_helper

// "time"
// "unc/services/unc-service/domain/repository"
// "github.com/gofiber/fiber/v2"

type HttpResponse struct {
	status  int
	message string
	error   error
}
