package paste

import (
	"net/http"
	"strconv"

	"github.com/mrumyantsev/pastebin-app/internal/reqerrors"
)

var (
	ErrContentOfZeroSize                  = reqerrors.New(http.StatusBadRequest, "unable to store the content of zero size")
	ErrContentLargerThanMaxAnonimousLimit = reqerrors.New(http.StatusBadRequest, "unable anonimously to store the content larger than "+strconv.Itoa(ContentLimitMaxAnonimous)+" bytes")
	ErrContentLargerThanTotalMaxLimit     = reqerrors.New(http.StatusBadRequest, "unable to store the content larger than "+strconv.Itoa(ContentLimitMaxTotal)+" bytes")
	ErrHeaderExpiresAtIsEmpty             = reqerrors.New(http.StatusBadRequest, "header 'Expires-At' is empty")

	ErrPasteNotFound = reqerrors.New(http.StatusNotFound, "paste not found")
)
