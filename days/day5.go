package days

import (
	"fmt"
	"math"
)

type cord struct {
	x int
	y int
}

type line struct {
	start cord
	end   cord
}

func getInput5() []line {
	ret := []line{}

	ret = append(ret, line{start: cord{x: 578, y: 391}, end: cord{x: 578, y: 322}})
	ret = append(ret, line{start: cord{x: 274, y: 585}, end: cord{x: 651, y: 962}})
	ret = append(ret, line{start: cord{x: 482, y: 348}, end: cord{x: 294, y: 348}})
	ret = append(ret, line{start: cord{x: 682, y: 514}, end: cord{x: 367, y: 829}})
	ret = append(ret, line{start: cord{x: 180, y: 243}, end: cord{x: 800, y: 863}})
	ret = append(ret, line{start: cord{x: 850, y: 828}, end: cord{x: 850, y: 98}})
	ret = append(ret, line{start: cord{x: 698, y: 439}, end: cord{x: 460, y: 677}})
	ret = append(ret, line{start: cord{x: 518, y: 379}, end: cord{x: 518, y: 176}})
	ret = append(ret, line{start: cord{x: 486, y: 437}, end: cord{x: 486, y: 640}})
	ret = append(ret, line{start: cord{x: 730, y: 420}, end: cord{x: 374, y: 420}})
	ret = append(ret, line{start: cord{x: 738, y: 726}, end: cord{x: 632, y: 726}})
	ret = append(ret, line{start: cord{x: 48, y: 959}, end: cord{x: 468, y: 539}})
	ret = append(ret, line{start: cord{x: 246, y: 526}, end: cord{x: 246, y: 174}})
	ret = append(ret, line{start: cord{x: 490, y: 438}, end: cord{x: 291, y: 239}})
	ret = append(ret, line{start: cord{x: 975, y: 116}, end: cord{x: 272, y: 116}})
	ret = append(ret, line{start: cord{x: 695, y: 883}, end: cord{x: 476, y: 883}})
	ret = append(ret, line{start: cord{x: 129, y: 393}, end: cord{x: 129, y: 300}})
	ret = append(ret, line{start: cord{x: 658, y: 556}, end: cord{x: 658, y: 436}})
	ret = append(ret, line{start: cord{x: 860, y: 777}, end: cord{x: 860, y: 365}})
	ret = append(ret, line{start: cord{x: 229, y: 321}, end: cord{x: 422, y: 514}})
	ret = append(ret, line{start: cord{x: 814, y: 312}, end: cord{x: 752, y: 312}})
	ret = append(ret, line{start: cord{x: 886, y: 103}, end: cord{x: 783, y: 206}})
	ret = append(ret, line{start: cord{x: 860, y: 786}, end: cord{x: 701, y: 945}})
	ret = append(ret, line{start: cord{x: 551, y: 789}, end: cord{x: 479, y: 789}})
	ret = append(ret, line{start: cord{x: 103, y: 685}, end: cord{x: 687, y: 685}})
	ret = append(ret, line{start: cord{x: 649, y: 395}, end: cord{x: 758, y: 395}})
	ret = append(ret, line{start: cord{x: 48, y: 233}, end: cord{x: 48, y: 677}})
	ret = append(ret, line{start: cord{x: 385, y: 22}, end: cord{x: 385, y: 120}})
	ret = append(ret, line{start: cord{x: 731, y: 546}, end: cord{x: 731, y: 463}})
	ret = append(ret, line{start: cord{x: 570, y: 507}, end: cord{x: 930, y: 507}})
	ret = append(ret, line{start: cord{x: 92, y: 288}, end: cord{x: 780, y: 976}})
	ret = append(ret, line{start: cord{x: 270, y: 622}, end: cord{x: 270, y: 231}})
	ret = append(ret, line{start: cord{x: 791, y: 76}, end: cord{x: 791, y: 769}})
	ret = append(ret, line{start: cord{x: 926, y: 60}, end: cord{x: 25, y: 961}})
	ret = append(ret, line{start: cord{x: 972, y: 986}, end: cord{x: 47, y: 61}})
	ret = append(ret, line{start: cord{x: 382, y: 601}, end: cord{x: 345, y: 638}})
	ret = append(ret, line{start: cord{x: 536, y: 122}, end: cord{x: 536, y: 822}})
	ret = append(ret, line{start: cord{x: 963, y: 864}, end: cord{x: 532, y: 433}})
	ret = append(ret, line{start: cord{x: 590, y: 550}, end: cord{x: 590, y: 221}})
	ret = append(ret, line{start: cord{x: 768, y: 744}, end: cord{x: 768, y: 981}})
	ret = append(ret, line{start: cord{x: 842, y: 129}, end: cord{x: 842, y: 65}})
	ret = append(ret, line{start: cord{x: 521, y: 548}, end: cord{x: 777, y: 548}})
	ret = append(ret, line{start: cord{x: 897, y: 410}, end: cord{x: 773, y: 410}})
	ret = append(ret, line{start: cord{x: 433, y: 738}, end: cord{x: 802, y: 369}})
	ret = append(ret, line{start: cord{x: 498, y: 815}, end: cord{x: 498, y: 874}})
	ret = append(ret, line{start: cord{x: 93, y: 905}, end: cord{x: 837, y: 161}})
	ret = append(ret, line{start: cord{x: 552, y: 281}, end: cord{x: 552, y: 491}})
	ret = append(ret, line{start: cord{x: 274, y: 82}, end: cord{x: 274, y: 760}})
	ret = append(ret, line{start: cord{x: 566, y: 398}, end: cord{x: 78, y: 886}})
	ret = append(ret, line{start: cord{x: 602, y: 654}, end: cord{x: 256, y: 654}})
	ret = append(ret, line{start: cord{x: 204, y: 816}, end: cord{x: 818, y: 202}})
	ret = append(ret, line{start: cord{x: 488, y: 265}, end: cord{x: 330, y: 107}})
	ret = append(ret, line{start: cord{x: 359, y: 620}, end: cord{x: 71, y: 332}})
	ret = append(ret, line{start: cord{x: 915, y: 133}, end: cord{x: 915, y: 689}})
	ret = append(ret, line{start: cord{x: 698, y: 119}, end: cord{x: 316, y: 501}})
	ret = append(ret, line{start: cord{x: 347, y: 25}, end: cord{x: 415, y: 25}})
	ret = append(ret, line{start: cord{x: 835, y: 902}, end: cord{x: 835, y: 65}})
	ret = append(ret, line{start: cord{x: 900, y: 539}, end: cord{x: 474, y: 113}})
	ret = append(ret, line{start: cord{x: 693, y: 809}, end: cord{x: 245, y: 809}})
	ret = append(ret, line{start: cord{x: 16, y: 32}, end: cord{x: 964, y: 980}})
	ret = append(ret, line{start: cord{x: 177, y: 94}, end: cord{x: 637, y: 554}})
	ret = append(ret, line{start: cord{x: 824, y: 455}, end: cord{x: 346, y: 455}})
	ret = append(ret, line{start: cord{x: 800, y: 893}, end: cord{x: 264, y: 893}})
	ret = append(ret, line{start: cord{x: 109, y: 342}, end: cord{x: 109, y: 676}})
	ret = append(ret, line{start: cord{x: 204, y: 630}, end: cord{x: 281, y: 630}})
	ret = append(ret, line{start: cord{x: 798, y: 930}, end: cord{x: 154, y: 930}})
	ret = append(ret, line{start: cord{x: 287, y: 688}, end: cord{x: 287, y: 106}})
	ret = append(ret, line{start: cord{x: 67, y: 641}, end: cord{x: 970, y: 641}})
	ret = append(ret, line{start: cord{x: 988, y: 908}, end: cord{x: 362, y: 282}})
	ret = append(ret, line{start: cord{x: 411, y: 949}, end: cord{x: 781, y: 949}})
	ret = append(ret, line{start: cord{x: 43, y: 356}, end: cord{x: 187, y: 356}})
	ret = append(ret, line{start: cord{x: 331, y: 848}, end: cord{x: 178, y: 695}})
	ret = append(ret, line{start: cord{x: 513, y: 658}, end: cord{x: 513, y: 763}})
	ret = append(ret, line{start: cord{x: 313, y: 250}, end: cord{x: 605, y: 542}})
	ret = append(ret, line{start: cord{x: 514, y: 552}, end: cord{x: 185, y: 223}})
	ret = append(ret, line{start: cord{x: 652, y: 726}, end: cord{x: 869, y: 726}})
	ret = append(ret, line{start: cord{x: 291, y: 590}, end: cord{x: 291, y: 969}})
	ret = append(ret, line{start: cord{x: 861, y: 808}, end: cord{x: 861, y: 379}})
	ret = append(ret, line{start: cord{x: 842, y: 170}, end: cord{x: 842, y: 928}})
	ret = append(ret, line{start: cord{x: 570, y: 166}, end: cord{x: 570, y: 285}})
	ret = append(ret, line{start: cord{x: 764, y: 439}, end: cord{x: 764, y: 486}})
	ret = append(ret, line{start: cord{x: 200, y: 806}, end: cord{x: 910, y: 806}})
	ret = append(ret, line{start: cord{x: 199, y: 200}, end: cord{x: 876, y: 200}})
	ret = append(ret, line{start: cord{x: 323, y: 474}, end: cord{x: 323, y: 433}})
	ret = append(ret, line{start: cord{x: 258, y: 426}, end: cord{x: 258, y: 808}})
	ret = append(ret, line{start: cord{x: 568, y: 575}, end: cord{x: 568, y: 34}})
	ret = append(ret, line{start: cord{x: 979, y: 982}, end: cord{x: 12, y: 15}})
	ret = append(ret, line{start: cord{x: 424, y: 534}, end: cord{x: 649, y: 759}})
	ret = append(ret, line{start: cord{x: 763, y: 710}, end: cord{x: 147, y: 94}})
	ret = append(ret, line{start: cord{x: 339, y: 232}, end: cord{x: 832, y: 232}})
	ret = append(ret, line{start: cord{x: 10, y: 19}, end: cord{x: 450, y: 19}})
	ret = append(ret, line{start: cord{x: 241, y: 846}, end: cord{x: 45, y: 650}})
	ret = append(ret, line{start: cord{x: 727, y: 990}, end: cord{x: 727, y: 273}})
	ret = append(ret, line{start: cord{x: 596, y: 555}, end: cord{x: 781, y: 370}})
	ret = append(ret, line{start: cord{x: 431, y: 950}, end: cord{x: 431, y: 627}})
	ret = append(ret, line{start: cord{x: 259, y: 415}, end: cord{x: 259, y: 358}})
	ret = append(ret, line{start: cord{x: 803, y: 236}, end: cord{x: 515, y: 236}})
	ret = append(ret, line{start: cord{x: 239, y: 735}, end: cord{x: 603, y: 735}})
	ret = append(ret, line{start: cord{x: 982, y: 377}, end: cord{x: 982, y: 581}})
	ret = append(ret, line{start: cord{x: 779, y: 221}, end: cord{x: 405, y: 595}})
	ret = append(ret, line{start: cord{x: 517, y: 288}, end: cord{x: 414, y: 288}})
	ret = append(ret, line{start: cord{x: 376, y: 688}, end: cord{x: 376, y: 892}})
	ret = append(ret, line{start: cord{x: 450, y: 300}, end: cord{x: 293, y: 143}})
	ret = append(ret, line{start: cord{x: 147, y: 217}, end: cord{x: 871, y: 217}})
	ret = append(ret, line{start: cord{x: 40, y: 144}, end: cord{x: 156, y: 144}})
	ret = append(ret, line{start: cord{x: 913, y: 873}, end: cord{x: 632, y: 592}})
	ret = append(ret, line{start: cord{x: 14, y: 415}, end: cord{x: 274, y: 155}})
	ret = append(ret, line{start: cord{x: 21, y: 987}, end: cord{x: 950, y: 58}})
	ret = append(ret, line{start: cord{x: 979, y: 960}, end: cord{x: 37, y: 18}})
	ret = append(ret, line{start: cord{x: 50, y: 903}, end: cord{x: 890, y: 63}})
	ret = append(ret, line{start: cord{x: 32, y: 523}, end: cord{x: 426, y: 523}})
	ret = append(ret, line{start: cord{x: 625, y: 491}, end: cord{x: 625, y: 692}})
	ret = append(ret, line{start: cord{x: 46, y: 47}, end: cord{x: 899, y: 900}})
	ret = append(ret, line{start: cord{x: 226, y: 633}, end: cord{x: 226, y: 318}})
	ret = append(ret, line{start: cord{x: 24, y: 136}, end: cord{x: 24, y: 693}})
	ret = append(ret, line{start: cord{x: 870, y: 675}, end: cord{x: 850, y: 675}})
	ret = append(ret, line{start: cord{x: 883, y: 862}, end: cord{x: 883, y: 421}})
	ret = append(ret, line{start: cord{x: 581, y: 97}, end: cord{x: 219, y: 97}})
	ret = append(ret, line{start: cord{x: 537, y: 743}, end: cord{x: 434, y: 743}})
	ret = append(ret, line{start: cord{x: 977, y: 77}, end: cord{x: 957, y: 77}})
	ret = append(ret, line{start: cord{x: 139, y: 720}, end: cord{x: 139, y: 403}})
	ret = append(ret, line{start: cord{x: 248, y: 14}, end: cord{x: 394, y: 14}})
	ret = append(ret, line{start: cord{x: 88, y: 55}, end: cord{x: 866, y: 833}})
	ret = append(ret, line{start: cord{x: 562, y: 652}, end: cord{x: 987, y: 227}})
	ret = append(ret, line{start: cord{x: 265, y: 54}, end: cord{x: 958, y: 747}})
	ret = append(ret, line{start: cord{x: 322, y: 161}, end: cord{x: 322, y: 573}})
	ret = append(ret, line{start: cord{x: 574, y: 236}, end: cord{x: 311, y: 236}})
	ret = append(ret, line{start: cord{x: 919, y: 393}, end: cord{x: 919, y: 587}})
	ret = append(ret, line{start: cord{x: 604, y: 906}, end: cord{x: 604, y: 156}})
	ret = append(ret, line{start: cord{x: 691, y: 468}, end: cord{x: 448, y: 225}})
	ret = append(ret, line{start: cord{x: 948, y: 167}, end: cord{x: 948, y: 516}})
	ret = append(ret, line{start: cord{x: 218, y: 238}, end: cord{x: 218, y: 92}})
	ret = append(ret, line{start: cord{x: 989, y: 229}, end: cord{x: 99, y: 229}})
	ret = append(ret, line{start: cord{x: 384, y: 481}, end: cord{x: 384, y: 15}})
	ret = append(ret, line{start: cord{x: 618, y: 681}, end: cord{x: 618, y: 815}})
	ret = append(ret, line{start: cord{x: 292, y: 956}, end: cord{x: 922, y: 326}})
	ret = append(ret, line{start: cord{x: 599, y: 967}, end: cord{x: 599, y: 250}})
	ret = append(ret, line{start: cord{x: 418, y: 648}, end: cord{x: 961, y: 105}})
	ret = append(ret, line{start: cord{x: 120, y: 791}, end: cord{x: 196, y: 791}})
	ret = append(ret, line{start: cord{x: 779, y: 559}, end: cord{x: 582, y: 362}})
	ret = append(ret, line{start: cord{x: 953, y: 941}, end: cord{x: 35, y: 23}})
	ret = append(ret, line{start: cord{x: 508, y: 934}, end: cord{x: 340, y: 934}})
	ret = append(ret, line{start: cord{x: 707, y: 752}, end: cord{x: 915, y: 752}})
	ret = append(ret, line{start: cord{x: 514, y: 958}, end: cord{x: 514, y: 926}})
	ret = append(ret, line{start: cord{x: 15, y: 945}, end: cord{x: 826, y: 134}})
	ret = append(ret, line{start: cord{x: 433, y: 921}, end: cord{x: 821, y: 533}})
	ret = append(ret, line{start: cord{x: 378, y: 80}, end: cord{x: 378, y: 407}})
	ret = append(ret, line{start: cord{x: 76, y: 957}, end: cord{x: 858, y: 175}})
	ret = append(ret, line{start: cord{x: 791, y: 617}, end: cord{x: 662, y: 488}})
	ret = append(ret, line{start: cord{x: 891, y: 897}, end: cord{x: 52, y: 58}})
	ret = append(ret, line{start: cord{x: 786, y: 841}, end: cord{x: 786, y: 973}})
	ret = append(ret, line{start: cord{x: 774, y: 799}, end: cord{x: 348, y: 373}})
	ret = append(ret, line{start: cord{x: 812, y: 48}, end: cord{x: 40, y: 820}})
	ret = append(ret, line{start: cord{x: 57, y: 749}, end: cord{x: 57, y: 767}})
	ret = append(ret, line{start: cord{x: 68, y: 750}, end: cord{x: 68, y: 891}})
	ret = append(ret, line{start: cord{x: 774, y: 920}, end: cord{x: 156, y: 302}})
	ret = append(ret, line{start: cord{x: 598, y: 400}, end: cord{x: 116, y: 882}})
	ret = append(ret, line{start: cord{x: 34, y: 285}, end: cord{x: 856, y: 285}})
	ret = append(ret, line{start: cord{x: 14, y: 473}, end: cord{x: 14, y: 134}})
	ret = append(ret, line{start: cord{x: 594, y: 877}, end: cord{x: 594, y: 333}})
	ret = append(ret, line{start: cord{x: 38, y: 989}, end: cord{x: 964, y: 63}})
	ret = append(ret, line{start: cord{x: 631, y: 209}, end: cord{x: 631, y: 121}})
	ret = append(ret, line{start: cord{x: 45, y: 296}, end: cord{x: 468, y: 296}})
	ret = append(ret, line{start: cord{x: 708, y: 904}, end: cord{x: 11, y: 904}})
	ret = append(ret, line{start: cord{x: 960, y: 20}, end: cord{x: 99, y: 881}})
	ret = append(ret, line{start: cord{x: 412, y: 557}, end: cord{x: 345, y: 557}})
	ret = append(ret, line{start: cord{x: 29, y: 389}, end: cord{x: 504, y: 864}})
	ret = append(ret, line{start: cord{x: 397, y: 713}, end: cord{x: 251, y: 713}})
	ret = append(ret, line{start: cord{x: 350, y: 548}, end: cord{x: 350, y: 61}})
	ret = append(ret, line{start: cord{x: 134, y: 610}, end: cord{x: 579, y: 165}})
	ret = append(ret, line{start: cord{x: 675, y: 947}, end: cord{x: 789, y: 947}})
	ret = append(ret, line{start: cord{x: 12, y: 986}, end: cord{x: 949, y: 49}})
	ret = append(ret, line{start: cord{x: 765, y: 601}, end: cord{x: 765, y: 627}})
	ret = append(ret, line{start: cord{x: 817, y: 701}, end: cord{x: 817, y: 305}})
	ret = append(ret, line{start: cord{x: 508, y: 532}, end: cord{x: 538, y: 502}})
	ret = append(ret, line{start: cord{x: 383, y: 136}, end: cord{x: 383, y: 700}})
	ret = append(ret, line{start: cord{x: 771, y: 549}, end: cord{x: 443, y: 549}})
	ret = append(ret, line{start: cord{x: 283, y: 134}, end: cord{x: 987, y: 838}})
	ret = append(ret, line{start: cord{x: 171, y: 855}, end: cord{x: 171, y: 248}})
	ret = append(ret, line{start: cord{x: 841, y: 858}, end: cord{x: 620, y: 858}})
	ret = append(ret, line{start: cord{x: 512, y: 26}, end: cord{x: 912, y: 26}})
	ret = append(ret, line{start: cord{x: 425, y: 39}, end: cord{x: 180, y: 39}})
	ret = append(ret, line{start: cord{x: 116, y: 279}, end: cord{x: 121, y: 279}})
	ret = append(ret, line{start: cord{x: 282, y: 482}, end: cord{x: 282, y: 939}})
	ret = append(ret, line{start: cord{x: 58, y: 937}, end: cord{x: 980, y: 15}})
	ret = append(ret, line{start: cord{x: 376, y: 641}, end: cord{x: 376, y: 503}})
	ret = append(ret, line{start: cord{x: 548, y: 17}, end: cord{x: 249, y: 17}})
	ret = append(ret, line{start: cord{x: 730, y: 411}, end: cord{x: 427, y: 714}})
	ret = append(ret, line{start: cord{x: 600, y: 73}, end: cord{x: 541, y: 73}})
	ret = append(ret, line{start: cord{x: 656, y: 619}, end: cord{x: 656, y: 810}})
	ret = append(ret, line{start: cord{x: 467, y: 237}, end: cord{x: 467, y: 255}})
	ret = append(ret, line{start: cord{x: 694, y: 946}, end: cord{x: 446, y: 946}})
	ret = append(ret, line{start: cord{x: 168, y: 646}, end: cord{x: 395, y: 646}})
	ret = append(ret, line{start: cord{x: 731, y: 265}, end: cord{x: 731, y: 20}})
	ret = append(ret, line{start: cord{x: 12, y: 172}, end: cord{x: 286, y: 446}})
	ret = append(ret, line{start: cord{x: 385, y: 762}, end: cord{x: 244, y: 903}})
	ret = append(ret, line{start: cord{x: 941, y: 366}, end: cord{x: 941, y: 807}})
	ret = append(ret, line{start: cord{x: 125, y: 383}, end: cord{x: 367, y: 383}})
	ret = append(ret, line{start: cord{x: 341, y: 177}, end: cord{x: 341, y: 809}})
	ret = append(ret, line{start: cord{x: 544, y: 830}, end: cord{x: 544, y: 192}})
	ret = append(ret, line{start: cord{x: 801, y: 943}, end: cord{x: 731, y: 873}})
	ret = append(ret, line{start: cord{x: 862, y: 436}, end: cord{x: 950, y: 436}})
	ret = append(ret, line{start: cord{x: 484, y: 422}, end: cord{x: 484, y: 267}})
	ret = append(ret, line{start: cord{x: 883, y: 155}, end: cord{x: 328, y: 155}})
	ret = append(ret, line{start: cord{x: 499, y: 321}, end: cord{x: 499, y: 449}})
	ret = append(ret, line{start: cord{x: 128, y: 310}, end: cord{x: 778, y: 960}})
	ret = append(ret, line{start: cord{x: 788, y: 571}, end: cord{x: 788, y: 795}})
	ret = append(ret, line{start: cord{x: 523, y: 765}, end: cord{x: 319, y: 765}})
	ret = append(ret, line{start: cord{x: 267, y: 928}, end: cord{x: 267, y: 665}})
	ret = append(ret, line{start: cord{x: 227, y: 829}, end: cord{x: 797, y: 829}})
	ret = append(ret, line{start: cord{x: 96, y: 972}, end: cord{x: 733, y: 335}})
	ret = append(ret, line{start: cord{x: 178, y: 364}, end: cord{x: 178, y: 425}})
	ret = append(ret, line{start: cord{x: 793, y: 201}, end: cord{x: 848, y: 201}})
	ret = append(ret, line{start: cord{x: 975, y: 242}, end: cord{x: 497, y: 720}})
	ret = append(ret, line{start: cord{x: 673, y: 242}, end: cord{x: 513, y: 242}})
	ret = append(ret, line{start: cord{x: 199, y: 163}, end: cord{x: 862, y: 826}})
	ret = append(ret, line{start: cord{x: 988, y: 51}, end: cord{x: 225, y: 814}})
	ret = append(ret, line{start: cord{x: 631, y: 928}, end: cord{x: 631, y: 567}})
	ret = append(ret, line{start: cord{x: 22, y: 474}, end: cord{x: 854, y: 474}})
	ret = append(ret, line{start: cord{x: 717, y: 607}, end: cord{x: 717, y: 514}})
	ret = append(ret, line{start: cord{x: 436, y: 753}, end: cord{x: 905, y: 753}})
	ret = append(ret, line{start: cord{x: 581, y: 343}, end: cord{x: 581, y: 641}})
	ret = append(ret, line{start: cord{x: 128, y: 912}, end: cord{x: 964, y: 76}})
	ret = append(ret, line{start: cord{x: 706, y: 634}, end: cord{x: 843, y: 634}})
	ret = append(ret, line{start: cord{x: 89, y: 826}, end: cord{x: 89, y: 667}})
	ret = append(ret, line{start: cord{x: 766, y: 268}, end: cord{x: 103, y: 268}})
	ret = append(ret, line{start: cord{x: 229, y: 131}, end: cord{x: 229, y: 138}})
	ret = append(ret, line{start: cord{x: 138, y: 112}, end: cord{x: 388, y: 362}})
	ret = append(ret, line{start: cord{x: 434, y: 117}, end: cord{x: 434, y: 387}})
	ret = append(ret, line{start: cord{x: 313, y: 746}, end: cord{x: 313, y: 941}})
	ret = append(ret, line{start: cord{x: 517, y: 944}, end: cord{x: 145, y: 944}})
	ret = append(ret, line{start: cord{x: 611, y: 945}, end: cord{x: 611, y: 872}})
	ret = append(ret, line{start: cord{x: 400, y: 869}, end: cord{x: 329, y: 869}})
	ret = append(ret, line{start: cord{x: 444, y: 701}, end: cord{x: 700, y: 957}})
	ret = append(ret, line{start: cord{x: 894, y: 975}, end: cord{x: 426, y: 975}})
	ret = append(ret, line{start: cord{x: 722, y: 544}, end: cord{x: 722, y: 55}})
	ret = append(ret, line{start: cord{x: 692, y: 927}, end: cord{x: 692, y: 874}})
	ret = append(ret, line{start: cord{x: 451, y: 211}, end: cord{x: 145, y: 211}})
	ret = append(ret, line{start: cord{x: 562, y: 850}, end: cord{x: 562, y: 252}})
	ret = append(ret, line{start: cord{x: 833, y: 154}, end: cord{x: 703, y: 284}})
	ret = append(ret, line{start: cord{x: 700, y: 911}, end: cord{x: 700, y: 738}})
	ret = append(ret, line{start: cord{x: 32, y: 982}, end: cord{x: 891, y: 123}})
	ret = append(ret, line{start: cord{x: 512, y: 512}, end: cord{x: 403, y: 512}})
	ret = append(ret, line{start: cord{x: 444, y: 963}, end: cord{x: 40, y: 559}})
	ret = append(ret, line{start: cord{x: 866, y: 53}, end: cord{x: 866, y: 733}})
	ret = append(ret, line{start: cord{x: 395, y: 90}, end: cord{x: 603, y: 90}})
	ret = append(ret, line{start: cord{x: 781, y: 175}, end: cord{x: 506, y: 175}})
	ret = append(ret, line{start: cord{x: 649, y: 569}, end: cord{x: 210, y: 130}})
	ret = append(ret, line{start: cord{x: 861, y: 926}, end: cord{x: 79, y: 144}})
	ret = append(ret, line{start: cord{x: 160, y: 953}, end: cord{x: 735, y: 953}})
	ret = append(ret, line{start: cord{x: 138, y: 837}, end: cord{x: 138, y: 166}})
	ret = append(ret, line{start: cord{x: 659, y: 683}, end: cord{x: 659, y: 656}})
	ret = append(ret, line{start: cord{x: 198, y: 587}, end: cord{x: 725, y: 60}})
	ret = append(ret, line{start: cord{x: 290, y: 36}, end: cord{x: 785, y: 36}})
	ret = append(ret, line{start: cord{x: 481, y: 228}, end: cord{x: 785, y: 532}})
	ret = append(ret, line{start: cord{x: 721, y: 152}, end: cord{x: 192, y: 681}})
	ret = append(ret, line{start: cord{x: 162, y: 445}, end: cord{x: 162, y: 476}})
	ret = append(ret, line{start: cord{x: 286, y: 93}, end: cord{x: 286, y: 611}})
	ret = append(ret, line{start: cord{x: 882, y: 393}, end: cord{x: 770, y: 393}})
	ret = append(ret, line{start: cord{x: 194, y: 703}, end: cord{x: 194, y: 714}})
	ret = append(ret, line{start: cord{x: 172, y: 505}, end: cord{x: 153, y: 524}})
	ret = append(ret, line{start: cord{x: 989, y: 986}, end: cord{x: 48, y: 45}})
	ret = append(ret, line{start: cord{x: 946, y: 334}, end: cord{x: 946, y: 864}})
	ret = append(ret, line{start: cord{x: 543, y: 48}, end: cord{x: 485, y: 48}})
	ret = append(ret, line{start: cord{x: 276, y: 520}, end: cord{x: 184, y: 612}})
	ret = append(ret, line{start: cord{x: 879, y: 488}, end: cord{x: 665, y: 488}})
	ret = append(ret, line{start: cord{x: 706, y: 312}, end: cord{x: 706, y: 300}})
	ret = append(ret, line{start: cord{x: 859, y: 958}, end: cord{x: 533, y: 958}})
	ret = append(ret, line{start: cord{x: 345, y: 591}, end: cord{x: 345, y: 685}})
	ret = append(ret, line{start: cord{x: 201, y: 734}, end: cord{x: 310, y: 734}})
	ret = append(ret, line{start: cord{x: 610, y: 781}, end: cord{x: 610, y: 250}})
	ret = append(ret, line{start: cord{x: 25, y: 702}, end: cord{x: 25, y: 470}})
	ret = append(ret, line{start: cord{x: 127, y: 802}, end: cord{x: 46, y: 802}})
	ret = append(ret, line{start: cord{x: 899, y: 330}, end: cord{x: 899, y: 942}})
	ret = append(ret, line{start: cord{x: 266, y: 118}, end: cord{x: 266, y: 978}})
	ret = append(ret, line{start: cord{x: 871, y: 535}, end: cord{x: 871, y: 230}})
	ret = append(ret, line{start: cord{x: 346, y: 290}, end: cord{x: 346, y: 138}})
	ret = append(ret, line{start: cord{x: 411, y: 171}, end: cord{x: 911, y: 671}})
	ret = append(ret, line{start: cord{x: 104, y: 427}, end: cord{x: 500, y: 31}})
	ret = append(ret, line{start: cord{x: 531, y: 115}, end: cord{x: 531, y: 861}})
	ret = append(ret, line{start: cord{x: 164, y: 699}, end: cord{x: 529, y: 699}})
	ret = append(ret, line{start: cord{x: 215, y: 560}, end: cord{x: 97, y: 442}})
	ret = append(ret, line{start: cord{x: 331, y: 323}, end: cord{x: 331, y: 321}})
	ret = append(ret, line{start: cord{x: 74, y: 969}, end: cord{x: 74, y: 57}})
	ret = append(ret, line{start: cord{x: 894, y: 743}, end: cord{x: 739, y: 588}})
	ret = append(ret, line{start: cord{x: 913, y: 895}, end: cord{x: 160, y: 895}})
	ret = append(ret, line{start: cord{x: 868, y: 291}, end: cord{x: 868, y: 987}})
	ret = append(ret, line{start: cord{x: 913, y: 390}, end: cord{x: 913, y: 144}})
	ret = append(ret, line{start: cord{x: 548, y: 812}, end: cord{x: 889, y: 812}})
	ret = append(ret, line{start: cord{x: 978, y: 819}, end: cord{x: 673, y: 514}})
	ret = append(ret, line{start: cord{x: 989, y: 130}, end: cord{x: 989, y: 589}})
	ret = append(ret, line{start: cord{x: 986, y: 12}, end: cord{x: 10, y: 988}})
	ret = append(ret, line{start: cord{x: 48, y: 18}, end: cord{x: 974, y: 944}})
	ret = append(ret, line{start: cord{x: 511, y: 336}, end: cord{x: 736, y: 111}})
	ret = append(ret, line{start: cord{x: 61, y: 609}, end: cord{x: 61, y: 742}})
	ret = append(ret, line{start: cord{x: 536, y: 650}, end: cord{x: 773, y: 650}})
	ret = append(ret, line{start: cord{x: 924, y: 691}, end: cord{x: 307, y: 74}})
	ret = append(ret, line{start: cord{x: 49, y: 988}, end: cord{x: 912, y: 125}})
	ret = append(ret, line{start: cord{x: 128, y: 692}, end: cord{x: 128, y: 969}})
	ret = append(ret, line{start: cord{x: 569, y: 837}, end: cord{x: 916, y: 837}})
	ret = append(ret, line{start: cord{x: 849, y: 745}, end: cord{x: 849, y: 105}})
	ret = append(ret, line{start: cord{x: 524, y: 926}, end: cord{x: 357, y: 926}})
	ret = append(ret, line{start: cord{x: 110, y: 827}, end: cord{x: 661, y: 827}})
	ret = append(ret, line{start: cord{x: 911, y: 36}, end: cord{x: 49, y: 898}})
	ret = append(ret, line{start: cord{x: 967, y: 15}, end: cord{x: 23, y: 959}})
	ret = append(ret, line{start: cord{x: 969, y: 166}, end: cord{x: 155, y: 980}})
	ret = append(ret, line{start: cord{x: 204, y: 684}, end: cord{x: 805, y: 83}})
	ret = append(ret, line{start: cord{x: 230, y: 960}, end: cord{x: 230, y: 556}})
	ret = append(ret, line{start: cord{x: 309, y: 718}, end: cord{x: 522, y: 931}})
	ret = append(ret, line{start: cord{x: 121, y: 208}, end: cord{x: 121, y: 443}})
	ret = append(ret, line{start: cord{x: 733, y: 797}, end: cord{x: 710, y: 820}})
	ret = append(ret, line{start: cord{x: 813, y: 780}, end: cord{x: 813, y: 909}})
	ret = append(ret, line{start: cord{x: 154, y: 97}, end: cord{x: 375, y: 318}})
	ret = append(ret, line{start: cord{x: 117, y: 916}, end: cord{x: 984, y: 49}})
	ret = append(ret, line{start: cord{x: 573, y: 525}, end: cord{x: 573, y: 980}})
	ret = append(ret, line{start: cord{x: 442, y: 636}, end: cord{x: 383, y: 695}})
	ret = append(ret, line{start: cord{x: 938, y: 21}, end: cord{x: 938, y: 50}})
	ret = append(ret, line{start: cord{x: 38, y: 672}, end: cord{x: 196, y: 672}})
	ret = append(ret, line{start: cord{x: 52, y: 829}, end: cord{x: 52, y: 835}})
	ret = append(ret, line{start: cord{x: 661, y: 278}, end: cord{x: 157, y: 782}})
	ret = append(ret, line{start: cord{x: 525, y: 347}, end: cord{x: 285, y: 347}})
	ret = append(ret, line{start: cord{x: 339, y: 468}, end: cord{x: 339, y: 42}})
	ret = append(ret, line{start: cord{x: 10, y: 20}, end: cord{x: 976, y: 986}})
	ret = append(ret, line{start: cord{x: 953, y: 812}, end: cord{x: 445, y: 304}})
	ret = append(ret, line{start: cord{x: 328, y: 327}, end: cord{x: 711, y: 327}})
	ret = append(ret, line{start: cord{x: 750, y: 820}, end: cord{x: 750, y: 172}})
	ret = append(ret, line{start: cord{x: 244, y: 935}, end: cord{x: 244, y: 360}})
	ret = append(ret, line{start: cord{x: 842, y: 36}, end: cord{x: 181, y: 697}})
	ret = append(ret, line{start: cord{x: 559, y: 730}, end: cord{x: 320, y: 730}})
	ret = append(ret, line{start: cord{x: 149, y: 510}, end: cord{x: 524, y: 510}})
	ret = append(ret, line{start: cord{x: 713, y: 913}, end: cord{x: 262, y: 462}})
	ret = append(ret, line{start: cord{x: 703, y: 957}, end: cord{x: 643, y: 957}})
	ret = append(ret, line{start: cord{x: 170, y: 930}, end: cord{x: 767, y: 930}})
	ret = append(ret, line{start: cord{x: 804, y: 259}, end: cord{x: 635, y: 90}})
	ret = append(ret, line{start: cord{x: 117, y: 948}, end: cord{x: 932, y: 133}})
	ret = append(ret, line{start: cord{x: 263, y: 806}, end: cord{x: 981, y: 806}})
	ret = append(ret, line{start: cord{x: 307, y: 665}, end: cord{x: 307, y: 743}})
	ret = append(ret, line{start: cord{x: 697, y: 164}, end: cord{x: 665, y: 132}})
	ret = append(ret, line{start: cord{x: 589, y: 568}, end: cord{x: 872, y: 285}})
	ret = append(ret, line{start: cord{x: 865, y: 189}, end: cord{x: 417, y: 637}})
	ret = append(ret, line{start: cord{x: 77, y: 76}, end: cord{x: 951, y: 950}})
	ret = append(ret, line{start: cord{x: 546, y: 350}, end: cord{x: 769, y: 350}})
	ret = append(ret, line{start: cord{x: 533, y: 479}, end: cord{x: 566, y: 446}})
	ret = append(ret, line{start: cord{x: 689, y: 79}, end: cord{x: 689, y: 417}})
	ret = append(ret, line{start: cord{x: 132, y: 666}, end: cord{x: 888, y: 666}})
	ret = append(ret, line{start: cord{x: 661, y: 88}, end: cord{x: 155, y: 88}})
	ret = append(ret, line{start: cord{x: 93, y: 27}, end: cord{x: 852, y: 786}})
	ret = append(ret, line{start: cord{x: 536, y: 366}, end: cord{x: 815, y: 366}})
	ret = append(ret, line{start: cord{x: 97, y: 649}, end: cord{x: 97, y: 214}})
	ret = append(ret, line{start: cord{x: 50, y: 784}, end: cord{x: 691, y: 143}})
	ret = append(ret, line{start: cord{x: 523, y: 687}, end: cord{x: 523, y: 881}})
	ret = append(ret, line{start: cord{x: 720, y: 825}, end: cord{x: 865, y: 825}})
	ret = append(ret, line{start: cord{x: 103, y: 985}, end: cord{x: 939, y: 149}})
	ret = append(ret, line{start: cord{x: 135, y: 94}, end: cord{x: 91, y: 50}})
	ret = append(ret, line{start: cord{x: 959, y: 26}, end: cord{x: 18, y: 967}})
	ret = append(ret, line{start: cord{x: 391, y: 617}, end: cord{x: 391, y: 147}})
	ret = append(ret, line{start: cord{x: 522, y: 103}, end: cord{x: 522, y: 202}})
	ret = append(ret, line{start: cord{x: 161, y: 774}, end: cord{x: 742, y: 193}})
	ret = append(ret, line{start: cord{x: 125, y: 291}, end: cord{x: 125, y: 513}})
	ret = append(ret, line{start: cord{x: 449, y: 436}, end: cord{x: 726, y: 436}})
	ret = append(ret, line{start: cord{x: 438, y: 127}, end: cord{x: 499, y: 66}})
	ret = append(ret, line{start: cord{x: 804, y: 577}, end: cord{x: 804, y: 385}})
	ret = append(ret, line{start: cord{x: 714, y: 112}, end: cord{x: 714, y: 90}})
	ret = append(ret, line{start: cord{x: 111, y: 184}, end: cord{x: 907, y: 980}})
	ret = append(ret, line{start: cord{x: 218, y: 209}, end: cord{x: 53, y: 209}})
	ret = append(ret, line{start: cord{x: 343, y: 949}, end: cord{x: 73, y: 679}})
	ret = append(ret, line{start: cord{x: 50, y: 205}, end: cord{x: 828, y: 983}})
	ret = append(ret, line{start: cord{x: 416, y: 664}, end: cord{x: 416, y: 213}})
	ret = append(ret, line{start: cord{x: 300, y: 902}, end: cord{x: 300, y: 137}})
	ret = append(ret, line{start: cord{x: 563, y: 366}, end: cord{x: 307, y: 366}})
	ret = append(ret, line{start: cord{x: 302, y: 750}, end: cord{x: 572, y: 750}})
	ret = append(ret, line{start: cord{x: 436, y: 59}, end: cord{x: 512, y: 59}})
	ret = append(ret, line{start: cord{x: 363, y: 299}, end: cord{x: 363, y: 471}})
	ret = append(ret, line{start: cord{x: 969, y: 988}, end: cord{x: 10, y: 29}})
	ret = append(ret, line{start: cord{x: 15, y: 349}, end: cord{x: 15, y: 424}})
	ret = append(ret, line{start: cord{x: 855, y: 231}, end: cord{x: 855, y: 241}})
	ret = append(ret, line{start: cord{x: 93, y: 771}, end: cord{x: 540, y: 324}})
	ret = append(ret, line{start: cord{x: 360, y: 363}, end: cord{x: 360, y: 481}})
	ret = append(ret, line{start: cord{x: 890, y: 391}, end: cord{x: 890, y: 824}})
	ret = append(ret, line{start: cord{x: 603, y: 916}, end: cord{x: 780, y: 916}})
	ret = append(ret, line{start: cord{x: 686, y: 776}, end: cord{x: 165, y: 255}})
	ret = append(ret, line{start: cord{x: 905, y: 64}, end: cord{x: 37, y: 932}})
	ret = append(ret, line{start: cord{x: 937, y: 607}, end: cord{x: 937, y: 846}})
	ret = append(ret, line{start: cord{x: 634, y: 108}, end: cord{x: 971, y: 108}})
	ret = append(ret, line{start: cord{x: 118, y: 419}, end: cord{x: 292, y: 419}})
	ret = append(ret, line{start: cord{x: 724, y: 241}, end: cord{x: 724, y: 663}})
	ret = append(ret, line{start: cord{x: 118, y: 327}, end: cord{x: 688, y: 327}})
	ret = append(ret, line{start: cord{x: 728, y: 316}, end: cord{x: 507, y: 316}})
	ret = append(ret, line{start: cord{x: 824, y: 652}, end: cord{x: 744, y: 652}})
	ret = append(ret, line{start: cord{x: 985, y: 72}, end: cord{x: 93, y: 964}})
	ret = append(ret, line{start: cord{x: 791, y: 652}, end: cord{x: 791, y: 621}})
	ret = append(ret, line{start: cord{x: 475, y: 488}, end: cord{x: 475, y: 448}})
	ret = append(ret, line{start: cord{x: 289, y: 386}, end: cord{x: 648, y: 386}})
	ret = append(ret, line{start: cord{x: 833, y: 925}, end: cord{x: 120, y: 925}})
	ret = append(ret, line{start: cord{x: 323, y: 813}, end: cord{x: 652, y: 813}})
	ret = append(ret, line{start: cord{x: 631, y: 615}, end: cord{x: 248, y: 615}})
	ret = append(ret, line{start: cord{x: 191, y: 222}, end: cord{x: 603, y: 634}})
	ret = append(ret, line{start: cord{x: 445, y: 322}, end: cord{x: 964, y: 322}})
	ret = append(ret, line{start: cord{x: 238, y: 672}, end: cord{x: 142, y: 672}})
	ret = append(ret, line{start: cord{x: 170, y: 370}, end: cord{x: 439, y: 370}})
	ret = append(ret, line{start: cord{x: 158, y: 77}, end: cord{x: 491, y: 410}})
	ret = append(ret, line{start: cord{x: 165, y: 737}, end: cord{x: 816, y: 737}})
	ret = append(ret, line{start: cord{x: 420, y: 957}, end: cord{x: 709, y: 668}})
	ret = append(ret, line{start: cord{x: 936, y: 283}, end: cord{x: 681, y: 283}})
	ret = append(ret, line{start: cord{x: 76, y: 781}, end: cord{x: 291, y: 781}})
	ret = append(ret, line{start: cord{x: 197, y: 575}, end: cord{x: 656, y: 116}})
	ret = append(ret, line{start: cord{x: 577, y: 746}, end: cord{x: 577, y: 748}})
	ret = append(ret, line{start: cord{x: 435, y: 198}, end: cord{x: 435, y: 803}})
	ret = append(ret, line{start: cord{x: 787, y: 623}, end: cord{x: 787, y: 153}})
	ret = append(ret, line{start: cord{x: 476, y: 176}, end: cord{x: 670, y: 176}})
	ret = append(ret, line{start: cord{x: 107, y: 581}, end: cord{x: 107, y: 167}})
	ret = append(ret, line{start: cord{x: 575, y: 495}, end: cord{x: 186, y: 106}})
	ret = append(ret, line{start: cord{x: 283, y: 760}, end: cord{x: 19, y: 760}})
	ret = append(ret, line{start: cord{x: 910, y: 483}, end: cord{x: 871, y: 483}})
	ret = append(ret, line{start: cord{x: 550, y: 99}, end: cord{x: 550, y: 94}})
	ret = append(ret, line{start: cord{x: 338, y: 522}, end: cord{x: 589, y: 522}})
	ret = append(ret, line{start: cord{x: 856, y: 435}, end: cord{x: 856, y: 388}})
	ret = append(ret, line{start: cord{x: 890, y: 380}, end: cord{x: 392, y: 878}})
	ret = append(ret, line{start: cord{x: 524, y: 885}, end: cord{x: 315, y: 676}})
	ret = append(ret, line{start: cord{x: 23, y: 34}, end: cord{x: 769, y: 780}})
	ret = append(ret, line{start: cord{x: 686, y: 647}, end: cord{x: 545, y: 647}})
	ret = append(ret, line{start: cord{x: 760, y: 442}, end: cord{x: 564, y: 246}})
	ret = append(ret, line{start: cord{x: 535, y: 264}, end: cord{x: 61, y: 264}})
	ret = append(ret, line{start: cord{x: 709, y: 168}, end: cord{x: 709, y: 33}})
	ret = append(ret, line{start: cord{x: 89, y: 230}, end: cord{x: 604, y: 230}})
	ret = append(ret, line{start: cord{x: 476, y: 558}, end: cord{x: 82, y: 558}})
	ret = append(ret, line{start: cord{x: 905, y: 48}, end: cord{x: 294, y: 48}})
	ret = append(ret, line{start: cord{x: 695, y: 882}, end: cord{x: 695, y: 153}})
	ret = append(ret, line{start: cord{x: 785, y: 716}, end: cord{x: 94, y: 716}})
	ret = append(ret, line{start: cord{x: 390, y: 990}, end: cord{x: 390, y: 757}})
	ret = append(ret, line{start: cord{x: 775, y: 699}, end: cord{x: 783, y: 699}})
	ret = append(ret, line{start: cord{x: 965, y: 126}, end: cord{x: 425, y: 126}})
	ret = append(ret, line{start: cord{x: 572, y: 45}, end: cord{x: 482, y: 45}})
	ret = append(ret, line{start: cord{x: 399, y: 391}, end: cord{x: 399, y: 827}})
	ret = append(ret, line{start: cord{x: 310, y: 660}, end: cord{x: 947, y: 23}})
	ret = append(ret, line{start: cord{x: 418, y: 813}, end: cord{x: 72, y: 467}})
	ret = append(ret, line{start: cord{x: 292, y: 911}, end: cord{x: 506, y: 697}})
	ret = append(ret, line{start: cord{x: 177, y: 685}, end: cord{x: 177, y: 100}})
	ret = append(ret, line{start: cord{x: 749, y: 294}, end: cord{x: 749, y: 927}})
	ret = append(ret, line{start: cord{x: 304, y: 832}, end: cord{x: 833, y: 303}})
	ret = append(ret, line{start: cord{x: 237, y: 759}, end: cord{x: 923, y: 73}})
	ret = append(ret, line{start: cord{x: 834, y: 95}, end: cord{x: 15, y: 914}})
	ret = append(ret, line{start: cord{x: 233, y: 99}, end: cord{x: 822, y: 99}})
	ret = append(ret, line{start: cord{x: 462, y: 841}, end: cord{x: 462, y: 845}})
	ret = append(ret, line{start: cord{x: 968, y: 70}, end: cord{x: 815, y: 70}})
	ret = append(ret, line{start: cord{x: 820, y: 565}, end: cord{x: 241, y: 565}})
	ret = append(ret, line{start: cord{x: 849, y: 469}, end: cord{x: 648, y: 670}})
	ret = append(ret, line{start: cord{x: 10, y: 825}, end: cord{x: 906, y: 825}})
	ret = append(ret, line{start: cord{x: 105, y: 105}, end: cord{x: 526, y: 526}})
	ret = append(ret, line{start: cord{x: 977, y: 173}, end: cord{x: 711, y: 173}})
	ret = append(ret, line{start: cord{x: 347, y: 66}, end: cord{x: 347, y: 959}})
	ret = append(ret, line{start: cord{x: 921, y: 42}, end: cord{x: 41, y: 42}})
	ret = append(ret, line{start: cord{x: 100, y: 264}, end: cord{x: 100, y: 580}})
	ret = append(ret, line{start: cord{x: 608, y: 211}, end: cord{x: 166, y: 653}})
	ret = append(ret, line{start: cord{x: 826, y: 171}, end: cord{x: 509, y: 171}})
	ret = append(ret, line{start: cord{x: 346, y: 541}, end: cord{x: 802, y: 85}})
	ret = append(ret, line{start: cord{x: 351, y: 70}, end: cord{x: 872, y: 70}})
	ret = append(ret, line{start: cord{x: 649, y: 79}, end: cord{x: 590, y: 79}})
	ret = append(ret, line{start: cord{x: 974, y: 31}, end: cord{x: 24, y: 981}})
	ret = append(ret, line{start: cord{x: 876, y: 145}, end: cord{x: 227, y: 794}})
	ret = append(ret, line{start: cord{x: 855, y: 903}, end: cord{x: 855, y: 891}})
	ret = append(ret, line{start: cord{x: 621, y: 734}, end: cord{x: 621, y: 930}})
	ret = append(ret, line{start: cord{x: 190, y: 184}, end: cord{x: 727, y: 721}})
	ret = append(ret, line{start: cord{x: 210, y: 855}, end: cord{x: 564, y: 855}})
	ret = append(ret, line{start: cord{x: 612, y: 919}, end: cord{x: 612, y: 628}})
	ret = append(ret, line{start: cord{x: 258, y: 851}, end: cord{x: 573, y: 851}})
	ret = append(ret, line{start: cord{x: 842, y: 85}, end: cord{x: 140, y: 787}})
	ret = append(ret, line{start: cord{x: 252, y: 312}, end: cord{x: 252, y: 17}})
	ret = append(ret, line{start: cord{x: 82, y: 352}, end: cord{x: 135, y: 352}})
	ret = append(ret, line{start: cord{x: 365, y: 583}, end: cord{x: 854, y: 583}})
	ret = append(ret, line{start: cord{x: 939, y: 666}, end: cord{x: 525, y: 252}})
	ret = append(ret, line{start: cord{x: 257, y: 481}, end: cord{x: 257, y: 591}})
	ret = append(ret, line{start: cord{x: 382, y: 725}, end: cord{x: 382, y: 786}})
	ret = append(ret, line{start: cord{x: 326, y: 111}, end: cord{x: 38, y: 399}})
	ret = append(ret, line{start: cord{x: 476, y: 480}, end: cord{x: 476, y: 544}})
	ret = append(ret, line{start: cord{x: 592, y: 49}, end: cord{x: 592, y: 473}})
	ret = append(ret, line{start: cord{x: 626, y: 748}, end: cord{x: 626, y: 477}})
	ret = append(ret, line{start: cord{x: 612, y: 574}, end: cord{x: 19, y: 574}})
	ret = append(ret, line{start: cord{x: 638, y: 734}, end: cord{x: 604, y: 734}})
	ret = append(ret, line{start: cord{x: 240, y: 794}, end: cord{x: 770, y: 794}})
	ret = append(ret, line{start: cord{x: 598, y: 931}, end: cord{x: 37, y: 370}})
	ret = append(ret, line{start: cord{x: 666, y: 559}, end: cord{x: 573, y: 559}})
	ret = append(ret, line{start: cord{x: 208, y: 337}, end: cord{x: 784, y: 913}})
	ret = append(ret, line{start: cord{x: 24, y: 17}, end: cord{x: 988, y: 981}})
	ret = append(ret, line{start: cord{x: 324, y: 267}, end: cord{x: 332, y: 267}})
	ret = append(ret, line{start: cord{x: 233, y: 589}, end: cord{x: 300, y: 589}})
	ret = append(ret, line{start: cord{x: 53, y: 46}, end: cord{x: 986, y: 979}})
	ret = append(ret, line{start: cord{x: 193, y: 649}, end: cord{x: 243, y: 649}})
	ret = append(ret, line{start: cord{x: 873, y: 600}, end: cord{x: 873, y: 618}})
	ret = append(ret, line{start: cord{x: 461, y: 102}, end: cord{x: 638, y: 102}})
	ret = append(ret, line{start: cord{x: 468, y: 574}, end: cord{x: 507, y: 535}})
	ret = append(ret, line{start: cord{x: 261, y: 521}, end: cord{x: 658, y: 521}})
	ret = append(ret, line{start: cord{x: 540, y: 234}, end: cord{x: 769, y: 234}})
	ret = append(ret, line{start: cord{x: 975, y: 337}, end: cord{x: 975, y: 478}})
	ret = append(ret, line{start: cord{x: 724, y: 982}, end: cord{x: 585, y: 982}})
	ret = append(ret, line{start: cord{x: 449, y: 639}, end: cord{x: 449, y: 255}})
	ret = append(ret, line{start: cord{x: 47, y: 296}, end: cord{x: 751, y: 296}})
	ret = append(ret, line{start: cord{x: 700, y: 262}, end: cord{x: 903, y: 262}})
	ret = append(ret, line{start: cord{x: 838, y: 833}, end: cord{x: 838, y: 626}})
	ret = append(ret, line{start: cord{x: 956, y: 17}, end: cord{x: 24, y: 949}})
	return ret
}

