package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/abelwhite/pigstydash/internal/models"
	"github.com/justinas/nosurf"
)

// shared data store
var dataStore = struct {
	sync.RWMutex
	data map[string]int64
}{data: make(map[string]int64)}

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	// remove the entry from the session manager
	flash := app.sessionManager.PopString(r.Context(), "flash")
	data := &templateData{ //putting flash into template data
		Flash:     flash,
		CSRFToken: nosurf.Token(r),
	}
	RenderTemplate(w, "signup.page.tmpl", data)

}

func (app *application) signupSubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.PostForm.Get("name") //"name" is the name of the form
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")

	// // check the web form fields to validity
	// errors_user := make(map[string]string)

	// if strings.TrimSpace(name) == "" {
	// 	errors_user["name"] = "This field cannot be left blank"
	// } else if utf8.RuneCountInString(name) > 50 {
	// 	errors_user["name"] = "This field is too long(maximum is 50 characters)"
	// } else if utf8.RuneCountInString(name) < 5 {
	// 	errors_user["name"] = "This field is too short(minimum is 5 characters)"
	// }
	// if strings.TrimSpace(email) == "" {
	// 	errors_user["email"] = "This field cannot be left blank"
	// } else if utf8.RuneCountInString(email) > 50 {
	// 	errors_user["email"] = "This field is too long(maximum is 50 characters)"
	// }

	// // Check to see if the email is valid
	// emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	// if !emailRegex.MatchString(email) {
	// 	errors_user["email"] = "Invalid email"
	// }

	// // Check the length of the password
	// if strings.TrimSpace(password) == "" {
	// 	errors_user["password"] = "This field cannot be left blank"
	// } else if utf8.RuneCountInString(password) > 50 {
	// 	errors_user["password"] = "This field is too long(maximum is 50 characters)"
	// } else if utf8.RuneCountInString(password) < 10 {
	// 	errors_user["password"] = "This field is too short(minimum is 10 characters)"
	// }

	// // check if there are any errors in the map
	// if len(errors_user) > 0 {
	// 	ts, err := template.ParseFiles("./ui/html/signup.page.tmpl")
	// 	if err != nil {
	// 		app.serverError(w, err)
	// 		return
	// 	}
	// 	err = ts.Execute(w, &templateData{
	// 		ErrorsFromForm:  errors_user,
	// 		FormData:        r.PostForm,
	// 		IsAuthenticated: app.isAuthenticated(r),
	// 	CSRFToken:       nosurf.Token(r),
	// 	})

	// 	if err != nil {
	// 		app.serverError(w, err)
	// 		return
	// 	}
	// 	return
	// }

	//lets write the data to the table
	err := app.user.Insert(name, email, password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			RenderTemplate(w, "signup.page.tmpl", nil)
		}
	}
	app.sessionManager.Put(r.Context(), "flash", "Signup Was successful")
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	// remove the entry from the session manager
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash:     flash,
		CSRFToken: nosurf.Token(r),
	}
	RenderTemplate(w, "login.page.tmpl", data)
}

func (app *application) loginSubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	//lets write the data to the table
	id, err := app.user.Authenticate(email, password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			flash := app.sessionManager.PopString(r.Context(), "flash")
			//render
			data := &templateData{ //putting flash into template data
				Flash:     flash,
				CSRFToken: nosurf.Token(r),
			}
			RenderTemplate(w, "login.page.tmpl", data)
		}
		return
	}
	//add the users to the session cookie
	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		return
	}
	//add an authenticate entry
	app.sessionManager.Put(r.Context(), "authenticatedUserID", id)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

}

func (app *application) logoutSubmit(w http.ResponseWriter, r *http.Request) {
	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		return
	}
	app.sessionManager.Remove(r.Context(), "authenticatedUserID")
	http.Redirect(w, r, "/login", http.StatusSeeOther)

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

func (app *application) temperature(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash: flash,
	}
	RenderTemplate(w, "temperature.page.tmpl", data)
}

// ------------ROOM CRUD HANDLERS---------------------
func (app *application) roomCreateShow(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash:     flash,
		CSRFToken: nosurf.Token(r),
	}
	RenderTemplate(w, "roomform.page.tmpl", data)
}

