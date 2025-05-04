package http

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"kaspi-tz/internal/domain/person"
	"kaspi-tz/internal/service/contragent"
	"kaspi-tz/pkg/server/response"
)

type PeopleHandler struct {
	ContragentService *contragent.Service
}

func NewPeopleHandler(contragentService *contragent.Service) *PeopleHandler {
	return &PeopleHandler{
		ContragentService: contragentService,
	}
}

func (h *PeopleHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/iin_check", h.validateIIN)

	r.Route("/people/info", func(r chi.Router) {
		r.Post("/", h.createPerson)
		r.Get("/iin/{iin}", h.getPersonByIIN)
		r.Get("/name/{name}", h.getPeopleByNamePart)
	})

	return r
}

//	@Summary	Validate IIN
//	@Tags		validation
//	@Accept		json
//	@Produce	json
//	@Param		iin			query		string	true	"iin"	default(020513550507)
//	@Success	200			{object}	person.ValidateIINResponse
//	@Failure	400			{object}	response.Object
//	@Failure	500			{object}	response.Object
//	@Router		/iin_check 	[get]
func (h *PeopleHandler) validateIIN(w http.ResponseWriter, r *http.Request) {
	iin := r.URL.Query().Get("iin")

	res, err := h.ContragentService.ValidateIIN(iin)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

//	@Summary	Create person
//	@Tags		people
//	@Accept		json
//	@Produce	json
//	@Param		req				body		person.CreatePersonRequest	true	"req body"
//	@Success	200				{object}	response.Object
//	@Failure	400				{object}	response.Object
//	@Failure	500				{object}	response.Object
//	@Router		/people/info 	[post]
func (h *PeopleHandler) createPerson(w http.ResponseWriter, r *http.Request) {
	req := person.CreatePersonRequest{}

	if err := render.Decode(r, &req); err != nil {
		response.BadRequest(w, r, err)
		return
	}

	if _, err := h.ContragentService.ValidateIIN(req.IIN); err != nil {
		response.BadRequest(w, r, err)
		return
	}

	if err := h.ContragentService.CreatePerson(r.Context(), req); err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.NoContent(w, r)
}

//	@Summary	Get person by IIN
//	@Tags		people
//	@Accept		json
//	@Produce	json
//	@Param		iin						path		string	true	"path param"	default(020513550507)
//	@Success	200						{object}	person.GetPersonResponse
//	@Failure	400						{object}	response.Object
//	@Failure	500						{object}	response.Object
//	@Router		/people/info/iin/{iin} 	[get]
func (h *PeopleHandler) getPersonByIIN(w http.ResponseWriter, r *http.Request) {
	iin := chi.URLParam(r, "iin")

	if _, err := h.ContragentService.ValidateIIN(iin); err != nil {
		response.BadRequest(w, r, err)
		return
	}

	res, err := h.ContragentService.GetPersonByIIN(r.Context(), iin)
	if err != nil {
		switch {
		case errors.Is(err, person.ErrPersonNotFound):
			response.NotFound(w, r, err)
		default:
			response.InternalServerError(w, r, err)
		}

		return
	}

	response.OK(w, r, res)
}

//	@Summary	Get people by name part
//	@Tags		people
//	@Accept		json
//	@Produce	json
//	@Param		name						path		string	true	"path param"	default(жакс)
//	@Success	200							{array}		person.GetPersonResponse
//	@Failure	400							{object}	response.Object
//	@Failure	500							{object}	response.Object
//	@Router		/people/info/name/{name} 	[get]
func (h *PeopleHandler) getPeopleByNamePart(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	res, err := h.ContragentService.GetPeopleByNamePart(r.Context(), name)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}
