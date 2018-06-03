package svg

import (
	"fmt"

	"github.com/2enhance/placeholdr/config/providers"
	svgo "github.com/ajstarks/svgo"
)

func body(canvas *svgo.SVG, color *providers.Color) *svgo.SVG {
	styling := fmt.Sprintf("fill:%s", color.Format())

	// body
	canvas.Path("M66.4828109,115.845289 L66.4828109,115.851308 L0.820746269,115.851308 L0.820746269,115.845289 C0.820746269,112.137632 1.43357214,108.566627 2.57253731,105.24304 C2.77661194,104.642856 2.99934826,104.048692 3.24074627,103.460547 L3.24074627,103.454527 C3.47612438,102.866383 3.73558209,102.290881 4.00828358,101.722 C5.09126368,99.4380498 6.44092537,97.302791 8.00670149,95.3469254 C8.40281095,94.8575075 8.80554229,94.3813333 9.21971144,93.9232189 C9.42378607,93.6878408 9.64050249,93.4590846 9.85721891,93.2297264 C10.0739353,93.00699 10.2906517,92.7842537 10.5133881,92.5615174 C11.6210498,91.4532537 12.8093781,90.4382985 14.0597114,89.50401 C14.6851791,89.0398756 15.3160647,88.6064428 15.9722338,88.1856517 C16.2076119,88.0369602 16.4423881,87.8882687 16.6837861,87.746199 C16.7891343,87.684194 16.9005025,87.6221891 17.0058507,87.5601841 C17.3279154,87.3681493 17.6493781,87.1827363 17.9774627,87.0033433 C17.9901045,86.9907015 18.0021443,86.9907015 18.0021443,86.9907015 C18.2682239,86.8480299 18.5282836,86.6933184 18.7883433,86.5325871 C19.5619005,86.0684527 20.3047562,85.5543532 21.0102886,84.9854726 C23.281597,83.1782985 25.1634179,80.9003682 26.5124776,78.306995 C26.9398905,77.4840746 27.3173383,76.6298507 27.6327811,75.7509453 C28.4189801,73.5597015 28.846393,71.2017065 28.846393,68.7383632 L33.630408,68.7630448 L33.6430498,68.7630448 L38.5510746,68.7883284 C38.5510746,71.2203682 38.9658458,73.5542836 39.7327811,75.7202438 C40.042204,76.6057711 40.4136318,77.459393 40.8404428,78.2829154 C42.1834826,80.8883284 44.0592836,83.1782985 46.330592,84.9920945 C46.7760647,85.3569005 47.2281592,85.691607 47.7049353,86.0070498 C47.7542985,86.0443731 47.7976418,86.0750746 47.847005,86.1063781 C48.0697413,86.2550697 48.2924776,86.4037612 48.5278557,86.5398109 C48.7818955,86.6945224 49.0419552,86.8492338 49.3014129,86.9913035 C49.3447562,87.0039453 49.3820796,87.022005 49.4127811,87.0466866 C49.7035423,87.2080199 49.9949055,87.3747711 50.2790448,87.5481443 C50.5270647,87.6902139 50.7744826,87.8389055 51.0219005,87.9936169 C51.2753383,88.1483284 51.5293781,88.3090597 51.776796,88.4884527 C52.2902935,88.8165373 52.7917512,89.1759254 53.2805672,89.5467512 C53.2805672,89.5527711 53.2865871,89.5527711 53.2865871,89.5527711 L53.293209,89.5527711 C53.9180746,90.0169055 54.5248806,90.5123433 55.1190448,91.0192189 C55.713209,91.5266965 56.2826915,92.0594577 56.8335124,92.6102786 C62.7998358,98.5525224 66.4828109,106.759453 66.4828109,115.845289 Z", styling)
	// face
	canvas.Path("M58.8513831,33.5839502 L58.8513831,48.6914925 C58.8513831,48.963592 58.8453632,49.2362935 58.8393433,49.508393 C58.8333234,49.7184876 58.8267015,49.9291841 58.8146617,50.1392786 C58.8086418,50.3499751 58.796,50.5540498 58.7773383,50.7641443 C58.7713184,50.9062139 58.7586766,51.0488856 58.7466368,51.1915572 C58.7406169,51.2782438 58.734597,51.3583085 58.7219552,51.444995 C58.7099154,51.5997065 58.6972736,51.7604378 58.672592,51.9151493 C58.6599502,52.0638408 58.6479104,52.2125323 58.6232289,52.3606219 C58.6105871,52.5032935 58.5925274,52.6513831 58.5672438,52.8000746 C58.5118607,53.2455473 58.4372139,53.6910199 58.3565473,54.1304726 C58.3011642,54.4278557 58.2451791,54.7246368 58.1831741,55.016 C58.1771542,55.0286418 58.1771542,55.0406816 58.1711343,55.0533234 C58.1651144,55.0966667 58.1584925,55.14001 58.1464527,55.1833532 C58.1404328,55.195995 58.1404328,55.2080348 58.1338109,55.2206766 C58.1338109,55.2333184 58.1338109,55.2453582 58.127791,55.258 C58.1151493,55.320005 58.1031095,55.38201 58.0844478,55.4434129 C58.072408,55.5174577 58.0537463,55.5921045 58.0350846,55.6661493 C58.0230448,55.7468159 58.0037811,55.8208607 57.9790995,55.8955075 C57.9604378,56.0068756 57.9297363,56.1182438 57.9044527,56.2296119 C57.885791,56.2916169 57.8737512,56.3596418 57.8490697,56.4216468 C57.8364279,56.4776318 57.8183682,56.5330149 57.8057264,56.5829801 C57.6756965,57.0904577 57.527005,57.5852935 57.3722935,58.0807313 C57.2982488,58.3161095 57.223602,58.5448657 57.1435373,58.7802438 C57.1369154,58.8049254 57.1248756,58.829607 57.1182537,58.8542886 C57.044209,59.0649851 56.9695622,59.2750796 56.8955174,59.4857761 C56.8708358,59.5477811 56.8461542,59.6091841 56.8214726,59.6778109 C56.7468259,59.8758657 56.6727811,60.0739204 56.5927164,60.2719751 C56.50001,60.5133731 56.4006816,60.7487512 56.3019552,60.9841294 C56.2152687,61.1948259 56.1225622,61.4049204 56.0292537,61.6150149 C56.0172139,61.6523383 55.9985522,61.6830398 55.9859104,61.7203632 C55.8805622,61.9430995 55.7818358,62.1658358 55.6764876,62.3825522 C55.5711393,62.6113085 55.4597711,62.8406667 55.3423831,63.0700249 C55.1196468,63.5215174 54.8782488,63.9736119 54.6368507,64.4130647 L54.3027463,65.0072289 C53.349194,66.6476517 52.2908955,68.1887463 51.1398905,69.6118507 C50.9851791,69.8159254 50.8244478,70.0139801 50.6570945,70.2060149 C50.4277363,70.4847363 50.1929602,70.7568358 49.9575821,71.0229154 C49.7468856,71.2582935 49.536791,71.4930697 49.3200746,71.7224279 C48.7259104,72.3599353 48.1191045,72.9667413 47.5002587,73.5362239 C47.3762488,73.6602338 47.2468209,73.771602 47.1161891,73.88899 C46.9054925,74.0810249 46.695398,74.2664378 46.4786816,74.4458308 C46.3113284,74.5945224 46.1385572,74.736592 45.9645821,74.8792637 C45.5191095,75.2440697 45.0736368,75.5968358 44.6215423,75.9249204 C44.4355274,76.067592 44.2440945,76.2036418 44.0520597,76.3336716 C43.7426368,76.556408 43.426592,76.7671045 43.1111493,76.9711791 C42.9377761,77.0885672 42.7583831,77.1999353 42.57899,77.3052836 C42.0034876,77.6640697 41.4153433,77.9921542 40.8398408,78.2829154 C40.7844577,78.3136169 40.7350945,78.3382985 40.6791095,78.3629801 C40.524398,78.4436468 40.3696866,78.5176915 40.2149751,78.5923383 C40.0602637,78.6663831 39.8995323,78.7410299 39.7321791,78.8150746 C39.5281045,78.9077811 39.3240299,78.9944677 39.1193532,79.0751343 C39.0386866,79.1118557 38.9580199,79.1431592 38.8779552,79.1678408 C38.6365572,79.2665672 38.3885373,79.3598756 38.1477413,79.4399403 C37.8443383,79.5513085 37.5349154,79.6440149 37.2321144,79.7307015 C37.0214179,79.7927065 36.8053035,79.8480896 36.594607,79.8974527 C36.5819652,79.9040746 36.5699254,79.9100945 36.5512637,79.9100945 C36.3405672,79.9594577 36.1304726,80.0088209 35.9197761,80.0461443 C35.863791,80.0587861 35.808408,80.0708259 35.7524229,80.0768458 C35.5790498,80.1075473 35.4056766,80.1388507 35.2323035,80.1575124 C34.9848856,80.1942338 34.7374677,80.2255373 34.4900498,80.2381791 C34.4033632,80.2508209 34.3226965,80.2568408 34.2366119,80.2568408 C34.0385572,80.2688806 33.8405025,80.2755025 33.6484677,80.2755025 L33.629806,80.2755025 C33.4377711,80.2755025 33.2463383,80.2694826 33.0543035,80.2568408 C32.9796567,80.2568408 32.9056119,80.2508209 32.8315672,80.2381791 C32.658194,80.2315572 32.478801,80.2134975 32.2927861,80.1888159 C32.2554627,80.1888159 32.2121194,80.182796 32.175398,80.1761741 C31.9466418,80.1454726 31.7172836,80.1081493 31.4819055,80.064806 C31.2838507,80.0341045 31.085796,79.9907612 30.8817214,79.941398 C29.8234229,79.7 28.7344229,79.334592 27.6321791,78.8457761 C27.2607512,78.6784229 26.8833035,78.5056517 26.5118756,78.306995 C25.9363731,78.0222537 25.3608706,77.7001891 24.7913881,77.3474229 C24.3146119,77.0566617 23.8378358,76.7472388 23.3676816,76.4191542 C23.1323035,76.251801 22.8975274,76.0850498 22.6621493,75.9116766 C21.9259154,75.3728955 21.1950995,74.7853532 20.4775274,74.1538657 C20.2794726,73.9744726 20.0814179,73.7950796 19.8833632,73.6090647 C19.6853085,73.4296716 19.4872537,73.2442587 19.2952189,73.0642637 C18.8124229,72.6001294 18.3422687,72.1173333 17.8841542,71.6158756 C17.7047612,71.4238408 17.5253682,71.2263881 17.3453731,71.0217114 C16.8505373,70.4648706 16.3671393,69.8893682 15.9096269,69.2885821 C15.7302338,69.0718657 15.5628806,68.8491294 15.3961294,68.626393 C14.7712637,67.7974527 14.1770995,66.9305871 13.6196567,66.0330199 C13.4709652,65.7916219 13.3228756,65.543602 13.180204,65.302806 C13.0068308,65.0120448 12.8400796,64.7146617 12.6727264,64.4178806 C12.0538806,63.303597 11.484398,62.1465721 10.9889602,60.9516219 C10.9149154,60.7782488 10.8402687,60.6048756 10.7782637,60.4321045 C10.6795373,60.2093682 10.5928507,59.9806119 10.5121841,59.7572736 C10.3881741,59.4291891 10.2707861,59.1011045 10.1594179,58.7730199 C10.0480498,58.4449353 9.94270149,58.1108308 9.83735323,57.7767264 C9.73802488,57.4546617 9.64531841,57.133199 9.55863184,56.8051144 C9.45990547,56.4583682 9.367199,56.105602 9.28653234,55.7468159 C9.28051244,55.7094925 9.26787065,55.6721692 9.26185075,55.6354478 C9.14446269,55.1406119 9.03911443,54.6391542 8.94640796,54.1310746 C8.82239801,53.4628657 8.72367164,52.7820149 8.64902488,52.0951443 C8.61832338,51.8784279 8.59966169,51.6617114 8.581,51.444995 C8.55631841,51.2162388 8.53765672,50.9929005 8.52561692,50.7641443 C8.49491542,50.3493731 8.47625373,49.9285821 8.46361194,49.507791 C8.45759204,49.2350896 8.45097015,48.96299 8.45097015,48.6908905 L8.45097015,33.5839502 C8.45097015,30.1243134 9.15048259,26.825408 10.4128557,23.8232836 C10.4435572,23.736597 10.4808806,23.6565323 10.518204,23.5758657 C11.0076219,22.4435224 11.5765025,21.3539204 12.2266517,20.3142836 C12.3259801,20.1475323 12.4307264,19.986199 12.5360746,19.8254677 C17.0293284,12.9429154 24.797408,8.39367662 33.629204,8.38765672 L33.6538856,8.38765672 C35.5730299,8.38765672 37.441607,8.60437313 39.2367413,9.00650249 C40.8771642,9.37793035 42.4549801,9.91008955 43.9527313,10.5849204 C45.9952836,11.5005473 47.8951642,12.6768358 49.6036119,14.0818806 C51.0393582,15.2575672 52.3456766,16.5819453 53.4840398,18.0429751 C53.9295124,18.6124577 54.3503035,19.2 54.7470149,19.806806 C54.8583831,19.9735572 54.9697512,20.1409104 55.0690796,20.3142836 C55.712607,21.3479005 56.2881095,22.4314826 56.7769254,23.5638259 L56.8882935,23.8238856 C58.1259851,26.764005 58.8194776,29.9822438 58.8441592,33.3618159 C58.8513831,33.4352587 58.8513831,33.5093035 58.8513831,33.5839502 Z", styling)

	return canvas
}

