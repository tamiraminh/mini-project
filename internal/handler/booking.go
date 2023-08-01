package handler

import (
	"encoding/json"
	"fmt"
	"mini-project/internal/domain"
	"mini-project/internal/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
)


// @Summary Create a booking
// @Description Create a new booking.
// @ID create-booking
// @Accept json
// @Produce json
// @Param booking body domain.BookingRequestFormat true "Booking data in JSON format"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.ResponseString
// @Failure 500 {object} models.ResponseString
// @Router /bookings [post]
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookings := domain.BookingRequestFormat{}
	resErr := models.ResponseString{}
	res := models.Response{}

	if err := json.NewDecoder(r.Body).Decode(&bookings); err != nil {
		fmt.Println(err)
		resErr.Data = "error"
		resErr.Message = "bad request"

		jsonData, err := json.Marshal(resErr)
		if err != nil {
			fmt.Println(err.Error())
	
		}
		w.WriteHeader(400)
		w.Write(jsonData)
		return	
	}

	

	bookingId, _ := uuid.NewV4()
	bookings.Data[0].Id = bookingId

	bookingInserted, err := domain.CreateBooking(bookings.Data[0])
	if err != nil {
		fmt.Println(err)
		resErr.Data = "error"
		resErr.Message = "Failed to insert to Database"
		jsonData, err := json.Marshal(resErr)
		if err != nil {
			fmt.Println(err.Error())
	
		}
		w.WriteHeader(500)
		w.Write(jsonData)
		return	
	}

	res.Data = []domain.Booking{bookingInserted}
	jsonData, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())

	}
	w.WriteHeader(201)
	w.Write(jsonData)
	return	

}

// @Summary Get bookings
// @Description Get a list of bookings with pagination and sorting options.
// @ID get-bookings
// @Accept json
// @Produce json
// @Param order query string false "Sorting order: asc or desc"
// @Param orderBy query string false "Sort by column: id, clientName, photographerName, package, dateTime, location, status"
// @Param page query int false "Page number"
// @Param limit query int false "Number of records per page"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ResponseString
// @Failure 500 {object} models.ResponseString
// @Router /bookings [get]
func ReadBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resErr := models.ResponseString{}
	res := models.Response{}


	order := r.URL.Query().Get("order")
	orderBy := r.URL.Query().Get("orderBy")
	pageStr := r.URL.Query().Get("page")
	pageInt, err := strconv.Atoi(pageStr)
	if err != nil {
		fmt.Println(err.Error())
		resErr.Data = "error"
		resErr.Message = "bad request"
		jsonData, err := json.Marshal(resErr)
		if err != nil {
			fmt.Println(err.Error())
	
		}
		w.WriteHeader(400)
		w.Write(jsonData)
		return	
	}
	limitStr := r.URL.Query().Get("limit")
	limitInt, err := strconv.Atoi(limitStr)
	if err != nil {
		fmt.Println(err.Error())
		resErr.Data = "error"
		resErr.Message = "bad request"
		jsonData, err := json.Marshal(resErr)
		if err != nil {
			fmt.Println(err.Error())
	
		}
		w.WriteHeader(400)
		w.Write(jsonData)
		return	
	}

	bookings, err := domain.ReadBookings(order, orderBy, pageInt, limitInt)
	if err != nil {
		fmt.Println(err.Error())
		resErr.Data = "error"
		resErr.Message = "Failed to Read Database"
		jsonData, err := json.Marshal(resErr)
		if err != nil {
			fmt.Println(err.Error())
	
		}
		w.WriteHeader(500)
		w.Write(jsonData)
		return	
	}

	res.Data = bookings
	jsonData, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())

	}
	w.WriteHeader(200)
	w.Write(jsonData)
	return	

}

// @Summary Update a booking
// @Description Update an existing booking.
// @ID update-booking
// @Accept json
// @Produce json
// @Param id path string true "Booking ID"
// @Param booking body domain.BookingRequestFormat true "Booking data in JSON format"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ResponseString
// @Failure 500 {object} models.ResponseString
// @Router /bookings/{id} [put]
func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookingId := chi.URLParam(r, "id")
	bookings := domain.BookingRequestFormat{}
	resErr := models.ResponseString{}
	res := models.Response{}

	if err := json.NewDecoder(r.Body).Decode(&bookings); err != nil {
		fmt.Println(err)
		resErr.Data = "error"
		resErr.Message = "Bad Request"

		jsonData, err := json.Marshal(resErr)
		if err != nil {
			fmt.Println(err.Error())
	
		}
		w.WriteHeader(400)
		w.Write(jsonData)
		return	
	}

	bookings.Data[0].Id = uuid.Must(uuid.FromString(bookingId))
	bookingUpdated, err := domain.UpdateBooking(bookings.Data[0])
	if err != nil {
		fmt.Println(err)
		resErr.Data = "error"
		resErr.Message = "Failed to Update on Database"
		jsonData, err := json.Marshal(resErr)
		if err != nil {
			fmt.Println(err.Error())
	
		}
		w.WriteHeader(500)
		w.Write(jsonData)
		return	
	}

	res.Data = []domain.Booking{bookingUpdated}
	jsonData, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())

	}
	w.WriteHeader(200)
	w.Write(jsonData)
	return	

}

// @Summary Delete a booking
// @Description Delete an existing booking.
// @ID delete-booking
// @Param id path string true "Booking ID"
// @Success 200 {object} models.ResponseString
// @Failure 500 {object} models.ResponseString
// @Router /bookings/{id} [delete]
func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookingId := chi.URLParam(r, "id")
	resErr := models.ResponseString{}
	res := models.ResponseString{}

	
	deletedMessage, err := domain.DeleteBooking(bookingId)
	if err != nil {
		fmt.Println(err)
		resErr.Data = "error"
		resErr.Message = "Failed to Update on Database"
		jsonData, err := json.Marshal(resErr)
		if err != nil {
			fmt.Println(err.Error())
	
		}
		w.WriteHeader(500)
		w.Write(jsonData)
		return	
	}

	res.Message = deletedMessage
	res.Data = "success"
	jsonData, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())

	}
	w.WriteHeader(200)
	w.Write(jsonData)
	return	

}