package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

var db *sql.DB

// @title			Your API Title
// @version		1.0
// @description	Your API description
// @termsOfService	https://example.com/terms
// @contact.name	Your Name
// @contact.email	youremail@example.com
// @BasePath		/
func main() {
	var err error
	db, err = sql.Open("mysql", "root:PdDHQ0DSPi2YS57ZnxAY@tcp(containers-us-west-186.railway.app:6588)/railway")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()

	// ATMs routes
	r.HandleFunc("/atm", getATMsHandler).Methods("GET")
	r.HandleFunc("/atm", createATMHandler).Methods("POST")

	// ATM Filters routes
	r.HandleFunc("/atm_filters", getATMFiltersHandler).Methods("GET")
	r.HandleFunc("/atm_filters", createATMFilterHandler).Methods("POST")

	// SalePoints routes
	r.HandleFunc("/salepoint", getSalePointsHandler).Methods("GET")
	r.HandleFunc("/salepoint", createSalePointHandler).Methods("POST")

	// SalePoint Filters routes
	r.HandleFunc("/salepoint_filters", getSalePointFiltersHandler).Methods("GET")
	r.HandleFunc("/salepoint_filters", createSalePointFilterHandler).Methods("POST")

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("docs/*any"),
	))

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	fmt.Println("Swagger UI is now available at http://localhost:8080/swagger/index.html")
	fmt.Println("Server is running on :8080")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

type ATM struct {
	// ID of the ATM
	// required: true
	ID int `json:"id_atms"`
	// ATM address
	// required: true
	Address string `json:"address"`
	// Latitude of the ATM
	// required: true
	Latitude float64 `json:"latitude"`
	// Longitude of the ATM
	// required: true
	Longitude float64 `json:"longitude"`
	// Indicates if the ATM operates 24/7
	// required: true
	AllDay string `json:"allDay"`
	// Services provided by the ATM
	// required: true
	Services string `json:"services"`
}

type ATMFilter struct {
	// ID of the ATM Filter
	// required: true
	ID int `json:"id_atms"`
	// Cash filter for the ATM
	// required: true
	Cash int `json:"cash"`
}

type SalePoint struct {
	// ID of the SalePoint
	// required: true
	ID int `json:"offices_id"`
	// SalePoint Name
	// required: true
	SalePointName string `json:"salePointName"`
	// SalePoint Address
	// required: true
	Address string `json:"address"`
	// SalePoint Status
	// required: true
	Status string `json:"status"`
	// SalePoint Open Hours
	// required: true
	OpenHours string `json:"openHours"`
	// RKO
	// required: true
	RKO string `json:"rko"`
	// Open Hours Individual
	// required: true
	OpenHoursIndividual string `json:"openHoursIndividual"`
	// Office Type
	// required: true
	OfficeType string `json:"officeType"`
	// SalePoint Format
	// required: true
	SalePointFormat string `json:"salePointFormat"`
	// SUO Availability
	// required: true
	SUOAvailability string `json:"suoAvailability"`
	// Has Ramp
	// required: true
	HasRamp string `json:"hasRamp"`
	// Latitude of the SalePoint
	// required: true
	Latitude float64 `json:"latitude"`
	// Longitude of the SalePoint
	// required: true
	Longitude float64 `json:"longitude"`
	// Metro Station
	// required: true
	MetroStation string `json:"metroStation"`
	// Distance
	// required: true
	Distance int `json:"distance"`
	// Kep
	// required: true
	Kep string `json:"kep"`
	// My Branch
	// required: true
	MyBranch string `json:"myBranch"`
	// Network
	// required: true
	Network string `json:"network"`
	// SalePoint Code
	// required: true
	SalePointCode string `json:"salePointCode"`
}

type SalePointFilter struct {
	// ID of the SalePoint Filter
	// required: true
	ID int `json:"offices_id"`
	// Current Workload filter for SalePoint
	// required: true
	CurrentWorkload int `json:"current_workload"`
	// Rating filter for SalePoint
	// required: true
	Rating int `json:"rating"`
}

func sendErrorResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorResponse := map[string]string{"error": message}
	json.NewEncoder(w).Encode(errorResponse)
}

// getATMsHandler retrieves a list of ATMs.
//
//	@Summary		Get a list of ATMs
//	@Description	Retrieve a list of ATMs
//	@Produce		json
//	@Success		200	{array}	ATM
//	@Router			/atm [get]
func getATMsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM ATM")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve ATMs")
		return
	}
	defer rows.Close()

	atms := []ATM{}
	for rows.Next() {
		var atm ATM
		err := rows.Scan(&atm.ID, &atm.Address, &atm.Latitude, &atm.Longitude, &atm.AllDay, &atm.Services)
		if err != nil {
			sendErrorResponse(w, http.StatusInternalServerError, "Failed to scan ATM data")
			return
		}
		atms = append(atms, atm)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(atms)
}