func maleHair(canvas *svgo.SVG, color *providers.Color) *svgo.SVG {
	styling := fmt.Sprintf("fill:%s", color.Format())

	canvas.Gtransform("translate(8.303483, 0.019900)")
	canvas.Gstyle(styling)
	canvas.Path("M50.6242921,24 C50.6198354,24 33.7946352,24 0.148691542,24 C0.148691542,14.3904455 2.62668241,10.5133452 7.18648072,8.54921272 C14.2968745,5.48641167 19.3801066,7.5428145 25.5672473,8.23105292 C26.5238237,8.33730858 37.3041384,9.68525728 39.9210915,8.54921272 C46.7287627,5.59393697 48.6571077,0.976376394 48.5964298,0.976376394 C53.8625076,8.83280868 56.5600491,19.9859648 50.6242921,24 Z")
	canvas.Path("M50.6731144,23.8051891 L41.8870697,22.7565224 L42.0237214,23.1809254 C43.1608806,26.7176169 45.0782189,29.9280299 47.4747413,32.7670149 C49.0742289,34.6620796 50.6737164,37.7647363 50.6737164,42.5932985 C50.6731144,35.9786318 50.6731144,23.8051891 50.6731144,23.8051891 Z")
	canvas.Path("M8.93714428,22.7511045 L8.80109453,23.1779154 C8.73306965,23.38801 8.66504478,23.5987065 8.59039801,23.8027811 C8.52839303,23.9701343 8.47300995,24.1368856 8.40438308,24.2982189 C7.24735821,27.3930498 5.49556716,30.2278209 3.34766667,32.765209 C1.83787562,34.5537214 0.321462687,37.4258159 0.166751244,41.8137214 C0.154109453,42.0671592 0.148089552,42.321199 0.148089552,42.5872786 L0.148089552,33.5634478 C0.148089552,28.4513483 0.148089552,23.8027811 0.148089552,23.8027811 L2.21532338,23.5553632 L4.23319403,23.3139652 L6.60985075,23.0292239 L8.93714428,22.7511045 Z")
	canvas.Gend()
	canvas.Gend()

	return canvas
}

