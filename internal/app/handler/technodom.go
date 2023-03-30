package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"technodom/internal/entity"
)

// Список адресов
func (h *Handler) getList(w http.ResponseWriter, r *http.Request) {
	l := r.FormValue("limit")
	o := r.FormValue("offset")

	limit, _ := strconv.Atoi(l)
	offset, _ := strconv.Atoi(o)

	urls, err := h.service.GetList(limit, offset)
	if err != nil {
		errMsg := fmt.Errorf("error in GetList service method %w", err).Error()
		h.logger.Error(errMsg, map[string]interface{}{})
		respond(w, r, http.StatusInternalServerError, nil)
		return
	}
	respond(w, r, 200, urls)
}

// Получение адреса по заданному ключу
func (h *Handler) getByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	activeLink, ok := h.cache.Get(id)
	if ok {
		msg := "getting data from cache"
		h.logger.Info(msg, map[string]interface{}{})
		respond(w, r, 200, entity.Url{HistoryLink: id, ActiveLink: activeLink})
		return
	}
	url, err := h.service.GetByID(id)
	if err != nil {
		errMsg := fmt.Errorf("error in GetByID service method %w", err).Error()
		h.logger.Error(errMsg, map[string]interface{}{})
		respond(w, r, http.StatusInternalServerError, nil)
		return
	}
	if h.cache.Len() < 1000 {
		msg := "adding data to cache"
		h.logger.Info(msg, map[string]interface{}{})
		h.cache.Add(id, url.ActiveLink, h.config.Technodom.Ttl)
	}
	respond(w, r, 200, url)

}

// Метод для создания новой записи в базе
func (h *Handler) post(w http.ResponseWriter, r *http.Request) {
	var url entity.Url
	jsonData, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(jsonData, &url)
	if err != nil {
		errMsg := fmt.Errorf("error in parsing from Post request %w", err).Error()
		h.logger.Error(errMsg, map[string]interface{}{})
		respond(w, r, http.StatusBadRequest, nil)
		return
	}
	err = h.service.Post(url)
	if err != nil {
		errMsg := fmt.Errorf("error in Post service method %w", err).Error()
		h.logger.Error(errMsg, map[string]interface{}{})
		respond(w, r, http.StatusInternalServerError, nil)
		return
	}
	respond(w, r, 200, "successfully insert")
}

// Метод для записи текущей актуальной ссылки в список исторических, и изменения активной ссылки
func (h *Handler) patch(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	activeLink := r.FormValue("activeLink")
	//var activeLink string
	//jsonData, _ := ioutil.ReadAll(r.Body)
	//err := json.Unmarshal(jsonData, &activeLink)
	//if err != nil {
	//	errMsg := fmt.Errorf("error in parsing Patch request %w", err).Error()
	//	h.logger.Error(errMsg, map[string]interface{}{})
	//	respond(w, r, http.StatusBadRequest, nil)
	//	return
	//}
	if h.cache.Len() < 1000 {
		msg := "adding data to cache"
		h.logger.Info(msg, map[string]interface{}{})
		h.cache.Add(id, activeLink, h.config.Technodom.Ttl)
	}
	err := h.service.Patch(id, activeLink)
	if err != nil {
		errMsg := fmt.Errorf("error in Patch service method %w", err).Error()
		h.logger.Error(errMsg, map[string]interface{}{})
		respond(w, r, http.StatusInternalServerError, nil)
	}
	respond(w, r, 200, "successfully update")
}

// Метод для удаления записи из базы
func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.service.Delete(id)
	if err != nil {
		errMsg := fmt.Errorf("error in Delete service method %w", err).Error()
		h.logger.Error(errMsg, map[string]interface{}{})
		respond(w, r, http.StatusInternalServerError, nil)
		return
	}
	respond(w, r, 200, "successfully delete")
}

// Метод для перенаправления пользовательского запроса на корректную ссылку.
func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	activeLinkCache, ok := h.cache.Get(id)
	if ok {
		msg := "getting data from cache"
		h.logger.Info(msg, map[string]interface{}{})
		respond(w, r, 301, activeLinkCache)
		return
	}
	url, err := h.service.Get(id)
	if err != nil {
		errMsg := fmt.Errorf("error in Get service method %w", err).Error()
		h.logger.Error(errMsg, map[string]interface{}{})
		respond(w, r, http.StatusInternalServerError, nil)
		return
	}
	if url.HistoryLink == id && h.cache.Len() < 1000 {
		msg := "adding data to cache"
		h.logger.Info(msg, map[string]interface{}{})
		h.cache.Add(url.HistoryLink, url.ActiveLink, h.config.Technodom.Ttl)
	}
	// Если заданная ссылка является активной
	if url.ActiveLink == id {
		respond(w, r, 200, nil)
		return
	}

	// Перенаправление
	respond(w, r, 301, url.ActiveLink)
}