func (app *application) roomCreateSubmit(w http.ResponseWriter, r *http.Request) {
	//get the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	name := r.PostForm.Get("name")
	num_of_pigsty, _ := strconv.ParseInt(r.PostForm.Get("num_of_pigsty"), 10, 64)
	temperature := r.PostForm.Get("temperature")
	humidity := r.PostForm.Get("humidity")

	log.Printf("%s %d %s %s\n", name, num_of_pigsty, temperature, humidity)
	id, err := app.room.Insert(name, num_of_pigsty, temperature, humidity)
	log.Printf("%s %d %s %s %d\n", name, num_of_pigsty, temperature, humidity, id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/room/show", http.StatusSeeOther)
}

func (app *application) roomShow(w http.ResponseWriter, r *http.Request) {
	q, err := app.room.Read()
	p, _ := app.pigsty.Read()
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//an instance of templateData
	data := &templateData{
		Room:   q,
		Pigsty: p,
	} //this allows us to pass in mutliple data into the template

	//display pigs using tmpl
	ts, err := template.ParseFiles("./ui/html/viewroom.page.tmpl")
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

func (app *application) roomDelete(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the room ID from the URL query parameters
	roomIDStr := r.URL.Query().Get("room_id")
	// Convert the room ID string to an integer
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	// Call the Delete method to remove the rrom from the database
	err = app.room.Delete(roomID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect the user back to the room list page
	http.Redirect(w, r, "/room/show", http.StatusSeeOther)

}

func (app *application) roomUpdate(w http.ResponseWriter, r *http.Request) {
	roomIDStr := r.URL.Query().Get("room_id")
	// Convert the pig ID string to an integer
	roomID, _ := strconv.Atoi(roomIDStr)
	q, err := app.room.Readd(roomID) //padd room_id from read, (needs to pass pig_id)

	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//an instance of templateData
	data := &templateData{
		Room:      q,
		CSRFToken: nosurf.Token(r),
	} //this allows us to pass in mutliple data into the template

	//display pigform using tmpl
	ts, err := template.ParseFiles("./ui/html/roomformupdate.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//assuming no error
	dataStore.Lock()
	dataStore.data["key"] = int64(roomID)
	dataStore.Unlock()

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}

func (app *application) roomUpdateQuery(w http.ResponseWriter, r *http.Request) {
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
	name := r.PostForm.Get("name")
	num_of_pigsty, _ := strconv.ParseInt(r.PostForm.Get("num_of_pigsty"), 10, 64)
	// num_of_pigsty := r.PostForm.Get("num_of_pigsty")
	temperature := r.PostForm.Get("temperature")
	humidity := r.PostForm.Get("humidity")

	dataStore.RLock()
	room_id := dataStore.data["key"]
	dataStore.RUnlock()

	data := &models.Room{
		RoomID:        room_id,
		Name:          name,
		NumOfPigsties: num_of_pigsty,
		Temperature:   temperature,
		Humidity:      humidity,
	}
	Test, err := app.room.Update(data)
	fmt.Println(Test)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/room/show", http.StatusSeeOther)

}

// ------------Pigsty (pen) CRUD HANDLERS---------------------
func (app *application) pigstyCreateShow(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash:     flash,
		CSRFToken: nosurf.Token(r),
	}
	RenderTemplate(w, "pigstyform.page.tmpl", data)
}

func (app *application) pigstyCreateSubmit(w http.ResponseWriter, r *http.Request) {
	//get the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	room := r.PostForm.Get("room")
	name := r.PostForm.Get("name")
	num_of_pigs, _ := strconv.ParseInt(r.PostForm.Get("num_of_pigs"), 10, 64)

	log.Printf("%s %s %d \n", room, name, num_of_pigs)
	id, err := app.pigsty.Insert(room, name, num_of_pigs)
	log.Printf("%s %s %d %d\n", room, name, num_of_pigs, id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/pigsty/show", http.StatusSeeOther)
}

func (app *application) pigstyShow(w http.ResponseWriter, r *http.Request) {
	q, err := app.pigsty.Read()
	p, _ := app.pig.Read()
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//an instance of templateData
	data := &templateData{
		Pigsty:    q,
		Pig:       p,
		CSRFToken: nosurf.Token(r),
	} //this allows us to pass in mutliple data into the template

	//display pigsty using tmpl
	ts, err := template.ParseFiles("./ui/html/viewpigsty.page.tmpl")
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

func (app *application) pigstyDelete(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the pig ID from the URL query parameters
	pigstyIDStr := r.URL.Query().Get("pigsty_id")
	// Convert the pig ID string to an integer
	pigstyID, err := strconv.Atoi(pigstyIDStr)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	// Call the Delete method to remove the pig from the database
	err = app.pigsty.Delete(pigstyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect the user back to the pig list page
	http.Redirect(w, r, "/pigsty/show", http.StatusSeeOther)

}

func (app *application) pigstyUpdate(w http.ResponseWriter, r *http.Request) {
	pigstyIDStr := r.URL.Query().Get("pigsty_id")

	// Convert the pig ID string to an integer
	pigstyID, _ := strconv.Atoi(pigstyIDStr)
	q, err := app.pigsty.Readd(pigstyID) //padd pig_id from read, (needs to pass pig_id)

	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//an instance of templateData

	data := &templateData{
		Pigsty:    q,
		CSRFToken: nosurf.Token(r),
	} //this allows us to pass in mutliple data into the template

	//display pigform using tmpl
	ts, err := template.ParseFiles("./ui/html/pigstyformupdate.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//assuming no error
	dataStore.Lock()
	dataStore.data["key"] = int64(pigstyID)
	dataStore.Unlock()

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}

func (app *application) pigstyUpdateQuery(w http.ResponseWriter, r *http.Request) {
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
	room := r.PostForm.Get("room")
	name := r.PostForm.Get("name")
	num_of_pigs, _ := strconv.ParseInt(r.PostForm.Get("num_of_pigs"), 10, 64)

	dataStore.RLock()
	pigsty_id := dataStore.data["key"]
	dataStore.RUnlock()

	data := &models.Pigsty{
		PigstyID:  pigsty_id,
		Room:      room,
		Name:      name,
		NumOfPigs: num_of_pigs,
	}
	Test, err := app.pigsty.Update(data)
	fmt.Println(Test)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/pigsty/show", http.StatusSeeOther)

}

// ------------Pig CRUD HANDLERS---------------------
func (app *application) pigCreateShow(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{ //putting flash into template data
		Flash:     flash,
		CSRFToken: nosurf.Token(r),
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
	dob, _ := time.Parse("2006-01-02", r.PostForm.Get("dob"))
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
		Pig:       q,
		CSRFToken: nosurf.Token(r),
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

	// Retrieve the pig ID from the URL query parameters
	pigIDStr := r.URL.Query().Get("pig_id")
	// Convert the pig ID string to an integer
	pigID, err := strconv.Atoi(pigIDStr)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	// Call the Delete method to remove the pig from the database
	err = app.pig.Delete(pigID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect the user back to the pig list page
	http.Redirect(w, r, "/pig/show", http.StatusSeeOther)

}

func (app *application) pigUpdate(w http.ResponseWriter, r *http.Request) {
	// pig_id := 1

	pigIDStr := r.URL.Query().Get("pig_id")

	// Convert the pig ID string to an integer
	pigID, _ := strconv.Atoi(pigIDStr)
	q, err := app.pig.Readd(pigID) //padd pig_id from read, (needs to pass pig_id)

	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//an instance of templateData

	data := &templateData{
		Pig:       q,
		CSRFToken: nosurf.Token(r),
	} //this allows us to pass in mutliple data into the template

	//display pigform using tmpl
	ts, err := template.ParseFiles("./ui/html/pigformupdate.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//assuming no error
	dataStore.Lock()
	dataStore.data["key"] = int64(pigID)
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
	room := r.PostForm.Get("room")
	pigsty := r.PostForm.Get("pigsty")
	breed := r.PostForm.Get("breed")
	age := r.PostForm.Get("age")
	dob, _ := time.Parse("2006-01-02", r.PostForm.Get("dob"))

	// dobStr := r.PostForm.Get("dob")
	// layout := "2006-01-02" // the layout string for the date format, e.g. "2006-01-02"

	// dob, err := time.Parse(layout, dobStr)
	// if err != nil {
	// 	log.Println("Error parsing date:", err)
	// }

	weight := r.PostForm.Get("weight")
	gender := r.PostForm.Get("gender")

	dataStore.RLock()
	pig_id := dataStore.data["key"]
	dataStore.RUnlock()

	data := &models.Pig{
		PigID:  pig_id,
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
	http.Redirect(w, r, "/pig/show", http.StatusSeeOther)

}
