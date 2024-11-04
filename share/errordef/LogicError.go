package error

type LogicError struct {
    Msg string
    Code int
}

func (err *LogicError) Error() string {
    return err.Msg
}