// createATMHandler creates a new ATM entry.
//
//	@Summary		Create a new ATM
//	@Description	Create a new ATM entry
//	@Accept			json
//	@Produce		json
//	@Param			newATM	body		ATM	true	"New ATM data"
//	@Success		201		{object}	ATM
//	@Router			/atm [post]
func createATMHandler(w http.ResponseWriter, r *http.Request) {
	var newATM ATM
	err := json.NewDecoder(r.Body).Decode(&newATM)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	result, err := db.Exec("INSERT INTO ATM (address, latitude, longitude, allDay, services) VALUES (?, ?, ?, ?, ?)",
		newATM.Address, newATM.Latitude, newATM.Longitude, newATM.AllDay, newATM.Services)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to create ATM")
		return
	}

	insertedID, _ := result.LastInsertId()
	newATM.ID = int(insertedID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newATM)
}

func updateATMHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	atmID, err := strconv.Atoi(params["atmID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Сначала получаем существующие данные ATM
	var existingATM ATM
	err = db.QueryRow("SELECT id_atms, address, latitude, longitude, allDay, services FROM ATM WHERE id_atms = ?", atmID).Scan(
		&existingATM.ID, &existingATM.Address, &existingATM.Latitude, &existingATM.Longitude, &existingATM.AllDay, &existingATM.Services)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Затем декодируем новую информацию
	var newATM ATM
	err = json.NewDecoder(r.Body).Decode(&newATM)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Обновляем информацию
	_, err = db.Exec("UPDATE ATM SET address = ?, latitude = ?, longitude = ?, allDay = ?, services = ? WHERE id_atms = ?",
		newATM.Address, newATM.Latitude, newATM.Longitude, newATM.AllDay, newATM.Services, atmID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newATM)
}


// deleteATMHandler deletes an ATM entry by ID.
func deleteATMHandler(w http.ResponseWriter, r *http.Request) {
    ATMID := mux.Vars(r)["id"]

    _, err := db.Exec("DELETE FROM ATM WHERE id_atms=?", ATMID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

// getATMFiltersHandler retrieves a list of ATM filters.
//
//	@Summary		Get a list of ATM filters
//	@Description	Retrieve a list of ATM filters
//	@Produce		json
//	@Success		200	{array}	ATMFilter
//	@Router			/atm_filters [get]
func getATMFiltersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM ATM_Filters")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve ATM filters")
		return
	}
	defer rows.Close()

	atmFilters := []ATMFilter{}
	for rows.Next() {
		var atmFilter ATMFilter
		err := rows.Scan(&atmFilter.ID, &atmFilter.Cash)
		if err != nil {
			sendErrorResponse(w, http.StatusInternalServerError, "Failed to scan ATM filter data")
			return
		}
		atmFilters = append(atmFilters, atmFilter)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(atmFilters)
}

// createATMFilterHandler creates a new ATM filter entry.
//
//	@Summary		Create a new ATM filter
//	@Description	Create a new ATM filter entry
//	@Accept			json
//	@Produce		json
//	@Param			newATMFilter	body		ATMFilter	true	"New ATM filter data"
//	@Success		201				{object}	ATMFilter
//	@Router			/atm_filters [post]
func createATMFilterHandler(w http.ResponseWriter, r *http.Request) {
	var newATMFilter ATMFilter
	err := json.NewDecoder(r.Body).Decode(&newATMFilter)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	result, err := db.Exec("INSERT INTO ATM_Filters (id_atms, cash) VALUES (?, ?)", newATMFilter.ID, newATMFilter.Cash)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to create ATM filter")
		return
	}

	insertedID, _ := result.LastInsertId()
	newATMFilter.ID = int(insertedID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newATMFilter)
}