func DayFiveP1() {
	input := getInput5()
	board := [][]int{}

	for i := 0; i < 1000; i++ {
		row := []int{}
		for j := 0; j < 1000; j++ {
			row = append(row, 0)
		}
		board = append(board, row)
	}

	for _, smoke := range input {
		if smoke.start.y == smoke.end.y {
			smaller, larger := 0, 0

			if smoke.start.x > smoke.end.x {
				smaller = smoke.end.x
				larger = smoke.start.x
			} else {
				larger = smoke.end.x
				smaller = smoke.start.x
			}

			for i := smaller; i <= larger; i++ {
				board[i][smoke.end.y]++
			}

		} else if smoke.start.x == smoke.end.x {
			smaller, larger := 0, 0

			if smoke.start.y > smoke.end.y {
				smaller = smoke.end.y
				larger = smoke.start.y
			} else {
				larger = smoke.end.y
				smaller = smoke.start.y
			}

			for i := smaller; i <= larger; i++ {
				board[smoke.start.x][i]++
			}
		}
	}

	count := 0

	for _, row := range board {
		for _, cell := range row {
			if cell >= 2 {
				count++
			}
		}
	}

	// for _, row := range board {
	// 	fmt.Println(row)
	// }

	fmt.Println("Part 1: ", count)
}

