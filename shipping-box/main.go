package shipping_box

type Product struct {
	Name string
	Len  int //ml
	Wid  int
	Hei  int
}

type Box struct {
	Len int //ml
	Wid int
	Hei int
}

//use and retrieve the best (smallest) box for given set of products

func getBestBox(availableBoxes []Box, products []Product) Box {

	//TODO: Complete this challenge!
	return Box{}
}
