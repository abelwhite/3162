package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/abelwhite/3162/test-1/internal/models"
)

// shared data store
var dataStore = struct {
	sync.RWMutex
	data map[string]int64
}{data: make(map[string]int64)}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	// remove the entry from the session manager
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash: flash,
	}
	RenderTemplate(w, "login.page.tmpl", data)
}

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	// remove the entry from the session manager
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash: flash,
	}
	RenderTemplate(w, "signup.page.tmpl", data)

}

func (app *application) dashboard(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash: flash,
	}
	RenderTemplate(w, "dashboard.page.tmpl", data)
}

func (app *application) settings(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash: flash,
	}
	RenderTemplate(w, "settings.page.tmpl", data)
}

func (app *application) profile(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash: flash,
	}
	RenderTemplate(w, "profile.page.tmpl", data)
}

// ------------ROOM CRUD HANDLERS---------------------
func (app *application) roomCreateShow(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash: flash,
	}
	RenderTemplate(w, "roomform.page.tmpl", data)
}

func (app *application) roomShow(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash: flash,
	}
	RenderTemplate(w, "viewroom.page.tmpl", data)
}

// ------------Pigsty (pen) CRUD HANDLERS---------------------
func (app *application) pigstyCreateShow(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash: flash,
	}
	RenderTemplate(w, "pigstyform.page.tmpl", data)
}

func (app *application) pigstyShow(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash: flash,
	}
	RenderTemplate(w, "viewpigsty.page.tmpl", data)
}

// ------------Pig CRUD HANDLERS---------------------
func (app *application) pigCreateShow(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash: flash,
	}
	RenderTemplate(w, "pigform.page.tmpl", data)
}

func (app *application) pigCreateSubmit(w http.ResponseWriter, r *http.Request) {
	//get the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	room := r.PostForm.Get("room") //insert question into the database
	pigsty := r.PostForm.Get("pigsty")
	breed := r.PostForm.Get("breed")
	age := r.PostForm.Get("age")
	//dob := r.Time.ParseTime.Get("dob")

	dobStr := r.PostForm.Get("dob")
	layout := "2006-01-02" // the layout string for the date format, e.g. "2006-01-02"

	dob, err := time.Parse(layout, dobStr)
	if err != nil {
		log.Println("Error parsing date:", err)
	}

	weight := r.PostForm.Get("weight")
	gender := r.PostForm.Get("gender")
	log.Printf("%s %s %s %s %s %s %s\n", room, pigsty, breed, age, dob, weight, gender)
	id, err := app.pig.Insert(room, pigsty, breed, age, dob, weight, gender)
	log.Printf("%s %s %s %s %s %s %s %d\n", room, pigsty, breed, age, dob, weight, gender, id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/pig/show", http.StatusSeeOther)
}
func (app *application) pigShow(w http.ResponseWriter, r *http.Request) {
	// pig_id := 1
	q, err := app.pig.Read()
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//an instance of templateData
	data := &templateData{
		Pig: q,
	} //this allows us to pass in mutliple data into the template

	//display pigs using tmpl
	ts, err := template.ParseFiles("./ui/html/viewpigs.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//assuming no error
	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}

}

func (app *application) pigDelete(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the quote ID from the URL query parameters
	quoteIDStr := r.URL.Query().Get("pig_id")

	// Convert the quote ID string to an integer
	quoteID, err := strconv.Atoi(quoteIDStr)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	// Call the Delete method to remove the quote from the database
	err = app.pig.Delete(quoteID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect the user back to the quote list page
	http.Redirect(w, r, "/pig/show", http.StatusSeeOther)

}

func (app *application) pigUpdate(w http.ResponseWriter, r *http.Request) {
	pig_id := 1
	q, err := app.pig.Read() //padd pig_id from read, (needs to pass pig_id)

	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//an instance of templateData

	data := &templateData{
		Pig: q,
	} //this allows us to pass in mutliple data into the template

	//display quotes using tmpl
	ts, err := template.ParseFiles("./ui/html/pigform.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//auming no error
	dataStore.Lock()
	dataStore.data["key"] = int64(pig_id)
	dataStore.Unlock()

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}

func (app *application) pigUpdateQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	//get the form data

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return

	}

	//collect values from pigs
	room := r.PostForm.Get("room") //insert question into the database
	pigsty := r.PostForm.Get("pigsty")
	breed := r.PostForm.Get("breed")
	age := r.PostForm.Get("age")
	//dob := r.Time.ParseTime.Get("dob")

	dobStr := r.PostForm.Get("dob")
	layout := "2006-01-02" // the layout string for the date format, e.g. "2006-01-02"

	dob, err := time.Parse(layout, dobStr)
	if err != nil {
		log.Println("Error parsing date:", err)
	}

	weight := r.PostForm.Get("weight")
	gender := r.PostForm.Get("gender")

	data := &models.Pig{
		Room:   room,
		Pigsty: pigsty,
		Breed:  breed,
		Age:    age,
		Dob:    dob,
		Weight: weight,
		Gender: gender,
	}
	Test, err := app.pig.Update(data)
	fmt.Println(Test)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// func (app *application) pigUpdate(w http.ResponseWriter, r *http.Request) {
// 	// Parse the form data
// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, "bad request", http.StatusBadRequest)
// 		return
// 	}

// 	// Retrieve the pig ID from the URL query parameters
// 	pigIDStr := r.URL.Query().Get("pig_id")

// 	// Convert the pig ID string to an integer
// 	pigID, err := strconv.Atoi(pigIDStr)
// 	if err != nil {
// 		log.Println(err.Error())
// 		http.Error(w,
// 			http.StatusText(http.StatusBadRequest),
// 			http.StatusBadRequest)
// 		return
// 	}

// 	// Retrieve the pig data from the form data
// 	room := r.PostForm.Get("room")
// 	pigsty := r.PostForm.Get("pigsty")
// 	breed := r.PostForm.Get("breed")
// 	age := r.PostForm.Get("age")
// 	dobStr := r.PostForm.Get("dob")
// 	layout := "2006-01-02" // the layout string for the date format, e.g. "2006-01-02"
// 	dob, err := time.Parse(layout, dobStr)
// 	if err != nil {
// 		log.Println("Error parsing date:", err)
// 	}
// 	weight := r.PostForm.Get("weight")
// 	gender := r.PostForm.Get("gender")

// 	// Create a pig object with the updated data
// 	weightVal, err := strconv.ParseFloat(weight, 64)
// 	if err != nil {
// 		log.Println("Error parsing weight:", err)
// 		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
// 		return
// 	}
// 	updatedPig := &models.Pig{
// 		PigID:  int64(pigID),
// 		Room:   room,
// 		Pigsty: pigsty,
// 		Breed:  breed,
// 		Age:    age,
// 		Dob:    dob,
// 		Weight: weightVal,
// 		Gender: gender,
// 	}

// 	// Call the Update method to update the pig in the database
// 	err = app.pig.Update(updatedPig)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Redirect the user back to the pig list page
// 	http.Redirect(w, r, "/pig/show", http.StatusSeeOther)
// }
