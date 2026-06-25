package response

//Entrega as informações
type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
}