// updateATMFilterHandler обновляет существующий фильтр ATM по его ID.
func updateATMFilterHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	atmFilterID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ATM Filter ID", http.StatusBadRequest)
		return
	}

	var updatedATMFilter ATMFilter
	err = json.NewDecoder(r.Body).Decode(&updatedATMFilter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Сначала получаем существующие данные фильтра ATM
	var existingATMFilter ATMFilter
	err = db.QueryRow("SELECT id_atms, cash FROM ATM_Filters WHERE id_atms = ?", atmFilterID).Scan(
		&existingATMFilter.ID, &existingATMFilter.Cash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Обновляем только необходимые поля
	if updatedATMFilter.Cash != 0 {
		existingATMFilter.Cash = updatedATMFilter.Cash
	}

	// Выполняем обновление фильтра ATM в базе данных
	_, err = db.Exec("UPDATE ATM_Filters SET cash = ? WHERE id_atms = ?", existingATMFilter.Cash, atmFilterID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// deleteATMFilterHandler удаляет фильтр ATM по его ID.
func deleteATMFilterHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	atmFilterID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ATM Filter ID", http.StatusBadRequest)
		return
	}

	// Удаляем фильтр ATM из базы данных
	_, err = db.Exec("DELETE FROM ATM_Filters WHERE id_atms = ?", atmFilterID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// getSalePointsHandler retrieves a list of SalePoints.
//
//	@Summary		Get a list of SalePoints
//	@Description	Retrieve a list of SalePoints
//	@Produce		json
//	@Success		200	{array}	SalePoint
//	@Router			/salepoint [get]
func getSalePointsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM SalePoint")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve SalePoints")
		return
	}
	defer rows.Close()

	salePoints := []SalePoint{}
	for rows.Next() {
		var salePoint SalePoint
		err := rows.Scan(&salePoint.ID, &salePoint.SalePointName, &salePoint.Address, &salePoint.Status, &salePoint.OpenHours, &salePoint.RKO,
			&salePoint.OpenHoursIndividual, &salePoint.OfficeType, &salePoint.SalePointFormat, &salePoint.SUOAvailability, &salePoint.HasRamp,
			&salePoint.Latitude, &salePoint.Longitude, &salePoint.MetroStation, &salePoint.Distance, &salePoint.Kep, &salePoint.MyBranch,
			&salePoint.Network, &salePoint.SalePointCode)
		if err != nil {
			sendErrorResponse(w, http.StatusInternalServerError, "Failed to scan SalePoint data")
			return
		}
		salePoints = append(salePoints, salePoint)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(salePoints)
}

// createSalePointHandler creates a new SalePoint entry.
//
//	@Summary		Create a new SalePoint
//	@Description	Create a new SalePoint entry
//	@Accept			json
//	@Produce		json
//	@Param			newSalePoint	body		SalePoint	true	"New SalePoint data"
//	@Success		201				{object}	SalePoint
//	@Router			/salepoint [post]
func createSalePointHandler(w http.ResponseWriter, r *http.Request) {
	var newSalePoint SalePoint
	err := json.NewDecoder(r.Body).Decode(&newSalePoint)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	result, err := db.Exec("INSERT INTO SalePoint (offices_id, salePointName, address, status, openHours, rko, openHoursIndividual, officeType, salePointFormat, suoAvailability, hasRamp, latitude, longitude, metroStation, distance, kep, myBranch, network, salePointCode) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		newSalePoint.ID, newSalePoint.SalePointName, newSalePoint.Address, newSalePoint.Status, newSalePoint.OpenHours, newSalePoint.RKO, newSalePoint.OpenHoursIndividual, newSalePoint.OfficeType, newSalePoint.SalePointFormat, newSalePoint.SUOAvailability, newSalePoint.HasRamp, newSalePoint.Latitude, newSalePoint.Longitude, newSalePoint.MetroStation, newSalePoint.Distance, newSalePoint.Kep, newSalePoint.MyBranch, newSalePoint.Network, newSalePoint.SalePointCode)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to create SalePoint")
		return
	}

	insertedID, _ := result.LastInsertId()
	newSalePoint.ID = int(insertedID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newSalePoint)
}

