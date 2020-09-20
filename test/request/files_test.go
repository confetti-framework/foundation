package request

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilesNotFoundInRequest(t *testing.T) {
	request := requestByFiles()

	files, err := request.FilesE("photos")
	assert.EqualError(t, err, "files not found by key: photos")
	assert.Len(t, files, 0)
}

func requestByFiles() inter.Request {
	options := http.Options{}
	return http.NewRequest(options)
}
