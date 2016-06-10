package txtparser

import(
  "encoding/csv"
  "os"
)

type SignatureDB struct{
  Description string
  Signature string
  Extenion string
}

func loadDBtoMemory() (error , [][]string) {
  file,err:=os.Open("FileSigs_20151213/customsigs_GCK.csv")
  defer file.Close()
  if err!=nil{
    return err , make([][]string,0)
  }
  r := csv.NewReader(file)
  r.Comma = ','

  records, err := r.ReadAll()

  if err != nil {
    return err , make([][]string,0)
  }
  //fmt.Println(reflect.TypeOf(records))
  return nil , records
}

func GetFileSigFromMem(fs string) (error , string){
  err, data:=loadDBtoMemory()
  var d string
  if err !=nil{
    return err, ""
  }

  for i:=0;i<len(data); i++{
    for x:=0; x<len(data[i]);x++{
      //fmt.Println(data[i][1])
      if fs==data[i][1]{
        d=data[i][2]
        break
      }
    }
  }
  return nil, d
}
