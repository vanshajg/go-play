package offline

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vanshajg/go-play/container"
	"github.com/vanshajg/go-play/models/dto"
)

func fetchComments(container container.Container) {
	logger := container.GetLogger().GetZapLogger()
	baseURL := "https://hacker-news.firebaseio.com/v0/item/%d.json"
	postURL := fmt.Sprintf(baseURL, 29782096)
	resp, err := http.Get(postURL)
	if err != nil {
		logger.Errorf("error getting %s error:  %s", postURL, err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("error reading response body %s", err.Error())
	}
	postdto := dto.NewCommentDto()
	if err := json.Unmarshal(body, &postdto); err != nil {
		logger.Errorf("failed to unmarshal data %s", err.Error())
	}
	commentDto := dto.NewCommentDto()
	for i := 0; i < len(postdto.Kids); i++ {
		commentURL := fmt.Sprintf(baseURL, postdto.Kids[i])
		logger.Debugf("fetching data for comment url: %s", commentURL)
		resp, err := http.Get(commentURL)
		if err != nil {
			logger.Errorf("error getting %s error:  %s", commentURL, err.Error())
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Errorf("error reading response body %s", err.Error())
		}
		if err := json.Unmarshal(body, &commentDto); err != nil {
			logger.Errorf("failed to unmarshal data %s", err.Error())
		}
		logger.Infof(commentDto.Text)
	}
}