func femaleHairTop(canvas *svgo.SVG, color *providers.Color) *svgo.SVG {
	styling := fmt.Sprintf("fill:%s", color.Format())

	canvas.Gtransform("translate(-8, 0)")
	canvas.Path("M75.8943284,94.079403 C73.6687355,97.3251023 59.7940299,79.6220896 59.6650746,63.5402985 C59.6591045,62.2447761 59.7265672,60.9432836 59.8925373,59.6364179 C60.3223881,56.1683582 61.9922388,52.8352239 63.3791045,49.4776119 C64.4716418,46.8197015 65.398209,44.1558209 65.398209,41.4059701 C65.398209,26.6746269 61.1265672,31.4746269 49.2125373,23.5253731 C46.3337313,21.6101493 44.2895522,20.038806 42.6871642,18.761791 C38.2620896,15.2507463 37.1385075,13.9928358 31.0680597,13.9928358 C23.3456716,13.9928358 20.2770149,27.4041791 19.399403,33.1373134 C18.4543284,39.3062687 18.441791,49.3050746 22.2907463,58.8250746 C25.2728358,66.201791 25.2734328,73.4340299 19.399403,81.4686567 C17.1958209,84.4829851 12.7710586,102.6076 9.62746269,97.9492537 C-5.01159751,76.2563227 12.0819043,83.9098826 4.89391872,63.5402985 C-2.29406688,43.1707144 6.29086509,43.2001902 11.8680597,23.8125373 C15.9102726,9.76084555 27.2370149,0.618507463 32.2214925,4.17910448 C33.8847761,1.1958209 38.2370149,0.0788059701 42.6865672,0.0788059701 C45.9462687,0.0788059701 49.2668657,0.686567164 51.6238806,1.60119403 C69.8232836,8.66626866 70.1307463,25.5337313 71.1002985,32.5492537 C71.8985075,38.3438806 74.7522388,43.7576119 75.7528358,49.478209 C75.9677612,50.6997015 76.0967164,51.9271642 76.1032836,53.1671642 C76.1456716,59.8638806 76.6665932,61.9922039 81.7555484,67.5713084 C85.615847,71.8005621 79.199524,89.259262 75.8943284,94.079403 Z", styling)
	canvas.Gend()

	return canvas
}

