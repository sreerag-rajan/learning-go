package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model for courses - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice float64 `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

//fake db

var courses []Course

//middlewares, helper -file
func (c *Course) isEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Build API in golang")

	courses = append(courses,
		Course{CourseId: "2", CourseName: "ReactJs", CoursePrice: 299, Author: &Author{FullName: "Sreerag Rajan", Website: "sr.learning.com"}},
		Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{FullName: "Sreerag Rajan", Website: "sr.learning.com"}},
	)
	// courses = append(courses, Course{CourseId: "2",CourseName: "ReactJs", CoursePrice: 299, Author: &Author{FullName: "Sreerag Rajan", Website: "sr.learning.com"}})
	// courses = append(courses, Course{CourseId: "2",CourseName: "ReactJs", CoursePrice: 299, Author: &Author{FullName: "Sreerag Rajan", Website: "sr.learning.com"}})

	r := mux.NewRouter()

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", addCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PATCH")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}

// controller - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by learn code online</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop through courses to find the matching id amd return response
	fmt.Printf("%v/n", params["id"])

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("No Course Found for the given id")
	return
}

func addCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create a course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	fmt.Printf("%#v", course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	for _, c := range courses {
		if c.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("A course with this name already exists")
			return
		}
	}

	// generate a unique id, convert to string
	// add or append this course to new course

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(courses)
	return

}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	if r.Body == nil {
		json.NewEncoder(w).Encode("Body is empty")
	}

	// loop, id, remove, add with same id
	for i, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			var newcourse Course
			_ = json.NewDecoder(r.Body).Decode(&newcourse)
			newcourse.CourseId = params["id"]
			courses = append(courses, newcourse)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}
	json.NewEncoder(w).Encode("Course with the given id was not found")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Course")

	params := mux.Vars(r)

	for i, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(courses)
	return
}
