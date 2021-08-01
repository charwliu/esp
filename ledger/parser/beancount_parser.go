// Code generated from BeanCountParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // BeanCountParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import httperror
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 57, 463,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 117, 10, 4, 12, 4, 14, 4,
	120, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 128, 10, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 170, 10, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 178, 10, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26,
	277, 10, 26, 3, 27, 7, 27, 280, 10, 27, 12, 27, 14, 27, 283, 11, 27, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 294,
	10, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 5, 29, 306, 10, 29, 3, 30, 7, 30, 309, 10, 30, 12, 30, 14, 30, 312,
	11, 30, 3, 31, 7, 31, 315, 10, 31, 12, 31, 14, 31, 318, 11, 31, 3, 32,
	3, 32, 3, 32, 3, 32, 5, 32, 324, 10, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3,
	35, 3, 35, 5, 35, 332, 10, 35, 7, 35, 334, 10, 35, 12, 35, 14, 35, 337,
	11, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 345, 10, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 7, 36, 353, 10, 36, 12, 36, 14,
	36, 356, 11, 36, 3, 37, 3, 37, 3, 38, 5, 38, 361, 10, 38, 3, 39, 5, 39,
	364, 10, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 7, 40, 371, 10, 40, 12,
	40, 14, 40, 374, 11, 40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42,
	3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3,
	42, 3, 42, 3, 42, 3, 42, 5, 42, 397, 10, 42, 3, 43, 3, 43, 3, 44, 3, 44,
	3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 5, 45, 410, 10, 45, 3,
	46, 3, 46, 3, 46, 3, 47, 5, 47, 416, 10, 47, 3, 48, 5, 48, 419, 10, 48,
	3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 5, 49, 430,
	10, 49, 3, 50, 3, 50, 5, 50, 434, 10, 50, 3, 50, 3, 50, 3, 50, 7, 50, 439,
	10, 50, 12, 50, 14, 50, 442, 11, 50, 3, 51, 3, 51, 3, 51, 3, 51, 5, 51,
	448, 10, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3,
	52, 3, 52, 3, 52, 5, 52, 461, 10, 52, 3, 52, 2, 6, 6, 70, 78, 98, 53, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76,
	78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 2, 8, 3, 2, 54, 55,
	3, 2, 22, 23, 3, 2, 20, 21, 6, 2, 20, 20, 31, 31, 33, 33, 51, 51, 5, 2,
	20, 20, 31, 31, 51, 51, 3, 2, 12, 13, 2, 471, 2, 104, 3, 2, 2, 2, 4, 107,
	3, 2, 2, 2, 6, 109, 3, 2, 2, 2, 8, 127, 3, 2, 2, 2, 10, 129, 3, 2, 2, 2,
	12, 133, 3, 2, 2, 2, 14, 137, 3, 2, 2, 2, 16, 143, 3, 2, 2, 2, 18, 148,
	3, 2, 2, 2, 20, 153, 3, 2, 2, 2, 22, 169, 3, 2, 2, 2, 24, 171, 3, 2, 2,
	2, 26, 179, 3, 2, 2, 2, 28, 185, 3, 2, 2, 2, 30, 193, 3, 2, 2, 2, 32, 201,
	3, 2, 2, 2, 34, 210, 3, 2, 2, 2, 36, 217, 3, 2, 2, 2, 38, 224, 3, 2, 2,
	2, 40, 232, 3, 2, 2, 2, 42, 240, 3, 2, 2, 2, 44, 248, 3, 2, 2, 2, 46, 256,
	3, 2, 2, 2, 48, 264, 3, 2, 2, 2, 50, 276, 3, 2, 2, 2, 52, 281, 3, 2, 2,
	2, 54, 293, 3, 2, 2, 2, 56, 305, 3, 2, 2, 2, 58, 310, 3, 2, 2, 2, 60, 316,
	3, 2, 2, 2, 62, 323, 3, 2, 2, 2, 64, 325, 3, 2, 2, 2, 66, 327, 3, 2, 2,
	2, 68, 335, 3, 2, 2, 2, 70, 344, 3, 2, 2, 2, 72, 357, 3, 2, 2, 2, 74, 360,
	3, 2, 2, 2, 76, 363, 3, 2, 2, 2, 78, 365, 3, 2, 2, 2, 80, 375, 3, 2, 2,
	2, 82, 396, 3, 2, 2, 2, 84, 398, 3, 2, 2, 2, 86, 400, 3, 2, 2, 2, 88, 409,
	3, 2, 2, 2, 90, 411, 3, 2, 2, 2, 92, 415, 3, 2, 2, 2, 94, 418, 3, 2, 2,
	2, 96, 429, 3, 2, 2, 2, 98, 433, 3, 2, 2, 2, 100, 447, 3, 2, 2, 2, 102,
	460, 3, 2, 2, 2, 104, 105, 5, 6, 4, 2, 105, 106, 7, 2, 2, 3, 106, 3, 3,
	2, 2, 2, 107, 108, 7, 5, 2, 2, 108, 5, 3, 2, 2, 2, 109, 118, 8, 4, 1, 2,
	110, 111, 12, 6, 2, 2, 111, 117, 7, 5, 2, 2, 112, 113, 12, 5, 2, 2, 113,
	117, 5, 22, 12, 2, 114, 115, 12, 4, 2, 2, 115, 117, 5, 8, 5, 2, 116, 110,
	3, 2, 2, 2, 116, 112, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 117, 120, 3, 2,
	2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 7, 3, 2, 2, 2, 120,
	118, 3, 2, 2, 2, 121, 128, 5, 10, 6, 2, 122, 128, 5, 12, 7, 2, 123, 128,
	5, 14, 8, 2, 124, 128, 5, 16, 9, 2, 125, 128, 5, 18, 10, 2, 126, 128, 5,
	20, 11, 2, 127, 121, 3, 2, 2, 2, 127, 122, 3, 2, 2, 2, 127, 123, 3, 2,
	2, 2, 127, 124, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127, 126, 3, 2, 2, 2,
	128, 9, 3, 2, 2, 2, 129, 130, 7, 45, 2, 2, 130, 131, 7, 55, 2, 2, 131,
	132, 5, 4, 3, 2, 132, 11, 3, 2, 2, 2, 133, 134, 7, 46, 2, 2, 134, 135,
	7, 55, 2, 2, 135, 136, 5, 4, 3, 2, 136, 13, 3, 2, 2, 2, 137, 138, 7, 47,
	2, 2, 138, 139, 7, 56, 2, 2, 139, 140, 7, 24, 2, 2, 140, 141, 5, 56, 29,
	2, 141, 142, 5, 4, 3, 2, 142, 15, 3, 2, 2, 2, 143, 144, 7, 48, 2, 2, 144,
	145, 7, 56, 2, 2, 145, 146, 7, 24, 2, 2, 146, 147, 5, 4, 3, 2, 147, 17,
	3, 2, 2, 2, 148, 149, 7, 50, 2, 2, 149, 150, 7, 57, 2, 2, 150, 151, 7,
	57, 2, 2, 151, 152, 5, 4, 3, 2, 152, 19, 3, 2, 2, 2, 153, 154, 7, 49, 2,
	2, 154, 155, 7, 57, 2, 2, 155, 156, 5, 4, 3, 2, 156, 21, 3, 2, 2, 2, 157,
	170, 5, 24, 13, 2, 158, 170, 5, 28, 15, 2, 159, 170, 5, 30, 16, 2, 160,
	170, 5, 32, 17, 2, 161, 170, 5, 34, 18, 2, 162, 170, 5, 36, 19, 2, 163,
	170, 5, 38, 20, 2, 164, 170, 5, 40, 21, 2, 165, 170, 5, 42, 22, 2, 166,
	170, 5, 44, 23, 2, 167, 170, 5, 46, 24, 2, 168, 170, 5, 48, 25, 2, 169,
	157, 3, 2, 2, 2, 169, 158, 3, 2, 2, 2, 169, 159, 3, 2, 2, 2, 169, 160,
	3, 2, 2, 2, 169, 161, 3, 2, 2, 2, 169, 162, 3, 2, 2, 2, 169, 163, 3, 2,
	2, 2, 169, 164, 3, 2, 2, 2, 169, 165, 3, 2, 2, 2, 169, 166, 3, 2, 2, 2,
	169, 167, 3, 2, 2, 2, 169, 168, 3, 2, 2, 2, 170, 23, 3, 2, 2, 2, 171, 177,
	5, 26, 14, 2, 172, 173, 7, 3, 2, 2, 173, 174, 5, 52, 27, 2, 174, 175, 5,
	78, 40, 2, 175, 176, 7, 4, 2, 2, 176, 178, 3, 2, 2, 2, 177, 172, 3, 2,
	2, 2, 177, 178, 3, 2, 2, 2, 178, 25, 3, 2, 2, 2, 179, 180, 7, 52, 2, 2,
	180, 181, 5, 72, 37, 2, 181, 182, 5, 62, 32, 2, 182, 183, 5, 60, 31, 2,
	183, 184, 5, 4, 3, 2, 184, 27, 3, 2, 2, 2, 185, 186, 7, 52, 2, 2, 186,
	187, 7, 42, 2, 2, 187, 188, 7, 10, 2, 2, 188, 189, 5, 86, 44, 2, 189, 190,
	5, 60, 31, 2, 190, 191, 5, 4, 3, 2, 191, 192, 5, 50, 26, 2, 192, 29, 3,
	2, 2, 2, 193, 194, 7, 52, 2, 2, 194, 195, 7, 34, 2, 2, 195, 196, 5, 64,
	33, 2, 196, 197, 5, 88, 45, 2, 197, 198, 5, 60, 31, 2, 198, 199, 5, 4,
	3, 2, 199, 200, 5, 50, 26, 2, 200, 31, 3, 2, 2, 2, 201, 202, 7, 52, 2,
	2, 202, 203, 7, 35, 2, 2, 203, 204, 5, 64, 33, 2, 204, 205, 5, 68, 35,
	2, 205, 206, 5, 76, 39, 2, 206, 207, 5, 60, 31, 2, 207, 208, 5, 4, 3, 2,
	208, 209, 5, 50, 26, 2, 209, 33, 3, 2, 2, 2, 210, 211, 7, 52, 2, 2, 211,
	212, 7, 36, 2, 2, 212, 213, 5, 64, 33, 2, 213, 214, 5, 60, 31, 2, 214,
	215, 5, 4, 3, 2, 215, 216, 5, 50, 26, 2, 216, 35, 3, 2, 2, 2, 217, 218,
	7, 52, 2, 2, 218, 219, 7, 37, 2, 2, 219, 220, 7, 10, 2, 2, 220, 221, 5,
	60, 31, 2, 221, 222, 5, 4, 3, 2, 222, 223, 5, 50, 26, 2, 223, 37, 3, 2,
	2, 2, 224, 225, 7, 52, 2, 2, 225, 226, 7, 38, 2, 2, 226, 227, 5, 64, 33,
	2, 227, 228, 5, 64, 33, 2, 228, 229, 5, 60, 31, 2, 229, 230, 5, 4, 3, 2,
	230, 231, 5, 50, 26, 2, 231, 39, 3, 2, 2, 2, 232, 233, 7, 52, 2, 2, 233,
	234, 7, 44, 2, 2, 234, 235, 5, 64, 33, 2, 235, 236, 5, 66, 34, 2, 236,
	237, 5, 60, 31, 2, 237, 238, 5, 4, 3, 2, 238, 239, 5, 50, 26, 2, 239, 41,
	3, 2, 2, 2, 240, 241, 7, 52, 2, 2, 241, 242, 7, 43, 2, 2, 242, 243, 5,
	64, 33, 2, 243, 244, 7, 57, 2, 2, 244, 245, 5, 60, 31, 2, 245, 246, 5,
	4, 3, 2, 246, 247, 5, 50, 26, 2, 247, 43, 3, 2, 2, 2, 248, 249, 7, 52,
	2, 2, 249, 250, 7, 39, 2, 2, 250, 251, 7, 57, 2, 2, 251, 252, 7, 57, 2,
	2, 252, 253, 5, 60, 31, 2, 253, 254, 5, 4, 3, 2, 254, 255, 5, 50, 26, 2,
	255, 45, 3, 2, 2, 2, 256, 257, 7, 52, 2, 2, 257, 258, 7, 40, 2, 2, 258,
	259, 7, 57, 2, 2, 259, 260, 7, 57, 2, 2, 260, 261, 5, 60, 31, 2, 261, 262,
	5, 4, 3, 2, 262, 263, 5, 50, 26, 2, 263, 47, 3, 2, 2, 2, 264, 265, 7, 52,
	2, 2, 265, 266, 7, 41, 2, 2, 266, 267, 7, 57, 2, 2, 267, 268, 5, 58, 30,
	2, 268, 269, 5, 4, 3, 2, 269, 270, 5, 50, 26, 2, 270, 49, 3, 2, 2, 2, 271,
	272, 7, 3, 2, 2, 272, 273, 5, 52, 27, 2, 273, 274, 7, 4, 2, 2, 274, 277,
	3, 2, 2, 2, 275, 277, 3, 2, 2, 2, 276, 271, 3, 2, 2, 2, 276, 275, 3, 2,
	2, 2, 277, 51, 3, 2, 2, 2, 278, 280, 5, 54, 28, 2, 279, 278, 3, 2, 2, 2,
	280, 283, 3, 2, 2, 2, 281, 279, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282,
	53, 3, 2, 2, 2, 283, 281, 3, 2, 2, 2, 284, 285, 7, 56, 2, 2, 285, 286,
	7, 24, 2, 2, 286, 287, 5, 56, 29, 2, 287, 288, 5, 4, 3, 2, 288, 294, 3,
	2, 2, 2, 289, 290, 7, 55, 2, 2, 290, 294, 5, 4, 3, 2, 291, 292, 7, 54,
	2, 2, 292, 294, 5, 4, 3, 2, 293, 284, 3, 2, 2, 2, 293, 289, 3, 2, 2, 2,
	293, 291, 3, 2, 2, 2, 294, 55, 3, 2, 2, 2, 295, 306, 7, 57, 2, 2, 296,
	306, 7, 10, 2, 2, 297, 306, 5, 64, 33, 2, 298, 306, 7, 55, 2, 2, 299, 306,
	7, 54, 2, 2, 300, 306, 7, 52, 2, 2, 301, 306, 7, 8, 2, 2, 302, 306, 7,
	9, 2, 2, 303, 306, 5, 70, 36, 2, 304, 306, 5, 86, 44, 2, 305, 295, 3, 2,
	2, 2, 305, 296, 3, 2, 2, 2, 305, 297, 3, 2, 2, 2, 305, 298, 3, 2, 2, 2,
	305, 299, 3, 2, 2, 2, 305, 300, 3, 2, 2, 2, 305, 301, 3, 2, 2, 2, 305,
	302, 3, 2, 2, 2, 305, 303, 3, 2, 2, 2, 305, 304, 3, 2, 2, 2, 306, 57, 3,
	2, 2, 2, 307, 309, 5, 56, 29, 2, 308, 307, 3, 2, 2, 2, 309, 312, 3, 2,
	2, 2, 310, 308, 3, 2, 2, 2, 310, 311, 3, 2, 2, 2, 311, 59, 3, 2, 2, 2,
	312, 310, 3, 2, 2, 2, 313, 315, 9, 2, 2, 2, 314, 313, 3, 2, 2, 2, 315,
	318, 3, 2, 2, 2, 316, 314, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 61, 3,
	2, 2, 2, 318, 316, 3, 2, 2, 2, 319, 320, 7, 57, 2, 2, 320, 324, 7, 57,
	2, 2, 321, 324, 7, 57, 2, 2, 322, 324, 3, 2, 2, 2, 323, 319, 3, 2, 2, 2,
	323, 321, 3, 2, 2, 2, 323, 322, 3, 2, 2, 2, 324, 63, 3, 2, 2, 2, 325, 326,
	7, 53, 2, 2, 326, 65, 3, 2, 2, 2, 327, 328, 7, 57, 2, 2, 328, 67, 3, 2,
	2, 2, 329, 331, 7, 10, 2, 2, 330, 332, 7, 18, 2, 2, 331, 330, 3, 2, 2,
	2, 331, 332, 3, 2, 2, 2, 332, 334, 3, 2, 2, 2, 333, 329, 3, 2, 2, 2, 334,
	337, 3, 2, 2, 2, 335, 333, 3, 2, 2, 2, 335, 336, 3, 2, 2, 2, 336, 69, 3,
	2, 2, 2, 337, 335, 3, 2, 2, 2, 338, 339, 8, 36, 1, 2, 339, 340, 7, 25,
	2, 2, 340, 341, 5, 70, 36, 2, 341, 342, 7, 26, 2, 2, 342, 345, 3, 2, 2,
	2, 343, 345, 7, 32, 2, 2, 344, 338, 3, 2, 2, 2, 344, 343, 3, 2, 2, 2, 345,
	354, 3, 2, 2, 2, 346, 347, 12, 6, 2, 2, 347, 348, 9, 3, 2, 2, 348, 353,
	5, 70, 36, 7, 349, 350, 12, 5, 2, 2, 350, 351, 9, 4, 2, 2, 351, 353, 5,
	70, 36, 6, 352, 346, 3, 2, 2, 2, 352, 349, 3, 2, 2, 2, 353, 356, 3, 2,
	2, 2, 354, 352, 3, 2, 2, 2, 354, 355, 3, 2, 2, 2, 355, 71, 3, 2, 2, 2,
	356, 354, 3, 2, 2, 2, 357, 358, 9, 5, 2, 2, 358, 73, 3, 2, 2, 2, 359, 361,
	9, 6, 2, 2, 360, 359, 3, 2, 2, 2, 360, 361, 3, 2, 2, 2, 361, 75, 3, 2,
	2, 2, 362, 364, 7, 57, 2, 2, 363, 362, 3, 2, 2, 2, 363, 364, 3, 2, 2, 2,
	364, 77, 3, 2, 2, 2, 365, 372, 8, 40, 1, 2, 366, 367, 12, 4, 2, 2, 367,
	371, 5, 4, 3, 2, 368, 369, 12, 3, 2, 2, 369, 371, 5, 80, 41, 2, 370, 366,
	3, 2, 2, 2, 370, 368, 3, 2, 2, 2, 371, 374, 3, 2, 2, 2, 372, 370, 3, 2,
	2, 2, 372, 373, 3, 2, 2, 2, 373, 79, 3, 2, 2, 2, 374, 372, 3, 2, 2, 2,
	375, 376, 5, 82, 42, 2, 376, 377, 5, 50, 26, 2, 377, 81, 3, 2, 2, 2, 378,
	379, 5, 74, 38, 2, 379, 380, 5, 64, 33, 2, 380, 381, 5, 90, 46, 2, 381,
	382, 5, 96, 49, 2, 382, 383, 9, 7, 2, 2, 383, 384, 5, 84, 43, 2, 384, 385,
	5, 4, 3, 2, 385, 397, 3, 2, 2, 2, 386, 387, 5, 74, 38, 2, 387, 388, 5,
	64, 33, 2, 388, 389, 5, 90, 46, 2, 389, 390, 5, 96, 49, 2, 390, 391, 5,
	4, 3, 2, 391, 397, 3, 2, 2, 2, 392, 393, 5, 74, 38, 2, 393, 394, 5, 64,
	33, 2, 394, 395, 5, 4, 3, 2, 395, 397, 3, 2, 2, 2, 396, 378, 3, 2, 2, 2,
	396, 386, 3, 2, 2, 2, 396, 392, 3, 2, 2, 2, 397, 83, 3, 2, 2, 2, 398, 399,
	5, 90, 46, 2, 399, 85, 3, 2, 2, 2, 400, 401, 5, 70, 36, 2, 401, 402, 7,
	10, 2, 2, 402, 87, 3, 2, 2, 2, 403, 410, 5, 86, 44, 2, 404, 405, 5, 70,
	36, 2, 405, 406, 7, 19, 2, 2, 406, 407, 5, 70, 36, 2, 407, 408, 7, 10,
	2, 2, 408, 410, 3, 2, 2, 2, 409, 403, 3, 2, 2, 2, 409, 404, 3, 2, 2, 2,
	410, 89, 3, 2, 2, 2, 411, 412, 5, 92, 47, 2, 412, 413, 5, 94, 48, 2, 413,
	91, 3, 2, 2, 2, 414, 416, 5, 70, 36, 2, 415, 414, 3, 2, 2, 2, 415, 416,
	3, 2, 2, 2, 416, 93, 3, 2, 2, 2, 417, 419, 7, 10, 2, 2, 418, 417, 3, 2,
	2, 2, 418, 419, 3, 2, 2, 2, 419, 95, 3, 2, 2, 2, 420, 421, 7, 16, 2, 2,
	421, 422, 5, 98, 50, 2, 422, 423, 7, 17, 2, 2, 423, 430, 3, 2, 2, 2, 424,
	425, 7, 14, 2, 2, 425, 426, 5, 98, 50, 2, 426, 427, 7, 15, 2, 2, 427, 430,
	3, 2, 2, 2, 428, 430, 3, 2, 2, 2, 429, 420, 3, 2, 2, 2, 429, 424, 3, 2,
	2, 2, 429, 428, 3, 2, 2, 2, 430, 97, 3, 2, 2, 2, 431, 434, 8, 50, 1, 2,
	432, 434, 5, 100, 51, 2, 433, 431, 3, 2, 2, 2, 433, 432, 3, 2, 2, 2, 434,
	440, 3, 2, 2, 2, 435, 436, 12, 3, 2, 2, 436, 437, 7, 18, 2, 2, 437, 439,
	5, 100, 51, 2, 438, 435, 3, 2, 2, 2, 439, 442, 3, 2, 2, 2, 440, 438, 3,
	2, 2, 2, 440, 441, 3, 2, 2, 2, 441, 99, 3, 2, 2, 2, 442, 440, 3, 2, 2,
	2, 443, 448, 5, 102, 52, 2, 444, 448, 7, 52, 2, 2, 445, 448, 7, 57, 2,
	2, 446, 448, 7, 20, 2, 2, 447, 443, 3, 2, 2, 2, 447, 444, 3, 2, 2, 2, 447,
	445, 3, 2, 2, 2, 447, 446, 3, 2, 2, 2, 448, 101, 3, 2, 2, 2, 449, 450,
	5, 92, 47, 2, 450, 451, 7, 10, 2, 2, 451, 461, 3, 2, 2, 2, 452, 453, 5,
	70, 36, 2, 453, 454, 5, 94, 48, 2, 454, 461, 3, 2, 2, 2, 455, 456, 5, 92,
	47, 2, 456, 457, 7, 51, 2, 2, 457, 458, 5, 92, 47, 2, 458, 459, 7, 10,
	2, 2, 459, 461, 3, 2, 2, 2, 460, 449, 3, 2, 2, 2, 460, 452, 3, 2, 2, 2,
	460, 455, 3, 2, 2, 2, 461, 103, 3, 2, 2, 2, 32, 116, 118, 127, 169, 177,
	276, 281, 293, 305, 310, 316, 323, 331, 335, 344, 352, 354, 360, 363, 370,
	372, 396, 409, 415, 418, 429, 433, 440, 447, 460,
}
var literalNames = []string{
	"", "", "", "", "", "", "", "'NULL'", "", "'|'", "'@@'", "'@'", "'{{'",
	"'}}'", "'{'", "'}'", "','", "'~'", "'*'", "'/'", "'+'", "'-'", "':'",
	"'('", "')'", "'['", "']'", "';'", "'.'", "", "", "'txn'", "'balance'",
	"'open'", "'close'", "'commodity'", "'pad'", "'event'", "'query'", "'custom'",
	"'price'", "'note'", "'document'", "'pushtag'", "'poptag'", "'pushmeta'",
	"'popmeta'", "'include'", "'option'", "'#'",
}
var symbolicNames = []string{
	"", "INDENT", "DEDENT", "NEWLINE", "WHITESPACE", "COMMENTS", "BOOL", "NONE",
	"CURRENCY", "PIPE", "ATAT", "AT", "LCURLCURL", "RCURLCURL", "LCURL", "RCURL",
	"COMMA", "TILDE", "ASTERISK", "SLASH", "PLUS", "MINUS", "COLON", "LPAREN",
	"RPAREN", "LBRACK", "RBRACK", "SEMI", "DOT", "FLAG", "NUMBER", "TXN", "BALANCE",
	"OPEN", "CLOSE", "COMMODITY", "PAD", "EVENT", "QUERY", "CUSTOM", "PRICE",
	"NOTE", "DOCUMENT", "PUSHTAG", "POPTAG", "PUSHMETA", "POPMETA", "INCLUDE",
	"OPTION", "HASH", "DATETIME", "ACCOUNT", "LINK", "TAG", "KEY", "STRING",
}

