package handlers

import (
	"WEB/internal/config"
	"WEB/internal/forms"
	"WEB/internal/models"
	"WEB/internal/render"
	"WEB/internal/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	//form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)

}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability handles post
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and sends JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Redir(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "payment-redirect.page.tmpl", &models.TemplateData{})
}

// Contact renders the contact page
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)

	if !ok {
		m.App.Session.Put(r.Context(), "error", "Server is busy 。!!!!")
		http.Redirect(w, r, "/make-reservation", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	m.App.Session.Put(r.Context(), "flash", "Redirect to Payment")

	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

func (m *Repository) RedirectToPayment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	currentTime := time.Now()
	formattedTime := currentTime.Format("20060102150405")

	reservation := models.Payment{
		OrderNo:      formattedTime,
		Desc:         "IPhone 14 pro 512GB",
		CustomerName: fmt.Sprint(r.Form.Get("first_name"), "-", r.Form.Get("last_name")),
		Email:        r.Form.Get("email"),
		Phone:        r.Form.Get("phone"),
		Amount:       1000,
		Currency:     "TWD",
	}

	formHtml := services.ECPay(reservation)

	//formHtml = "<script  languge='javascript'>location.assign('" + "http://yahoo.com.tw" + "')</script>"

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, formHtml)

}

// Contact renders the contact page
func (m *Repository) PaymentSuccess(w http.ResponseWriter, r *http.Request) {
	orderNo := r.PostFormValue("MerchantTradeNo")
	paymentAmount, err := strconv.ParseFloat(r.PostFormValue("TradeAmt"), 32)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	data := make(map[string]interface{})
	data["payment"] = models.Payment{
		OrderNo: orderNo,
		Amount:  float32(paymentAmount),
	}

	render.RenderTemplate(w, r, "pay-success.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// Contact renders the contact page
func (m *Repository) ECPayCallback(w http.ResponseWriter, r *http.Request) {
	orderNo := r.PostFormValue("MerchantTradeNo")
	fmt.Println("OrderNo:", orderNo)

	for key, values := range r.PostForm {
		fmt.Println("Key:", key)
		for _, value := range values {
			fmt.Println("Value:", value)
		}
	}
}
