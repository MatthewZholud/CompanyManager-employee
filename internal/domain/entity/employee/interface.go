package employee

type Reader interface {
	Get(id ID) (*Employee, error)
	Search(query string) ([]*Employee, error)
	List() ([]*Employee, error)
}


type Writer interface {
	Create(e *Employee) (ID, error)
	Update(e *Employee) error
	Delete(id ID) error
}

//repository interface
type Repository interface {
	Reader
	Writer
}

type employeeRepository struct {
	repo Repository
}

func NewEmployeeRepository() *employeeRepository {
	return &employeeRepository{}
}
