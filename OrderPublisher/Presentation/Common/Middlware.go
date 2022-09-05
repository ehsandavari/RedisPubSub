package Common

import (
	"github.com/kataras/iris/v12"
)

func ErrorHandlerMiddleware(ctx iris.Context) {
	defer func() {
		recorder := ctx.Recorder()
		var err error
		if value := recover(); value != nil {
			err = value.(error)
		} else {
			err = ctx.GetErr()
		}

		if err != nil {
			recorder.ResetBody()
			recorder.ResetHeaders()
			ctx.StopWithJSON(iris.StatusInternalServerError, errorResponse(err))
		}
	}()

	ctx.Next()
}
