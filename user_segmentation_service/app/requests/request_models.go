package requests

type TestCreate struct {
	Name string `validate:"required,max=255,min=3"`
}
