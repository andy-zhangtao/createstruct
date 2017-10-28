/*
 *
 * Copyright (c) 2017. The gojsonschema Author ztao8607@gmail.com
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *
 */

package service

import (
	"fmt"
	"net/http"

	"encoding/base64"
	"encoding/json"
	"io/ioutil"

	"github.com/andy-zhangtao/gojsonschema/generation"
	"github.com/andy-zhangtao/gojsonschema/parse"
	"github.com/julienschmidt/httprouter"
)

const (
	ERROR     = 500
	JSONEMPTY = "JSON cannot be empty"
	NAMEEMPEY = "NAME cannot be empty"
)

type JsonString struct {
	Name  string `json:"name"`
	Jsons string `json:"json"`
}

func Ping(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Hello")
}

func GenerateAPI(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(ERROR)
		fmt.Fprint(w, err.Error())
	}

	var js JsonString

	err = json.Unmarshal(content, &js)
	if err != nil {
		w.WriteHeader(ERROR)
		fmt.Fprint(w, err.Error())
	}

	if js.Jsons == "" {
		w.WriteHeader(ERROR)
		fmt.Fprintf(w, JSONEMPTY)
		return
	}

	if js.Name == "" {
		w.WriteHeader(ERROR)
		fmt.Fprint(w, NAMEEMPEY)
	}

	jsons, err := base64.StdEncoding.DecodeString(js.Jsons)
	if err != nil {
		w.WriteHeader(ERROR)
		fmt.Fprint(w, err.Error())
	}

	//log.Println(string(jsons))
	mapStruct, mapTag, err := parse.ParseJsonBytes(js.Name, jsons)
	if err != nil {
		w.WriteHeader(ERROR)
		fmt.Fprint(w, err.Error())
	}

	data := generation.Generate(js.Name, mapStruct, mapTag)

	fmt.Fprint(w, data)
}
