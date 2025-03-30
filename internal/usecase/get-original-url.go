package usecase

func (uc *Usecase) GetOriginalURL(shortedUrl string) (string, error) {
	longURL, err := uc.repo.GetLongURL(shortedUrl)

	if err != nil {
		return "", err
	}

	return longURL, nil
}