var ruleNames = []string{
	"ledger", "eol", "declarations", "pragma", "pushtag", "poptag", "pushmeta",
	"popmeta", "option", "include", "directive", "transaction", "transactionLine",
	"price", "balance", "open", "closeDirective", "commodity", "pad", "document",
	"note", "event", "query", "custom", "indentedMetadata", "metadata", "metadataLine",
	"metaValue", "metaValueList", "tagsOrLinks", "payeeNarration", "account",
	"filename", "currencyList", "numberExpr", "txn", "optflag", "booking",
	"postingList", "postingAndMetadata", "posting", "priceAnnotation", "amount",
	"amountTolerance", "partialAmount", "maybeNumber", "maybeCurrency", "costSpec",
	"costCompList", "costComp", "compoundAmount",
}

type BeanCountParser struct {
	*antlr.BaseParser
}

// NewBeanCountParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *BeanCountParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewBeanCountParser(input antlr.TokenStream) *BeanCountParser {
	this := new(BeanCountParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "BeanCountParser.g4"

	return this
}

// BeanCountParser tokens.
const (
	BeanCountParserEOF        = antlr.TokenEOF
	BeanCountParserINDENT     = 1
	BeanCountParserDEDENT     = 2
	BeanCountParserNEWLINE    = 3
	BeanCountParserWHITESPACE = 4
	BeanCountParserCOMMENTS   = 5
	BeanCountParserBOOL       = 6
	BeanCountParserNONE       = 7
	BeanCountParserCURRENCY   = 8
	BeanCountParserPIPE       = 9
	BeanCountParserATAT       = 10
	BeanCountParserAT         = 11
	BeanCountParserLCURLCURL  = 12
	BeanCountParserRCURLCURL  = 13
	BeanCountParserLCURL      = 14
	BeanCountParserRCURL      = 15
	BeanCountParserCOMMA      = 16
	BeanCountParserTILDE      = 17
	BeanCountParserASTERISK   = 18
	BeanCountParserSLASH      = 19
	BeanCountParserPLUS       = 20
	BeanCountParserMINUS      = 21
	BeanCountParserCOLON      = 22
	BeanCountParserLPAREN     = 23
	BeanCountParserRPAREN     = 24
	BeanCountParserLBRACK     = 25
	BeanCountParserRBRACK     = 26
	BeanCountParserSEMI       = 27
	BeanCountParserDOT        = 28
	BeanCountParserFLAG       = 29
	BeanCountParserNUMBER     = 30
	BeanCountParserTXN        = 31
	BeanCountParserBALANCE    = 32
	BeanCountParserOPEN       = 33
	BeanCountParserCLOSE      = 34
	BeanCountParserCOMMODITY  = 35
	BeanCountParserPAD        = 36
	BeanCountParserEVENT      = 37
	BeanCountParserQUERY      = 38
	BeanCountParserCUSTOM     = 39
	BeanCountParserPRICE      = 40
	BeanCountParserNOTE       = 41
	BeanCountParserDOCUMENT   = 42
	BeanCountParserPUSHTAG    = 43
	BeanCountParserPOPTAG     = 44
	BeanCountParserPUSHMETA   = 45
	BeanCountParserPOPMETA    = 46
	BeanCountParserINCLUDE    = 47
	BeanCountParserOPTION     = 48
	BeanCountParserHASH       = 49
	BeanCountParserDATETIME   = 50
	BeanCountParserACCOUNT    = 51
	BeanCountParserLINK       = 52
	BeanCountParserTAG        = 53
	BeanCountParserKEY        = 54
	BeanCountParserSTRING     = 55
)

