package mesg

type NonParamRequest struct {
	CommonRequest
}

type QueryCarRequest struct {
	CommonRequest
	CarID string
}

type CreateCarRequest struct {
	CommonRequest
	CarID  string
	Make   string
	Model  string
	Colour string
	Owner  string
}

type ChangeCarOwnerRequest struct {
	CommonRequest
	CarID string
	To    string
}

func (n NonParamRequest) Set(args []string) error {
	if len(args) != 0 {
		return ErrInvalidArgment
	}

	return nil
}

func (q *QueryCarRequest) Set(args []string) error {
	if len(args) != q.FieldCount() {
		return ErrInvalidArgment
	}

	q.CarID = args[0]

	return nil
}

func (c *CreateCarRequest) Set(args []string) error {
	if len(args) != c.FieldCount() {
		return ErrInvalidArgment
	}

	c.Make = args[0]
	c.Model = args[1]
	c.Colour = args[2]
	c.Owner = args[3]

	return nil
}

func (c *ChangeCarOwnerRequest) Set(args []string) error {
	if len(args) != c.FieldCount() {
		return ErrInvalidArgment
	}

	c.CarID = args[0]
	c.To = args[1]

	return nil
}