// updateSalePointHandler обновляет данные о SalePoint по его ID.
// @Summary		Обновить данные SalePoint
// @Description	Обновить данные SalePoint по его ID
// @Accept		json
// @Produce		json
// @Param		officeID	path	int	true	"ID SalePoint"
// @Param		updatedSalePoint	body	SalePoint	true	"Обновленные данные SalePoint"
// @Success		200	{object}	SalePoint
// @Failure		400	{string}	string
// @Router		/salepoint/{officeID} [put]
func updateSalePointHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	officeID, err := strconv.Atoi(params["officeID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedSalePoint SalePoint
	err = json.NewDecoder(r.Body).Decode(&updatedSalePoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE SalePoint SET salePointName = ?, address = ?, status = ?, openHours = ?, rko = ?, openHoursIndividual = ?, officeType = ?, salePointFormat = ?, suoAvailability = ?, hasRamp = ?, latitude = ?, longitude = ?, metroStation = ?, distance = ?, kep = ?, myBranch = ?, network = ?, salePointCode = ? WHERE offices_id = ?",
		updatedSalePoint.SalePointName, updatedSalePoint.Address, updatedSalePoint.Status, updatedSalePoint.OpenHours, updatedSalePoint.RKO, updatedSalePoint.OpenHoursIndividual, updatedSalePoint.OfficeType, updatedSalePoint.SalePointFormat, updatedSalePoint.SUOAvailability, updatedSalePoint.HasRamp, updatedSalePoint.Latitude, updatedSalePoint.Longitude, updatedSalePoint.MetroStation, updatedSalePoint.Distance, updatedSalePoint.Kep, updatedSalePoint.MyBranch, updatedSalePoint.Network, updatedSalePoint.SalePointCode, officeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedSalePoint.ID = officeID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedSalePoint)
}

// deleteSalePointHandler удаляет SalePoint по его ID.
// @Summary		Удалить SalePoint
// @Description	Удалить SalePoint по его ID
// @Param		officeID	path	int	true	"ID SalePoint"
// @Success		204	{string}	string
// @Failure		400	{string}	string
// @Router		/salepoint/{officeID} [delete]
func deleteSalePointHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	officeID, err := strconv.Atoi(params["officeID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM SalePoint WHERE offices_id = ?", officeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// getSalePointFiltersHandler retrieves a list of SalePoint filters.
//
//	@Summary		Get a list of SalePoint filters
//	@Description	Retrieve a list of SalePoint filters
//	@Produce		json
//	@Success		200	{array}	SalePointFilter
//	@Router			/salepoint_filters [get]
func getSalePointFiltersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM SalePointFilter")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve SalePoint filters")
		return
	}
	defer rows.Close()

	salePointFilters := []SalePointFilter{}
	for rows.Next() {
		var salePointFilter SalePointFilter
		err := rows.Scan(&salePointFilter.ID, &salePointFilter.CurrentWorkload, &salePointFilter.Rating)
		if err != nil {
			sendErrorResponse(w, http.StatusInternalServerError, "Failed to scan SalePointFilter data")
			return
		}
		salePointFilters = append(salePointFilters, salePointFilter)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(salePointFilters)
}

// createSalePointFilterHandler creates a new SalePoint filter entry.
//
//	@Summary		Create a new SalePoint filter
//	@Description	Create a new SalePoint filter entry
//	@Accept			json
//	@Produce		json
//	@Param			newSalePointFilter	body		SalePointFilter	true	"New SalePoint filter data"
//	@Success		201					{object}	SalePointFilter
//	@Router			/salepoint_filters [post]
func createSalePointFilterHandler(w http.ResponseWriter, r *http.Request) {
	var newSalePointFilter SalePointFilter
	err := json.NewDecoder(r.Body).Decode(&newSalePointFilter)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	result, err := db.Exec("INSERT INTO SalePointFilter (offices_id, current_workload, rating) VALUES (?, ?, ?)", newSalePointFilter.ID, newSalePointFilter.CurrentWorkload, newSalePointFilter.Rating)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to create SalePointFilter")
		return
	}

	insertedID, _ := result.LastInsertId()
	newSalePointFilter.ID = int(insertedID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newSalePointFilter)
}

// updateSalePointFilterHandler обновляет информацию о SalePointFilter.
// @Summary		Обновить SalePointFilter
// @Description	Обновить информацию о SalePointFilter
// @Accept		json
// @Param		filterID	path	int	true	"ID SalePointFilter"
// @Param		newFilter	body	SalePointFilter	true	"Новая информация о SalePointFilter"
// @Success		200	{object}	SalePointFilter
// @Failure		400	{string}	string
// @Router		/salepoint_filters/{filterID} [put]
func updateSalePointFilterHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	filterID, err := strconv.Atoi(params["filterID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var newFilter SalePointFilter
	err = json.NewDecoder(r.Body).Decode(&newFilter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE SalePointFilter SET current_workload = ?, rating = ? WHERE offices_id = ?", newFilter.CurrentWorkload, newFilter.Rating, filterID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newFilter)
}

// deleteSalePointFilterHandler удаляет SalePointFilter по его ID.
// @Summary		Удалить SalePointFilter
// @Description	Удалить SalePointFilter по его ID
// @Param		filterID	path	int	true	"ID SalePointFilter"
// @Success		204	{string}	string
// @Failure		400	{string}	string
// @Router		/salepoint_filters/{filterID} [delete]
func deleteSalePointFilterHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	filterID, err := strconv.Atoi(params["filterID"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM SalePointFilter WHERE offices_id = ?", filterID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
