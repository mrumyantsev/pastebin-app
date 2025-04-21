package paste

import (
	"context"

	"github.com/mrumyantsev/go-base64conv"
)

type Service struct {
	database DatabaseAdapterer
	storage  StorageAdapterer
	http     HttpAdapterer
}

func NewService(
	databaseAdapterer DatabaseAdapterer,
	storageAdapterer StorageAdapterer,
	httpAdapterer HttpAdapterer,
) *Service {
	return &Service{
		database: databaseAdapterer,
		storage:  storageAdapterer,
		http:     httpAdapterer,
	}
}

func (s *Service) CreatePaste(ctx context.Context, outerPaste OuterPaste) (string, error) {
	var err error

	outerPaste.Base64Id, err = s.http.GetGeneratedPasteId(ctx)
	if err != nil {
		return "", err
	}

	paste, err := outerPaste.ToPaste()
	if err != nil {
		return "", err
	}

	var storageContent []byte

	paste.Content, storageContent = splitContents(paste.Content)

	if storageContent != nil {
		if err = s.storage.CreatePasteContentById(ctx, paste.Id, storageContent); err != nil {
			return "", err
		}
	}

	if err = s.database.CreatePaste(ctx, paste); err != nil {
		return "", err
	}

	return outerPaste.Base64Id, nil
}

func (s *Service) GetAllPastes(ctx context.Context) ([]OuterPaste, error) {
	pastes, err := s.database.GetAllPastes(ctx)
	if err != nil {
		return nil, err
	}

	outerPastes := make([]OuterPaste, len(pastes))

	for i := range outerPastes {
		outerPastes[i], err = pastes[i].ToOuterPaste()
		if err != nil {
			return nil, err
		}
	}

	return outerPastes, nil
}

func (s *Service) GetPasteById(ctx context.Context, base64Id string) (OuterPaste, error) {
	id, err := base64conv.BtoiRawUrl(base64Id)
	if err != nil {
		return OuterPaste{}, err
	}

	paste, err := s.database.GetPasteById(ctx, id)
	if err != nil {
		return OuterPaste{}, err
	}

	outerPaste, err := paste.ToOuterPaste()
	if err != nil {
		return OuterPaste{}, err
	}

	return outerPaste, nil
}

func (s *Service) UpdatePasteById(ctx context.Context, base64Id string, outerPaste OuterPaste) error {
	outerPaste.Base64Id = base64Id

	paste, err := outerPaste.ToPaste()
	if err != nil {
		return err
	}

	var storageContent []byte

	paste.Content, storageContent = splitContents(paste.Content)

	if storageContent != nil {
		if err = s.storage.CreateOrUpdatePasteContentById(ctx, paste.Id, storageContent); err != nil {
			return err
		}
	} else {
		if err = s.storage.DeletePasteContentById(ctx, paste.Id); err != nil {
			return err
		}
	}

	if err = s.database.UpdatePasteById(ctx, paste.Id, paste); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeletePasteById(ctx context.Context, base64Id string) error {
	id, err := base64conv.BtoiRawUrl(base64Id)
	if err != nil {
		return err
	}

	if err = s.storage.DeletePasteContentById(ctx, id); err != nil {
		return err
	}

	if err = s.database.DeletePasteById(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *Service) IsPasteContentExistsById(ctx context.Context, base64Id string) (bool, error) {
	id, err := base64conv.BtoiRawUrl(base64Id)
	if err != nil {
		return false, err
	}

	isExists, err := s.storage.IsPasteContentExistsById(ctx, id)
	if err != nil {
		return false, err
	}

	return isExists, nil
}

func splitContents(fullContent []byte) (databaseContent, storageContent []byte) {
	if len(fullContent) > ContentLimitMaxDatabase {
		return fullContent[:ContentLimitMaxDatabase], fullContent[ContentLimitMaxDatabase:]
	}

	return fullContent, nil
}
