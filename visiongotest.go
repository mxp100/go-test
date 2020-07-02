package main

import (
	"./bst"
	"./logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
)

type any interface{}

var tree = &bst.TreeNode{}

func main() {
	err := initTree()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/delete", deleteHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Default port %s", port)
	}

	log.Printf("Open http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), logRequest(http.DefaultServeMux)))
}

func initTree() error {
	jsonFile, err := os.Open("bst.json")
	if err != nil {
		return err
	}
	log.Println("File bst.json opened")
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var vars []int
	err = json.Unmarshal(byteValue, &vars)
	if err != nil {
		return err
	}

	for i := 0; i < len(vars); i++ {
		if i == 0 {
			tree = &bst.TreeNode{Value: vars[0]}
		} else {
			err = tree.Insert(vars[i])
			if err != nil {
				return err
			}
		}
	}
	log.Println("Tree loaded")
	return nil
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.JSON(map[string]any{
			"Type":       "Request",
			"RemoteAddr": r.RemoteAddr,
			"Method":     r.Method,
			"Url":        r.URL.String(),
		})

		record := httptest.NewRecorder()
		handler.ServeHTTP(record, r)

		for k, v := range record.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(record.Code)

		body, _ := ioutil.ReadAll(record.Body)

		if _, err := w.Write(body); err != nil {
			log.Fatal(err)
		}

		logger.JSON(map[string]any{
			"Type":      "Response",
			"Code":      record.Code,
			"HeaderMap": record.Header(),
			"Body":      string(body),
		})
	})
}

func responseJSON(w http.ResponseWriter, data any) {
	w.Header().Add("Content-Type", "application/json")
	bytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
	}
	_, err = w.Write(bytes)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var val = r.URL.Query().Get("val")
	if searchVal, err := strconv.Atoi(val); err == nil {
		tree, found := tree.Find(searchVal)
		if found {
			responseJSON(w, map[string]any{
				"Found": found,
				"Tree":  tree,
			})
		} else {
			responseJSON(w, map[string]any{
				"Found": found,
				"Tree":  tree,
			})
		}
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var vars []int

	err := json.NewDecoder(r.Body).Decode(&vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(vars); i++ {
		err = tree.Insert(vars[i])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.WriteHeader(http.StatusCreated)
	responseJSON(w, map[string]any{
		"Inserted": vars,
		"Tree":     tree,
	})
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	val := r.URL.Query().Get("val")
	if deleteVal, err := strconv.Atoi(val); err == nil {
		tree.Remove(deleteVal)

		responseJSON(w, map[string]any{
			"Deleted": deleteVal,
			"Tree":    tree,
		})
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
