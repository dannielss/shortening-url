package usecase

import "fmt"

func (uc *usecase) GetOriginalURL(shortedUrl string) (string, error) {
	longURL, err := uc.repo.GetLongURL(shortedUrl)

	fmt.Printf("%v", longURL)
	if err != nil {
		return "", err
	}
	return longURL, nil
}
