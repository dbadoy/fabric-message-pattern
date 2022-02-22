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

	c.CarID = args[0]
	c.Make = args[1]
	c.Model = args[2]
	c.Colour = args[3]
	c.Owner = args[4]

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