// BeanCountParser rules.
const (
	BeanCountParserRULE_ledger             = 0
	BeanCountParserRULE_eol                = 1
	BeanCountParserRULE_declarations       = 2
	BeanCountParserRULE_pragma             = 3
	BeanCountParserRULE_pushtag            = 4
	BeanCountParserRULE_poptag             = 5
	BeanCountParserRULE_pushmeta           = 6
	BeanCountParserRULE_popmeta            = 7
	BeanCountParserRULE_option             = 8
	BeanCountParserRULE_include            = 9
	BeanCountParserRULE_directive          = 10
	BeanCountParserRULE_transaction        = 11
	BeanCountParserRULE_transactionLine    = 12
	BeanCountParserRULE_price              = 13
	BeanCountParserRULE_balance            = 14
	BeanCountParserRULE_open               = 15
	BeanCountParserRULE_closeDirective     = 16
	BeanCountParserRULE_commodity          = 17
	BeanCountParserRULE_pad                = 18
	BeanCountParserRULE_document           = 19
	BeanCountParserRULE_note               = 20
	BeanCountParserRULE_event              = 21
	BeanCountParserRULE_query              = 22
	BeanCountParserRULE_custom             = 23
	BeanCountParserRULE_indentedMetadata   = 24
	BeanCountParserRULE_metadata           = 25
	BeanCountParserRULE_metadataLine       = 26
	BeanCountParserRULE_metaValue          = 27
	BeanCountParserRULE_metaValueList      = 28
	BeanCountParserRULE_tagsOrLinks        = 29
	BeanCountParserRULE_payeeNarration     = 30
	BeanCountParserRULE_account            = 31
	BeanCountParserRULE_filename           = 32
	BeanCountParserRULE_currencyList       = 33
	BeanCountParserRULE_numberExpr         = 34
	BeanCountParserRULE_txn                = 35
	BeanCountParserRULE_optflag            = 36
	BeanCountParserRULE_booking            = 37
	BeanCountParserRULE_postingList        = 38
	BeanCountParserRULE_postingAndMetadata = 39
	BeanCountParserRULE_posting            = 40
	BeanCountParserRULE_priceAnnotation    = 41
	BeanCountParserRULE_amount             = 42
	BeanCountParserRULE_amountTolerance    = 43
	BeanCountParserRULE_partialAmount      = 44
	BeanCountParserRULE_maybeNumber        = 45
	BeanCountParserRULE_maybeCurrency      = 46
	BeanCountParserRULE_costSpec           = 47
	BeanCountParserRULE_costCompList       = 48
	BeanCountParserRULE_costComp           = 49
	BeanCountParserRULE_compoundAmount     = 50
)

// ILedgerContext is an interface to support dynamic dispatch.
type ILedgerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLedgerContext differentiates from other interfaces.
	IsLedgerContext()
}

type LedgerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLedgerContext() *LedgerContext {
	var p = new(LedgerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_ledger
	return p
}

func (*LedgerContext) IsLedgerContext() {}

func NewLedgerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LedgerContext {
	var p = new(LedgerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_ledger

	return p
}

func (s *LedgerContext) GetParser() antlr.Parser { return s.parser }

func (s *LedgerContext) Declarations() IDeclarationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationsContext)
}

func (s *LedgerContext) EOF() antlr.TerminalNode {
	return s.GetToken(BeanCountParserEOF, 0)
}

func (s *LedgerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LedgerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LedgerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitLedger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Ledger() (localctx ILedgerContext) {
	localctx = NewLedgerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BeanCountParserRULE_ledger)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.declarations(0)
	}
	{
		p.SetState(103)
		p.Match(BeanCountParserEOF)
	}

	return localctx
}

// IEolContext is an interface to support dynamic dispatch.
type IEolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEolContext differentiates from other interfaces.
	IsEolContext()
}

type EolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEolContext() *EolContext {
	var p = new(EolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_eol
	return p
}

func (*EolContext) IsEolContext() {}

func NewEolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EolContext {
	var p = new(EolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_eol

	return p
}

func (s *EolContext) GetParser() antlr.Parser { return s.parser }

func (s *EolContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeanCountParserNEWLINE, 0)
}

func (s *EolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitEol(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Eol() (localctx IEolContext) {
	localctx = NewEolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BeanCountParserRULE_eol)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(BeanCountParserNEWLINE)
	}

	return localctx
}

// IDeclarationsContext is an interface to support dynamic dispatch.
type IDeclarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationsContext differentiates from other interfaces.
	IsDeclarationsContext()
}

type DeclarationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationsContext() *DeclarationsContext {
	var p = new(DeclarationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_declarations
	return p
}

func (*DeclarationsContext) IsDeclarationsContext() {}

func NewDeclarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationsContext {
	var p = new(DeclarationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_declarations

	return p
}

func (s *DeclarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationsContext) Declarations() IDeclarationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationsContext)
}

func (s *DeclarationsContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeanCountParserNEWLINE, 0)
}

func (s *DeclarationsContext) Directive() IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *DeclarationsContext) Pragma() IPragmaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPragmaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPragmaContext)
}

func (s *DeclarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitDeclarations(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Declarations() (localctx IDeclarationsContext) {
	return p.declarations(0)
}

func (p *BeanCountParser) declarations(_p int) (localctx IDeclarationsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDeclarationsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDeclarationsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, BeanCountParserRULE_declarations, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(114)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDeclarationsContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, BeanCountParserRULE_declarations)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(109)
					p.Match(BeanCountParserNEWLINE)
				}

			case 2:
				localctx = NewDeclarationsContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, BeanCountParserRULE_declarations)
				p.SetState(110)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(111)
					p.Directive()
				}

			case 3:
				localctx = NewDeclarationsContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, BeanCountParserRULE_declarations)
				p.SetState(112)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(113)
					p.Pragma()
				}

			}

		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IPragmaContext is an interface to support dynamic dispatch.
type IPragmaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPragmaContext differentiates from other interfaces.
	IsPragmaContext()
}

type PragmaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPragmaContext() *PragmaContext {
	var p = new(PragmaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_pragma
	return p
}

func (*PragmaContext) IsPragmaContext() {}

func NewPragmaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PragmaContext {
	var p = new(PragmaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_pragma

	return p
}

func (s *PragmaContext) GetParser() antlr.Parser { return s.parser }

func (s *PragmaContext) Pushtag() IPushtagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPushtagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPushtagContext)
}

func (s *PragmaContext) Poptag() IPoptagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPoptagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPoptagContext)
}

func (s *PragmaContext) Pushmeta() IPushmetaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPushmetaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPushmetaContext)
}

func (s *PragmaContext) Popmeta() IPopmetaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPopmetaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPopmetaContext)
}

func (s *PragmaContext) Option() IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *PragmaContext) Include() IIncludeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncludeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncludeContext)
}

func (s *PragmaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PragmaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PragmaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPragma(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Pragma() (localctx IPragmaContext) {
	localctx = NewPragmaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BeanCountParserRULE_pragma)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(125)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeanCountParserPUSHTAG:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Pushtag()
		}

	case BeanCountParserPOPTAG:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Poptag()
		}

	case BeanCountParserPUSHMETA:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(121)
			p.Pushmeta()
		}

	case BeanCountParserPOPMETA:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(122)
			p.Popmeta()
		}

	case BeanCountParserOPTION:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(123)
			p.Option()
		}

	case BeanCountParserINCLUDE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(124)
			p.Include()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPushtagContext is an interface to support dynamic dispatch.
type IPushtagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPushtagContext differentiates from other interfaces.
	IsPushtagContext()
}

type PushtagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPushtagContext() *PushtagContext {
	var p = new(PushtagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_pushtag
	return p
}

func (*PushtagContext) IsPushtagContext() {}

func NewPushtagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PushtagContext {
	var p = new(PushtagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_pushtag

	return p
}

func (s *PushtagContext) GetParser() antlr.Parser { return s.parser }

func (s *PushtagContext) PUSHTAG() antlr.TerminalNode {
	return s.GetToken(BeanCountParserPUSHTAG, 0)
}

func (s *PushtagContext) TAG() antlr.TerminalNode {
	return s.GetToken(BeanCountParserTAG, 0)
}

func (s *PushtagContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *PushtagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PushtagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PushtagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPushtag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Pushtag() (localctx IPushtagContext) {
	localctx = NewPushtagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BeanCountParserRULE_pushtag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(BeanCountParserPUSHTAG)
	}
	{
		p.SetState(128)
		p.Match(BeanCountParserTAG)
	}
	{
		p.SetState(129)
		p.Eol()
	}

	return localctx
}

// IPoptagContext is an interface to support dynamic dispatch.
type IPoptagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPoptagContext differentiates from other interfaces.
	IsPoptagContext()
}

type PoptagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPoptagContext() *PoptagContext {
	var p = new(PoptagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_poptag
	return p
}

func (*PoptagContext) IsPoptagContext() {}

func NewPoptagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PoptagContext {
	var p = new(PoptagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_poptag

	return p
}

func (s *PoptagContext) GetParser() antlr.Parser { return s.parser }

func (s *PoptagContext) POPTAG() antlr.TerminalNode {
	return s.GetToken(BeanCountParserPOPTAG, 0)
}

func (s *PoptagContext) TAG() antlr.TerminalNode {
	return s.GetToken(BeanCountParserTAG, 0)
}

func (s *PoptagContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *PoptagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PoptagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PoptagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPoptag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Poptag() (localctx IPoptagContext) {
	localctx = NewPoptagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BeanCountParserRULE_poptag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(BeanCountParserPOPTAG)
	}
	{
		p.SetState(132)
		p.Match(BeanCountParserTAG)
	}
	{
		p.SetState(133)
		p.Eol()
	}

	return localctx
}

// IPushmetaContext is an interface to support dynamic dispatch.
type IPushmetaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPushmetaContext differentiates from other interfaces.
	IsPushmetaContext()
}

type PushmetaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPushmetaContext() *PushmetaContext {
	var p = new(PushmetaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_pushmeta
	return p
}

func (*PushmetaContext) IsPushmetaContext() {}

func NewPushmetaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PushmetaContext {
	var p = new(PushmetaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_pushmeta

	return p
}

func (s *PushmetaContext) GetParser() antlr.Parser { return s.parser }

func (s *PushmetaContext) PUSHMETA() antlr.TerminalNode {
	return s.GetToken(BeanCountParserPUSHMETA, 0)
}

func (s *PushmetaContext) KEY() antlr.TerminalNode {
	return s.GetToken(BeanCountParserKEY, 0)
}

func (s *PushmetaContext) COLON() antlr.TerminalNode {
	return s.GetToken(BeanCountParserCOLON, 0)
}

func (s *PushmetaContext) MetaValue() IMetaValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetaValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMetaValueContext)
}

func (s *PushmetaContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *PushmetaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PushmetaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PushmetaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPushmeta(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Pushmeta() (localctx IPushmetaContext) {
	localctx = NewPushmetaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BeanCountParserRULE_pushmeta)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(BeanCountParserPUSHMETA)
	}
	{
		p.SetState(136)
		p.Match(BeanCountParserKEY)
	}
	{
		p.SetState(137)
		p.Match(BeanCountParserCOLON)
	}
	{
		p.SetState(138)
		p.MetaValue()
	}
	{
		p.SetState(139)
		p.Eol()
	}

	return localctx
}

// IPopmetaContext is an interface to support dynamic dispatch.
type IPopmetaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPopmetaContext differentiates from other interfaces.
	IsPopmetaContext()
}

type PopmetaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPopmetaContext() *PopmetaContext {
	var p = new(PopmetaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_popmeta
	return p
}

func (*PopmetaContext) IsPopmetaContext() {}

func NewPopmetaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PopmetaContext {
	var p = new(PopmetaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_popmeta

	return p
}

func (s *PopmetaContext) GetParser() antlr.Parser { return s.parser }

func (s *PopmetaContext) POPMETA() antlr.TerminalNode {
	return s.GetToken(BeanCountParserPOPMETA, 0)
}

func (s *PopmetaContext) KEY() antlr.TerminalNode {
	return s.GetToken(BeanCountParserKEY, 0)
}

func (s *PopmetaContext) COLON() antlr.TerminalNode {
	return s.GetToken(BeanCountParserCOLON, 0)
}

func (s *PopmetaContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *PopmetaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PopmetaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PopmetaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPopmeta(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Popmeta() (localctx IPopmetaContext) {
	localctx = NewPopmetaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BeanCountParserRULE_popmeta)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(BeanCountParserPOPMETA)
	}
	{
		p.SetState(142)
		p.Match(BeanCountParserKEY)
	}
	{
		p.SetState(143)
		p.Match(BeanCountParserCOLON)
	}
	{
		p.SetState(144)
		p.Eol()
	}

	return localctx
}

// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
	value  antlr.Token
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_option
	return p
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) GetKey() antlr.Token { return s.key }

func (s *OptionContext) GetValue() antlr.Token { return s.value }

func (s *OptionContext) SetKey(v antlr.Token) { s.key = v }

func (s *OptionContext) SetValue(v antlr.Token) { s.value = v }

func (s *OptionContext) OPTION() antlr.TerminalNode {
	return s.GetToken(BeanCountParserOPTION, 0)
}

func (s *OptionContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *OptionContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(BeanCountParserSTRING)
}

func (s *OptionContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(BeanCountParserSTRING, i)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Option() (localctx IOptionContext) {
	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BeanCountParserRULE_option)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(BeanCountParserOPTION)
	}
	{
		p.SetState(147)

		var _m = p.Match(BeanCountParserSTRING)

		localctx.(*OptionContext).key = _m
	}
	{
		p.SetState(148)

		var _m = p.Match(BeanCountParserSTRING)

		localctx.(*OptionContext).value = _m
	}
	{
		p.SetState(149)
		p.Eol()
	}

	return localctx
}

// IIncludeContext is an interface to support dynamic dispatch.
type IIncludeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncludeContext differentiates from other interfaces.
	IsIncludeContext()
}

type IncludeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludeContext() *IncludeContext {
	var p = new(IncludeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_include
	return p
}

func (*IncludeContext) IsIncludeContext() {}

func NewIncludeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludeContext {
	var p = new(IncludeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_include

	return p
}

func (s *IncludeContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludeContext) INCLUDE() antlr.TerminalNode {
	return s.GetToken(BeanCountParserINCLUDE, 0)
}

func (s *IncludeContext) STRING() antlr.TerminalNode {
	return s.GetToken(BeanCountParserSTRING, 0)
}

func (s *IncludeContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *IncludeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitInclude(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Include() (localctx IIncludeContext) {
	localctx = NewIncludeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BeanCountParserRULE_include)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(BeanCountParserINCLUDE)
	}
	{
		p.SetState(152)
		p.Match(BeanCountParserSTRING)
	}
	{
		p.SetState(153)
		p.Eol()
	}

	return localctx
}

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_directive
	return p
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) CopyFrom(ctx *DirectiveContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DocumentDirContext struct {
	*DirectiveContext
}

func NewDocumentDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DocumentDirContext {
	var p = new(DocumentDirContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *DocumentDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentDirContext) Document() IDocumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocumentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDocumentContext)
}

func (s *DocumentDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitDocumentDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type EventDirContext struct {
	*DirectiveContext
}

func NewEventDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EventDirContext {
	var p = new(EventDirContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *EventDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventDirContext) Event() IEventContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEventContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEventContext)
}

func (s *EventDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitEventDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type QueryDirContext struct {
	*DirectiveContext
}

func NewQueryDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QueryDirContext {
	var p = new(QueryDirContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *QueryDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryDirContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *QueryDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitQueryDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type CustomDirContext struct {
	*DirectiveContext
}

func NewCustomDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CustomDirContext {
	var p = new(CustomDirContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *CustomDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CustomDirContext) Custom() ICustomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICustomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICustomContext)
}

func (s *CustomDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitCustomDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type NoteDirContext struct {
	*DirectiveContext
}

func NewNoteDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NoteDirContext {
	var p = new(NoteDirContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *NoteDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoteDirContext) Note() INoteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoteContext)
}

func (s *NoteDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitNoteDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type TransactionDirContext struct {
	*DirectiveContext
}

func NewTransactionDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TransactionDirContext {
	var p = new(TransactionDirContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *TransactionDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransactionDirContext) Transaction() ITransactionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransactionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransactionContext)
}

func (s *TransactionDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitTransactionDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpenDirContext struct {
	*DirectiveContext
}

func NewOpenDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpenDirContext {
	var p = new(OpenDirContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *OpenDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpenDirContext) Open() IOpenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpenContext)
}

func (s *OpenDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitOpenDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type CloseDirContext struct {
	*DirectiveContext
}

func NewCloseDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CloseDirContext {
	var p = new(CloseDirContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *CloseDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CloseDirContext) CloseDirective() ICloseDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICloseDirectiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICloseDirectiveContext)
}

func (s *CloseDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitCloseDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type PriceDirContext struct {
	*DirectiveContext
}

func NewPriceDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PriceDirContext {
	var p = new(PriceDirContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *PriceDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PriceDirContext) Price() IPriceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPriceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPriceContext)
}

func (s *PriceDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPriceDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type BalanceDirContext struct {
	*DirectiveContext
}

func NewBalanceDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BalanceDirContext {
	var p = new(BalanceDirContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *BalanceDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BalanceDirContext) Balance() IBalanceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBalanceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBalanceContext)
}

func (s *BalanceDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitBalanceDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type CommodityDirContext struct {
	*DirectiveContext
}

func NewCommodityDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CommodityDirContext {
	var p = new(CommodityDirContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *CommodityDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommodityDirContext) Commodity() ICommodityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommodityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommodityContext)
}

func (s *CommodityDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitCommodityDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type PadDirContext struct {
	*DirectiveContext
}

func NewPadDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PadDirContext {
	var p = new(PadDirContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *PadDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PadDirContext) Pad() IPadContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPadContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPadContext)
}

func (s *PadDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPadDir(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BeanCountParserRULE_directive)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTransactionDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.Transaction()
		}

	case 2:
		localctx = NewPriceDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Price()
		}

	case 3:
		localctx = NewBalanceDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)
			p.Balance()
		}

	case 4:
		localctx = NewOpenDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(158)
			p.Open()
		}

	case 5:
		localctx = NewCloseDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(159)
			p.CloseDirective()
		}

	case 6:
		localctx = NewCommodityDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(160)
			p.Commodity()
		}

	case 7:
		localctx = NewPadDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(161)
			p.Pad()
		}

	case 8:
		localctx = NewDocumentDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(162)
			p.Document()
		}

	case 9:
		localctx = NewNoteDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(163)
			p.Note()
		}

	case 10:
		localctx = NewEventDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(164)
			p.Event()
		}

	case 11:
		localctx = NewQueryDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(165)
			p.Query()
		}

	case 12:
		localctx = NewCustomDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(166)
			p.Custom()
		}

	}

	return localctx
}

// ITransactionContext is an interface to support dynamic dispatch.
type ITransactionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTransactionContext differentiates from other interfaces.
	IsTransactionContext()
}

type TransactionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransactionContext() *TransactionContext {
	var p = new(TransactionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_transaction
	return p
}

func (*TransactionContext) IsTransactionContext() {}

func NewTransactionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransactionContext {
	var p = new(TransactionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_transaction

	return p
}

func (s *TransactionContext) GetParser() antlr.Parser { return s.parser }

func (s *TransactionContext) TransactionLine() ITransactionLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransactionLineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransactionLineContext)
}

func (s *TransactionContext) INDENT() antlr.TerminalNode {
	return s.GetToken(BeanCountParserINDENT, 0)
}

func (s *TransactionContext) Metadata() IMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMetadataContext)
}

func (s *TransactionContext) PostingList() IPostingListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostingListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostingListContext)
}

func (s *TransactionContext) DEDENT() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDEDENT, 0)
}

func (s *TransactionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransactionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransactionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitTransaction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Transaction() (localctx ITransactionContext) {
	localctx = NewTransactionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BeanCountParserRULE_transaction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.TransactionLine()
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(170)
			p.Match(BeanCountParserINDENT)
		}
		{
			p.SetState(171)
			p.Metadata()
		}
		{
			p.SetState(172)
			p.postingList(0)
		}
		{
			p.SetState(173)
			p.Match(BeanCountParserDEDENT)
		}

	}

	return localctx
}

// ITransactionLineContext is an interface to support dynamic dispatch.
type ITransactionLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTransactionLineContext differentiates from other interfaces.
	IsTransactionLineContext()
}

type TransactionLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransactionLineContext() *TransactionLineContext {
	var p = new(TransactionLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_transactionLine
	return p
}

func (*TransactionLineContext) IsTransactionLineContext() {}

func NewTransactionLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransactionLineContext {
	var p = new(TransactionLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_transactionLine

	return p
}

func (s *TransactionLineContext) GetParser() antlr.Parser { return s.parser }

func (s *TransactionLineContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDATETIME, 0)
}

func (s *TransactionLineContext) Txn() ITxnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITxnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITxnContext)
}

func (s *TransactionLineContext) PayeeNarration() IPayeeNarrationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPayeeNarrationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPayeeNarrationContext)
}

func (s *TransactionLineContext) TagsOrLinks() ITagsOrLinksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsOrLinksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagsOrLinksContext)
}

func (s *TransactionLineContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *TransactionLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransactionLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransactionLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitTransactionLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) TransactionLine() (localctx ITransactionLineContext) {
	localctx = NewTransactionLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BeanCountParserRULE_transactionLine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(BeanCountParserDATETIME)
	}
	{
		p.SetState(178)
		p.Txn()
	}
	{
		p.SetState(179)
		p.PayeeNarration()
	}
	{
		p.SetState(180)
		p.TagsOrLinks()
	}
	{
		p.SetState(181)
		p.Eol()
	}

	return localctx
}

// IPriceContext is an interface to support dynamic dispatch.
type IPriceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPriceContext differentiates from other interfaces.
	IsPriceContext()
}

type PriceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPriceContext() *PriceContext {
	var p = new(PriceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_price
	return p
}

func (*PriceContext) IsPriceContext() {}

func NewPriceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PriceContext {
	var p = new(PriceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_price

	return p
}

func (s *PriceContext) GetParser() antlr.Parser { return s.parser }

func (s *PriceContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDATETIME, 0)
}

func (s *PriceContext) PRICE() antlr.TerminalNode {
	return s.GetToken(BeanCountParserPRICE, 0)
}

func (s *PriceContext) CURRENCY() antlr.TerminalNode {
	return s.GetToken(BeanCountParserCURRENCY, 0)
}

func (s *PriceContext) Amount() IAmountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAmountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAmountContext)
}

func (s *PriceContext) TagsOrLinks() ITagsOrLinksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsOrLinksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagsOrLinksContext)
}

func (s *PriceContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *PriceContext) IndentedMetadata() IIndentedMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndentedMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndentedMetadataContext)
}

func (s *PriceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PriceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PriceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPrice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Price() (localctx IPriceContext) {
	localctx = NewPriceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BeanCountParserRULE_price)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(BeanCountParserDATETIME)
	}
	{
		p.SetState(184)
		p.Match(BeanCountParserPRICE)
	}
	{
		p.SetState(185)
		p.Match(BeanCountParserCURRENCY)
	}
	{
		p.SetState(186)
		p.Amount()
	}
	{
		p.SetState(187)
		p.TagsOrLinks()
	}
	{
		p.SetState(188)
		p.Eol()
	}
	{
		p.SetState(189)
		p.IndentedMetadata()
	}

	return localctx
}

// IBalanceContext is an interface to support dynamic dispatch.
type IBalanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBalanceContext differentiates from other interfaces.
	IsBalanceContext()
}

type BalanceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBalanceContext() *BalanceContext {
	var p = new(BalanceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_balance
	return p
}

func (*BalanceContext) IsBalanceContext() {}

func NewBalanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BalanceContext {
	var p = new(BalanceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_balance

	return p
}

func (s *BalanceContext) GetParser() antlr.Parser { return s.parser }

func (s *BalanceContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDATETIME, 0)
}

func (s *BalanceContext) BALANCE() antlr.TerminalNode {
	return s.GetToken(BeanCountParserBALANCE, 0)
}

func (s *BalanceContext) Account() IAccountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccountContext)
}

func (s *BalanceContext) AmountTolerance() IAmountToleranceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAmountToleranceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAmountToleranceContext)
}

func (s *BalanceContext) TagsOrLinks() ITagsOrLinksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsOrLinksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagsOrLinksContext)
}

func (s *BalanceContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *BalanceContext) IndentedMetadata() IIndentedMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndentedMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndentedMetadataContext)
}

func (s *BalanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BalanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BalanceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitBalance(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Balance() (localctx IBalanceContext) {
	localctx = NewBalanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BeanCountParserRULE_balance)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Match(BeanCountParserDATETIME)
	}
	{
		p.SetState(192)
		p.Match(BeanCountParserBALANCE)
	}
	{
		p.SetState(193)
		p.Account()
	}
	{
		p.SetState(194)
		p.AmountTolerance()
	}
	{
		p.SetState(195)
		p.TagsOrLinks()
	}
	{
		p.SetState(196)
		p.Eol()
	}
	{
		p.SetState(197)
		p.IndentedMetadata()
	}

	return localctx
}

// IOpenContext is an interface to support dynamic dispatch.
type IOpenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpenContext differentiates from other interfaces.
	IsOpenContext()
}

type OpenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpenContext() *OpenContext {
	var p = new(OpenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_open
	return p
}

func (*OpenContext) IsOpenContext() {}

func NewOpenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpenContext {
	var p = new(OpenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_open

	return p
}

func (s *OpenContext) GetParser() antlr.Parser { return s.parser }

func (s *OpenContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDATETIME, 0)
}

func (s *OpenContext) OPEN() antlr.TerminalNode {
	return s.GetToken(BeanCountParserOPEN, 0)
}

func (s *OpenContext) Account() IAccountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccountContext)
}

func (s *OpenContext) CurrencyList() ICurrencyListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICurrencyListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICurrencyListContext)
}

func (s *OpenContext) Booking() IBookingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBookingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBookingContext)
}

func (s *OpenContext) TagsOrLinks() ITagsOrLinksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsOrLinksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagsOrLinksContext)
}

func (s *OpenContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *OpenContext) IndentedMetadata() IIndentedMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndentedMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndentedMetadataContext)
}

func (s *OpenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitOpen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Open() (localctx IOpenContext) {
	localctx = NewOpenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BeanCountParserRULE_open)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(BeanCountParserDATETIME)
	}
	{
		p.SetState(200)
		p.Match(BeanCountParserOPEN)
	}
	{
		p.SetState(201)
		p.Account()
	}
	{
		p.SetState(202)
		p.CurrencyList()
	}
	{
		p.SetState(203)
		p.Booking()
	}
	{
		p.SetState(204)
		p.TagsOrLinks()
	}
	{
		p.SetState(205)
		p.Eol()
	}
	{
		p.SetState(206)
		p.IndentedMetadata()
	}

	return localctx
}

// ICloseDirectiveContext is an interface to support dynamic dispatch.
type ICloseDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCloseDirectiveContext differentiates from other interfaces.
	IsCloseDirectiveContext()
}

type CloseDirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCloseDirectiveContext() *CloseDirectiveContext {
	var p = new(CloseDirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_closeDirective
	return p
}

func (*CloseDirectiveContext) IsCloseDirectiveContext() {}

func NewCloseDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CloseDirectiveContext {
	var p = new(CloseDirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_closeDirective

	return p
}

func (s *CloseDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *CloseDirectiveContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDATETIME, 0)
}

func (s *CloseDirectiveContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(BeanCountParserCLOSE, 0)
}

func (s *CloseDirectiveContext) Account() IAccountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccountContext)
}

func (s *CloseDirectiveContext) TagsOrLinks() ITagsOrLinksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsOrLinksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagsOrLinksContext)
}

func (s *CloseDirectiveContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *CloseDirectiveContext) IndentedMetadata() IIndentedMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndentedMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndentedMetadataContext)
}

func (s *CloseDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CloseDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CloseDirectiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitCloseDirective(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) CloseDirective() (localctx ICloseDirectiveContext) {
	localctx = NewCloseDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BeanCountParserRULE_closeDirective)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(BeanCountParserDATETIME)
	}
	{
		p.SetState(209)
		p.Match(BeanCountParserCLOSE)
	}
	{
		p.SetState(210)
		p.Account()
	}
	{
		p.SetState(211)
		p.TagsOrLinks()
	}
	{
		p.SetState(212)
		p.Eol()
	}
	{
		p.SetState(213)
		p.IndentedMetadata()
	}

	return localctx
}

