package txtparser

import(
  "encoding/csv"
  "os"
)

type SignatureInfo struct{
  Description string
  Signature string
  Extenion string
}

type SigDB struct{
  db []SignatureInfo
}

var signatrueInfos []SignatureInfo

func loadDBtoMemory() (error , [][]string) {
  file,err:=os.Open("db/customsigs_GCK.csv")
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
  return nil , records
}

func prepareSigDBWithData() (error , *SigDB){
  err, data:=loadDBtoMemory()
  if err !=nil{
    return err, nil
  }

  signatrueInfos = make([]SignatureInfo , len(data))

  for i:=0;i<len(data); i++{
    for x:=0; x<len(data[i]);x++{
      signatrueInfos[i].Description=data[i][0]
      signatrueInfos[i].Signature=data[i][1]
      signatrueInfos[i].Extenion=data[i][2]
    }
  }
  return nil, &SigDB{
    db:signatrueInfos,
  }
}

func GetFileSigFromMem(fs string) (error , string){
  err,sData:=prepareSigDBWithData()
  if err !=nil{
    return err, ""
  }
  var d string
  for i:=0;i<len(sData.db); i++{
    if fs==sData.db[i].Signature{
      d=sData.db[i].Extenion
      break
    }
  }
  return nil, d
}
