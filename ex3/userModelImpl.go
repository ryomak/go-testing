package main


type db struct{
  users map[uint]user
}


func (d db)getByID(id uint)user{
  return d.users[id]
}