// ICommodityContext is an interface to support dynamic dispatch.
type ICommodityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommodityContext differentiates from other interfaces.
	IsCommodityContext()
}

type CommodityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommodityContext() *CommodityContext {
	var p = new(CommodityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_commodity
	return p
}

func (*CommodityContext) IsCommodityContext() {}

func NewCommodityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommodityContext {
	var p = new(CommodityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_commodity

	return p
}

func (s *CommodityContext) GetParser() antlr.Parser { return s.parser }

func (s *CommodityContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDATETIME, 0)
}

func (s *CommodityContext) COMMODITY() antlr.TerminalNode {
	return s.GetToken(BeanCountParserCOMMODITY, 0)
}

func (s *CommodityContext) CURRENCY() antlr.TerminalNode {
	return s.GetToken(BeanCountParserCURRENCY, 0)
}

func (s *CommodityContext) TagsOrLinks() ITagsOrLinksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsOrLinksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagsOrLinksContext)
}

func (s *CommodityContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *CommodityContext) IndentedMetadata() IIndentedMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndentedMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndentedMetadataContext)
}

func (s *CommodityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommodityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommodityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitCommodity(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Commodity() (localctx ICommodityContext) {
	localctx = NewCommodityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BeanCountParserRULE_commodity)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(BeanCountParserDATETIME)
	}
	{
		p.SetState(216)
		p.Match(BeanCountParserCOMMODITY)
	}
	{
		p.SetState(217)
		p.Match(BeanCountParserCURRENCY)
	}
	{
		p.SetState(218)
		p.TagsOrLinks()
	}
	{
		p.SetState(219)
		p.Eol()
	}
	{
		p.SetState(220)
		p.IndentedMetadata()
	}

	return localctx
}

// IPadContext is an interface to support dynamic dispatch.
type IPadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSource returns the source rule contexts.
	GetSource() IAccountContext

	// SetSource sets the source rule contexts.
	SetSource(IAccountContext)

	// IsPadContext differentiates from other interfaces.
	IsPadContext()
}

type PadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	source IAccountContext
}

func NewEmptyPadContext() *PadContext {
	var p = new(PadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_pad
	return p
}

func (*PadContext) IsPadContext() {}

func NewPadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PadContext {
	var p = new(PadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_pad

	return p
}

func (s *PadContext) GetParser() antlr.Parser { return s.parser }

func (s *PadContext) GetSource() IAccountContext { return s.source }

func (s *PadContext) SetSource(v IAccountContext) { s.source = v }

func (s *PadContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDATETIME, 0)
}

func (s *PadContext) PAD() antlr.TerminalNode {
	return s.GetToken(BeanCountParserPAD, 0)
}

func (s *PadContext) AllAccount() []IAccountContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAccountContext)(nil)).Elem())
	var tst = make([]IAccountContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAccountContext)
		}
	}

	return tst
}

func (s *PadContext) Account(i int) IAccountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccountContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAccountContext)
}

func (s *PadContext) TagsOrLinks() ITagsOrLinksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsOrLinksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagsOrLinksContext)
}

func (s *PadContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *PadContext) IndentedMetadata() IIndentedMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndentedMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndentedMetadataContext)
}

func (s *PadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPad(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Pad() (localctx IPadContext) {
	localctx = NewPadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BeanCountParserRULE_pad)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(BeanCountParserDATETIME)
	}
	{
		p.SetState(223)
		p.Match(BeanCountParserPAD)
	}
	{
		p.SetState(224)
		p.Account()
	}
	{
		p.SetState(225)

		var _x = p.Account()

		localctx.(*PadContext).source = _x
	}
	{
		p.SetState(226)
		p.TagsOrLinks()
	}
	{
		p.SetState(227)
		p.Eol()
	}
	{
		p.SetState(228)
		p.IndentedMetadata()
	}

	return localctx
}

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDATETIME, 0)
}

func (s *DocumentContext) DOCUMENT() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDOCUMENT, 0)
}

func (s *DocumentContext) Account() IAccountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccountContext)
}

func (s *DocumentContext) Filename() IFilenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilenameContext)
}

func (s *DocumentContext) TagsOrLinks() ITagsOrLinksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsOrLinksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagsOrLinksContext)
}

func (s *DocumentContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *DocumentContext) IndentedMetadata() IIndentedMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndentedMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndentedMetadataContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitDocument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BeanCountParserRULE_document)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(BeanCountParserDATETIME)
	}
	{
		p.SetState(231)
		p.Match(BeanCountParserDOCUMENT)
	}
	{
		p.SetState(232)
		p.Account()
	}
	{
		p.SetState(233)
		p.Filename()
	}
	{
		p.SetState(234)
		p.TagsOrLinks()
	}
	{
		p.SetState(235)
		p.Eol()
	}
	{
		p.SetState(236)
		p.IndentedMetadata()
	}

	return localctx
}

// INoteContext is an interface to support dynamic dispatch.
type INoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetComment returns the comment token.
	GetComment() antlr.Token

	// SetComment sets the comment token.
	SetComment(antlr.Token)

	// IsNoteContext differentiates from other interfaces.
	IsNoteContext()
}

type NoteContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	comment antlr.Token
}

func NewEmptyNoteContext() *NoteContext {
	var p = new(NoteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_note
	return p
}

func (*NoteContext) IsNoteContext() {}

func NewNoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoteContext {
	var p = new(NoteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_note

	return p
}

func (s *NoteContext) GetParser() antlr.Parser { return s.parser }

func (s *NoteContext) GetComment() antlr.Token { return s.comment }

func (s *NoteContext) SetComment(v antlr.Token) { s.comment = v }

func (s *NoteContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDATETIME, 0)
}

func (s *NoteContext) NOTE() antlr.TerminalNode {
	return s.GetToken(BeanCountParserNOTE, 0)
}

func (s *NoteContext) Account() IAccountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccountContext)
}

func (s *NoteContext) TagsOrLinks() ITagsOrLinksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsOrLinksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagsOrLinksContext)
}

func (s *NoteContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *NoteContext) IndentedMetadata() IIndentedMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndentedMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndentedMetadataContext)
}

func (s *NoteContext) STRING() antlr.TerminalNode {
	return s.GetToken(BeanCountParserSTRING, 0)
}

func (s *NoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitNote(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Note() (localctx INoteContext) {
	localctx = NewNoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BeanCountParserRULE_note)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(BeanCountParserDATETIME)
	}
	{
		p.SetState(239)
		p.Match(BeanCountParserNOTE)
	}
	{
		p.SetState(240)
		p.Account()
	}
	{
		p.SetState(241)

		var _m = p.Match(BeanCountParserSTRING)

		localctx.(*NoteContext).comment = _m
	}
	{
		p.SetState(242)
		p.TagsOrLinks()
	}
	{
		p.SetState(243)
		p.Eol()
	}
	{
		p.SetState(244)
		p.IndentedMetadata()
	}

	return localctx
}

// IEventContext is an interface to support dynamic dispatch.
type IEventContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEventType returns the eventType token.
	GetEventType() antlr.Token

	// GetDescription returns the description token.
	GetDescription() antlr.Token

	// SetEventType sets the eventType token.
	SetEventType(antlr.Token)

	// SetDescription sets the description token.
	SetDescription(antlr.Token)

	// IsEventContext differentiates from other interfaces.
	IsEventContext()
}

type EventContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	eventType   antlr.Token
	description antlr.Token
}

func NewEmptyEventContext() *EventContext {
	var p = new(EventContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_event
	return p
}

func (*EventContext) IsEventContext() {}

func NewEventContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EventContext {
	var p = new(EventContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_event

	return p
}

func (s *EventContext) GetParser() antlr.Parser { return s.parser }

func (s *EventContext) GetEventType() antlr.Token { return s.eventType }

func (s *EventContext) GetDescription() antlr.Token { return s.description }

func (s *EventContext) SetEventType(v antlr.Token) { s.eventType = v }

func (s *EventContext) SetDescription(v antlr.Token) { s.description = v }

func (s *EventContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDATETIME, 0)
}

func (s *EventContext) EVENT() antlr.TerminalNode {
	return s.GetToken(BeanCountParserEVENT, 0)
}

func (s *EventContext) TagsOrLinks() ITagsOrLinksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsOrLinksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagsOrLinksContext)
}

func (s *EventContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *EventContext) IndentedMetadata() IIndentedMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndentedMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndentedMetadataContext)
}

func (s *EventContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(BeanCountParserSTRING)
}

func (s *EventContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(BeanCountParserSTRING, i)
}

func (s *EventContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitEvent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Event() (localctx IEventContext) {
	localctx = NewEventContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BeanCountParserRULE_event)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(BeanCountParserDATETIME)
	}
	{
		p.SetState(247)
		p.Match(BeanCountParserEVENT)
	}
	{
		p.SetState(248)

		var _m = p.Match(BeanCountParserSTRING)

		localctx.(*EventContext).eventType = _m
	}
	{
		p.SetState(249)

		var _m = p.Match(BeanCountParserSTRING)

		localctx.(*EventContext).description = _m
	}
	{
		p.SetState(250)
		p.TagsOrLinks()
	}
	{
		p.SetState(251)
		p.Eol()
	}
	{
		p.SetState(252)
		p.IndentedMetadata()
	}

	return localctx
}

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetQstr returns the qstr token.
	GetQstr() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetQstr sets the qstr token.
	SetQstr(antlr.Token)

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	qstr   antlr.Token
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) GetName() antlr.Token { return s.name }

func (s *QueryContext) GetQstr() antlr.Token { return s.qstr }

func (s *QueryContext) SetName(v antlr.Token) { s.name = v }

func (s *QueryContext) SetQstr(v antlr.Token) { s.qstr = v }

func (s *QueryContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDATETIME, 0)
}

func (s *QueryContext) QUERY() antlr.TerminalNode {
	return s.GetToken(BeanCountParserQUERY, 0)
}

func (s *QueryContext) TagsOrLinks() ITagsOrLinksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsOrLinksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagsOrLinksContext)
}

func (s *QueryContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *QueryContext) IndentedMetadata() IIndentedMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndentedMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndentedMetadataContext)
}

func (s *QueryContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(BeanCountParserSTRING)
}

func (s *QueryContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(BeanCountParserSTRING, i)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BeanCountParserRULE_query)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Match(BeanCountParserDATETIME)
	}
	{
		p.SetState(255)
		p.Match(BeanCountParserQUERY)
	}
	{
		p.SetState(256)

		var _m = p.Match(BeanCountParserSTRING)

		localctx.(*QueryContext).name = _m
	}
	{
		p.SetState(257)

		var _m = p.Match(BeanCountParserSTRING)

		localctx.(*QueryContext).qstr = _m
	}
	{
		p.SetState(258)
		p.TagsOrLinks()
	}
	{
		p.SetState(259)
		p.Eol()
	}
	{
		p.SetState(260)
		p.IndentedMetadata()
	}

	return localctx
}

// ICustomContext is an interface to support dynamic dispatch.
type ICustomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCustomType returns the customType token.
	GetCustomType() antlr.Token

	// SetCustomType sets the customType token.
	SetCustomType(antlr.Token)

	// IsCustomContext differentiates from other interfaces.
	IsCustomContext()
}

type CustomContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	customType antlr.Token
}

func NewEmptyCustomContext() *CustomContext {
	var p = new(CustomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_custom
	return p
}

func (*CustomContext) IsCustomContext() {}

func NewCustomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CustomContext {
	var p = new(CustomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_custom

	return p
}

func (s *CustomContext) GetParser() antlr.Parser { return s.parser }

func (s *CustomContext) GetCustomType() antlr.Token { return s.customType }

func (s *CustomContext) SetCustomType(v antlr.Token) { s.customType = v }

func (s *CustomContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDATETIME, 0)
}

func (s *CustomContext) CUSTOM() antlr.TerminalNode {
	return s.GetToken(BeanCountParserCUSTOM, 0)
}

func (s *CustomContext) MetaValueList() IMetaValueListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetaValueListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMetaValueListContext)
}

func (s *CustomContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *CustomContext) IndentedMetadata() IIndentedMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndentedMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndentedMetadataContext)
}

func (s *CustomContext) STRING() antlr.TerminalNode {
	return s.GetToken(BeanCountParserSTRING, 0)
}

func (s *CustomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CustomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CustomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitCustom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Custom() (localctx ICustomContext) {
	localctx = NewCustomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BeanCountParserRULE_custom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.Match(BeanCountParserDATETIME)
	}
	{
		p.SetState(263)
		p.Match(BeanCountParserCUSTOM)
	}
	{
		p.SetState(264)

		var _m = p.Match(BeanCountParserSTRING)

		localctx.(*CustomContext).customType = _m
	}
	{
		p.SetState(265)
		p.MetaValueList()
	}
	{
		p.SetState(266)
		p.Eol()
	}
	{
		p.SetState(267)
		p.IndentedMetadata()
	}

	return localctx
}

// IIndentedMetadataContext is an interface to support dynamic dispatch.
type IIndentedMetadataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndentedMetadataContext differentiates from other interfaces.
	IsIndentedMetadataContext()
}

type IndentedMetadataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndentedMetadataContext() *IndentedMetadataContext {
	var p = new(IndentedMetadataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_indentedMetadata
	return p
}

func (*IndentedMetadataContext) IsIndentedMetadataContext() {}

func NewIndentedMetadataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndentedMetadataContext {
	var p = new(IndentedMetadataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_indentedMetadata

	return p
}

func (s *IndentedMetadataContext) GetParser() antlr.Parser { return s.parser }

func (s *IndentedMetadataContext) INDENT() antlr.TerminalNode {
	return s.GetToken(BeanCountParserINDENT, 0)
}

func (s *IndentedMetadataContext) Metadata() IMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMetadataContext)
}

func (s *IndentedMetadataContext) DEDENT() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDEDENT, 0)
}

