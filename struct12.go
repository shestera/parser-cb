package main

import "encoding/xml"

// sbFile12 Автосгенерированная схема
type sbFile12 struct {
	XMLName   xml.Name `xml:"СообщОтказ_115ФЗ"`
	Text      string   `xml:",chardata"`
	СлужЧасть struct {
		Text             string `xml:",chardata"`
		ВерсияФормата    string `xml:"ВерсияФормата"`
		ДатаСообщения    string `xml:"ДатаСообщения"`
		ТелОператор      string `xml:"ТелОператор"`
		ЭлектроннаяПочта string `xml:"ЭлектроннаяПочта"`
	} `xml:"СлужЧасть"`
	ИнформЧасть struct {
		Text           string `xml:",chardata"`
		ТипОрганизация string `xml:"ТипОрганизация"`
		Раздел2        []struct {
			Text         string `xml:",chardata"`
			НомерЗаписи  string `xml:"НомерЗаписи"`
			ФорматЗаписи string `xml:"ФорматЗаписи"`
			ТипЗаписи    string `xml:"ТипЗаписи"`
			СведОперация struct {
				Text        string `xml:",chardata"`
				ПризнакИнф  string `xml:"ПризнакИнф"`
				КодОтказа   string `xml:"КодОтказа"`
				ДатаОтказа  string `xml:"ДатаОтказа"`
				ОснованиеОп []struct {
					Text        string `xml:",chardata"`
					КодДок      string `xml:"КодДок"`
					НаимДок     string `xml:"НаимДок"`
					ИноеНаимДок string `xml:"ИноеНаимДок"`
					ДатаДок     string `xml:"ДатаДок"`
					НомДок      string `xml:"НомДок"`
				} `xml:"ОснованиеОп"`
				КодПризнОперации  string   `xml:"КодПризнОперации"`
				КодДенежСредств   string   `xml:"КодДенежСредств"`
				ПризнНеобОперации []string `xml:"ПризнНеобОперации"`
			} `xml:"СведОперация"`
			Участник []struct {
				Text            string `xml:",chardata"`
				СтатусУчастника string `xml:"СтатусУчастника"`
				ТипУчастника    string `xml:"ТипУчастника"`
				ПризнУчастника  string `xml:"ПризнУчастника"`
				СведФЛИП        struct {
					Text    string `xml:",chardata"`
					ФИОФЛИП struct {
						Text string `xml:",chardata"`
						Фам  string `xml:"Фам"`
						Имя  string `xml:"Имя"`
						Отч  string `xml:"Отч"`
					} `xml:"ФИОФЛИП"`
					ИННФЛИП       string `xml:"ИННФЛИП"`
					СведДокУдЛичн struct {
						Text               string `xml:",chardata"`
						ВидДокКод          string `xml:"ВидДокКод"`
						ВидДокНаименование string `xml:"ВидДокНаименование"`
						СерияДок           string `xml:"СерияДок"`
						НомДок             string `xml:"НомДок"`
						ДатВыдачиДок       string `xml:"ДатВыдачиДок"`
						КемВыданДок        string `xml:"КемВыданДок"`
						КодПодр            string `xml:"КодПодр"`
					} `xml:"СведДокУдЛичн"`
					ДатаРождения string `xml:"ДатаРождения"`
					МестоРожд    struct {
						Text               string `xml:",chardata"`
						КодОКСМ            string `xml:"КодОКСМ"`
						СтранаНаименование string `xml:"СтранаНаименование"`
						КодСубъектаПоОКАТО string `xml:"КодСубъектаПоОКАТО"`
						Район              string `xml:"Район"`
						Пункт              string `xml:"Пункт"`
					} `xml:"МестоРожд"`
					КодОКСМ            string `xml:"КодОКСМ"`
					СтранаНаименование string `xml:"СтранаНаименование"`
					ПризнакПубЛицо     string `xml:"ПризнакПубЛицо"`
				} `xml:"СведФЛИП"`
				СведЮЛ struct {
					Text     string `xml:",chardata"`
					НаимЮЛ   string `xml:"НаимЮЛ"`
					ИННЮЛ    string `xml:"ИННЮЛ"`
					КППЮЛ    string `xml:"КППЮЛ"`
					ОГРНЮЛ   string `xml:"ОГРНЮЛ"`
					АдрРегЮЛ struct {
						Text               string `xml:",chardata"`
						КодОКСМ            string `xml:"КодОКСМ"`
						СтранаНаименование string `xml:"СтранаНаименование"`
						КодСубъектаПоОКАТО string `xml:"КодСубъектаПоОКАТО"`
						Район              string `xml:"Район"`
						Пункт              string `xml:"Пункт"`
						Улица              string `xml:"Улица"`
						Дом                string `xml:"Дом"`
						Корп               string `xml:"Корп"`
						Оф                 string `xml:"Оф"`
					} `xml:"АдрРегЮЛ"`
					АдрЮЛ struct {
						Text               string `xml:",chardata"`
						КодОКСМ            string `xml:"КодОКСМ"`
						СтранаНаименование string `xml:"СтранаНаименование"`
						КодСубъектаПоОКАТО string `xml:"КодСубъектаПоОКАТО"`
						Район              string `xml:"Район"`
						Пункт              string `xml:"Пункт"`
						Улица              string `xml:"Улица"`
						Дом                string `xml:"Дом"`
						Корп               string `xml:"Корп"`
						Оф                 string `xml:"Оф"`
					} `xml:"АдрЮЛ"`
				} `xml:"СведЮЛ"`
			} `xml:"Участник"`
		} `xml:"Раздел2"`
	} `xml:"ИнформЧасть"`
}
