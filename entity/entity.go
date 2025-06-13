package entity

type PreCardapio struct {
	ID        int     `json:"id"`
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	CHO       float64 `json:"cho"`
	LIP       float64 `json:"lip"`
	PTN       float64 `json:"ptn"`
	KCAL      float64 `json:"kcal"`
	PesoPrato int     `json:"peso_prato"`
}

type QuantidadeAlmoco struct {
	ID         int    `json:"id"`
	CardapioID int    `json:"cardapio_id"`
	Local      string `json:"local"` // ALDEIA, FLOR DO CAMARA, COOPERATIVA
	Quantidade int    `json:"quantidade"`
}

type Cardapio struct {
	ID        int     `json:"id"`
	Data      string  `json:"data"`
	Descricao string  `json:"descricao"`
	CHO       float64 `json:"cho"`
	LIP       float64 `json:"lip"`
	PTN       float64 `json:"ptn"`
	KCAL      float64 `json:"kcal"`
	PesoPrato int     `json:"peso_prato"`
}
