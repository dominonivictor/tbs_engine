package utils

import (
    "encoding/csv"
    "log"
    "os"
    comp "github.com/dominonivictor/tbs_engine/components"
)

func Read_csv(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func Load_materials() map[comp.MAT_ID]comp.Material{
  materials_csv := Read_csv("./data/materials.csv")
  materials := map[comp.MAT_ID]comp.Material{}
  for _, mat := range materials_csv {
    id := comp.MAT_ID(mat[0])
   materials[id] = comp.NewMaterialFromCSV(mat...)
  }
  return materials
}
