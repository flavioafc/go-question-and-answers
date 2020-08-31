package router

import (
	"encoding/json"
	"net/http"

	. "github.com/flavioafc/go-question-and-answers/dao"
	. "github.com/flavioafc/go-question-and-answers/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type AnswerRouter struct {
}

var answer_dao = AnswerDAO{}

// GetAll godoc
// @Summary Get a list to all questions and answers from API
// @Description Get a list of all questions and answers
// @Tags faq
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Answer
// @Router /api/v1/answer [get]
// GetAll list all questions and answers from QA API
func (a *AnswerRouter) GetAll(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	faq, err := answer_dao.GetAll(params["idquestion"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	answer_dao.ClearMemory()
	respondWithJSON(w, http.StatusOK, faq)
}

// GetByID godoc
// @Summary Get one answer item from the API
// @Description Get and answer by ID
// @Tags faq
// @Accept  json
// @Produce  json
// @Param id path string true "ObjectId"
// @Success 200 {object} models.Answer
// @Router /api/v1/answer/{id} [get]
// GetByID retrieve the question and answer by Id
func (a *AnswerRouter) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	faq, err := answer_dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid QA ID")
		return
	}
	respondWithJSON(w, http.StatusOK, faq)
}

// Create godoc
// @Summary Create a new Answer item
// @Description Create a new Answer with the input paylod
// @Tags faq
// @Accept  json
// @Produce  json
// @Param faq body models.Answer true "Create"
// @Success 200 {object} models.Answer
// @Router /api/v1/answer/{idparent} [post]
// Create method insert in the mongo database a new question and answer
func (a *AnswerRouter) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var faq AnswerRequest

	if err := json.NewDecoder(r.Body).Decode(&faq); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	faq.ID = bson.NewObjectId()
	if err := answer_dao.Create(faq); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, faq)
}

// Update godoc
// @Summary Update a new Question and Answer item
// @Description Update a new Question and Answer with the input paylod
// @Tags faq
// @Accept  json
// @Produce  json
// @Param id path string true "ObjectId"
// @Param faq body models.Answer true "Update"
// @Success 200 "ObjectIdHex(id), was successful updated!"
// @Router /api/v1/answer/{id} [put]
// Update have to update the question
func (a *AnswerRouter) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	var faq Answer
	if err := json.NewDecoder(r.Body).Decode(&faq); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := answer_dao.Update(params["id"], faq); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": faq.ID.String() + ", was successful updated!"})
}

// Delete godoc
// @Summary Delete one question and answer item from the API
// @Description Delete a question and answer
// @Tags faq
// @Accept  json
// @Produce  json
// @Param id path string true "ObjectId"
// @Success 200 "result: success"
// @Router /api/v1/answer/{id} [delete]
// Delete must delete the question
func (a *AnswerRouter) Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := answer_dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
