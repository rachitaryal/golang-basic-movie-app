package main

type Movie struct {
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `json: "title`
	Director *Director `json: "director"`
}

type Director struct {
	ID string `json: "id"`
	FirstName string `json: "firstName"`
	LastName string `json: "lastName"`
}

func main(){
	// main function

}