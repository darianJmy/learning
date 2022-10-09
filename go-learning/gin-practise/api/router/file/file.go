package file

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/darianJmy/learning/go-learning/gin-practise/pkg/httputils"
)

func UpLoad(c *gin.Context) {
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		httputils.Failed(c, "Bad request")
		return
	}

	filename := header.Filename

	out, err := os.Create(fmt.Sprintf("/etc/gin-practise/%s", filename))
	if err != nil {
		httputils.Failed(c, "Failed to create to file")
		return
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		httputils.Failed(c, "Failed to write to file")
		return
	}
	httputils.Success(c, "success to write to file")
}
