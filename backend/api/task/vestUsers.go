package task

import (
	"bytes"
)

var vests = [100]string{
	"17600000001",
	"17600000002",
	"17600000003",
	"17600000004",
	"17600000005",
	"17600000006",
	"17600000007",
	"17600000008",
	"17600000009",
	"17600000010",
	"17600000011",
	"17600000012",
	"17600000013",
	"17600000014",
	"17600000015",
	"17600000016",
	"17600000017",
	"17600000018",
	"17600000019",
	"17600000020",
	"17600000021",
	"17600000022",
	"17600000023",
	"17600000024",
	"17600000029",
	"17600000026",
	"17600000027",
	"17600000028",
	"17600000029",
	"17600000030",
	"17600000031",
	"17600000032",
	"17600000033",
	"17600000034",
	"17600000035",
	"17600000036",
	"17600000037",
	"17600000038",
	"17600000039",
	"17600000040",
	"17600000041",
	"17600000042",
	"17600000043",
	"17600000044",
	"17600000045",
	"17600000046",
	"17600000047",
	"17600000048",
	"17600000049",
	"17600000050",
	"17600000051",
	"17600000052",
	"17600000053",
	"17600000054",
	"17600000055",
	"17600000056",
	"17600000057",
	"17600000058",
	"17600000059",
	"17600000060",
	"17600000061",
	"17600000062",
	"17600000063",
	"17600000064",
	"17600000065",
	"17600000066",
	"17600000067",
	"17600000068",
	"17600000069",
	"17600000070",
	"17600000071",
	"17600000072",
	"17600000073",
	"17600000074",
	"17600000075",
	"17600000076",
	"17600000077",
	"17600000078",
	"17600000079",
	"17600000080",
	"17600000081",
	"17600000082",
	"17600000083",
	"17600000084",
	"17600000085",
	"17600000086",
	"17600000087",
	"17600000088",
	"17600000089",
	"17600000090",
	"17600000091",
	"17600000092",
	"17600000093",
	"17600000094",
	"17600000095",
	"17600000096",
	"17600000097",
	"17600000098",
	"17600000099",
	"17600000100",
}
// ('','','')

func GetVestStr() string {
	var buffer bytes.Buffer
	buffer.WriteString("('")
	for k, v := range vests {
		buffer.WriteString(v)
		if k < len(vests)-1 {
			buffer.WriteString("','")
		} else {
			buffer.WriteString("'")
		}
	}
	buffer.WriteString(")")
	return buffer.String()
}