func DayFiveP2() {
	input := getInput5()
	board := [][]int{}

	for i := 0; i < 1000; i++ {
		row := []int{}
		for j := 0; j < 1000; j++ {
			row = append(row, 0)
		}
		board = append(board, row)
	}

	for _, smoke := range input {
		smaller, larger := 0, 0
		if smoke.start.y == smoke.end.y {

			if smoke.start.x > smoke.end.x {
				smaller = smoke.end.x
				larger = smoke.start.x
			} else {
				larger = smoke.end.x
				smaller = smoke.start.x
			}

			for i := smaller; i <= larger; i++ {
				board[i][smoke.end.y]++
			}

		} else if smoke.start.x == smoke.end.x {
			if smoke.start.y > smoke.end.y {
				smaller = smoke.end.y
				larger = smoke.start.y
			} else {
				larger = smoke.end.y
				smaller = smoke.start.y
			}

			fmt.Println("smaller: ", smaller, " larger: ", larger, " point: ", smoke)

			for i := smaller; i <= larger; i++ {
				board[smoke.start.x][i]++
			}
		}

		spreadX, spreadY := math.Abs(float64(smoke.end.x-smoke.start.x)), math.Abs(float64(smoke.end.y-smoke.start.y))

		if spreadX == spreadY {
			if smoke.start.y < smoke.end.y {

			} else {

			}
		}
	}

	count := 0

	for _, row := range board {
		for _, cell := range row {
			if cell >= 2 {
				count++
			}
		}
	}

	// for _, row := range board {
	// 	fmt.Println(row)
	// }

	fmt.Println("Part 1: ", count)
}