func femaleHairBottom(canvas *svgo.SVG, color *providers.Color) *svgo.SVG {
	styling := fmt.Sprintf("fill:%s", color.Format())

	canvas.Gtransform("translate(-8, 0)")
	canvas.Path("M74.5634508,60.6853373 C83.9271158,80.972035 83.2761148,94.2693236 72.610448,100.577203 C54.8834024,111.061317 21.5229576,112.030621 7.11235917,100.577203 C-2.24019805,93.1438718 -1.31614919,78.4343538 9.88450577,56.4486487 L74.5634508,60.6853373 Z", styling)
	canvas.Gend()

	return canvas
}

func shirt(canvas *svgo.SVG, color *providers.Color) *svgo.SVG {
	styling := fmt.Sprintf("fill:%s", color.Format())

	canvas.Path("M66.4828109,115.845289 L0.820746269,115.845289 C0.820746269,112.137632 1.43357214,108.566627 2.57253731,105.24304 C2.77661194,104.642856 2.99934826,104.048692 3.24074627,103.460547 L3.24074627,103.454527 C3.47612438,102.866383 3.73558209,102.290881 4.00828358,101.722 C5.09126368,99.4380498 6.44092537,97.302791 8.00670149,95.3469254 C8.40281095,94.8575075 8.80554229,94.3813333 9.21971144,93.9232189 C9.42378607,93.6878408 9.64050249,93.4590846 9.85721891,93.2297264 C10.0739353,93.00699 10.2906517,92.7842537 10.5133881,92.5615174 C11.6210498,91.4532537 12.8093781,90.4382985 14.0597114,89.50401 C14.6851791,89.0398756 15.3160647,88.6064428 15.9722338,88.1856517 C16.2076119,88.0369602 16.4423881,87.8882687 16.6837861,87.746199 C16.7891343,87.6781741 16.8944826,87.6161692 16.9992289,87.5541642 C17.3212935,87.3621294 17.6493781,87.1767164 17.9774627,86.9973234 C17.9901045,86.9846816 18.0021443,86.9913035 18.0021443,86.9913035 L18.0021443,86.9852836 C18.2682239,86.836592 18.5282836,86.6879005 18.7883433,86.533791 C19.5619005,86.0630348 20.3047562,85.5435174 20.9976468,84.9740348 C21.0036667,84.9800547 21.0036667,84.9800547 21.0102886,84.9860746 C24.2164876,88.2663184 28.6790398,90.3028507 33.630408,90.3148905 L33.6797711,90.3148905 C38.6311393,90.3148905 43.1123532,88.2723383 46.330592,84.9920945 L46.3366119,84.9860746 C46.7760647,85.3514826 47.2341791,85.691607 47.7043333,86.0070498 C47.7536965,86.0443731 47.7970398,86.0750746 47.846403,86.1063781 C48.0691393,86.2550697 48.2978955,86.3971393 48.5272537,86.5398109 C48.7812935,86.6945224 49.0413532,86.8432139 49.3008109,86.9852836 C49.3441542,86.9979254 49.3748557,87.022607 49.4121791,87.0406667 C49.7029403,87.202 50.0003234,87.3687512 50.2850647,87.5421244 C50.5324826,87.6908159 50.7805025,87.8395075 51.0219005,87.9942189 C51.2753383,88.1489303 51.5293781,88.3096617 51.776796,88.4890547 C52.2902935,88.8171393 52.7917512,89.1765274 53.2805672,89.5473532 L53.293209,89.5473532 C53.9180746,90.0175075 54.5309005,90.5069254 55.1190448,91.0204229 C55.713209,91.5279005 56.2826915,92.0606617 56.8335124,92.6114826 C62.7998358,98.5525224 66.4828109,106.759453 66.4828109,115.845289 Z", styling)

	return canvas
}