package services

var ItemsService itemsServiceInterface = &itemsService{}

type (
	itemsService          struct{}
	itemsServiceInterface interface {
		GetItem()
		SaveItem()
	}
)

func (i *itemsService) GetItem() {

}

func (i *itemsService) SaveItem() {

}
