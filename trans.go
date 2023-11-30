package main

var transTups = map[string][][2]string{
	"ka": {
		{"ა", "a"},
		{"ბ", "b"},
		{"გ", "g"},
		{"დ", "d"},
		{"ე", "e"},
		{"ვ", "v"},
		{"ზ", "z"},
		{"თ", "t"},
		{"ი", "i"},
		{"კ", "k'"},
		{"ლ", "l"},
		{"მ", "m"},
		{"ნ", "n"},
		{"ო", "o"},
		{"პ", "p'"},
		{"ჟ", "zh"},
		{"რ", "r"},
		{"ს", "s"},
		{"ტ", "t'"},
		{"უ", "u"},
		{"ფ", "p"},
		{"ქ", "k"},
		{"ღ", "gh"},
		{"ყ", "q'"},
		{"შ", "sh"},
		{"ჩ", "ch"},
		{"ც", "ts"},
		{"ძ", "dz"},
		{"წ", "ts'"},
		{"ჭ", "ch'"},
		{"ხ", "kh"},
		{"ჯ", "j"},
		{"ჰ", "h"},
	},
	"hy": {
		{"Ա", "A"},
		{"Բ", "B"},
		{"Գ", "G"},
		{"Դ", "D"},
		{"Ե", "E"},
		{"Զ", "Z"},
		{"Է", "E"},
		{"Ը", "Ə"},
		{"Թ", "T'"},
		{"Ժ", "Zh"},
		{"Ի", "I"},
		{"Լ", "L"},
		{"Խ", "Kh"},
		{"Ծ", "Ts"},
		{"Կ", "K"},
		{"Հ", "H"},
		{"Ձ", "Dz"},
		{"Ղ", "Gh"},
		{"Ճ", "Ch'"},
		{"Մ", "M"},
		{"Յ", "Y"},
		{"Ն", "N"},
		{"Շ", "Sh"},
		{"Ո", "O"},
		{"Չ", "Ch"},
		{"Պ", "P'"},
		{"Ջ", "J"},
		{"Ռ", "R"},
		{"Ս", "S"},
		{"Վ", "V"},
		{"Տ", "T"},
		{"Ր", "R"},
		{"Ց", "Ts'"},
		{"Ւ", "W"},
		{"Փ", "P'"},
		{"Ք", "K'"},
		{"Օ", "O"},
		{"Ֆ", "F"},
		{"ԵՎ", "Ew"}, // Ligature

		// Lowercase
		{"ա", "a"},
		{"բ", "b"},
		{"գ", "g"},
		{"դ", "d"},
		{"ե", "e"},
		{"զ", "z"},
		{"է", "e"},
		{"ը", "ə"},
		{"թ", "t'"},
		{"ժ", "zh"},
		{"ի", "i"},
		{"լ", "l"},
		{"խ", "kh"},
		{"ծ", "ts"},
		{"կ", "k"},
		{"հ", "h"},
		{"ձ", "dz"},
		{"ղ", "gh"},
		{"ճ", "ch'"},
		{"մ", "m"},
		{"յ", "y"},
		{"ն", "n"},
		{"շ", "sh"},
		{"ո", "o"},
		{"չ", "ch"},
		{"պ", "p'"},
		{"ջ", "j"},
		{"ռ", "r"},
		{"ս", "s"},
		{"վ", "v"},
		{"տ", "t"},
		{"ր", "r"},
		{"ց", "ts'"},
		{"ւ", "w"},
		{"փ", "p'"},
		{"ք", "k'"},
		{"օ", "o"},
		{"ֆ", "f"},
		{"և", "ew"}, // Ligature
	},
	"am": {
		// 'ha' series
		{"ሀ", "ha"},
		{"ሁ", "hu"},
		{"ሂ", "hi"},
		{"ሃ", "haa"},
		{"ሄ", "hee"},
		{"ህ", "h"},
		{"ሆ", "ho"},

		// 'le' series
		{"ለ", "le"},
		{"ሉ", "lu"},
		{"ሊ", "li"},
		{"ላ", "laa"},
		{"ሌ", "lee"},
		{"ል", "l"},
		{"ሎ", "lo"},

		// 'ḥe' series
		{"ሐ", "ḥe"},
		{"ሑ", "ḥu"},
		{"ሒ", "ḥi"},
		{"ሓ", "ḥaa"},
		{"ሔ", "ḥee"},
		{"ሕ", "ḥ"},
		{"ሖ", "ḥo"},

		// 'me' series
		{"መ", "me"},
		{"ሙ", "mu"},
		{"ሚ", "mi"},
		{"ማ", "maa"},
		{"ሜ", "mee"},
		{"ም", "m"},
		{"ሞ", "mo"},

		// 'śe' series (or 'se' in some transliterations)
		{"ሠ", "śe"},
		{"ሡ", "śu"},
		{"ሢ", "śi"},
		{"ሣ", "śaa"},
		{"ሤ", "śee"},
		{"ሥ", "ś"},
		{"ሦ", "śo"},

		// 're' series
		{"ረ", "re"},
		{"ሩ", "ru"},
		{"ሪ", "ri"},
		{"ራ", "raa"},
		{"ሬ", "ree"},
		{"ር", "r"},
		{"ሮ", "ro"},

		// 'se' series
		{"ሰ", "se"},
		{"ሱ", "su"},
		{"ሲ", "si"},
		{"ሳ", "saa"},
		{"ሴ", "see"},
		{"ስ", "s"},
		{"ሶ", "so"},

		// 'še' series (or 'she' in some transliterations)
		{"ሸ", "še"},
		{"ሹ", "šu"},
		{"ሺ", "ši"},
		{"ሻ", "šaa"},
		{"ሼ", "šee"},
		{"ሽ", "š"},
		{"ሾ", "šo"},

		// 'qe' series
		{"ቀ", "qe"},
		{"ቁ", "qu"},
		{"ቂ", "qi"},
		{"ቃ", "qaa"},
		{"ቄ", "qee"},
		{"ቅ", "q"},
		{"ቆ", "qo"},

		// 'be' series
		{"በ", "be"},
		{"ቡ", "bu"},
		{"ቢ", "bi"},
		{"ባ", "baa"},
		{"ቤ", "bee"},
		{"ብ", "b"},
		{"ቦ", "bo"},

		// 've' series
		{"ቨ", "ve"},
		{"ቩ", "vu"},
		{"ቪ", "vi"},
		{"ቫ", "vaa"},
		{"ቬ", "vee"},
		{"ቭ", "v"},
		{"ቮ", "vo"},

		// 'te' series
		{"ተ", "te"},
		{"ቱ", "tu"},
		{"ቲ", "ti"},
		{"ታ", "taa"},
		{"ቴ", "tee"},
		{"ት", "t"},
		{"ቶ", "to"},

		// 'che' series
		{"ቸ", "che"},
		{"ቹ", "chu"},
		{"ቺ", "chi"},
		{"ቻ", "chaa"},
		{"ቼ", "chee"},
		{"ች", "ch"},
		{"ቾ", "cho"},

		// 'ne' series
		{"ነ", "ne"},
		{"ኑ", "nu"},
		{"ኒ", "ni"},
		{"ና", "naa"},
		{"ኔ", "nee"},
		{"ን", "n"},
		{"ኖ", "no"},

		// 'ñe' series
		{"ኘ", "ñe"},
		{"ኙ", "ñu"},
		{"ኚ", "ñi"},
		{"ኛ", "ñaa"},
		{"ኜ", "ñee"},
		{"ኝ", "ñ"},
		{"ኞ", "ño"},

		// 'ʾä' series
		{"አ", "'a"},
		{"ኡ", "'u"},
		{"ኢ", "'i"},
		{"ኣ", "'aa"},
		{"ኤ", "'ee"},
		{"እ", "'"},
		{"ኦ", "'o"},

		// 'ke' series
		{"ከ", "ke"},
		{"ኩ", "ku"},
		{"ኪ", "ki"},
		{"ካ", "kaa"},
		{"ኬ", "kee"},
		{"ክ", "k"},
		{"ኮ", "ko"},

		// 'xe' series
		{"ኸ", "xe"},
		{"ኹ", "xu"},
		{"ኺ", "xi"},
		{"ኻ", "xaa"},
		{"ኼ", "xee"},
		{"ኽ", "x"},
		{"ኾ", "xo"},

		// 'we' series
		{"ወ", "we"},
		{"ዉ", "wu"},
		{"ዊ", "wi"},
		{"ዋ", "waa"},
		{"ዌ", "wee"},
		{"ው", "w"},
		{"ዎ", "wo"},

		// '`e' series
		{"ዐ", "`a"},
		{"ዑ", "`u"},
		{"ዒ", "`i"},
		{"ዓ", "`aa"},
		{"ዔ", "`ee"},
		{"ዕ", "`"},
		{"ዖ", "`o"},

		// 'ze' series
		{"ዘ", "ze"},
		{"ዙ", "zu"},
		{"ዚ", "zi"},
		{"ዛ", "zaa"},
		{"ዜ", "zee"},
		{"ዝ", "z"},
		{"ዞ", "zo"},

		// 'že' series
		{"ዠ", "že"},
		{"ዡ", "žu"},
		{"ዢ", "ži"},
		{"ዣ", "žaa"},
		{"ዤ", "žee"},
		{"ዥ", "ž"},
		{"ዦ", "žo"},

		// 'ye' series
		{"የ", "ye"},
		{"ዩ", "yu"},
		{"ዪ", "yi"},
		{"ያ", "yaa"},
		{"ዬ", "yee"},
		{"ይ", "y"},
		{"ዮ", "yo"},

		// 'de' series
		{"ደ", "de"},
		{"ዱ", "du"},
		{"ዲ", "di"},
		{"ዳ", "daa"},
		{"ዴ", "dee"},
		{"ድ", "d"},
		{"ዶ", "do"},

		// 'ge' series
		{"ጀ", "ge"},
		{"ጁ", "gu"},
		{"ጂ", "gi"},
		{"ጃ", "gaa"},
		{"ጄ", "gee"},
		{"ጅ", "g"},
		{"ጆ", "go"},

		// 'ge' series
		{"ገ", "ge"},
		{"ጉ", "gu"},
		{"ጊ", "gi"},
		{"ጋ", "gaa"},
		{"ጌ", "gee"},
		{"ግ", "g"},
		{"ጎ", "go"},

		// 'te' series
		{"ጠ", "te"},
		{"ጡ", "tu"},
		{"ጢ", "ti"},
		{"ጣ", "taa"},
		{"ጤ", "tee"},
		{"ጥ", "t"},
		{"ጦ", "to"},

		// 'če' series
		{"ጨ", "če"},
		{"ጩ", "ču"},
		{"ጪ", "či"},
		{"ጫ", "čaa"},
		{"ጬ", "čee"},
		{"ጭ", "č"},
		{"ጮ", "čo"},

		// 'pe' series
		{"ጰ", "pe"},
		{"ጱ", "pu"},
		{"ጲ", "pi"},
		{"ጳ", "paa"},
		{"ጴ", "pee"},
		{"ጵ", "p"},
		{"ጶ", "po"},

		// 'se' series
		{"ጸ", "se"},
		{"ጹ", "su"},
		{"ጺ", "si"},
		{"ጻ", "saa"},
		{"ጼ", "see"},
		{"ጽ", "s"},
		{"ጾ", "so"},

		// 'se' series
		{"ፀ", "se"},
		{"ፁ", "su"},
		{"ፂ", "si"},
		{"ፃ", "saa"},
		{"ፄ", "see"},
		{"ፅ", "s"},
		{"ፆ", "so"},

		// 'fe' series
		{"ፈ", "fe"},
		{"ፉ", "fu"},
		{"ፊ", "fi"},
		{"ፋ", "faa"},
		{"ፌ", "fee"},
		{"ፍ", "f"},
		{"ፎ", "fo"},

		// 'pe' series
		{"ፐ", "pe"},
		{"ፑ", "pu"},
		{"ፒ", "pi"},
		{"ፓ", "paa"},
		{"ፔ", "pee"},
		{"ፕ", "p"},
		{"ፖ", "po"},

		// Labialized 'ka' series
		{"ኰ", "kwa"},
		//{"኱", "kwu"},
		{"ኲ", "kwi"},
		{"ኳ", "kwaa"},
		{"ኴ", "kwee"},
		{"ኵ", "kw"},
		//{"኶", "kwo"},

		// Labialized 'ga' series
		{"ጐ", "gwa"},
		//{"጑", "gwu"},
		{"ጒ", "gwi"},
		{"ጓ", "gwaa"},
		{"ጔ", "gwee"},
		{"ጕ", "gw"},
		{"጖", "gwo"},

		// Additional series
		{"ቈ", "k'u"},
		{"ቊ", "k'i"},
		{"ቋ", "k'aa"},
		{"ቌ", "k'ee"},
		{"ቍ", "k'o"},

		{"ቘ", "kʰu"},
		{"ቚ", "kʰi"},
		{"ቛ", "kʰaa"},
		{"ቜ", "kʰee"},
		{"ቝ", "kʰo"},

		{"ቨ", "fya"},
		{"ቩ", "fyu"},
		{"ቪ", "fyi"},
		{"ቫ", "fyaa"},
		{"ቬ", "fyee"},
		{"ቭ", "fy"},
		{"ቮ", "fyo"},

		{"ኘ", "ña"},
		{"ኙ", "ñu"},
		{"ኚ", "ñi"},
		{"ኛ", "ñaa"},
		{"ኜ", "ñee"},
		{"ኝ", "ñ"},
		{"ኞ", "ño"},

		{"ኸ", "fy'a"},
		{"ኹ", "fy'u"},
		{"ኺ", "fy'i"},
		{"ኻ", "fy'aa"},
		{"ኼ", "fy'ee"},
		{"ኽ", "fy'"},
		{"ኾ", "fy'o"},
	},
}