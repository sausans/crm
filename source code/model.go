package main 


type seproduct struct {
        name string `json: "name, omitempty"`
        promotion string `json: "promotion, omitempty"`
}

type user struct {
    Username string `json: "username, omitempty"`
    Email string `json: "email, omitempty"`
    Productsbought string `json: "productsbought, omitempty"`
    Productsname string `json: "productsname, omitempty"`
}

