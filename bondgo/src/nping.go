package main

import (
	"bondgo"
)

func pong(c chan uint8) {
	var ball uint8
	for {
		ball = <-c
		ball++
		c <- ball
	}
	bondgo.Void(ball)
}

func main() {

	var ch0 chan uint8
	var ch1 chan uint8
	var ch2 chan uint8
	var ch3 chan uint8
	var ch4 chan uint8
	var ch5 chan uint8
	var ch6 chan uint8
	var ch7 chan uint8
	var ch8 chan uint8
	var ch9 chan uint8
	var ch10 chan uint8
	var ch11 chan uint8
	var ch12 chan uint8
	var ch13 chan uint8
	var ch14 chan uint8
	var ch15 chan uint8
	var ch16 chan uint8
	var ch17 chan uint8
	var ch18 chan uint8
	var ch19 chan uint8
	var ch20 chan uint8
	var ch21 chan uint8
	var ch22 chan uint8
	var ch23 chan uint8
	var ch24 chan uint8
	var ch25 chan uint8
	var ch26 chan uint8
	var ch27 chan uint8
	var ch28 chan uint8
	var ch29 chan uint8
	var ch30 chan uint8
	var ch31 chan uint8
	var ch32 chan uint8
	var ch33 chan uint8
	var ch34 chan uint8
	var ch35 chan uint8
	var ch36 chan uint8
	var ch37 chan uint8
	var ch38 chan uint8
	var ch39 chan uint8
	var ch40 chan uint8
	var ch41 chan uint8
	var ch42 chan uint8
	var ch43 chan uint8
	var ch44 chan uint8
	var ch45 chan uint8
	var ch46 chan uint8
	var ch47 chan uint8
	var ch48 chan uint8
	var ch49 chan uint8
	var ch50 chan uint8
	var ch51 chan uint8
	var ch52 chan uint8
	var ch53 chan uint8
	var ch54 chan uint8
	var ch55 chan uint8
	var ch56 chan uint8
	var ch57 chan uint8
	var ch58 chan uint8
	var ch59 chan uint8
	var ch60 chan uint8
	var ch61 chan uint8
	var ch62 chan uint8
	var ch63 chan uint8
	var ch64 chan uint8
	var ch65 chan uint8
	var ch66 chan uint8
	var ch67 chan uint8
	var ch68 chan uint8
	var ch69 chan uint8
	var ch70 chan uint8
	var ch71 chan uint8
	var ch72 chan uint8
	var ch73 chan uint8
	var ch74 chan uint8
	var ch75 chan uint8
	var ch76 chan uint8
	var ch77 chan uint8
	var ch78 chan uint8
	var ch79 chan uint8
	var ch80 chan uint8
	var ch81 chan uint8
	var ch82 chan uint8
	var ch83 chan uint8
	var ch84 chan uint8
	var ch85 chan uint8
	var ch86 chan uint8
	var ch87 chan uint8
	var ch88 chan uint8
	var ch89 chan uint8
	var ch90 chan uint8
	var ch91 chan uint8
	var ch92 chan uint8
	var ch93 chan uint8
	var ch94 chan uint8
	var ch95 chan uint8
	var ch96 chan uint8
	var ch97 chan uint8
	var ch98 chan uint8
	var ch99 chan uint8
	var ch100 chan uint8
	var ch101 chan uint8
	var ch102 chan uint8
	var ch103 chan uint8
	var ch104 chan uint8
	var ch105 chan uint8
	var ch106 chan uint8
	var ch107 chan uint8
	var ch108 chan uint8
	var ch109 chan uint8
	var ch110 chan uint8
	var ch111 chan uint8
	var ch112 chan uint8
	var ch113 chan uint8
	var ch114 chan uint8
	var ch115 chan uint8
	var ch116 chan uint8
	var ch117 chan uint8
	var ch118 chan uint8
	var ch119 chan uint8
	var ch120 chan uint8
	var ch121 chan uint8
	var ch122 chan uint8
	var ch123 chan uint8
	var ch124 chan uint8
	var ch125 chan uint8
	var ch126 chan uint8
	var ch127 chan uint8
	var ch128 chan uint8
	var ch129 chan uint8
	var ch130 chan uint8
	var ch131 chan uint8
	var ch132 chan uint8
	var ch133 chan uint8
	var ch134 chan uint8
	var ch135 chan uint8
	var ch136 chan uint8
	var ch137 chan uint8
	var ch138 chan uint8
	var ch139 chan uint8
	var ch140 chan uint8
	var ch141 chan uint8
	var ch142 chan uint8
	var ch143 chan uint8
	var ch144 chan uint8
	var ch145 chan uint8
	var ch146 chan uint8
	var ch147 chan uint8
	var ch148 chan uint8
	var ch149 chan uint8
	var ch150 chan uint8
	var ch151 chan uint8
	var ch152 chan uint8
	var ch153 chan uint8
	var ch154 chan uint8
	var ch155 chan uint8
	var ch156 chan uint8
	var ch157 chan uint8
	var ch158 chan uint8
	var ch159 chan uint8
	var ch160 chan uint8
	var ch161 chan uint8
	var ch162 chan uint8
	var ch163 chan uint8
	var ch164 chan uint8
	var ch165 chan uint8
	var ch166 chan uint8
	var ch167 chan uint8
	var ch168 chan uint8
	var ch169 chan uint8
	var ch170 chan uint8
	var ch171 chan uint8
	var ch172 chan uint8
	var ch173 chan uint8
	var ch174 chan uint8
	var ch175 chan uint8
	var ch176 chan uint8
	var ch177 chan uint8
	var ch178 chan uint8
	var ch179 chan uint8
	var ch180 chan uint8
	var ch181 chan uint8
	var ch182 chan uint8
	var ch183 chan uint8
	var ch184 chan uint8
	var ch185 chan uint8
	var ch186 chan uint8
	var ch187 chan uint8
	var ch188 chan uint8
	var ch189 chan uint8
	var ch190 chan uint8
	var ch191 chan uint8
	var ch192 chan uint8
	var ch193 chan uint8
	var ch194 chan uint8
	var ch195 chan uint8
	var ch196 chan uint8
	var ch197 chan uint8
	var ch198 chan uint8
	var ch199 chan uint8
	var ch200 chan uint8
	var ch201 chan uint8
	var ch202 chan uint8
	var ch203 chan uint8
	var ch204 chan uint8
	var ch205 chan uint8
	var ch206 chan uint8
	var ch207 chan uint8
	var ch208 chan uint8
	var ch209 chan uint8
	var ch210 chan uint8
	var ch211 chan uint8
	var ch212 chan uint8
	var ch213 chan uint8
	var ch214 chan uint8
	var ch215 chan uint8
	var ch216 chan uint8
	var ch217 chan uint8
	var ch218 chan uint8
	var ch219 chan uint8
	var ch220 chan uint8
	var ch221 chan uint8
	var ch222 chan uint8
	var ch223 chan uint8
	var ch224 chan uint8
	var ch225 chan uint8
	var ch226 chan uint8
	var ch227 chan uint8
	var ch228 chan uint8
	var ch229 chan uint8
	var ch230 chan uint8
	var ch231 chan uint8
	var ch232 chan uint8
	var ch233 chan uint8
	var ch234 chan uint8
	var ch235 chan uint8
	var ch236 chan uint8
	var ch237 chan uint8
	var ch238 chan uint8
	var ch239 chan uint8
	var ch240 chan uint8
	var ch241 chan uint8
	var ch242 chan uint8
	var ch243 chan uint8
	var ch244 chan uint8
	var ch245 chan uint8
	var ch246 chan uint8
	var ch247 chan uint8
	var ch248 chan uint8
	var ch249 chan uint8
	var ch250 chan uint8
	var ch251 chan uint8
	var ch252 chan uint8
	var ch253 chan uint8
	var ch254 chan uint8
	var ch255 chan uint8
	var ball0 uint8
	var ball1 uint8
	var ball2 uint8
	var ball3 uint8
	var ball4 uint8
	var ball5 uint8
	var ball6 uint8
	var ball7 uint8
	var ball8 uint8
	var ball9 uint8
	var ball10 uint8
	var ball11 uint8
	var ball12 uint8
	var ball13 uint8
	var ball14 uint8
	var ball15 uint8
	var ball16 uint8
	var ball17 uint8
	var ball18 uint8
	var ball19 uint8
	var ball20 uint8
	var ball21 uint8
	var ball22 uint8
	var ball23 uint8
	var ball24 uint8
	var ball25 uint8
	var ball26 uint8
	var ball27 uint8
	var ball28 uint8
	var ball29 uint8
	var ball30 uint8
	var ball31 uint8
	var ball32 uint8
	var ball33 uint8
	var ball34 uint8
	var ball35 uint8
	var ball36 uint8
	var ball37 uint8
	var ball38 uint8
	var ball39 uint8
	var ball40 uint8
	var ball41 uint8
	var ball42 uint8
	var ball43 uint8
	var ball44 uint8
	var ball45 uint8
	var ball46 uint8
	var ball47 uint8
	var ball48 uint8
	var ball49 uint8
	var ball50 uint8
	var ball51 uint8
	var ball52 uint8
	var ball53 uint8
	var ball54 uint8
	var ball55 uint8
	var ball56 uint8
	var ball57 uint8
	var ball58 uint8
	var ball59 uint8
	var ball60 uint8
	var ball61 uint8
	var ball62 uint8
	var ball63 uint8
	var ball64 uint8
	var ball65 uint8
	var ball66 uint8
	var ball67 uint8
	var ball68 uint8
	var ball69 uint8
	var ball70 uint8
	var ball71 uint8
	var ball72 uint8
	var ball73 uint8
	var ball74 uint8
	var ball75 uint8
	var ball76 uint8
	var ball77 uint8
	var ball78 uint8
	var ball79 uint8
	var ball80 uint8
	var ball81 uint8
	var ball82 uint8
	var ball83 uint8
	var ball84 uint8
	var ball85 uint8
	var ball86 uint8
	var ball87 uint8
	var ball88 uint8
	var ball89 uint8
	var ball90 uint8
	var ball91 uint8
	var ball92 uint8
	var ball93 uint8
	var ball94 uint8
	var ball95 uint8
	var ball96 uint8
	var ball97 uint8
	var ball98 uint8
	var ball99 uint8
	var ball100 uint8
	var ball101 uint8
	var ball102 uint8
	var ball103 uint8
	var ball104 uint8
	var ball105 uint8
	var ball106 uint8
	var ball107 uint8
	var ball108 uint8
	var ball109 uint8
	var ball110 uint8
	var ball111 uint8
	var ball112 uint8
	var ball113 uint8
	var ball114 uint8
	var ball115 uint8
	var ball116 uint8
	var ball117 uint8
	var ball118 uint8
	var ball119 uint8
	var ball120 uint8
	var ball121 uint8
	var ball122 uint8
	var ball123 uint8
	var ball124 uint8
	var ball125 uint8
	var ball126 uint8
	var ball127 uint8
	var ball128 uint8
	var ball129 uint8
	var ball130 uint8
	var ball131 uint8
	var ball132 uint8
	var ball133 uint8
	var ball134 uint8
	var ball135 uint8
	var ball136 uint8
	var ball137 uint8
	var ball138 uint8
	var ball139 uint8
	var ball140 uint8
	var ball141 uint8
	var ball142 uint8
	var ball143 uint8
	var ball144 uint8
	var ball145 uint8
	var ball146 uint8
	var ball147 uint8
	var ball148 uint8
	var ball149 uint8
	var ball150 uint8
	var ball151 uint8
	var ball152 uint8
	var ball153 uint8
	var ball154 uint8
	var ball155 uint8
	var ball156 uint8
	var ball157 uint8
	var ball158 uint8
	var ball159 uint8
	var ball160 uint8
	var ball161 uint8
	var ball162 uint8
	var ball163 uint8
	var ball164 uint8
	var ball165 uint8
	var ball166 uint8
	var ball167 uint8
	var ball168 uint8
	var ball169 uint8
	var ball170 uint8
	var ball171 uint8
	var ball172 uint8
	var ball173 uint8
	var ball174 uint8
	var ball175 uint8
	var ball176 uint8
	var ball177 uint8
	var ball178 uint8
	var ball179 uint8
	var ball180 uint8
	var ball181 uint8
	var ball182 uint8
	var ball183 uint8
	var ball184 uint8
	var ball185 uint8
	var ball186 uint8
	var ball187 uint8
	var ball188 uint8
	var ball189 uint8
	var ball190 uint8
	var ball191 uint8
	var ball192 uint8
	var ball193 uint8
	var ball194 uint8
	var ball195 uint8
	var ball196 uint8
	var ball197 uint8
	var ball198 uint8
	var ball199 uint8
	var ball200 uint8
	var ball201 uint8
	var ball202 uint8
	var ball203 uint8
	var ball204 uint8
	var ball205 uint8
	var ball206 uint8
	var ball207 uint8
	var ball208 uint8
	var ball209 uint8
	var ball210 uint8
	var ball211 uint8
	var ball212 uint8
	var ball213 uint8
	var ball214 uint8
	var ball215 uint8
	var ball216 uint8
	var ball217 uint8
	var ball218 uint8
	var ball219 uint8
	var ball220 uint8
	var ball221 uint8
	var ball222 uint8
	var ball223 uint8
	var ball224 uint8
	var ball225 uint8
	var ball226 uint8
	var ball227 uint8
	var ball228 uint8
	var ball229 uint8
	var ball230 uint8
	var ball231 uint8
	var ball232 uint8
	var ball233 uint8
	var ball234 uint8
	var ball235 uint8
	var ball236 uint8
	var ball237 uint8
	var ball238 uint8
	var ball239 uint8
	var ball240 uint8
	var ball241 uint8
	var ball242 uint8
	var ball243 uint8
	var ball244 uint8
	var ball245 uint8
	var ball246 uint8
	var ball247 uint8
	var ball248 uint8
	var ball249 uint8
	var ball250 uint8
	var ball251 uint8
	var ball252 uint8
	var ball253 uint8
	var ball254 uint8
	var ball255 uint8
	ball0 = 1
	ball1 = 1
	ball2 = 1
	ball3 = 1
	ball4 = 1
	ball5 = 1
	ball6 = 1
	ball7 = 1
	ball8 = 1
	ball9 = 1
	ball10 = 1
	ball11 = 1
	ball12 = 1
	ball13 = 1
	ball14 = 1
	ball15 = 1
	ball16 = 1
	ball17 = 1
	ball18 = 1
	ball19 = 1
	ball20 = 1
	ball21 = 1
	ball22 = 1
	ball23 = 1
	ball24 = 1
	ball25 = 1
	ball26 = 1
	ball27 = 1
	ball28 = 1
	ball29 = 1
	ball30 = 1
	ball31 = 1
	ball32 = 1
	ball33 = 1
	ball34 = 1
	ball35 = 1
	ball36 = 1
	ball37 = 1
	ball38 = 1
	ball39 = 1
	ball40 = 1
	ball41 = 1
	ball42 = 1
	ball43 = 1
	ball44 = 1
	ball45 = 1
	ball46 = 1
	ball47 = 1
	ball48 = 1
	ball49 = 1
	ball50 = 1
	ball51 = 1
	ball52 = 1
	ball53 = 1
	ball54 = 1
	ball55 = 1
	ball56 = 1
	ball57 = 1
	ball58 = 1
	ball59 = 1
	ball60 = 1
	ball61 = 1
	ball62 = 1
	ball63 = 1
	ball64 = 1
	ball65 = 1
	ball66 = 1
	ball67 = 1
	ball68 = 1
	ball69 = 1
	ball70 = 1
	ball71 = 1
	ball72 = 1
	ball73 = 1
	ball74 = 1
	ball75 = 1
	ball76 = 1
	ball77 = 1
	ball78 = 1
	ball79 = 1
	ball80 = 1
	ball81 = 1
	ball82 = 1
	ball83 = 1
	ball84 = 1
	ball85 = 1
	ball86 = 1
	ball87 = 1
	ball88 = 1
	ball89 = 1
	ball90 = 1
	ball91 = 1
	ball92 = 1
	ball93 = 1
	ball94 = 1
	ball95 = 1
	ball96 = 1
	ball97 = 1
	ball98 = 1
	ball99 = 1
	ball100 = 1
	ball101 = 1
	ball102 = 1
	ball103 = 1
	ball104 = 1
	ball105 = 1
	ball106 = 1
	ball107 = 1
	ball108 = 1
	ball109 = 1
	ball110 = 1
	ball111 = 1
	ball112 = 1
	ball113 = 1
	ball114 = 1
	ball115 = 1
	ball116 = 1
	ball117 = 1
	ball118 = 1
	ball119 = 1
	ball120 = 1
	ball121 = 1
	ball122 = 1
	ball123 = 1
	ball124 = 1
	ball125 = 1
	ball126 = 1
	ball127 = 1
	ball128 = 1
	ball129 = 1
	ball130 = 1
	ball131 = 1
	ball132 = 1
	ball133 = 1
	ball134 = 1
	ball135 = 1
	ball136 = 1
	ball137 = 1
	ball138 = 1
	ball139 = 1
	ball140 = 1
	ball141 = 1
	ball142 = 1
	ball143 = 1
	ball144 = 1
	ball145 = 1
	ball146 = 1
	ball147 = 1
	ball148 = 1
	ball149 = 1
	ball150 = 1
	ball151 = 1
	ball152 = 1
	ball153 = 1
	ball154 = 1
	ball155 = 1
	ball156 = 1
	ball157 = 1
	ball158 = 1
	ball159 = 1
	ball160 = 1
	ball161 = 1
	ball162 = 1
	ball163 = 1
	ball164 = 1
	ball165 = 1
	ball166 = 1
	ball167 = 1
	ball168 = 1
	ball169 = 1
	ball170 = 1
	ball171 = 1
	ball172 = 1
	ball173 = 1
	ball174 = 1
	ball175 = 1
	ball176 = 1
	ball177 = 1
	ball178 = 1
	ball179 = 1
	ball180 = 1
	ball181 = 1
	ball182 = 1
	ball183 = 1
	ball184 = 1
	ball185 = 1
	ball186 = 1
	ball187 = 1
	ball188 = 1
	ball189 = 1
	ball190 = 1
	ball191 = 1
	ball192 = 1
	ball193 = 1
	ball194 = 1
	ball195 = 1
	ball196 = 1
	ball197 = 1
	ball198 = 1
	ball199 = 1
	ball200 = 1
	ball201 = 1
	ball202 = 1
	ball203 = 1
	ball204 = 1
	ball205 = 1
	ball206 = 1
	ball207 = 1
	ball208 = 1
	ball209 = 1
	ball210 = 1
	ball211 = 1
	ball212 = 1
	ball213 = 1
	ball214 = 1
	ball215 = 1
	ball216 = 1
	ball217 = 1
	ball218 = 1
	ball219 = 1
	ball220 = 1
	ball221 = 1
	ball222 = 1
	ball223 = 1
	ball224 = 1
	ball225 = 1
	ball226 = 1
	ball227 = 1
	ball228 = 1
	ball229 = 1
	ball230 = 1
	ball231 = 1
	ball232 = 1
	ball233 = 1
	ball234 = 1
	ball235 = 1
	ball236 = 1
	ball237 = 1
	ball238 = 1
	ball239 = 1
	ball240 = 1
	ball241 = 1
	ball242 = 1
	ball243 = 1
	ball244 = 1
	ball245 = 1
	ball246 = 1
	ball247 = 1
	ball248 = 1
	ball249 = 1
	ball250 = 1
	ball251 = 1
	ball252 = 1
	ball253 = 1
	ball254 = 1
	ball255 = 1

	var sel uint8

	sel = 0

	//ch1 = make(chan uint8)
	//ch2 = make(chan uint8)

	go pong(ch0)
	go pong(ch1)
	go pong(ch2)
	go pong(ch3)
	go pong(ch4)
	go pong(ch5)
	go pong(ch6)
	go pong(ch7)
	go pong(ch8)
	go pong(ch9)
	go pong(ch10)
	go pong(ch11)
	go pong(ch12)
	go pong(ch13)
	go pong(ch14)
	go pong(ch15)
	go pong(ch16)
	go pong(ch17)
	go pong(ch18)
	go pong(ch19)
	go pong(ch20)
	go pong(ch21)
	go pong(ch22)
	go pong(ch23)
	go pong(ch24)
	go pong(ch25)
	go pong(ch26)
	go pong(ch27)
	go pong(ch28)
	go pong(ch29)
	go pong(ch30)
	go pong(ch31)
	go pong(ch32)
	go pong(ch33)
	go pong(ch34)
	go pong(ch35)
	go pong(ch36)
	go pong(ch37)
	go pong(ch38)
	go pong(ch39)
	go pong(ch40)
	go pong(ch41)
	go pong(ch42)
	go pong(ch43)
	go pong(ch44)
	go pong(ch45)
	go pong(ch46)
	go pong(ch47)
	go pong(ch48)
	go pong(ch49)
	go pong(ch50)
	go pong(ch51)
	go pong(ch52)
	go pong(ch53)
	go pong(ch54)
	go pong(ch55)
	go pong(ch56)
	go pong(ch57)
	go pong(ch58)
	go pong(ch59)
	go pong(ch60)
	go pong(ch61)
	go pong(ch62)
	go pong(ch63)
	go pong(ch64)
	go pong(ch65)
	go pong(ch66)
	go pong(ch67)
	go pong(ch68)
	go pong(ch69)
	go pong(ch70)
	go pong(ch71)
	go pong(ch72)
	go pong(ch73)
	go pong(ch74)
	go pong(ch75)
	go pong(ch76)
	go pong(ch77)
	go pong(ch78)
	go pong(ch79)
	go pong(ch80)
	go pong(ch81)
	go pong(ch82)
	go pong(ch83)
	go pong(ch84)
	go pong(ch85)
	go pong(ch86)
	go pong(ch87)
	go pong(ch88)
	go pong(ch89)
	go pong(ch90)
	go pong(ch91)
	go pong(ch92)
	go pong(ch93)
	go pong(ch94)
	go pong(ch95)
	go pong(ch96)
	go pong(ch97)
	go pong(ch98)
	go pong(ch99)
	go pong(ch100)
	go pong(ch101)
	go pong(ch102)
	go pong(ch103)
	go pong(ch104)
	go pong(ch105)
	go pong(ch106)
	go pong(ch107)
	go pong(ch108)
	go pong(ch109)
	go pong(ch110)
	go pong(ch111)
	go pong(ch112)
	go pong(ch113)
	go pong(ch114)
	go pong(ch115)
	go pong(ch116)
	go pong(ch117)
	go pong(ch118)
	go pong(ch119)
	go pong(ch120)
	go pong(ch121)
	go pong(ch122)
	go pong(ch123)
	go pong(ch124)
	go pong(ch125)
	go pong(ch126)
	go pong(ch127)
	go pong(ch128)
	go pong(ch129)
	go pong(ch130)
	go pong(ch131)
	go pong(ch132)
	go pong(ch133)
	go pong(ch134)
	go pong(ch135)
	go pong(ch136)
	go pong(ch137)
	go pong(ch138)
	go pong(ch139)
	go pong(ch140)
	go pong(ch141)
	go pong(ch142)
	go pong(ch143)
	go pong(ch144)
	go pong(ch145)
	go pong(ch146)
	go pong(ch147)
	go pong(ch148)
	go pong(ch149)
	go pong(ch150)
	go pong(ch151)
	go pong(ch152)
	go pong(ch153)
	go pong(ch154)
	go pong(ch155)
	go pong(ch156)
	go pong(ch157)
	go pong(ch158)
	go pong(ch159)
	go pong(ch160)
	go pong(ch161)
	go pong(ch162)
	go pong(ch163)
	go pong(ch164)
	go pong(ch165)
	go pong(ch166)
	go pong(ch167)
	go pong(ch168)
	go pong(ch169)
	go pong(ch170)
	go pong(ch171)
	go pong(ch172)
	go pong(ch173)
	go pong(ch174)
	go pong(ch175)
	go pong(ch176)
	go pong(ch177)
	go pong(ch178)
	go pong(ch179)
	go pong(ch180)
	go pong(ch181)
	go pong(ch182)
	go pong(ch183)
	go pong(ch184)
	go pong(ch185)
	go pong(ch186)
	go pong(ch187)
	go pong(ch188)
	go pong(ch189)
	go pong(ch190)
	go pong(ch191)
	go pong(ch192)
	go pong(ch193)
	go pong(ch194)
	go pong(ch195)
	go pong(ch196)
	go pong(ch197)
	go pong(ch198)
	go pong(ch199)
	go pong(ch200)
	go pong(ch201)
	go pong(ch202)
	go pong(ch203)
	go pong(ch204)
	go pong(ch205)
	go pong(ch206)
	go pong(ch207)
	go pong(ch208)
	go pong(ch209)
	go pong(ch210)
	go pong(ch211)
	go pong(ch212)
	go pong(ch213)
	go pong(ch214)
	go pong(ch215)
	go pong(ch216)
	go pong(ch217)
	go pong(ch218)
	go pong(ch219)
	go pong(ch220)
	go pong(ch221)
	go pong(ch222)
	go pong(ch223)
	go pong(ch224)
	go pong(ch225)
	go pong(ch226)
	go pong(ch227)
	go pong(ch228)
	go pong(ch229)
	go pong(ch230)
	go pong(ch231)
	go pong(ch232)
	go pong(ch233)
	go pong(ch234)
	go pong(ch235)
	go pong(ch236)
	go pong(ch237)
	go pong(ch238)
	go pong(ch239)
	go pong(ch240)
	go pong(ch241)
	go pong(ch242)
	go pong(ch243)
	go pong(ch244)
	go pong(ch245)
	go pong(ch246)
	go pong(ch247)
	go pong(ch248)
	go pong(ch249)
	go pong(ch250)
	go pong(ch251)
	go pong(ch252)
	go pong(ch253)
	go pong(ch254)
	go pong(ch255)

	for {
		select {

		case ch0 <- ball0:
			sel = 1
		case ball0 = <-ch0:
			sel = 2
		case ch1 <- ball1:
			sel = 3
		case ball1 = <-ch1:
			sel = 4
		case ch2 <- ball2:
			sel = 5
		case ball2 = <-ch2:
			sel = 6
		case ch3 <- ball3:
			sel = 7
		case ball3 = <-ch3:
			sel = 8
		case ch4 <- ball4:
			sel = 9
		case ball4 = <-ch4:
			sel = 10
		case ch5 <- ball5:
			sel = 11
		case ball5 = <-ch5:
			sel = 12
		case ch6 <- ball6:
			sel = 13
		case ball6 = <-ch6:
			sel = 14
		case ch7 <- ball7:
			sel = 15
		case ball7 = <-ch7:
			sel = 16
		case ch8 <- ball8:
			sel = 17
		case ball8 = <-ch8:
			sel = 18
		case ch9 <- ball9:
			sel = 19
		case ball9 = <-ch9:
			sel = 20
		case ch10 <- ball10:
			sel = 21
		case ball10 = <-ch10:
			sel = 22
		case ch11 <- ball11:
			sel = 23
		case ball11 = <-ch11:
			sel = 24
		case ch12 <- ball12:
			sel = 25
		case ball12 = <-ch12:
			sel = 26
		case ch13 <- ball13:
			sel = 27
		case ball13 = <-ch13:
			sel = 28
		case ch14 <- ball14:
			sel = 29
		case ball14 = <-ch14:
			sel = 30
		case ch15 <- ball15:
			sel = 31
		case ball15 = <-ch15:
			sel = 32
		case ch16 <- ball16:
			sel = 33
		case ball16 = <-ch16:
			sel = 34
		case ch17 <- ball17:
			sel = 35
		case ball17 = <-ch17:
			sel = 36
		case ch18 <- ball18:
			sel = 37
		case ball18 = <-ch18:
			sel = 38
		case ch19 <- ball19:
			sel = 39
		case ball19 = <-ch19:
			sel = 40
		case ch20 <- ball20:
			sel = 41
		case ball20 = <-ch20:
			sel = 42
		case ch21 <- ball21:
			sel = 43
		case ball21 = <-ch21:
			sel = 44
		case ch22 <- ball22:
			sel = 45
		case ball22 = <-ch22:
			sel = 46
		case ch23 <- ball23:
			sel = 47
		case ball23 = <-ch23:
			sel = 48
		case ch24 <- ball24:
			sel = 49
		case ball24 = <-ch24:
			sel = 50
		case ch25 <- ball25:
			sel = 51
		case ball25 = <-ch25:
			sel = 52
		case ch26 <- ball26:
			sel = 53
		case ball26 = <-ch26:
			sel = 54
		case ch27 <- ball27:
			sel = 55
		case ball27 = <-ch27:
			sel = 56
		case ch28 <- ball28:
			sel = 57
		case ball28 = <-ch28:
			sel = 58
		case ch29 <- ball29:
			sel = 59
		case ball29 = <-ch29:
			sel = 60
		case ch30 <- ball30:
			sel = 61
		case ball30 = <-ch30:
			sel = 62
		case ch31 <- ball31:
			sel = 63
		case ball31 = <-ch31:
			sel = 64
		case ch32 <- ball32:
			sel = 65
		case ball32 = <-ch32:
			sel = 66
		case ch33 <- ball33:
			sel = 67
		case ball33 = <-ch33:
			sel = 68
		case ch34 <- ball34:
			sel = 69
		case ball34 = <-ch34:
			sel = 70
		case ch35 <- ball35:
			sel = 71
		case ball35 = <-ch35:
			sel = 72
		case ch36 <- ball36:
			sel = 73
		case ball36 = <-ch36:
			sel = 74
		case ch37 <- ball37:
			sel = 75
		case ball37 = <-ch37:
			sel = 76
		case ch38 <- ball38:
			sel = 77
		case ball38 = <-ch38:
			sel = 78
		case ch39 <- ball39:
			sel = 79
		case ball39 = <-ch39:
			sel = 80
		case ch40 <- ball40:
			sel = 81
		case ball40 = <-ch40:
			sel = 82
		case ch41 <- ball41:
			sel = 83
		case ball41 = <-ch41:
			sel = 84
		case ch42 <- ball42:
			sel = 85
		case ball42 = <-ch42:
			sel = 86
		case ch43 <- ball43:
			sel = 87
		case ball43 = <-ch43:
			sel = 88
		case ch44 <- ball44:
			sel = 89
		case ball44 = <-ch44:
			sel = 90
		case ch45 <- ball45:
			sel = 91
		case ball45 = <-ch45:
			sel = 92
		case ch46 <- ball46:
			sel = 93
		case ball46 = <-ch46:
			sel = 94
		case ch47 <- ball47:
			sel = 95
		case ball47 = <-ch47:
			sel = 96
		case ch48 <- ball48:
			sel = 97
		case ball48 = <-ch48:
			sel = 98
		case ch49 <- ball49:
			sel = 99
		case ball49 = <-ch49:
			sel = 100
		case ch50 <- ball50:
			sel = 101
		case ball50 = <-ch50:
			sel = 102
		case ch51 <- ball51:
			sel = 103
		case ball51 = <-ch51:
			sel = 104
		case ch52 <- ball52:
			sel = 105
		case ball52 = <-ch52:
			sel = 106
		case ch53 <- ball53:
			sel = 107
		case ball53 = <-ch53:
			sel = 108
		case ch54 <- ball54:
			sel = 109
		case ball54 = <-ch54:
			sel = 110
		case ch55 <- ball55:
			sel = 111
		case ball55 = <-ch55:
			sel = 112
		case ch56 <- ball56:
			sel = 113
		case ball56 = <-ch56:
			sel = 114
		case ch57 <- ball57:
			sel = 115
		case ball57 = <-ch57:
			sel = 116
		case ch58 <- ball58:
			sel = 117
		case ball58 = <-ch58:
			sel = 118
		case ch59 <- ball59:
			sel = 119
		case ball59 = <-ch59:
			sel = 120
		case ch60 <- ball60:
			sel = 121
		case ball60 = <-ch60:
			sel = 122
		case ch61 <- ball61:
			sel = 123
		case ball61 = <-ch61:
			sel = 124
		case ch62 <- ball62:
			sel = 125
		case ball62 = <-ch62:
			sel = 126
		case ch63 <- ball63:
			sel = 127
		case ball63 = <-ch63:
			sel = 128
		case ch64 <- ball64:
			sel = 129
		case ball64 = <-ch64:
			sel = 130
		case ch65 <- ball65:
			sel = 131
		case ball65 = <-ch65:
			sel = 132
		case ch66 <- ball66:
			sel = 133
		case ball66 = <-ch66:
			sel = 134
		case ch67 <- ball67:
			sel = 135
		case ball67 = <-ch67:
			sel = 136
		case ch68 <- ball68:
			sel = 137
		case ball68 = <-ch68:
			sel = 138
		case ch69 <- ball69:
			sel = 139
		case ball69 = <-ch69:
			sel = 140
		case ch70 <- ball70:
			sel = 141
		case ball70 = <-ch70:
			sel = 142
		case ch71 <- ball71:
			sel = 143
		case ball71 = <-ch71:
			sel = 144
		case ch72 <- ball72:
			sel = 145
		case ball72 = <-ch72:
			sel = 146
		case ch73 <- ball73:
			sel = 147
		case ball73 = <-ch73:
			sel = 148
		case ch74 <- ball74:
			sel = 149
		case ball74 = <-ch74:
			sel = 150
		case ch75 <- ball75:
			sel = 151
		case ball75 = <-ch75:
			sel = 152
		case ch76 <- ball76:
			sel = 153
		case ball76 = <-ch76:
			sel = 154
		case ch77 <- ball77:
			sel = 155
		case ball77 = <-ch77:
			sel = 156
		case ch78 <- ball78:
			sel = 157
		case ball78 = <-ch78:
			sel = 158
		case ch79 <- ball79:
			sel = 159
		case ball79 = <-ch79:
			sel = 160
		case ch80 <- ball80:
			sel = 161
		case ball80 = <-ch80:
			sel = 162
		case ch81 <- ball81:
			sel = 163
		case ball81 = <-ch81:
			sel = 164
		case ch82 <- ball82:
			sel = 165
		case ball82 = <-ch82:
			sel = 166
		case ch83 <- ball83:
			sel = 167
		case ball83 = <-ch83:
			sel = 168
		case ch84 <- ball84:
			sel = 169
		case ball84 = <-ch84:
			sel = 170
		case ch85 <- ball85:
			sel = 171
		case ball85 = <-ch85:
			sel = 172
		case ch86 <- ball86:
			sel = 173
		case ball86 = <-ch86:
			sel = 174
		case ch87 <- ball87:
			sel = 175
		case ball87 = <-ch87:
			sel = 176
		case ch88 <- ball88:
			sel = 177
		case ball88 = <-ch88:
			sel = 178
		case ch89 <- ball89:
			sel = 179
		case ball89 = <-ch89:
			sel = 180
		case ch90 <- ball90:
			sel = 181
		case ball90 = <-ch90:
			sel = 182
		case ch91 <- ball91:
			sel = 183
		case ball91 = <-ch91:
			sel = 184
		case ch92 <- ball92:
			sel = 185
		case ball92 = <-ch92:
			sel = 186
		case ch93 <- ball93:
			sel = 187
		case ball93 = <-ch93:
			sel = 188
		case ch94 <- ball94:
			sel = 189
		case ball94 = <-ch94:
			sel = 190
		case ch95 <- ball95:
			sel = 191
		case ball95 = <-ch95:
			sel = 192
		case ch96 <- ball96:
			sel = 193
		case ball96 = <-ch96:
			sel = 194
		case ch97 <- ball97:
			sel = 195
		case ball97 = <-ch97:
			sel = 196
		case ch98 <- ball98:
			sel = 197
		case ball98 = <-ch98:
			sel = 198
		case ch99 <- ball99:
			sel = 199
		case ball99 = <-ch99:
			sel = 200
		case ch100 <- ball100:
			sel = 201
		case ball100 = <-ch100:
			sel = 202
		case ch101 <- ball101:
			sel = 203
		case ball101 = <-ch101:
			sel = 204
		case ch102 <- ball102:
			sel = 205
		case ball102 = <-ch102:
			sel = 206
		case ch103 <- ball103:
			sel = 207
		case ball103 = <-ch103:
			sel = 208
		case ch104 <- ball104:
			sel = 209
		case ball104 = <-ch104:
			sel = 210
		case ch105 <- ball105:
			sel = 211
		case ball105 = <-ch105:
			sel = 212
		case ch106 <- ball106:
			sel = 213
		case ball106 = <-ch106:
			sel = 214
		case ch107 <- ball107:
			sel = 215
		case ball107 = <-ch107:
			sel = 216
		case ch108 <- ball108:
			sel = 217
		case ball108 = <-ch108:
			sel = 218
		case ch109 <- ball109:
			sel = 219
		case ball109 = <-ch109:
			sel = 220
		case ch110 <- ball110:
			sel = 221
		case ball110 = <-ch110:
			sel = 222
		case ch111 <- ball111:
			sel = 223
		case ball111 = <-ch111:
			sel = 224
		case ch112 <- ball112:
			sel = 225
		case ball112 = <-ch112:
			sel = 226
		case ch113 <- ball113:
			sel = 227
		case ball113 = <-ch113:
			sel = 228
		case ch114 <- ball114:
			sel = 229
		case ball114 = <-ch114:
			sel = 230
		case ch115 <- ball115:
			sel = 231
		case ball115 = <-ch115:
			sel = 232
		case ch116 <- ball116:
			sel = 233
		case ball116 = <-ch116:
			sel = 234
		case ch117 <- ball117:
			sel = 235
		case ball117 = <-ch117:
			sel = 236
		case ch118 <- ball118:
			sel = 237
		case ball118 = <-ch118:
			sel = 238
		case ch119 <- ball119:
			sel = 239
		case ball119 = <-ch119:
			sel = 240
		case ch120 <- ball120:
			sel = 241
		case ball120 = <-ch120:
			sel = 242
		case ch121 <- ball121:
			sel = 243
		case ball121 = <-ch121:
			sel = 244
		case ch122 <- ball122:
			sel = 245
		case ball122 = <-ch122:
			sel = 246
		case ch123 <- ball123:
			sel = 247
		case ball123 = <-ch123:
			sel = 248
		case ch124 <- ball124:
			sel = 249
		case ball124 = <-ch124:
			sel = 250
		case ch125 <- ball125:
			sel = 251
		case ball125 = <-ch125:
			sel = 252
		case ch126 <- ball126:
			sel = 253
		case ball126 = <-ch126:
			sel = 254
		case ch127 <- ball127:
			sel = 255
		case ball127 = <-ch127:
			sel = 256
		case ch128 <- ball128:
			sel = 257
		case ball128 = <-ch128:
			sel = 258
		case ch129 <- ball129:
			sel = 259
		case ball129 = <-ch129:
			sel = 260
		case ch130 <- ball130:
			sel = 261
		case ball130 = <-ch130:
			sel = 262
		case ch131 <- ball131:
			sel = 263
		case ball131 = <-ch131:
			sel = 264
		case ch132 <- ball132:
			sel = 265
		case ball132 = <-ch132:
			sel = 266
		case ch133 <- ball133:
			sel = 267
		case ball133 = <-ch133:
			sel = 268
		case ch134 <- ball134:
			sel = 269
		case ball134 = <-ch134:
			sel = 270
		case ch135 <- ball135:
			sel = 271
		case ball135 = <-ch135:
			sel = 272
		case ch136 <- ball136:
			sel = 273
		case ball136 = <-ch136:
			sel = 274
		case ch137 <- ball137:
			sel = 275
		case ball137 = <-ch137:
			sel = 276
		case ch138 <- ball138:
			sel = 277
		case ball138 = <-ch138:
			sel = 278
		case ch139 <- ball139:
			sel = 279
		case ball139 = <-ch139:
			sel = 280
		case ch140 <- ball140:
			sel = 281
		case ball140 = <-ch140:
			sel = 282
		case ch141 <- ball141:
			sel = 283
		case ball141 = <-ch141:
			sel = 284
		case ch142 <- ball142:
			sel = 285
		case ball142 = <-ch142:
			sel = 286
		case ch143 <- ball143:
			sel = 287
		case ball143 = <-ch143:
			sel = 288
		case ch144 <- ball144:
			sel = 289
		case ball144 = <-ch144:
			sel = 290
		case ch145 <- ball145:
			sel = 291
		case ball145 = <-ch145:
			sel = 292
		case ch146 <- ball146:
			sel = 293
		case ball146 = <-ch146:
			sel = 294
		case ch147 <- ball147:
			sel = 295
		case ball147 = <-ch147:
			sel = 296
		case ch148 <- ball148:
			sel = 297
		case ball148 = <-ch148:
			sel = 298
		case ch149 <- ball149:
			sel = 299
		case ball149 = <-ch149:
			sel = 300
		case ch150 <- ball150:
			sel = 301
		case ball150 = <-ch150:
			sel = 302
		case ch151 <- ball151:
			sel = 303
		case ball151 = <-ch151:
			sel = 304
		case ch152 <- ball152:
			sel = 305
		case ball152 = <-ch152:
			sel = 306
		case ch153 <- ball153:
			sel = 307
		case ball153 = <-ch153:
			sel = 308
		case ch154 <- ball154:
			sel = 309
		case ball154 = <-ch154:
			sel = 310
		case ch155 <- ball155:
			sel = 311
		case ball155 = <-ch155:
			sel = 312
		case ch156 <- ball156:
			sel = 313
		case ball156 = <-ch156:
			sel = 314
		case ch157 <- ball157:
			sel = 315
		case ball157 = <-ch157:
			sel = 316
		case ch158 <- ball158:
			sel = 317
		case ball158 = <-ch158:
			sel = 318
		case ch159 <- ball159:
			sel = 319
		case ball159 = <-ch159:
			sel = 320
		case ch160 <- ball160:
			sel = 321
		case ball160 = <-ch160:
			sel = 322
		case ch161 <- ball161:
			sel = 323
		case ball161 = <-ch161:
			sel = 324
		case ch162 <- ball162:
			sel = 325
		case ball162 = <-ch162:
			sel = 326
		case ch163 <- ball163:
			sel = 327
		case ball163 = <-ch163:
			sel = 328
		case ch164 <- ball164:
			sel = 329
		case ball164 = <-ch164:
			sel = 330
		case ch165 <- ball165:
			sel = 331
		case ball165 = <-ch165:
			sel = 332
		case ch166 <- ball166:
			sel = 333
		case ball166 = <-ch166:
			sel = 334
		case ch167 <- ball167:
			sel = 335
		case ball167 = <-ch167:
			sel = 336
		case ch168 <- ball168:
			sel = 337
		case ball168 = <-ch168:
			sel = 338
		case ch169 <- ball169:
			sel = 339
		case ball169 = <-ch169:
			sel = 340
		case ch170 <- ball170:
			sel = 341
		case ball170 = <-ch170:
			sel = 342
		case ch171 <- ball171:
			sel = 343
		case ball171 = <-ch171:
			sel = 344
		case ch172 <- ball172:
			sel = 345
		case ball172 = <-ch172:
			sel = 346
		case ch173 <- ball173:
			sel = 347
		case ball173 = <-ch173:
			sel = 348
		case ch174 <- ball174:
			sel = 349
		case ball174 = <-ch174:
			sel = 350
		case ch175 <- ball175:
			sel = 351
		case ball175 = <-ch175:
			sel = 352
		case ch176 <- ball176:
			sel = 353
		case ball176 = <-ch176:
			sel = 354
		case ch177 <- ball177:
			sel = 355
		case ball177 = <-ch177:
			sel = 356
		case ch178 <- ball178:
			sel = 357
		case ball178 = <-ch178:
			sel = 358
		case ch179 <- ball179:
			sel = 359
		case ball179 = <-ch179:
			sel = 360
		case ch180 <- ball180:
			sel = 361
		case ball180 = <-ch180:
			sel = 362
		case ch181 <- ball181:
			sel = 363
		case ball181 = <-ch181:
			sel = 364
		case ch182 <- ball182:
			sel = 365
		case ball182 = <-ch182:
			sel = 366
		case ch183 <- ball183:
			sel = 367
		case ball183 = <-ch183:
			sel = 368
		case ch184 <- ball184:
			sel = 369
		case ball184 = <-ch184:
			sel = 370
		case ch185 <- ball185:
			sel = 371
		case ball185 = <-ch185:
			sel = 372
		case ch186 <- ball186:
			sel = 373
		case ball186 = <-ch186:
			sel = 374
		case ch187 <- ball187:
			sel = 375
		case ball187 = <-ch187:
			sel = 376
		case ch188 <- ball188:
			sel = 377
		case ball188 = <-ch188:
			sel = 378
		case ch189 <- ball189:
			sel = 379
		case ball189 = <-ch189:
			sel = 380
		case ch190 <- ball190:
			sel = 381
		case ball190 = <-ch190:
			sel = 382
		case ch191 <- ball191:
			sel = 383
		case ball191 = <-ch191:
			sel = 384
		case ch192 <- ball192:
			sel = 385
		case ball192 = <-ch192:
			sel = 386
		case ch193 <- ball193:
			sel = 387
		case ball193 = <-ch193:
			sel = 388
		case ch194 <- ball194:
			sel = 389
		case ball194 = <-ch194:
			sel = 390
		case ch195 <- ball195:
			sel = 391
		case ball195 = <-ch195:
			sel = 392
		case ch196 <- ball196:
			sel = 393
		case ball196 = <-ch196:
			sel = 394
		case ch197 <- ball197:
			sel = 395
		case ball197 = <-ch197:
			sel = 396
		case ch198 <- ball198:
			sel = 397
		case ball198 = <-ch198:
			sel = 398
		case ch199 <- ball199:
			sel = 399
		case ball199 = <-ch199:
			sel = 400
		case ch200 <- ball200:
			sel = 401
		case ball200 = <-ch200:
			sel = 402
		case ch201 <- ball201:
			sel = 403
		case ball201 = <-ch201:
			sel = 404
		case ch202 <- ball202:
			sel = 405
		case ball202 = <-ch202:
			sel = 406
		case ch203 <- ball203:
			sel = 407
		case ball203 = <-ch203:
			sel = 408
		case ch204 <- ball204:
			sel = 409
		case ball204 = <-ch204:
			sel = 410
		case ch205 <- ball205:
			sel = 411
		case ball205 = <-ch205:
			sel = 412
		case ch206 <- ball206:
			sel = 413
		case ball206 = <-ch206:
			sel = 414
		case ch207 <- ball207:
			sel = 415
		case ball207 = <-ch207:
			sel = 416
		case ch208 <- ball208:
			sel = 417
		case ball208 = <-ch208:
			sel = 418
		case ch209 <- ball209:
			sel = 419
		case ball209 = <-ch209:
			sel = 420
		case ch210 <- ball210:
			sel = 421
		case ball210 = <-ch210:
			sel = 422
		case ch211 <- ball211:
			sel = 423
		case ball211 = <-ch211:
			sel = 424
		case ch212 <- ball212:
			sel = 425
		case ball212 = <-ch212:
			sel = 426
		case ch213 <- ball213:
			sel = 427
		case ball213 = <-ch213:
			sel = 428
		case ch214 <- ball214:
			sel = 429
		case ball214 = <-ch214:
			sel = 430
		case ch215 <- ball215:
			sel = 431
		case ball215 = <-ch215:
			sel = 432
		case ch216 <- ball216:
			sel = 433
		case ball216 = <-ch216:
			sel = 434
		case ch217 <- ball217:
			sel = 435
		case ball217 = <-ch217:
			sel = 436
		case ch218 <- ball218:
			sel = 437
		case ball218 = <-ch218:
			sel = 438
		case ch219 <- ball219:
			sel = 439
		case ball219 = <-ch219:
			sel = 440
		case ch220 <- ball220:
			sel = 441
		case ball220 = <-ch220:
			sel = 442
		case ch221 <- ball221:
			sel = 443
		case ball221 = <-ch221:
			sel = 444
		case ch222 <- ball222:
			sel = 445
		case ball222 = <-ch222:
			sel = 446
		case ch223 <- ball223:
			sel = 447
		case ball223 = <-ch223:
			sel = 448
		case ch224 <- ball224:
			sel = 449
		case ball224 = <-ch224:
			sel = 450
		case ch225 <- ball225:
			sel = 451
		case ball225 = <-ch225:
			sel = 452
		case ch226 <- ball226:
			sel = 453
		case ball226 = <-ch226:
			sel = 454
		case ch227 <- ball227:
			sel = 455
		case ball227 = <-ch227:
			sel = 456
		case ch228 <- ball228:
			sel = 457
		case ball228 = <-ch228:
			sel = 458
		case ch229 <- ball229:
			sel = 459
		case ball229 = <-ch229:
			sel = 460
		case ch230 <- ball230:
			sel = 461
		case ball230 = <-ch230:
			sel = 462
		case ch231 <- ball231:
			sel = 463
		case ball231 = <-ch231:
			sel = 464
		case ch232 <- ball232:
			sel = 465
		case ball232 = <-ch232:
			sel = 466
		case ch233 <- ball233:
			sel = 467
		case ball233 = <-ch233:
			sel = 468
		case ch234 <- ball234:
			sel = 469
		case ball234 = <-ch234:
			sel = 470
		case ch235 <- ball235:
			sel = 471
		case ball235 = <-ch235:
			sel = 472
		case ch236 <- ball236:
			sel = 473
		case ball236 = <-ch236:
			sel = 474
		case ch237 <- ball237:
			sel = 475
		case ball237 = <-ch237:
			sel = 476
		case ch238 <- ball238:
			sel = 477
		case ball238 = <-ch238:
			sel = 478
		case ch239 <- ball239:
			sel = 479
		case ball239 = <-ch239:
			sel = 480
		case ch240 <- ball240:
			sel = 481
		case ball240 = <-ch240:
			sel = 482
		case ch241 <- ball241:
			sel = 483
		case ball241 = <-ch241:
			sel = 484
		case ch242 <- ball242:
			sel = 485
		case ball242 = <-ch242:
			sel = 486
		case ch243 <- ball243:
			sel = 487
		case ball243 = <-ch243:
			sel = 488
		case ch244 <- ball244:
			sel = 489
		case ball244 = <-ch244:
			sel = 490
		case ch245 <- ball245:
			sel = 491
		case ball245 = <-ch245:
			sel = 492
		case ch246 <- ball246:
			sel = 493
		case ball246 = <-ch246:
			sel = 494
		case ch247 <- ball247:
			sel = 495
		case ball247 = <-ch247:
			sel = 496
		case ch248 <- ball248:
			sel = 497
		case ball248 = <-ch248:
			sel = 498
		case ch249 <- ball249:
			sel = 499
		case ball249 = <-ch249:
			sel = 500
		case ch250 <- ball250:
			sel = 501
		case ball250 = <-ch250:
			sel = 502
		case ch251 <- ball251:
			sel = 503
		case ball251 = <-ch251:
			sel = 504
		case ch252 <- ball252:
			sel = 505
		case ball252 = <-ch252:
			sel = 506
		case ch253 <- ball253:
			sel = 507
		case ball253 = <-ch253:
			sel = 508
		case ch254 <- ball254:
			sel = 509
		case ball254 = <-ch254:
			sel = 510
		case ch255 <- ball255:
			sel = 511
		case ball255 = <-ch255:
			sel = 512

		default:
			sel = 10000
		}
	}

	bondgo.Void(ch0)
	bondgo.Void(ch1)
	bondgo.Void(ch2)
	bondgo.Void(ch3)
	bondgo.Void(ch4)
	bondgo.Void(ch5)
	bondgo.Void(ch6)
	bondgo.Void(ch7)
	bondgo.Void(ch8)
	bondgo.Void(ch9)
	bondgo.Void(ch10)
	bondgo.Void(ch11)
	bondgo.Void(ch12)
	bondgo.Void(ch13)
	bondgo.Void(ch14)
	bondgo.Void(ch15)
	bondgo.Void(ch16)
	bondgo.Void(ch17)
	bondgo.Void(ch18)
	bondgo.Void(ch19)
	bondgo.Void(ch20)
	bondgo.Void(ch21)
	bondgo.Void(ch22)
	bondgo.Void(ch23)
	bondgo.Void(ch24)
	bondgo.Void(ch25)
	bondgo.Void(ch26)
	bondgo.Void(ch27)
	bondgo.Void(ch28)
	bondgo.Void(ch29)
	bondgo.Void(ch30)
	bondgo.Void(ch31)
	bondgo.Void(ch32)
	bondgo.Void(ch33)
	bondgo.Void(ch34)
	bondgo.Void(ch35)
	bondgo.Void(ch36)
	bondgo.Void(ch37)
	bondgo.Void(ch38)
	bondgo.Void(ch39)
	bondgo.Void(ch40)
	bondgo.Void(ch41)
	bondgo.Void(ch42)
	bondgo.Void(ch43)
	bondgo.Void(ch44)
	bondgo.Void(ch45)
	bondgo.Void(ch46)
	bondgo.Void(ch47)
	bondgo.Void(ch48)
	bondgo.Void(ch49)
	bondgo.Void(ch50)
	bondgo.Void(ch51)
	bondgo.Void(ch52)
	bondgo.Void(ch53)
	bondgo.Void(ch54)
	bondgo.Void(ch55)
	bondgo.Void(ch56)
	bondgo.Void(ch57)
	bondgo.Void(ch58)
	bondgo.Void(ch59)
	bondgo.Void(ch60)
	bondgo.Void(ch61)
	bondgo.Void(ch62)
	bondgo.Void(ch63)
	bondgo.Void(ch64)
	bondgo.Void(ch65)
	bondgo.Void(ch66)
	bondgo.Void(ch67)
	bondgo.Void(ch68)
	bondgo.Void(ch69)
	bondgo.Void(ch70)
	bondgo.Void(ch71)
	bondgo.Void(ch72)
	bondgo.Void(ch73)
	bondgo.Void(ch74)
	bondgo.Void(ch75)
	bondgo.Void(ch76)
	bondgo.Void(ch77)
	bondgo.Void(ch78)
	bondgo.Void(ch79)
	bondgo.Void(ch80)
	bondgo.Void(ch81)
	bondgo.Void(ch82)
	bondgo.Void(ch83)
	bondgo.Void(ch84)
	bondgo.Void(ch85)
	bondgo.Void(ch86)
	bondgo.Void(ch87)
	bondgo.Void(ch88)
	bondgo.Void(ch89)
	bondgo.Void(ch90)
	bondgo.Void(ch91)
	bondgo.Void(ch92)
	bondgo.Void(ch93)
	bondgo.Void(ch94)
	bondgo.Void(ch95)
	bondgo.Void(ch96)
	bondgo.Void(ch97)
	bondgo.Void(ch98)
	bondgo.Void(ch99)
	bondgo.Void(ch100)
	bondgo.Void(ch101)
	bondgo.Void(ch102)
	bondgo.Void(ch103)
	bondgo.Void(ch104)
	bondgo.Void(ch105)
	bondgo.Void(ch106)
	bondgo.Void(ch107)
	bondgo.Void(ch108)
	bondgo.Void(ch109)
	bondgo.Void(ch110)
	bondgo.Void(ch111)
	bondgo.Void(ch112)
	bondgo.Void(ch113)
	bondgo.Void(ch114)
	bondgo.Void(ch115)
	bondgo.Void(ch116)
	bondgo.Void(ch117)
	bondgo.Void(ch118)
	bondgo.Void(ch119)
	bondgo.Void(ch120)
	bondgo.Void(ch121)
	bondgo.Void(ch122)
	bondgo.Void(ch123)
	bondgo.Void(ch124)
	bondgo.Void(ch125)
	bondgo.Void(ch126)
	bondgo.Void(ch127)
	bondgo.Void(ch128)
	bondgo.Void(ch129)
	bondgo.Void(ch130)
	bondgo.Void(ch131)
	bondgo.Void(ch132)
	bondgo.Void(ch133)
	bondgo.Void(ch134)
	bondgo.Void(ch135)
	bondgo.Void(ch136)
	bondgo.Void(ch137)
	bondgo.Void(ch138)
	bondgo.Void(ch139)
	bondgo.Void(ch140)
	bondgo.Void(ch141)
	bondgo.Void(ch142)
	bondgo.Void(ch143)
	bondgo.Void(ch144)
	bondgo.Void(ch145)
	bondgo.Void(ch146)
	bondgo.Void(ch147)
	bondgo.Void(ch148)
	bondgo.Void(ch149)
	bondgo.Void(ch150)
	bondgo.Void(ch151)
	bondgo.Void(ch152)
	bondgo.Void(ch153)
	bondgo.Void(ch154)
	bondgo.Void(ch155)
	bondgo.Void(ch156)
	bondgo.Void(ch157)
	bondgo.Void(ch158)
	bondgo.Void(ch159)
	bondgo.Void(ch160)
	bondgo.Void(ch161)
	bondgo.Void(ch162)
	bondgo.Void(ch163)
	bondgo.Void(ch164)
	bondgo.Void(ch165)
	bondgo.Void(ch166)
	bondgo.Void(ch167)
	bondgo.Void(ch168)
	bondgo.Void(ch169)
	bondgo.Void(ch170)
	bondgo.Void(ch171)
	bondgo.Void(ch172)
	bondgo.Void(ch173)
	bondgo.Void(ch174)
	bondgo.Void(ch175)
	bondgo.Void(ch176)
	bondgo.Void(ch177)
	bondgo.Void(ch178)
	bondgo.Void(ch179)
	bondgo.Void(ch180)
	bondgo.Void(ch181)
	bondgo.Void(ch182)
	bondgo.Void(ch183)
	bondgo.Void(ch184)
	bondgo.Void(ch185)
	bondgo.Void(ch186)
	bondgo.Void(ch187)
	bondgo.Void(ch188)
	bondgo.Void(ch189)
	bondgo.Void(ch190)
	bondgo.Void(ch191)
	bondgo.Void(ch192)
	bondgo.Void(ch193)
	bondgo.Void(ch194)
	bondgo.Void(ch195)
	bondgo.Void(ch196)
	bondgo.Void(ch197)
	bondgo.Void(ch198)
	bondgo.Void(ch199)
	bondgo.Void(ch200)
	bondgo.Void(ch201)
	bondgo.Void(ch202)
	bondgo.Void(ch203)
	bondgo.Void(ch204)
	bondgo.Void(ch205)
	bondgo.Void(ch206)
	bondgo.Void(ch207)
	bondgo.Void(ch208)
	bondgo.Void(ch209)
	bondgo.Void(ch210)
	bondgo.Void(ch211)
	bondgo.Void(ch212)
	bondgo.Void(ch213)
	bondgo.Void(ch214)
	bondgo.Void(ch215)
	bondgo.Void(ch216)
	bondgo.Void(ch217)
	bondgo.Void(ch218)
	bondgo.Void(ch219)
	bondgo.Void(ch220)
	bondgo.Void(ch221)
	bondgo.Void(ch222)
	bondgo.Void(ch223)
	bondgo.Void(ch224)
	bondgo.Void(ch225)
	bondgo.Void(ch226)
	bondgo.Void(ch227)
	bondgo.Void(ch228)
	bondgo.Void(ch229)
	bondgo.Void(ch230)
	bondgo.Void(ch231)
	bondgo.Void(ch232)
	bondgo.Void(ch233)
	bondgo.Void(ch234)
	bondgo.Void(ch235)
	bondgo.Void(ch236)
	bondgo.Void(ch237)
	bondgo.Void(ch238)
	bondgo.Void(ch239)
	bondgo.Void(ch240)
	bondgo.Void(ch241)
	bondgo.Void(ch242)
	bondgo.Void(ch243)
	bondgo.Void(ch244)
	bondgo.Void(ch245)
	bondgo.Void(ch246)
	bondgo.Void(ch247)
	bondgo.Void(ch248)
	bondgo.Void(ch249)
	bondgo.Void(ch250)
	bondgo.Void(ch251)
	bondgo.Void(ch252)
	bondgo.Void(ch253)
	bondgo.Void(ch254)
	bondgo.Void(ch255)
	bondgo.Void(ball0)
	bondgo.Void(ball1)
	bondgo.Void(ball2)
	bondgo.Void(ball3)
	bondgo.Void(ball4)
	bondgo.Void(ball5)
	bondgo.Void(ball6)
	bondgo.Void(ball7)
	bondgo.Void(ball8)
	bondgo.Void(ball9)
	bondgo.Void(ball10)
	bondgo.Void(ball11)
	bondgo.Void(ball12)
	bondgo.Void(ball13)
	bondgo.Void(ball14)
	bondgo.Void(ball15)
	bondgo.Void(ball16)
	bondgo.Void(ball17)
	bondgo.Void(ball18)
	bondgo.Void(ball19)
	bondgo.Void(ball20)
	bondgo.Void(ball21)
	bondgo.Void(ball22)
	bondgo.Void(ball23)
	bondgo.Void(ball24)
	bondgo.Void(ball25)
	bondgo.Void(ball26)
	bondgo.Void(ball27)
	bondgo.Void(ball28)
	bondgo.Void(ball29)
	bondgo.Void(ball30)
	bondgo.Void(ball31)
	bondgo.Void(ball32)
	bondgo.Void(ball33)
	bondgo.Void(ball34)
	bondgo.Void(ball35)
	bondgo.Void(ball36)
	bondgo.Void(ball37)
	bondgo.Void(ball38)
	bondgo.Void(ball39)
	bondgo.Void(ball40)
	bondgo.Void(ball41)
	bondgo.Void(ball42)
	bondgo.Void(ball43)
	bondgo.Void(ball44)
	bondgo.Void(ball45)
	bondgo.Void(ball46)
	bondgo.Void(ball47)
	bondgo.Void(ball48)
	bondgo.Void(ball49)
	bondgo.Void(ball50)
	bondgo.Void(ball51)
	bondgo.Void(ball52)
	bondgo.Void(ball53)
	bondgo.Void(ball54)
	bondgo.Void(ball55)
	bondgo.Void(ball56)
	bondgo.Void(ball57)
	bondgo.Void(ball58)
	bondgo.Void(ball59)
	bondgo.Void(ball60)
	bondgo.Void(ball61)
	bondgo.Void(ball62)
	bondgo.Void(ball63)
	bondgo.Void(ball64)
	bondgo.Void(ball65)
	bondgo.Void(ball66)
	bondgo.Void(ball67)
	bondgo.Void(ball68)
	bondgo.Void(ball69)
	bondgo.Void(ball70)
	bondgo.Void(ball71)
	bondgo.Void(ball72)
	bondgo.Void(ball73)
	bondgo.Void(ball74)
	bondgo.Void(ball75)
	bondgo.Void(ball76)
	bondgo.Void(ball77)
	bondgo.Void(ball78)
	bondgo.Void(ball79)
	bondgo.Void(ball80)
	bondgo.Void(ball81)
	bondgo.Void(ball82)
	bondgo.Void(ball83)
	bondgo.Void(ball84)
	bondgo.Void(ball85)
	bondgo.Void(ball86)
	bondgo.Void(ball87)
	bondgo.Void(ball88)
	bondgo.Void(ball89)
	bondgo.Void(ball90)
	bondgo.Void(ball91)
	bondgo.Void(ball92)
	bondgo.Void(ball93)
	bondgo.Void(ball94)
	bondgo.Void(ball95)
	bondgo.Void(ball96)
	bondgo.Void(ball97)
	bondgo.Void(ball98)
	bondgo.Void(ball99)
	bondgo.Void(ball100)
	bondgo.Void(ball101)
	bondgo.Void(ball102)
	bondgo.Void(ball103)
	bondgo.Void(ball104)
	bondgo.Void(ball105)
	bondgo.Void(ball106)
	bondgo.Void(ball107)
	bondgo.Void(ball108)
	bondgo.Void(ball109)
	bondgo.Void(ball110)
	bondgo.Void(ball111)
	bondgo.Void(ball112)
	bondgo.Void(ball113)
	bondgo.Void(ball114)
	bondgo.Void(ball115)
	bondgo.Void(ball116)
	bondgo.Void(ball117)
	bondgo.Void(ball118)
	bondgo.Void(ball119)
	bondgo.Void(ball120)
	bondgo.Void(ball121)
	bondgo.Void(ball122)
	bondgo.Void(ball123)
	bondgo.Void(ball124)
	bondgo.Void(ball125)
	bondgo.Void(ball126)
	bondgo.Void(ball127)
	bondgo.Void(ball128)
	bondgo.Void(ball129)
	bondgo.Void(ball130)
	bondgo.Void(ball131)
	bondgo.Void(ball132)
	bondgo.Void(ball133)
	bondgo.Void(ball134)
	bondgo.Void(ball135)
	bondgo.Void(ball136)
	bondgo.Void(ball137)
	bondgo.Void(ball138)
	bondgo.Void(ball139)
	bondgo.Void(ball140)
	bondgo.Void(ball141)
	bondgo.Void(ball142)
	bondgo.Void(ball143)
	bondgo.Void(ball144)
	bondgo.Void(ball145)
	bondgo.Void(ball146)
	bondgo.Void(ball147)
	bondgo.Void(ball148)
	bondgo.Void(ball149)
	bondgo.Void(ball150)
	bondgo.Void(ball151)
	bondgo.Void(ball152)
	bondgo.Void(ball153)
	bondgo.Void(ball154)
	bondgo.Void(ball155)
	bondgo.Void(ball156)
	bondgo.Void(ball157)
	bondgo.Void(ball158)
	bondgo.Void(ball159)
	bondgo.Void(ball160)
	bondgo.Void(ball161)
	bondgo.Void(ball162)
	bondgo.Void(ball163)
	bondgo.Void(ball164)
	bondgo.Void(ball165)
	bondgo.Void(ball166)
	bondgo.Void(ball167)
	bondgo.Void(ball168)
	bondgo.Void(ball169)
	bondgo.Void(ball170)
	bondgo.Void(ball171)
	bondgo.Void(ball172)
	bondgo.Void(ball173)
	bondgo.Void(ball174)
	bondgo.Void(ball175)
	bondgo.Void(ball176)
	bondgo.Void(ball177)
	bondgo.Void(ball178)
	bondgo.Void(ball179)
	bondgo.Void(ball180)
	bondgo.Void(ball181)
	bondgo.Void(ball182)
	bondgo.Void(ball183)
	bondgo.Void(ball184)
	bondgo.Void(ball185)
	bondgo.Void(ball186)
	bondgo.Void(ball187)
	bondgo.Void(ball188)
	bondgo.Void(ball189)
	bondgo.Void(ball190)
	bondgo.Void(ball191)
	bondgo.Void(ball192)
	bondgo.Void(ball193)
	bondgo.Void(ball194)
	bondgo.Void(ball195)
	bondgo.Void(ball196)
	bondgo.Void(ball197)
	bondgo.Void(ball198)
	bondgo.Void(ball199)
	bondgo.Void(ball200)
	bondgo.Void(ball201)
	bondgo.Void(ball202)
	bondgo.Void(ball203)
	bondgo.Void(ball204)
	bondgo.Void(ball205)
	bondgo.Void(ball206)
	bondgo.Void(ball207)
	bondgo.Void(ball208)
	bondgo.Void(ball209)
	bondgo.Void(ball210)
	bondgo.Void(ball211)
	bondgo.Void(ball212)
	bondgo.Void(ball213)
	bondgo.Void(ball214)
	bondgo.Void(ball215)
	bondgo.Void(ball216)
	bondgo.Void(ball217)
	bondgo.Void(ball218)
	bondgo.Void(ball219)
	bondgo.Void(ball220)
	bondgo.Void(ball221)
	bondgo.Void(ball222)
	bondgo.Void(ball223)
	bondgo.Void(ball224)
	bondgo.Void(ball225)
	bondgo.Void(ball226)
	bondgo.Void(ball227)
	bondgo.Void(ball228)
	bondgo.Void(ball229)
	bondgo.Void(ball230)
	bondgo.Void(ball231)
	bondgo.Void(ball232)
	bondgo.Void(ball233)
	bondgo.Void(ball234)
	bondgo.Void(ball235)
	bondgo.Void(ball236)
	bondgo.Void(ball237)
	bondgo.Void(ball238)
	bondgo.Void(ball239)
	bondgo.Void(ball240)
	bondgo.Void(ball241)
	bondgo.Void(ball242)
	bondgo.Void(ball243)
	bondgo.Void(ball244)
	bondgo.Void(ball245)
	bondgo.Void(ball246)
	bondgo.Void(ball247)
	bondgo.Void(ball248)
	bondgo.Void(ball249)
	bondgo.Void(ball250)
	bondgo.Void(ball251)
	bondgo.Void(ball252)
	bondgo.Void(ball253)
	bondgo.Void(ball254)
	bondgo.Void(ball255)

	bondgo.Void(sel)
}
