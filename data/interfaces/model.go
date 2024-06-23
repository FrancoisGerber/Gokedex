package interfaces

type IModel[T any] interface {

	//Get an array of objects
	GetAll() ([]T, error)

	//Get a single object
	Get(int64) error

	//Create a new object in the DB
	Create() error

	//Update the object in the DB
	Update(int64) (int64, error)

	//Delete the object from the DB
	Delete(int64) (bool, error)
}
