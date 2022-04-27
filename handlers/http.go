package handlers

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/nedimdjulic/url-shortener/models"
	"github.com/nedimdjulic/url-shortener/service"

	"github.com/labstack/echo/v4"
)

// Handler --
type Handler struct {
	port string
	svc  service.Service
}

// NewHTTPSSvc --
func NewHTTPSvc(port string, svc service.Service) *Handler {
	return &Handler{port: port, svc: svc}
}

type message struct {
	Message string `json:"message,omitempty"`
}

type createReq struct {
	Url string `json:"url,omitempty"`
}

// Validate validates create request
func (r *createReq) Validate() error {
	_, err := url.ParseRequestURI(r.Url)
	if err != nil {
		return errors.New("not a valid URL")
	}

	return nil
}

type createRes struct {
	Shortened string `json:"shortened,omitempty"`
	Message   string `json:"message,omitempty"`
}

// HandleCreate godoc
// @Summary Creates a short URL
// @Description Returns a shortened URL mapped to the original URL sent in request
// @Accept json
// @Produce json
// @Param url-request body createReq true "Create new short URL"
// @Success 200 {object} createRes
// @Failure 500 {object} message
// @Router /create [post]
func (h *Handler) HandleCreate(c echo.Context) error {
	var req createReq

	if err := c.Bind(&req); err != nil {
		msg := message{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	if err := req.Validate(); err != nil {
		msg := message{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	newShort := models.Url{
		Original: req.Url,
	}

	newURL, err := h.svc.Create(&newShort)
	if err != nil {
		msg := message{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	res := createRes{
		Shortened: "http://localhost" + h.port + "/" + newURL.Shortened,
	}

	return c.JSON(http.StatusOK, res)
}

// HandleRedirect godoc
// @Summary Redirects to original URL
// @Description Redirection the to original URL, provided the shortened link
// @Accept json
// @Produce json
// @Param   short_url     path    string     true        "Shortened URL"
// @Success 303 {object} string
// @Failure 500 {object} message
// @Router /{short_url} [get]
func (h *Handler) HandleRedirect(c echo.Context) error {
	key := c.Param("short")

	res, err := h.svc.RetrieveFullLink(key)
	if err != nil {
		msg := message{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.Redirect(http.StatusSeeOther, res)
}

type countRes struct {
	Count int `json:"count,omitempty"`
}

// HandleGetCount godoc
// @Summary Get redirections count
// @Description Returns the count of redirections from shortened to original URL
// @Accept json
// @Produce json
// @Param   short_url     path    string     true        "Shortened URL"
// @Success 200 {object} countRes
// @Failure 500 {object} message
// @Router /count/{short_url} [get]
func (h *Handler) HandleGetCount(c echo.Context) error {
	key := c.Param("short")

	count, err := h.svc.RetrieveCount(key)
	if err != nil {
		msg := message{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	res := countRes{
		Count: count,
	}

	return c.JSON(http.StatusOK, res)
}

// HandleDelete godoc
// @Summary Delete short link
// @Description Removes the URL mapping of short to original
// @Accept json
// @Produce json
// @Param   short_url     path    string     true        "Shortened URL"
// @Success 200 {object} message
// @Failure 500 {object} message
// @Router /delete/{short_url} [delete]
func (h *Handler) HandleDelete(c echo.Context) error {
	key := c.Param("short")

	if err := h.svc.DeleteURL(key); err != nil {
		msg := message{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	res := message{
		Message: "URL removed",
	}

	return c.JSON(http.StatusOK, res)
}