func (s *IndentedMetadataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndentedMetadataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndentedMetadataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitIndentedMetadata(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) IndentedMetadata() (localctx IIndentedMetadataContext) {
	localctx = NewIndentedMetadataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BeanCountParserRULE_indentedMetadata)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(269)
			p.Match(BeanCountParserINDENT)
		}
		{
			p.SetState(270)
			p.Metadata()
		}
		{
			p.SetState(271)
			p.Match(BeanCountParserDEDENT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)

	}

	return localctx
}

// IMetadataContext is an interface to support dynamic dispatch.
type IMetadataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMetadataContext differentiates from other interfaces.
	IsMetadataContext()
}

type MetadataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetadataContext() *MetadataContext {
	var p = new(MetadataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_metadata
	return p
}

func (*MetadataContext) IsMetadataContext() {}

func NewMetadataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetadataContext {
	var p = new(MetadataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_metadata

	return p
}

func (s *MetadataContext) GetParser() antlr.Parser { return s.parser }

func (s *MetadataContext) AllMetadataLine() []IMetadataLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMetadataLineContext)(nil)).Elem())
	var tst = make([]IMetadataLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMetadataLineContext)
		}
	}

	return tst
}

func (s *MetadataContext) MetadataLine(i int) IMetadataLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetadataLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMetadataLineContext)
}

func (s *MetadataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetadataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetadataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitMetadata(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Metadata() (localctx IMetadataContext) {
	localctx = NewMetadataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, BeanCountParserRULE_metadata)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(276)
				p.MetadataLine()
			}

		}
		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IMetadataLineContext is an interface to support dynamic dispatch.
type IMetadataLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMetadataLineContext differentiates from other interfaces.
	IsMetadataLineContext()
}

type MetadataLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetadataLineContext() *MetadataLineContext {
	var p = new(MetadataLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_metadataLine
	return p
}

func (*MetadataLineContext) IsMetadataLineContext() {}

func NewMetadataLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetadataLineContext {
	var p = new(MetadataLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_metadataLine

	return p
}

func (s *MetadataLineContext) GetParser() antlr.Parser { return s.parser }

func (s *MetadataLineContext) KEY() antlr.TerminalNode {
	return s.GetToken(BeanCountParserKEY, 0)
}

func (s *MetadataLineContext) COLON() antlr.TerminalNode {
	return s.GetToken(BeanCountParserCOLON, 0)
}

func (s *MetadataLineContext) MetaValue() IMetaValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetaValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMetaValueContext)
}

func (s *MetadataLineContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *MetadataLineContext) TAG() antlr.TerminalNode {
	return s.GetToken(BeanCountParserTAG, 0)
}

func (s *MetadataLineContext) LINK() antlr.TerminalNode {
	return s.GetToken(BeanCountParserLINK, 0)
}

func (s *MetadataLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetadataLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetadataLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitMetadataLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) MetadataLine() (localctx IMetadataLineContext) {
	localctx = NewMetadataLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BeanCountParserRULE_metadataLine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(291)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeanCountParserKEY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(282)
			p.Match(BeanCountParserKEY)
		}
		{
			p.SetState(283)
			p.Match(BeanCountParserCOLON)
		}
		{
			p.SetState(284)
			p.MetaValue()
		}
		{
			p.SetState(285)
			p.Eol()
		}

	case BeanCountParserTAG:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(287)
			p.Match(BeanCountParserTAG)
		}
		{
			p.SetState(288)
			p.Eol()
		}

	case BeanCountParserLINK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(289)
			p.Match(BeanCountParserLINK)
		}
		{
			p.SetState(290)
			p.Eol()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMetaValueContext is an interface to support dynamic dispatch.
type IMetaValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMetaValueContext differentiates from other interfaces.
	IsMetaValueContext()
}

type MetaValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetaValueContext() *MetaValueContext {
	var p = new(MetaValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_metaValue
	return p
}

func (*MetaValueContext) IsMetaValueContext() {}

func NewMetaValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetaValueContext {
	var p = new(MetaValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_metaValue

	return p
}

func (s *MetaValueContext) GetParser() antlr.Parser { return s.parser }

func (s *MetaValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(BeanCountParserSTRING, 0)
}

func (s *MetaValueContext) CURRENCY() antlr.TerminalNode {
	return s.GetToken(BeanCountParserCURRENCY, 0)
}

func (s *MetaValueContext) Account() IAccountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccountContext)
}

func (s *MetaValueContext) TAG() antlr.TerminalNode {
	return s.GetToken(BeanCountParserTAG, 0)
}

func (s *MetaValueContext) LINK() antlr.TerminalNode {
	return s.GetToken(BeanCountParserLINK, 0)
}

func (s *MetaValueContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDATETIME, 0)
}

func (s *MetaValueContext) BOOL() antlr.TerminalNode {
	return s.GetToken(BeanCountParserBOOL, 0)
}

func (s *MetaValueContext) NONE() antlr.TerminalNode {
	return s.GetToken(BeanCountParserNONE, 0)
}

func (s *MetaValueContext) NumberExpr() INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *MetaValueContext) Amount() IAmountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAmountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAmountContext)
}

func (s *MetaValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetaValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetaValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitMetaValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) MetaValue() (localctx IMetaValueContext) {
	localctx = NewMetaValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BeanCountParserRULE_metaValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.Match(BeanCountParserSTRING)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(294)
			p.Match(BeanCountParserCURRENCY)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(295)
			p.Account()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(296)
			p.Match(BeanCountParserTAG)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(297)
			p.Match(BeanCountParserLINK)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(298)
			p.Match(BeanCountParserDATETIME)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(299)
			p.Match(BeanCountParserBOOL)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(300)
			p.Match(BeanCountParserNONE)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(301)
			p.numberExpr(0)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(302)
			p.Amount()
		}

	}

	return localctx
}

// IMetaValueListContext is an interface to support dynamic dispatch.
type IMetaValueListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMetaValueListContext differentiates from other interfaces.
	IsMetaValueListContext()
}

type MetaValueListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetaValueListContext() *MetaValueListContext {
	var p = new(MetaValueListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_metaValueList
	return p
}

func (*MetaValueListContext) IsMetaValueListContext() {}

func NewMetaValueListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetaValueListContext {
	var p = new(MetaValueListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_metaValueList

	return p
}

func (s *MetaValueListContext) GetParser() antlr.Parser { return s.parser }

func (s *MetaValueListContext) AllMetaValue() []IMetaValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMetaValueContext)(nil)).Elem())
	var tst = make([]IMetaValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMetaValueContext)
		}
	}

	return tst
}

func (s *MetaValueListContext) MetaValue(i int) IMetaValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetaValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMetaValueContext)
}

func (s *MetaValueListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetaValueListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetaValueListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitMetaValueList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) MetaValueList() (localctx IMetaValueListContext) {
	localctx = NewMetaValueListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BeanCountParserRULE_metaValueList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeanCountParserBOOL)|(1<<BeanCountParserNONE)|(1<<BeanCountParserCURRENCY)|(1<<BeanCountParserLPAREN)|(1<<BeanCountParserNUMBER))) != 0) || (((_la-50)&-(0x1f+1)) == 0 && ((1<<uint((_la-50)))&((1<<(BeanCountParserDATETIME-50))|(1<<(BeanCountParserACCOUNT-50))|(1<<(BeanCountParserLINK-50))|(1<<(BeanCountParserTAG-50))|(1<<(BeanCountParserSTRING-50)))) != 0) {
		{
			p.SetState(305)
			p.MetaValue()
		}

		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITagsOrLinksContext is an interface to support dynamic dispatch.
type ITagsOrLinksContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagsOrLinksContext differentiates from other interfaces.
	IsTagsOrLinksContext()
}

type TagsOrLinksContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagsOrLinksContext() *TagsOrLinksContext {
	var p = new(TagsOrLinksContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_tagsOrLinks
	return p
}

func (*TagsOrLinksContext) IsTagsOrLinksContext() {}

func NewTagsOrLinksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagsOrLinksContext {
	var p = new(TagsOrLinksContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_tagsOrLinks

	return p
}

func (s *TagsOrLinksContext) GetParser() antlr.Parser { return s.parser }

func (s *TagsOrLinksContext) AllTAG() []antlr.TerminalNode {
	return s.GetTokens(BeanCountParserTAG)
}

func (s *TagsOrLinksContext) TAG(i int) antlr.TerminalNode {
	return s.GetToken(BeanCountParserTAG, i)
}

func (s *TagsOrLinksContext) AllLINK() []antlr.TerminalNode {
	return s.GetTokens(BeanCountParserLINK)
}

func (s *TagsOrLinksContext) LINK(i int) antlr.TerminalNode {
	return s.GetToken(BeanCountParserLINK, i)
}

func (s *TagsOrLinksContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagsOrLinksContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagsOrLinksContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitTagsOrLinks(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) TagsOrLinks() (localctx ITagsOrLinksContext) {
	localctx = NewTagsOrLinksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BeanCountParserRULE_tagsOrLinks)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BeanCountParserLINK || _la == BeanCountParserTAG {
		{
			p.SetState(311)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BeanCountParserLINK || _la == BeanCountParserTAG) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPayeeNarrationContext is an interface to support dynamic dispatch.
type IPayeeNarrationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPayee returns the payee token.
	GetPayee() antlr.Token

	// GetNarration returns the narration token.
	GetNarration() antlr.Token

	// SetPayee sets the payee token.
	SetPayee(antlr.Token)

	// SetNarration sets the narration token.
	SetNarration(antlr.Token)

	// IsPayeeNarrationContext differentiates from other interfaces.
	IsPayeeNarrationContext()
}

type PayeeNarrationContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	payee     antlr.Token
	narration antlr.Token
}

func NewEmptyPayeeNarrationContext() *PayeeNarrationContext {
	var p = new(PayeeNarrationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_payeeNarration
	return p
}

func (*PayeeNarrationContext) IsPayeeNarrationContext() {}

func NewPayeeNarrationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PayeeNarrationContext {
	var p = new(PayeeNarrationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_payeeNarration

	return p
}

func (s *PayeeNarrationContext) GetParser() antlr.Parser { return s.parser }

func (s *PayeeNarrationContext) GetPayee() antlr.Token { return s.payee }

func (s *PayeeNarrationContext) GetNarration() antlr.Token { return s.narration }

func (s *PayeeNarrationContext) SetPayee(v antlr.Token) { s.payee = v }

func (s *PayeeNarrationContext) SetNarration(v antlr.Token) { s.narration = v }

func (s *PayeeNarrationContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(BeanCountParserSTRING)
}

func (s *PayeeNarrationContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(BeanCountParserSTRING, i)
}

func (s *PayeeNarrationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PayeeNarrationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PayeeNarrationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPayeeNarration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) PayeeNarration() (localctx IPayeeNarrationContext) {
	localctx = NewPayeeNarrationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, BeanCountParserRULE_payeeNarration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(317)

			var _m = p.Match(BeanCountParserSTRING)

			localctx.(*PayeeNarrationContext).payee = _m
		}
		{
			p.SetState(318)

			var _m = p.Match(BeanCountParserSTRING)

			localctx.(*PayeeNarrationContext).narration = _m
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(319)

			var _m = p.Match(BeanCountParserSTRING)

			localctx.(*PayeeNarrationContext).narration = _m
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)

	}

	return localctx
}

// IAccountContext is an interface to support dynamic dispatch.
type IAccountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccountContext differentiates from other interfaces.
	IsAccountContext()
}

type AccountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccountContext() *AccountContext {
	var p = new(AccountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_account
	return p
}

func (*AccountContext) IsAccountContext() {}

func NewAccountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccountContext {
	var p = new(AccountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_account

	return p
}

func (s *AccountContext) GetParser() antlr.Parser { return s.parser }

func (s *AccountContext) ACCOUNT() antlr.TerminalNode {
	return s.GetToken(BeanCountParserACCOUNT, 0)
}

func (s *AccountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccountContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitAccount(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Account() (localctx IAccountContext) {
	localctx = NewAccountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, BeanCountParserRULE_account)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
		p.Match(BeanCountParserACCOUNT)
	}

	return localctx
}

// IFilenameContext is an interface to support dynamic dispatch.
type IFilenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilenameContext differentiates from other interfaces.
	IsFilenameContext()
}

type FilenameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilenameContext() *FilenameContext {
	var p = new(FilenameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_filename
	return p
}

func (*FilenameContext) IsFilenameContext() {}

func NewFilenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilenameContext {
	var p = new(FilenameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_filename

	return p
}

func (s *FilenameContext) GetParser() antlr.Parser { return s.parser }

func (s *FilenameContext) STRING() antlr.TerminalNode {
	return s.GetToken(BeanCountParserSTRING, 0)
}

func (s *FilenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilenameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitFilename(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Filename() (localctx IFilenameContext) {
	localctx = NewFilenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, BeanCountParserRULE_filename)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		p.Match(BeanCountParserSTRING)
	}

	return localctx
}

// ICurrencyListContext is an interface to support dynamic dispatch.
type ICurrencyListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCurrencyListContext differentiates from other interfaces.
	IsCurrencyListContext()
}

type CurrencyListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurrencyListContext() *CurrencyListContext {
	var p = new(CurrencyListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_currencyList
	return p
}

func (*CurrencyListContext) IsCurrencyListContext() {}

func NewCurrencyListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CurrencyListContext {
	var p = new(CurrencyListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_currencyList

	return p
}

func (s *CurrencyListContext) GetParser() antlr.Parser { return s.parser }

func (s *CurrencyListContext) AllCURRENCY() []antlr.TerminalNode {
	return s.GetTokens(BeanCountParserCURRENCY)
}

func (s *CurrencyListContext) CURRENCY(i int) antlr.TerminalNode {
	return s.GetToken(BeanCountParserCURRENCY, i)
}

func (s *CurrencyListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BeanCountParserCOMMA)
}

func (s *CurrencyListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BeanCountParserCOMMA, i)
}

func (s *CurrencyListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurrencyListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CurrencyListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitCurrencyList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) CurrencyList() (localctx ICurrencyListContext) {
	localctx = NewCurrencyListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, BeanCountParserRULE_currencyList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BeanCountParserCURRENCY {
		{
			p.SetState(327)
			p.Match(BeanCountParserCURRENCY)
		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BeanCountParserCOMMA {
			{
				p.SetState(328)
				p.Match(BeanCountParserCOMMA)
			}

		}

		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INumberExprContext is an interface to support dynamic dispatch.
type INumberExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberExprContext differentiates from other interfaces.
	IsNumberExprContext()
}

type NumberExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberExprContext() *NumberExprContext {
	var p = new(NumberExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_numberExpr
	return p
}

func (*NumberExprContext) IsNumberExprContext() {}

func NewNumberExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberExprContext {
	var p = new(NumberExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_numberExpr

	return p
}

func (s *NumberExprContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberExprContext) CopyFrom(ctx *NumberExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *NumberExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UniaryOpContext struct {
	*NumberExprContext
}

func NewUniaryOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UniaryOpContext {
	var p = new(UniaryOpContext)

	p.NumberExprContext = NewEmptyNumberExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberExprContext))

	return p
}

func (s *UniaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UniaryOpContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BeanCountParserLPAREN, 0)
}

func (s *UniaryOpContext) NumberExpr() INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *UniaryOpContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BeanCountParserRPAREN, 0)
}

func (s *UniaryOpContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(BeanCountParserNUMBER, 0)
}

func (s *UniaryOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitUniaryOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryOpContext struct {
	*NumberExprContext
}

func NewBinaryOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryOpContext {
	var p = new(BinaryOpContext)

	p.NumberExprContext = NewEmptyNumberExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberExprContext))

	return p
}

func (s *BinaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOpContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *BinaryOpContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *BinaryOpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(BeanCountParserPLUS, 0)
}

func (s *BinaryOpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(BeanCountParserMINUS, 0)
}

func (s *BinaryOpContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(BeanCountParserASTERISK, 0)
}

func (s *BinaryOpContext) SLASH() antlr.TerminalNode {
	return s.GetToken(BeanCountParserSLASH, 0)
}

func (s *BinaryOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitBinaryOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) NumberExpr() (localctx INumberExprContext) {
	return p.numberExpr(0)
}

func (p *BeanCountParser) numberExpr(_p int) (localctx INumberExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewNumberExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx INumberExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 68
	p.EnterRecursionRule(localctx, 68, BeanCountParserRULE_numberExpr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(342)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeanCountParserLPAREN:
		localctx = NewUniaryOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(337)
			p.Match(BeanCountParserLPAREN)
		}
		{
			p.SetState(338)
			p.numberExpr(0)
		}
		{
			p.SetState(339)
			p.Match(BeanCountParserRPAREN)
		}

	case BeanCountParserNUMBER:
		localctx = NewUniaryOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(341)
			p.Match(BeanCountParserNUMBER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(350)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryOpContext(p, NewNumberExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeanCountParserRULE_numberExpr)
				p.SetState(344)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(345)
					_la = p.GetTokenStream().LA(1)

					if !(_la == BeanCountParserPLUS || _la == BeanCountParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(346)
					p.numberExpr(5)
				}

			case 2:
				localctx = NewBinaryOpContext(p, NewNumberExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, BeanCountParserRULE_numberExpr)
				p.SetState(347)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(348)
					_la = p.GetTokenStream().LA(1)

					if !(_la == BeanCountParserASTERISK || _la == BeanCountParserSLASH) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(349)
					p.numberExpr(4)
				}

			}

		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// ITxnContext is an interface to support dynamic dispatch.
type ITxnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTxnContext differentiates from other interfaces.
	IsTxnContext()
}

type TxnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTxnContext() *TxnContext {
	var p = new(TxnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_txn
	return p
}

func (*TxnContext) IsTxnContext() {}

func NewTxnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TxnContext {
	var p = new(TxnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_txn

	return p
}

func (s *TxnContext) GetParser() antlr.Parser { return s.parser }

func (s *TxnContext) TXN() antlr.TerminalNode {
	return s.GetToken(BeanCountParserTXN, 0)
}

func (s *TxnContext) FLAG() antlr.TerminalNode {
	return s.GetToken(BeanCountParserFLAG, 0)
}

func (s *TxnContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(BeanCountParserASTERISK, 0)
}

func (s *TxnContext) HASH() antlr.TerminalNode {
	return s.GetToken(BeanCountParserHASH, 0)
}

func (s *TxnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TxnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TxnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitTxn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Txn() (localctx ITxnContext) {
	localctx = NewTxnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, BeanCountParserRULE_txn)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-18)&-(0x1f+1)) == 0 && ((1<<uint((_la-18)))&((1<<(BeanCountParserASTERISK-18))|(1<<(BeanCountParserFLAG-18))|(1<<(BeanCountParserTXN-18))|(1<<(BeanCountParserHASH-18)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOptflagContext is an interface to support dynamic dispatch.
type IOptflagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptflagContext differentiates from other interfaces.
	IsOptflagContext()
}

type OptflagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptflagContext() *OptflagContext {
	var p = new(OptflagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_optflag
	return p
}

func (*OptflagContext) IsOptflagContext() {}

func NewOptflagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptflagContext {
	var p = new(OptflagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_optflag

	return p
}

func (s *OptflagContext) GetParser() antlr.Parser { return s.parser }

func (s *OptflagContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(BeanCountParserASTERISK, 0)
}

func (s *OptflagContext) HASH() antlr.TerminalNode {
	return s.GetToken(BeanCountParserHASH, 0)
}

func (s *OptflagContext) FLAG() antlr.TerminalNode {
	return s.GetToken(BeanCountParserFLAG, 0)
}

func (s *OptflagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptflagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptflagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitOptflag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Optflag() (localctx IOptflagContext) {
	localctx = NewOptflagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, BeanCountParserRULE_optflag)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-18)&-(0x1f+1)) == 0 && ((1<<uint((_la-18)))&((1<<(BeanCountParserASTERISK-18))|(1<<(BeanCountParserFLAG-18))|(1<<(BeanCountParserHASH-18)))) != 0 {
		{
			p.SetState(357)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-18)&-(0x1f+1)) == 0 && ((1<<uint((_la-18)))&((1<<(BeanCountParserASTERISK-18))|(1<<(BeanCountParserFLAG-18))|(1<<(BeanCountParserHASH-18)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// IBookingContext is an interface to support dynamic dispatch.
type IBookingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBookingContext differentiates from other interfaces.
	IsBookingContext()
}

type BookingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBookingContext() *BookingContext {
	var p = new(BookingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_booking
	return p
}

func (*BookingContext) IsBookingContext() {}

func NewBookingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BookingContext {
	var p = new(BookingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_booking

	return p
}

func (s *BookingContext) GetParser() antlr.Parser { return s.parser }

func (s *BookingContext) STRING() antlr.TerminalNode {
	return s.GetToken(BeanCountParserSTRING, 0)
}

func (s *BookingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BookingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BookingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitBooking(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Booking() (localctx IBookingContext) {
	localctx = NewBookingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, BeanCountParserRULE_booking)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BeanCountParserSTRING {
		{
			p.SetState(360)
			p.Match(BeanCountParserSTRING)
		}

	}

	return localctx
}

// IPostingListContext is an interface to support dynamic dispatch.
type IPostingListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostingListContext differentiates from other interfaces.
	IsPostingListContext()
}

type PostingListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostingListContext() *PostingListContext {
	var p = new(PostingListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_postingList
	return p
}

func (*PostingListContext) IsPostingListContext() {}

func NewPostingListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostingListContext {
	var p = new(PostingListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_postingList

	return p
}

func (s *PostingListContext) GetParser() antlr.Parser { return s.parser }

func (s *PostingListContext) PostingList() IPostingListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostingListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostingListContext)
}

func (s *PostingListContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *PostingListContext) PostingAndMetadata() IPostingAndMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostingAndMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostingAndMetadataContext)
}

func (s *PostingListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostingListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostingListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPostingList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) PostingList() (localctx IPostingListContext) {
	return p.postingList(0)
}

func (p *BeanCountParser) postingList(_p int) (localctx IPostingListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPostingListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPostingListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 76
	p.EnterRecursionRule(localctx, 76, BeanCountParserRULE_postingList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(370)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(368)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPostingListContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, BeanCountParserRULE_postingList)
				p.SetState(364)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(365)
					p.Eol()
				}

			case 2:
				localctx = NewPostingListContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, BeanCountParserRULE_postingList)
				p.SetState(366)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(367)
					p.PostingAndMetadata()
				}

			}

		}
		p.SetState(372)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IPostingAndMetadataContext is an interface to support dynamic dispatch.
type IPostingAndMetadataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostingAndMetadataContext differentiates from other interfaces.
	IsPostingAndMetadataContext()
}

type PostingAndMetadataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostingAndMetadataContext() *PostingAndMetadataContext {
	var p = new(PostingAndMetadataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_postingAndMetadata
	return p
}

func (*PostingAndMetadataContext) IsPostingAndMetadataContext() {}

func NewPostingAndMetadataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostingAndMetadataContext {
	var p = new(PostingAndMetadataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_postingAndMetadata

	return p
}

func (s *PostingAndMetadataContext) GetParser() antlr.Parser { return s.parser }

func (s *PostingAndMetadataContext) Posting() IPostingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostingContext)
}

func (s *PostingAndMetadataContext) IndentedMetadata() IIndentedMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndentedMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndentedMetadataContext)
}

func (s *PostingAndMetadataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostingAndMetadataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostingAndMetadataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPostingAndMetadata(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) PostingAndMetadata() (localctx IPostingAndMetadataContext) {
	localctx = NewPostingAndMetadataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, BeanCountParserRULE_postingAndMetadata)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)
		p.Posting()
	}
	{
		p.SetState(374)
		p.IndentedMetadata()
	}

	return localctx
}

// IPostingContext is an interface to support dynamic dispatch.
type IPostingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostingContext differentiates from other interfaces.
	IsPostingContext()
}

type PostingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostingContext() *PostingContext {
	var p = new(PostingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_posting
	return p
}

func (*PostingContext) IsPostingContext() {}

func NewPostingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostingContext {
	var p = new(PostingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_posting

	return p
}

func (s *PostingContext) GetParser() antlr.Parser { return s.parser }

func (s *PostingContext) CopyFrom(ctx *PostingContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PostingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PostingAccountContext struct {
	*PostingContext
}

func NewPostingAccountContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostingAccountContext {
	var p = new(PostingAccountContext)

	p.PostingContext = NewEmptyPostingContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PostingContext))

	return p
}

func (s *PostingAccountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostingAccountContext) Optflag() IOptflagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptflagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptflagContext)
}

func (s *PostingAccountContext) Account() IAccountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccountContext)
}

func (s *PostingAccountContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *PostingAccountContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPostingAccount(s)

	default:
		return t.VisitChildren(s)
	}
}

type PostingAmountContext struct {
	*PostingContext
}

func NewPostingAmountContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostingAmountContext {
	var p = new(PostingAmountContext)

	p.PostingContext = NewEmptyPostingContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PostingContext))

	return p
}

func (s *PostingAmountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostingAmountContext) Optflag() IOptflagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptflagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptflagContext)
}

func (s *PostingAmountContext) Account() IAccountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccountContext)
}

func (s *PostingAmountContext) PartialAmount() IPartialAmountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPartialAmountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPartialAmountContext)
}

func (s *PostingAmountContext) CostSpec() ICostSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICostSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICostSpecContext)
}

func (s *PostingAmountContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *PostingAmountContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPostingAmount(s)

	default:
		return t.VisitChildren(s)
	}
}

type PostingPriceContext struct {
	*PostingContext
}

func NewPostingPriceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostingPriceContext {
	var p = new(PostingPriceContext)

	p.PostingContext = NewEmptyPostingContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PostingContext))

	return p
}

func (s *PostingPriceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostingPriceContext) Optflag() IOptflagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptflagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptflagContext)
}

func (s *PostingPriceContext) Account() IAccountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccountContext)
}

func (s *PostingPriceContext) PartialAmount() IPartialAmountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPartialAmountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPartialAmountContext)
}

func (s *PostingPriceContext) CostSpec() ICostSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICostSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICostSpecContext)
}

func (s *PostingPriceContext) PriceAnnotation() IPriceAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPriceAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPriceAnnotationContext)
}

func (s *PostingPriceContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *PostingPriceContext) AT() antlr.TerminalNode {
	return s.GetToken(BeanCountParserAT, 0)
}

func (s *PostingPriceContext) ATAT() antlr.TerminalNode {
	return s.GetToken(BeanCountParserATAT, 0)
}

func (s *PostingPriceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPostingPrice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Posting() (localctx IPostingContext) {
	localctx = NewPostingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, BeanCountParserRULE_posting)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPostingPriceContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(376)
			p.Optflag()
		}
		{
			p.SetState(377)
			p.Account()
		}
		{
			p.SetState(378)
			p.PartialAmount()
		}
		{
			p.SetState(379)
			p.CostSpec()
		}
		{
			p.SetState(380)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BeanCountParserATAT || _la == BeanCountParserAT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(381)
			p.PriceAnnotation()
		}
		{
			p.SetState(382)
			p.Eol()
		}

	case 2:
		localctx = NewPostingAmountContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(384)
			p.Optflag()
		}
		{
			p.SetState(385)
			p.Account()
		}
		{
			p.SetState(386)
			p.PartialAmount()
		}
		{
			p.SetState(387)
			p.CostSpec()
		}
		{
			p.SetState(388)
			p.Eol()
		}

	case 3:
		localctx = NewPostingAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(390)
			p.Optflag()
		}
		{
			p.SetState(391)
			p.Account()
		}
		{
			p.SetState(392)
			p.Eol()
		}

	}

	return localctx
}

// IPriceAnnotationContext is an interface to support dynamic dispatch.
type IPriceAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPriceAnnotationContext differentiates from other interfaces.
	IsPriceAnnotationContext()
}

type PriceAnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPriceAnnotationContext() *PriceAnnotationContext {
	var p = new(PriceAnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_priceAnnotation
	return p
}

func (*PriceAnnotationContext) IsPriceAnnotationContext() {}

func NewPriceAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PriceAnnotationContext {
	var p = new(PriceAnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_priceAnnotation

	return p
}

func (s *PriceAnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *PriceAnnotationContext) PartialAmount() IPartialAmountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPartialAmountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPartialAmountContext)
}

