package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/motojouya/ddd_go/domain/basic/core"
)

// FIXME templateなのでないが、操作者としての人格(Userとか)も引数にとれるようにすべき
func Hand[C any, I any, O any](createControl func() (C, error), handleControl func(C, I) (O, error)) echo.HandlerFunc {
	return func(c echo.Context) error {

		var request I
		if err := c.Bind(&request); err != nil {
			return err
		}

		control, err := createControl()
		if err != nil {
			return err
		}

		response, err := handleControl(control, request)
		if err != nil {
			return err
		}

		if closable, ok := any(control).(core.Closable); ok {
			err := closable.Close()
			if err != nil {
				return err
			}
		}

		return c.JSON(200, response)
	}
}
