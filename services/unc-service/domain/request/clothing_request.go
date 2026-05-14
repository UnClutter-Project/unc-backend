package request

import "mime/multipart"

type CreateClothingRequest struct {
	MainColor1  string                `form:"main_color_1"  validate:"required"`
	MainColor2  string                `form:"main_color_2"  `
	AccentColor string                `form:"accent_color"  `
	Category    string                `form:"category" validate:"required,max=64"`
	Type        string                `form:"type"     validate:"required,max=64"`
	Brand       string                `form:"brand"       validate:"max=64"`
	Style       string                `form:"style"       validate:"max=64"`
	Image       *multipart.FileHeader `form:"image"       validate:"required"`
}
