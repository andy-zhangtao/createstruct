/*
 *
 * Copyright (c) 2017. The createstruct Author ztao8607@gmail.com
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

package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/andy-zhangtao/createstruct/service"
	"github.com/rs/cors"
	"log"
	"net/http"
)


const (
	APIVERSION ="/v1"
	POSTJSON = "/post/json"
	CONNECT = "/_ping"
)
func main(){
	router := httprouter.New()
	router.POST(getAPIPath(POSTJSON), service.GenerateAPI)
	router.GET(getAPIPath(CONNECT), service.Ping)
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func getAPIPath(path string) string{
	log.Println(APIVERSION+path)
	return APIVERSION+path
}

