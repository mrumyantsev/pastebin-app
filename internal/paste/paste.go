package paste

import (
	"context"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/base64conv-go"
	"github.com/mrumyantsev/errlib-go"
	"github.com/mrumyantsev/pastebin-app/internal/server"
)

type Servicer interface {
	CreatePaste(ctx context.Context, outerPaste OuterPaste) (string, error)
	GetAllPastes(ctx context.Context) ([]OuterPaste, error)
	GetPasteById(ctx context.Context, base64Id string) (OuterPaste, error)
	UpdatePasteById(ctx context.Context, base64Id string, outerPaste OuterPaste) error
	DeletePasteById(ctx context.Context, base64Id string) error
	IsPasteContentExistsById(ctx context.Context, base64Id string) (bool, error)
}

type DatabaseAdapterer interface {
	CreatePaste(ctx context.Context, paste Paste) error
	GetAllPastes(ctx context.Context) ([]Paste, error)
	GetPasteById(ctx context.Context, id int64) (Paste, error)
	UpdatePasteById(ctx context.Context, id int64, paste Paste) error
	DeletePasteById(ctx context.Context, id int64) error
}

type StorageAdapterer interface {
	CreatePasteContentById(ctx context.Context, id int64, content []byte) error
	CreateOrUpdatePasteContentById(ctx context.Context, id int64, content []byte) error
	GetPasteContentById(ctx context.Context, id int64) ([]byte, error)
	DeletePasteContentById(ctx context.Context, id int64) error
	IsPasteContentExistsById(ctx context.Context, id int64) (bool, error)
}

type HttpAdapterer interface {
	GetGeneratedPasteId(ctx context.Context) (string, error)
}

type Paste struct {
	Id        int64
	UserId    *int64
	Content   []byte
	CreatedAt time.Time
	ExpiresAt time.Time
}

func (p Paste) ToOuterPaste() (OuterPaste, error) {
	var userId int64 = -1

	if p.UserId != nil {
		userId = *p.UserId
	}

	return OuterPaste{
		Base64Id:  base64conv.ItobRawUrl(p.Id),
		UserId:    userId,
		Content:   p.Content,
		CreatedAt: p.CreatedAt.Format(time.DateTime),
		ExpiresAt: p.ExpiresAt.Format(time.DateTime),
	}, nil
}

type OuterPaste struct {
	Base64Id  string `json:"base64Id"`
	UserId    int64  `json:"userId"`
	Content   []byte `json:"content"`
	CreatedAt string `json:"createdAt"`
	ExpiresAt string `json:"expiresAt"`
}

func NewOuterPaste(c *gin.Context) (OuterPaste, error) {
	userId := server.GetUserId(c)

	if c.Request.ContentLength == 0 {
		return OuterPaste{}, ErrContentOfZeroSize
	}
	if userId < 0 && c.Request.ContentLength > ContentLimitMaxAnonimous {
		return OuterPaste{}, ErrContentLargerThanMaxAnonimousLimit
	}
	if c.Request.ContentLength > ContentLimitMaxTotal {
		return OuterPaste{}, ErrContentLargerThanTotalMaxLimit
	}

	expiresAt := c.Request.Header.Get("Expires-At")

	if expiresAt == "" {
		return OuterPaste{}, ErrHeaderExpiresAtIsEmpty
	}

	bodyData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return OuterPaste{}, errlib.Wrap(err, "could not read paste content from body")
	}

	return OuterPaste{
		UserId:    userId,
		Content:   bodyData,
		ExpiresAt: expiresAt,
	}, nil
}

func (p OuterPaste) ToPaste() (Paste, error) {
	id, err := base64conv.BtoiRawUrl(p.Base64Id)
	if err != nil {
		return Paste{}, errlib.Wrap(err, "could not convert base64 id to integer")
	}

	var ptrUserId *int64

	if p.UserId >= 0 {
		ptrUserId = new(int64)
		*ptrUserId = p.UserId
	}

	timeExpiresAt, err := time.Parse(time.DateTime, p.ExpiresAt)
	if err != nil {
		return Paste{}, errlib.Wrap(err, "could not parse expires-at time from string")
	}

	return Paste{
		Id:        id,
		UserId:    ptrUserId,
		Content:   p.Content,
		ExpiresAt: timeExpiresAt,
	}, nil
}
