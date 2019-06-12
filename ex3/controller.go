package main

import(
  "net/http"
  "strconv"
  "encoding/json"
)

var d = db{
  users: map[uint]user{
    1:{1,"test1",13},
    2:{2,"test2",15},
    3:{3,"test3",20},
  },
}

func userHandler(w http.ResponseWriter, r *http.Request){
  id ,err:= strconv.Atoi(r.URL.Query()["id"][0])
  if err != nil{
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("not found"))
    return
  }
  user := d.getByID(uint(id))
  w.WriteHeader(http.StatusOK)

  json.NewEncoder(w).Encode(user)
}
