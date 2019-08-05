package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/tealeg/xlsx"
)

const reportPach = "report"

// ReportList список данных для отчета
type ReportList struct {
	Reports []Report
}

// Report данные для отчета
type Report struct {
	ДатаСообщения     string
	КодОтказа         string
	ПризнНеобОперации []string
	СтатусУчастника   string
	Фам               string
	Имя               string
	Отч               string
	ДатаРождения      string
}

func main() {
	reportListgo := ReportList{}

	files, err := ioutil.ReadDir(reportPach)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		xmlFile, err := os.Open(path.Join(reportPach, file.Name()))
		if err != nil {
			log.Println(err)
		}
		defer xmlFile.Close()

		byteValue, _ := ioutil.ReadAll(xmlFile)
		var message sbFile12
		xml.Unmarshal(byteValue, &message)

		if message.СлужЧасть.ВерсияФормата == "1.2" {
			for i := 0; i < len(message.ИнформЧасть.Раздел2); i++ {
				for p := 0; p < len(message.ИнформЧасть.Раздел2[i].Участник); p++ {
					if message.ИнформЧасть.Раздел2[i].Участник[p].ТипУчастника == "2" {
						var r Report
						r.parse(message, i, p)
						reportListgo.addItem(r)
					}
				}
			}
		} else {
			log.Println(file.Name(), "файл версия не 1.2")
		}
	}
	reportListgo.createXls()
}

func (l *ReportList) addItem(r Report) {
	l.Reports = append(l.Reports, r)
}

func (l *ReportList) createXls() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	head := []string{"ДатаСообщения", "КодОтказа", "ПризнНеобОперации", "СтатусУчастника", "Фам", "Имя", "Отч", "ДатаРождения"}
	for _, i := range head {
		cell = row.AddCell()
		cell.Value = i
	}

	for _, item := range l.Reports {
		row = sheet.AddRow()
		values := []string{item.ДатаСообщения, item.КодОтказа, strings.Join(item.ПризнНеобОперации, ", "), item.СтатусУчастника, item.Фам, item.Имя, item.Отч, item.ДатаРождения}
		for _, i := range values {
			cell = row.AddCell()
			cell.Value = i
		}
	}

	err = file.Save("result.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func (r *Report) parse(m sbFile12, section int, participant int) {
	r.ДатаСообщения = m.СлужЧасть.ДатаСообщения
	r.КодОтказа = m.ИнформЧасть.Раздел2[section].СведОперация.КодОтказа
	r.ПризнНеобОперации = m.ИнформЧасть.Раздел2[section].СведОперация.ПризнНеобОперации
	r.СтатусУчастника = m.ИнформЧасть.Раздел2[section].Участник[participant].СтатусУчастника
	r.Фам = m.ИнформЧасть.Раздел2[section].Участник[participant].СведФЛИП.ФИОФЛИП.Фам
	r.Имя = m.ИнформЧасть.Раздел2[section].Участник[participant].СведФЛИП.ФИОФЛИП.Имя
	r.Отч = m.ИнформЧасть.Раздел2[section].Участник[participant].СведФЛИП.ФИОФЛИП.Отч
	r.ДатаРождения = m.ИнформЧасть.Раздел2[section].Участник[participant].СведФЛИП.ДатаРождения
}
