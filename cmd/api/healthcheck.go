package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// 1、使用fmt.Fprintf()函数
	// fmt.Fprintln(w, "status: aviable")
	// fmt.Fprintf(w, "environment : %s\n",app.config.env)
	// fmt.Fprintf(w, "version:%s\n", version)

	// 	2、使用字符串转换为JSON
	// js := `{"status":"available", "environment": %q, "version": %q}`
	// js = fmt.Sprintf(js, app.config.env, version)
	// w.Header().Set("Content-Type","application/json")

	// w.Write([]byte(js))

	// 3、使用json.Marshal()将Go的对象转化为JSON
	env := envelope{
		"status": "avaiable",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writerJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.logger.PrintError(err, nil)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}
