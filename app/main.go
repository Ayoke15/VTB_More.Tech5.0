package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var db *sql.DB

type ATM struct {
	ID          int    `json:"id_atms"`
	Address     string `json:"address"`
	Coordinates string `json:"coordinates"`
	AllDay      bool   `json:"all_day"`
	OfficesID   int    `json:"offices_id"`
}

type Service struct {
	ServiceID   int    `json:"service_id"`
	ServiceName string `json:"service_name"`
}

type ATMService struct {
	ID                int    `json:"id_atms_service"`
	ATMID             int    `json:"id_atms"`
	ServiceID         int    `json:"service_id"`
	ServiceCapability string `json:"service_capability"`
	ServiceActivity   string `json:"service_activity"`
}

type SalePoint struct {
	OfficesID       int    `json:"offices_id"`
	SalePointName   string `json:"sale_point_name"`
	Address         string `json:"address"`
	Status          string `json:"status"`
	RKO             string `json:"rko"`
	OfficeType      string `json:"office_type"`
	SalePointFormat string `json:"sale_point_format"`
	SUOAvailability string `json:"suo_availability"`
	HasRamp         string `json:"has_ramp"`
	Coordinates     string `json:"coordinates"`
	MetroStation    string `json:"metro_station"`
	Distance        int    `json:"distance"`
	Kep             bool   `json:"kep"`
	MyBranch        bool   `json:"my_branch"`
}

type SalePointOpenHours struct {
	OpenHourID int    `json:"open_hour_id"`
	OfficesID  int    `json:"offices_id"`
	Day        string `json:"day"`
	Hours      string `json:"hours"`
}

type SalePointFilter struct {
	OfficeFilterID  int     `json:"office_filter_id"`
	Rating          float64 `json:"rating"`
	CurrentWorkload int     `json:"current_workload"`
	OfficesID       []int   `json:"offices_id"`
}

type ATMFilter struct {
	ATMFilterID int     `json:"atm_filter_id"`
	Rating      float64 `json:"rating"`
	Cash        int     `json:"cash"`
	IDAtms      []int   `json:"id_atms"`
}

