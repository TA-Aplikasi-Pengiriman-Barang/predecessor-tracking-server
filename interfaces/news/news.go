package news

import (
	"time"
	"tracking-server/application"
	"tracking-server/shared"
	"tracking-server/shared/dto"
)

type (
	ViewService interface {
		CreateNews(data dto.CreateNewsDto) (dto.CreateNewsResponse, error)
		GetAllNews() (dto.GetAllNewsResponse, error)
		GetNewsDetail(id string) (dto.News, error)
		DeleteNews(id string) error
		EditNews(data dto.EditNewsDto, id string) (dto.News, error)
	}
	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)

func (v *viewService) CreateNews(data dto.CreateNewsDto) (dto.CreateNewsResponse, error) {
	var (
		news     *dto.News
		response dto.CreateNewsResponse
	)

	news = &dto.News{
		Title:     data.Title,
		Detail:    data.Detail,
		CreatedAt: time.Now(),
	}

	err := v.application.NewsService.Create(news)
	if err != nil {
		v.shared.Logger.Errorf("error when inserting news to database, err: %s", err.Error())
		return response, err
	}

	response = news.ToCreateNewsResponse()

	return response, nil
}

func (v *viewService) GetAllNews() (dto.GetAllNewsResponse, error) {
	var (
		news     = &dto.NewsSlice{}
		response dto.GetAllNewsResponse
	)

	err := v.application.NewsService.GetAll(news)
	if err != nil {
		v.shared.Logger.Errorf("error when getting all news, err: %s", err.Error())
		return response, err
	}

	response = news.ToGetAllNewsResponse()

	return response, nil
}

func (v *viewService) GetNewsDetail(id string) (dto.News, error) {
	var (
		news = &dto.News{}
	)

	err := v.application.NewsService.GetById(id, news)
	if err != nil {
		v.shared.Logger.Errorf("error when getting news by id, err: %s", err.Error())
		return *news, err
	}

	return *news, nil
}

func (v *viewService) DeleteNews(id string) error {
	err := v.application.NewsService.Delete(id)
	if err != nil {
		v.shared.Logger.Errorf("error when deleting news, err: %s", err.Error())
		return err
	}
	return nil
}

func (v *viewService) EditNews(data dto.EditNewsDto, id string) (dto.News, error) {
	var (
		news = &dto.News{}
	)

	err := v.application.NewsService.GetById(id, news)
	if err != nil {
		v.shared.Logger.Errorf("error when finding news by id, err: %s", err.Error())
		return *news, err
	}

	news.FillNewsEdit(data)

	err = v.application.NewsService.Save(news)
	if err != nil {
		v.shared.Logger.Errorf("error when updating news, err: %s", err.Error())
		return *news, err
	}

	return *news, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}
