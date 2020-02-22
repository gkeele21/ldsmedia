package main

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbstandardworkchapter"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbstandardworkverse"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbtargetentity"
)

type Book struct {
	StandardWorkID int64
	BookName       string
	DisplayOrder   int64
	ChapterTitle   string
	NumChapters    int64
}

func main() {

	m := make(map[int]int)
	bookName := "Art"
	bookId := int64(45278)
	m[1] = 13
	//m[2] = 25
	//m[3] = 28
	//m[4] = 31
	//m[5] = 21
	//m[6] = 68
	//m[7] = 69
	//m[8] = 30
	//m[9] = 14
	//m[10] = 70
	//m[11] = 30
	//m[12] = 9
	//m[13] = 1
	//m[14] = 11
	//m[15] = 6
	//m[16] = 6
	//m[17] = 9
	//m[18] = 47
	//m[19] = 41
	//m[20] = 84
	//m[21] = 12
	//m[22] = 4
	//m[23] = 7
	//m[24] = 19
	//m[25] = 16
	//m[26] = 2
	//m[27] = 18
	//m[28] = 16
	//m[29] = 50
	//m[30] = 11
	//m[31] = 13
	//m[32] = 5
	//m[33] = 18
	//m[34] = 12
	//m[35] = 27
	//m[36] = 8
	//m[37] = 4
	//m[38] = 42
	//m[39] = 24
	//m[40] = 3
	//m[41] = 12
	//m[42] = 93
	//m[43] = 35
	//m[44] = 6
	//m[45] = 75
	//m[46] = 33
	//m[47] = 4
	//m[48] = 6
	//m[49] = 28
	//m[50] = 46
	//m[51] = 20
	//m[52] = 44
	//m[53] = 7
	//m[54] = 10
	//m[55] = 6
	//m[56] = 20
	//m[57] = 16
	//m[58] = 65
	//m[59] = 24
	//m[60] = 17
	//m[61] = 39
	//m[62] = 9
	//m[63] = 66
	//m[64] = 43
	//m[65] = 6
	//m[66] = 13
	//m[67] = 14
	//m[68] = 35
	//m[69] = 8
	//m[70] = 18
	//m[71] = 11
	//m[72] = 26
	//m[73] = 6
	//m[74] = 7
	//m[75] = 36
	//m[76] = 119
	//m[77] = 15
	//m[78] = 22
	//m[79] = 4
	//m[80] = 5
	//m[81] = 7
	//m[82] = 24
	//m[83] = 6
	//m[84] = 120
	//m[85] = 12
	//m[86] = 11
	//m[87] = 8
	//m[88] = 141
	//m[89] = 21
	//m[90] = 37
	//m[91] = 6
	//m[92] = 2
	//m[93] = 53
	//m[94] = 17
	//m[95] = 17
	//m[96] = 9
	//m[97] = 28
	//m[98] = 48
	//m[99] = 8
	//m[100] = 17
	//m[101] = 101
	//m[102] = 34
	//m[103] = 40
	//m[104] = 86
	//m[105] = 41
	//m[106] = 8
	//m[107] = 100
	//m[108] = 8
	//m[109] = 80
	//m[110] = 16
	//m[111] = 11
	//m[112] = 34
	//m[113] = 10
	//m[114] = 2
	//m[115] = 19
	//m[116] = 1
	//m[117] = 16
	//m[118] = 6
	//m[119] = 7
	//m[120] = 1
	//m[121] = 46
	//m[122] = 9
	//m[123] = 17
	//m[124] = 145
	//m[125] = 4
	//m[126] = 3
	//m[127] = 12
	//m[128] = 25
	//m[129] = 9
	//m[130] = 23
	//m[131] = 8
	//m[132] = 66
	//m[133] = 74
	//m[134] = 12
	//m[135] = 7
	//m[136] = 42
	//m[137] = 10
	//m[138] = 60
	//m[139] = 24
	//m[140] = 13
	//m[141] = 10
	//m[142] = 7
	//m[143] = 12
	//m[144] = 15
	//m[145] = 21
	//m[146] = 10
	//m[147] = 20
	//m[148] = 14
	//m[149] = 9
	//m[150] = 6

	for k, v := range m {
		fmt.Printf("%s Chapter %s has %s verses\n", bookName, k, v)

		targetEntity := dbtargetentity.TargetEntity{
			Type: "standard_work_chapter",
		}

		err := dbtargetentity.Save(&targetEntity)
		if err != nil {
			fmt.Printf("Error saving TargetEntity : %s\n", err)
		}

		chapter := dbstandardworkchapter.StandardWorkChapter{
			StandardWorkChapterID: targetEntity.TargetEntityID,
			StandardWorkBookID:    bookId,
			ChapterNumber:         int64(k),
		}

		fmt.Printf("Saving Chapter : %#v\n", chapter)
		err = dbstandardworkchapter.Insert(&chapter)
		if err != nil {
			fmt.Printf("Error : %s\n", err)
		}

		for verse := 1; verse <= v; verse++ {

			targetEntityVerse := dbtargetentity.TargetEntity{
				Type: "standard_work_verse",
			}

			err := dbtargetentity.Save(&targetEntityVerse)
			if err != nil {
				fmt.Printf("Error saving TargetEntityVerse : %s\n", err)
			}

			verse := dbstandardworkverse.StandardWorkVerse{
				StandardWorkVerseID:   targetEntityVerse.TargetEntityID,
				StandardWorkChapterID: targetEntity.TargetEntityID,
				VerseNumber:           int64(verse),
			}

			fmt.Printf("Saving Verse : %#v\n", verse)
			err = dbstandardworkverse.Insert(&verse)
			if err != nil {
				fmt.Printf("Error : %s\n", err)
			}

		}
	}
}