func (s *PriceAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PriceAnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PriceAnnotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPriceAnnotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) PriceAnnotation() (localctx IPriceAnnotationContext) {
	localctx = NewPriceAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, BeanCountParserRULE_priceAnnotation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(396)
		p.PartialAmount()
	}

	return localctx
}

// IAmountContext is an interface to support dynamic dispatch.
type IAmountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAmountContext differentiates from other interfaces.
	IsAmountContext()
}

type AmountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAmountContext() *AmountContext {
	var p = new(AmountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_amount
	return p
}

func (*AmountContext) IsAmountContext() {}

func NewAmountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AmountContext {
	var p = new(AmountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_amount

	return p
}

func (s *AmountContext) GetParser() antlr.Parser { return s.parser }

func (s *AmountContext) NumberExpr() INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *AmountContext) CURRENCY() antlr.TerminalNode {
	return s.GetToken(BeanCountParserCURRENCY, 0)
}

func (s *AmountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AmountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AmountContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitAmount(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) Amount() (localctx IAmountContext) {
	localctx = NewAmountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, BeanCountParserRULE_amount)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(398)
		p.numberExpr(0)
	}
	{
		p.SetState(399)
		p.Match(BeanCountParserCURRENCY)
	}

	return localctx
}

// IAmountToleranceContext is an interface to support dynamic dispatch.
type IAmountToleranceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAmountToleranceContext differentiates from other interfaces.
	IsAmountToleranceContext()
}

type AmountToleranceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAmountToleranceContext() *AmountToleranceContext {
	var p = new(AmountToleranceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_amountTolerance
	return p
}

func (*AmountToleranceContext) IsAmountToleranceContext() {}

func NewAmountToleranceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AmountToleranceContext {
	var p = new(AmountToleranceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_amountTolerance

	return p
}

func (s *AmountToleranceContext) GetParser() antlr.Parser { return s.parser }

func (s *AmountToleranceContext) Amount() IAmountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAmountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAmountContext)
}

func (s *AmountToleranceContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *AmountToleranceContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *AmountToleranceContext) TILDE() antlr.TerminalNode {
	return s.GetToken(BeanCountParserTILDE, 0)
}

func (s *AmountToleranceContext) CURRENCY() antlr.TerminalNode {
	return s.GetToken(BeanCountParserCURRENCY, 0)
}

func (s *AmountToleranceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AmountToleranceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AmountToleranceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitAmountTolerance(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) AmountTolerance() (localctx IAmountToleranceContext) {
	localctx = NewAmountToleranceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, BeanCountParserRULE_amountTolerance)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(401)
			p.Amount()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(402)
			p.numberExpr(0)
		}
		{
			p.SetState(403)
			p.Match(BeanCountParserTILDE)
		}
		{
			p.SetState(404)
			p.numberExpr(0)
		}
		{
			p.SetState(405)
			p.Match(BeanCountParserCURRENCY)
		}

	}

	return localctx
}

// IPartialAmountContext is an interface to support dynamic dispatch.
type IPartialAmountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPartialAmountContext differentiates from other interfaces.
	IsPartialAmountContext()
}

type PartialAmountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPartialAmountContext() *PartialAmountContext {
	var p = new(PartialAmountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_partialAmount
	return p
}

func (*PartialAmountContext) IsPartialAmountContext() {}

func NewPartialAmountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PartialAmountContext {
	var p = new(PartialAmountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_partialAmount

	return p
}

func (s *PartialAmountContext) GetParser() antlr.Parser { return s.parser }

func (s *PartialAmountContext) MaybeNumber() IMaybeNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMaybeNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMaybeNumberContext)
}

func (s *PartialAmountContext) MaybeCurrency() IMaybeCurrencyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMaybeCurrencyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMaybeCurrencyContext)
}

func (s *PartialAmountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PartialAmountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PartialAmountContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitPartialAmount(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) PartialAmount() (localctx IPartialAmountContext) {
	localctx = NewPartialAmountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, BeanCountParserRULE_partialAmount)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(409)
		p.MaybeNumber()
	}
	{
		p.SetState(410)
		p.MaybeCurrency()
	}

	return localctx
}

// IMaybeNumberContext is an interface to support dynamic dispatch.
type IMaybeNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMaybeNumberContext differentiates from other interfaces.
	IsMaybeNumberContext()
}

type MaybeNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMaybeNumberContext() *MaybeNumberContext {
	var p = new(MaybeNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_maybeNumber
	return p
}

func (*MaybeNumberContext) IsMaybeNumberContext() {}

func NewMaybeNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MaybeNumberContext {
	var p = new(MaybeNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_maybeNumber

	return p
}

func (s *MaybeNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *MaybeNumberContext) NumberExpr() INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *MaybeNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaybeNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MaybeNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitMaybeNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) MaybeNumber() (localctx IMaybeNumberContext) {
	localctx = NewMaybeNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, BeanCountParserRULE_maybeNumber)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BeanCountParserLPAREN || _la == BeanCountParserNUMBER {
		{
			p.SetState(412)
			p.numberExpr(0)
		}

	}

	return localctx
}

// IMaybeCurrencyContext is an interface to support dynamic dispatch.
type IMaybeCurrencyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMaybeCurrencyContext differentiates from other interfaces.
	IsMaybeCurrencyContext()
}

type MaybeCurrencyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMaybeCurrencyContext() *MaybeCurrencyContext {
	var p = new(MaybeCurrencyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_maybeCurrency
	return p
}

func (*MaybeCurrencyContext) IsMaybeCurrencyContext() {}

func NewMaybeCurrencyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MaybeCurrencyContext {
	var p = new(MaybeCurrencyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_maybeCurrency

	return p
}

func (s *MaybeCurrencyContext) GetParser() antlr.Parser { return s.parser }

func (s *MaybeCurrencyContext) CURRENCY() antlr.TerminalNode {
	return s.GetToken(BeanCountParserCURRENCY, 0)
}

func (s *MaybeCurrencyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaybeCurrencyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MaybeCurrencyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitMaybeCurrency(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) MaybeCurrency() (localctx IMaybeCurrencyContext) {
	localctx = NewMaybeCurrencyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, BeanCountParserRULE_maybeCurrency)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(416)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(415)
			p.Match(BeanCountParserCURRENCY)
		}

	}

	return localctx
}

// ICostSpecContext is an interface to support dynamic dispatch.
type ICostSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCostSpecContext differentiates from other interfaces.
	IsCostSpecContext()
}

type CostSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCostSpecContext() *CostSpecContext {
	var p = new(CostSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_costSpec
	return p
}

func (*CostSpecContext) IsCostSpecContext() {}

func NewCostSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CostSpecContext {
	var p = new(CostSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_costSpec

	return p
}

func (s *CostSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *CostSpecContext) LCURL() antlr.TerminalNode {
	return s.GetToken(BeanCountParserLCURL, 0)
}

func (s *CostSpecContext) CostCompList() ICostCompListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICostCompListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICostCompListContext)
}

func (s *CostSpecContext) RCURL() antlr.TerminalNode {
	return s.GetToken(BeanCountParserRCURL, 0)
}

func (s *CostSpecContext) LCURLCURL() antlr.TerminalNode {
	return s.GetToken(BeanCountParserLCURLCURL, 0)
}

func (s *CostSpecContext) RCURLCURL() antlr.TerminalNode {
	return s.GetToken(BeanCountParserRCURLCURL, 0)
}

func (s *CostSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CostSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CostSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitCostSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) CostSpec() (localctx ICostSpecContext) {
	localctx = NewCostSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, BeanCountParserRULE_costSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(427)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeanCountParserLCURL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(418)
			p.Match(BeanCountParserLCURL)
		}
		{
			p.SetState(419)
			p.costCompList(0)
		}
		{
			p.SetState(420)
			p.Match(BeanCountParserRCURL)
		}

	case BeanCountParserLCURLCURL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(422)
			p.Match(BeanCountParserLCURLCURL)
		}
		{
			p.SetState(423)
			p.costCompList(0)
		}
		{
			p.SetState(424)
			p.Match(BeanCountParserRCURLCURL)
		}

	case BeanCountParserNEWLINE, BeanCountParserATAT, BeanCountParserAT:
		p.EnterOuterAlt(localctx, 3)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICostCompListContext is an interface to support dynamic dispatch.
type ICostCompListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCostCompListContext differentiates from other interfaces.
	IsCostCompListContext()
}

type CostCompListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCostCompListContext() *CostCompListContext {
	var p = new(CostCompListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_costCompList
	return p
}

func (*CostCompListContext) IsCostCompListContext() {}

func NewCostCompListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CostCompListContext {
	var p = new(CostCompListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_costCompList

	return p
}

func (s *CostCompListContext) GetParser() antlr.Parser { return s.parser }

func (s *CostCompListContext) CostComp() ICostCompContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICostCompContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICostCompContext)
}

func (s *CostCompListContext) CostCompList() ICostCompListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICostCompListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICostCompListContext)
}

func (s *CostCompListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(BeanCountParserCOMMA, 0)
}

func (s *CostCompListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CostCompListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CostCompListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitCostCompList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) CostCompList() (localctx ICostCompListContext) {
	return p.costCompList(0)
}

func (p *BeanCountParser) costCompList(_p int) (localctx ICostCompListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCostCompListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICostCompListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 96
	p.EnterRecursionRule(localctx, 96, BeanCountParserRULE_costCompList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:

	case 2:
		{
			p.SetState(430)
			p.CostComp()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCostCompListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, BeanCountParserRULE_costCompList)
			p.SetState(433)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(434)
				p.Match(BeanCountParserCOMMA)
			}
			{
				p.SetState(435)
				p.CostComp()
			}

		}
		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}

	return localctx
}

// ICostCompContext is an interface to support dynamic dispatch.
type ICostCompContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCostCompContext differentiates from other interfaces.
	IsCostCompContext()
}

type CostCompContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCostCompContext() *CostCompContext {
	var p = new(CostCompContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_costComp
	return p
}

func (*CostCompContext) IsCostCompContext() {}

func NewCostCompContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CostCompContext {
	var p = new(CostCompContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_costComp

	return p
}

func (s *CostCompContext) GetParser() antlr.Parser { return s.parser }

func (s *CostCompContext) CompoundAmount() ICompoundAmountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundAmountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundAmountContext)
}

func (s *CostCompContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(BeanCountParserDATETIME, 0)
}

func (s *CostCompContext) STRING() antlr.TerminalNode {
	return s.GetToken(BeanCountParserSTRING, 0)
}

func (s *CostCompContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(BeanCountParserASTERISK, 0)
}

func (s *CostCompContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CostCompContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CostCompContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitCostComp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) CostComp() (localctx ICostCompContext) {
	localctx = NewCostCompContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, BeanCountParserRULE_costComp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(445)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeanCountParserCURRENCY, BeanCountParserLPAREN, BeanCountParserNUMBER, BeanCountParserHASH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(441)
			p.CompoundAmount()
		}

	case BeanCountParserDATETIME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(442)
			p.Match(BeanCountParserDATETIME)
		}

	case BeanCountParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(443)
			p.Match(BeanCountParserSTRING)
		}

	case BeanCountParserASTERISK:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(444)
			p.Match(BeanCountParserASTERISK)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICompoundAmountContext is an interface to support dynamic dispatch.
type ICompoundAmountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompoundAmountContext differentiates from other interfaces.
	IsCompoundAmountContext()
}

type CompoundAmountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompoundAmountContext() *CompoundAmountContext {
	var p = new(CompoundAmountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeanCountParserRULE_compoundAmount
	return p
}

func (*CompoundAmountContext) IsCompoundAmountContext() {}

func NewCompoundAmountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompoundAmountContext {
	var p = new(CompoundAmountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeanCountParserRULE_compoundAmount

	return p
}

func (s *CompoundAmountContext) GetParser() antlr.Parser { return s.parser }

func (s *CompoundAmountContext) AllMaybeNumber() []IMaybeNumberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMaybeNumberContext)(nil)).Elem())
	var tst = make([]IMaybeNumberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMaybeNumberContext)
		}
	}

	return tst
}

func (s *CompoundAmountContext) MaybeNumber(i int) IMaybeNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMaybeNumberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMaybeNumberContext)
}

func (s *CompoundAmountContext) CURRENCY() antlr.TerminalNode {
	return s.GetToken(BeanCountParserCURRENCY, 0)
}

func (s *CompoundAmountContext) NumberExpr() INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *CompoundAmountContext) MaybeCurrency() IMaybeCurrencyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMaybeCurrencyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMaybeCurrencyContext)
}

func (s *CompoundAmountContext) HASH() antlr.TerminalNode {
	return s.GetToken(BeanCountParserHASH, 0)
}

func (s *CompoundAmountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundAmountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompoundAmountContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeanCountParserVisitor:
		return t.VisitCompoundAmount(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeanCountParser) CompoundAmount() (localctx ICompoundAmountContext) {
	localctx = NewCompoundAmountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, BeanCountParserRULE_compoundAmount)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(447)
			p.MaybeNumber()
		}
		{
			p.SetState(448)
			p.Match(BeanCountParserCURRENCY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(450)
			p.numberExpr(0)
		}
		{
			p.SetState(451)
			p.MaybeCurrency()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(453)
			p.MaybeNumber()
		}
		{
			p.SetState(454)
			p.Match(BeanCountParserHASH)
		}
		{
			p.SetState(455)
			p.MaybeNumber()
		}
		{
			p.SetState(456)
			p.Match(BeanCountParserCURRENCY)
		}

	}

	return localctx
}

func (p *BeanCountParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *DeclarationsContext = nil
		if localctx != nil {
			t = localctx.(*DeclarationsContext)
		}
		return p.Declarations_Sempred(t, predIndex)

	case 34:
		var t *NumberExprContext = nil
		if localctx != nil {
			t = localctx.(*NumberExprContext)
		}
		return p.NumberExpr_Sempred(t, predIndex)

	case 38:
		var t *PostingListContext = nil
		if localctx != nil {
			t = localctx.(*PostingListContext)
		}
		return p.PostingList_Sempred(t, predIndex)

	case 48:
		var t *CostCompListContext = nil
		if localctx != nil {
			t = localctx.(*CostCompListContext)
		}
		return p.CostCompList_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *BeanCountParser) Declarations_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *BeanCountParser) NumberExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *BeanCountParser) PostingList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *BeanCountParser) CostCompList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