func main() {
	db, _ = sql.Open("sqlite3", "your_database.db")
	router := mux.NewRouter()

	router.HandleFunc("/api/atms", GetATMs).Methods("GET")
	router.HandleFunc("/api/atms", CreateATM).Methods("POST")

	router.HandleFunc("/api/services", GetServices).Methods("GET")
	router.HandleFunc("/api/services", CreateService).Methods("POST")

	router.HandleFunc("/api/atm_services", GetATMServices).Methods("GET")
	router.HandleFunc("/api/atm_services", CreateATMService).Methods("POST")

	router.HandleFunc("/api/sale_points", GetSalePoints).Methods("GET")
	router.HandleFunc("/api/sale_points", CreateSalePoint).Methods("POST")

	router.HandleFunc("/api/sale_point_open_hours", GetSalePointOpenHours).Methods("GET")
	router.HandleFunc("/api/sale_point_open_hours", CreateSalePointOpenHours).Methods("POST")

	router.HandleFunc("/api/sale_point_filters", GetSalePointFilters).Methods("GET")
	router.HandleFunc("/api/sale_point_filters", CreateSalePointFilter).Methods("POST")

	router.HandleFunc("/api/atm_filters", GetATMFilters).Methods("GET")
	router.HandleFunc("/api/atm_filters", CreateATMFilter).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetATMs(w http.ResponseWriter, r *http.Request) {
	var atms []ATM
	rows, err := db.Query("SELECT * FROM ATM")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var atm ATM
		err := rows.Scan(&atm.ID, &atm.Address, &atm.Coordinates, &atm.AllDay, &atm.OfficesID)
		if err != nil {
			fmt.Println(err)
			return
		}
		atms = append(atms, atm)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(atms)
}

func CreateATM(w http.ResponseWriter, r *http.Request) {
	var atm ATM
	json.NewDecoder(r.Body).Decode(&atm)
	stmt, err := db.Prepare("INSERT INTO ATM (address, coordinates, all_day, offices_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = stmt.Exec(atm.Address, atm.Coordinates, atm.AllDay, atm.OfficesID)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(atm)
}

func GetServices(w http.ResponseWriter, r *http.Request) {
	var services []Service
	rows, err := db.Query("SELECT * FROM Service")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var service Service
		err := rows.Scan(&service.ServiceID, &service.ServiceName)
		if err != nil {
			fmt.Println(err)
			return
		}
		services = append(services, service)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func CreateService(w http.ResponseWriter, r *http.Request) {
	var service Service
	json.NewDecoder(r.Body).Decode(&service)
	stmt, err := db.Prepare("INSERT INTO Service (service_name) VALUES (?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = stmt.Exec(service.ServiceName)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
}

func GetATMServices(w http.ResponseWriter, r *http.Request) {
	var atmServices []ATMService
	rows, err := db.Query("SELECT * FROM ATM_Service")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var atmService ATMService
		err := rows.Scan(&atmService.ID, &atmService.ATMID, &atmService.ServiceID, &atmService.ServiceCapability, &atmService.ServiceActivity)
		if err != nil {
			fmt.Println(err)
			return
		}
		atmServices = append(atmServices, atmService)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(atmServices)
}

func CreateATMService(w http.ResponseWriter, r *http.Request) {
	var atmService ATMService
	json.NewDecoder(r.Body).Decode(&atmService)
	stmt, err := db.Prepare("INSERT INTO ATM_Service (id_atms, service_id, service_capability, service_activity) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = stmt.Exec(atmService.ATMID, atmService.ServiceID, atmService.ServiceCapability, atmService.ServiceActivity)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(atmService)
}

func GetSalePoints(w http.ResponseWriter, r *http.Request) {
	var salePoints []SalePoint
	rows, err := db.Query("SELECT * FROM SalePoint")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var salePoint SalePoint
		err := rows.Scan(&salePoint.OfficesID, &salePoint.SalePointName, &salePoint.Address, &salePoint.Status, &salePoint.RKO, &salePoint.OfficeType, &salePoint.SalePointFormat, &salePoint.SUOAvailability, &salePoint.HasRamp, &salePoint.Coordinates, &salePoint.MetroStation, &salePoint.Distance, &salePoint.Kep, &salePoint.MyBranch)
		if err != nil {
			fmt.Println(err)
			return
		}
		salePoints = append(salePoints, salePoint)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(salePoints)
}

func CreateSalePoint(w http.ResponseWriter, r *http.Request) {
	var salePoint SalePoint
	json.NewDecoder(r.Body).Decode(&salePoint)
	stmt, err := db.Prepare("INSERT INTO SalePoint (sale_point_name, address, status, rko, office_type, sale_point_format, suo_availability, has_ramp, coordinates, metro_station, distance, kep, my_branch) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = stmt.Exec(salePoint.SalePointName, salePoint.Address, salePoint.Status, salePoint.RKO, salePoint.OfficeType, salePoint.SalePointFormat, salePoint.SUOAvailability, salePoint.HasRamp, salePoint.Coordinates, salePoint.MetroStation, salePoint.Distance, salePoint.Kep, salePoint.MyBranch)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(salePoint)
}

func GetSalePointOpenHours(w http.ResponseWriter, r *http.Request) {
	var openHours []SalePointOpenHours
	rows, err := db.Query("SELECT * FROM SalePointOpenHours")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var openHour SalePointOpenHours
		err := rows.Scan(&openHour.OpenHourID, &openHour.OfficesID, &openHour.Day, &openHour.Hours)
		if err != nil {
			fmt.Println(err)
			return
		}
		openHours = append(openHours, openHour)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(openHours)
}

func CreateSalePointOpenHours(w http.ResponseWriter, r *http.Request) {
	var openHour SalePointOpenHours
	json.NewDecoder(r.Body).Decode(&openHour)
	stmt, err := db.Prepare("INSERT INTO SalePointOpenHours (offices_id, day, hours) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = stmt.Exec(openHour.OfficesID, openHour.Day, openHour.Hours)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(openHour)
}

func GetSalePointFilters(w http.ResponseWriter, r *http.Request) {
	var filters []SalePointFilter
	rows, err := db.Query("SELECT * FROM SalePointFilter")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var filter SalePointFilter
		err := rows.Scan(&filter.OfficeFilterID, &filter.Rating, &filter.CurrentWorkload)
		if err != nil {
			fmt.Println(err)
			return
		}
		filter.OfficesID = getOfficeIDsForFilter(filter.OfficeFilterID)
		filters = append(filters, filter)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filters)
}

func getOfficeIDsForFilter(filterID int) []int {
	var officeIDs []int
	rows, err := db.Query("SELECT offices_id FROM SalePointFilter_Offices WHERE office_filter_id = ?", filterID)
	if err != nil {
		fmt.Println(err)
		return officeIDs
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
			return officeIDs
		}
		officeIDs = append(officeIDs, id)
	}
	return officeIDs
}

func CreateSalePointFilter(w http.ResponseWriter, r *http.Request) {
	var filter SalePointFilter
	json.NewDecoder(r.Body).Decode(&filter)
	stmt, err := db.Prepare("INSERT INTO SalePointFilter (rating, current_workload) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := stmt.Exec(filter.Rating, filter.CurrentWorkload)
	if err != nil {
		fmt.Println(err)
		return
	}
	filterID, _ := result.LastInsertId()
	for _, officeID := range filter.OfficesID {
		stmt, err = db.Prepare("INSERT INTO SalePointFilter_Offices (office_filter_id, offices_id) VALUES (?, ?)")
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = stmt.Exec(filterID, officeID)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filter)
}

func GetATMFilters(w http.ResponseWriter, r *http.Request) {
	var filters []ATMFilter
	rows, err := db.Query("SELECT * FROM ATMFilter")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var filter ATMFilter
		err := rows.Scan(&filter.ATMFilterID, &filter.Rating, &filter.Cash)
		if err != nil {
			fmt.Println(err)
			return
		}
		filter.IDAtms = getATMIDsForFilter(filter.ATMFilterID)
		filters = append(filters, filter)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filters)
}

func getATMIDsForFilter(filterID int) []int {
	var atmIDs []int
	rows, err := db.Query("SELECT id_atms FROM ATMFilter_ATMs WHERE atm_filter_id = ?", filterID)
	if err != nil {
		fmt.Println(err)
		return atmIDs
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
			return atmIDs
		}
		atmIDs = append(atmIDs, id)
	}
	return atmIDs
}

func CreateATMFilter(w http.ResponseWriter, r *http.Request) {
	var filter ATMFilter
	json.NewDecoder(r.Body).Decode(&filter)
	stmt, err := db.Prepare("INSERT INTO ATMFilter (rating, cash) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := stmt.Exec(filter.Rating, filter.Cash)
	if err != nil {
		fmt.Println(err)
		return
	}
	filterID, _ := result.LastInsertId()
	for _, atmID := range filter.IDAtms {
		stmt, err = db.Prepare("INSERT INTO ATMFilter_ATMs (atm_filter_id, id_atms) VALUES (?, ?)")
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = stmt.Exec(filterID, atmID)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filter)
}
