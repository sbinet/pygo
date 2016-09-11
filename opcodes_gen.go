// Copyright 2016 The pygo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pygo

import "fmt"

// Python-3 Opcodes
const (
	Op_000                        Opcode = 0
	Op_POP_TOP                    Opcode = 1
	Op_ROT_TWO                    Opcode = 2
	Op_ROT_THREE                  Opcode = 3
	Op_DUP_TOP                    Opcode = 4
	Op_DUP_TOP_TWO                Opcode = 5
	Op_006                        Opcode = 6
	Op_007                        Opcode = 7
	Op_008                        Opcode = 8
	Op_NOP                        Opcode = 9
	Op_UNARY_POSITIVE             Opcode = 10
	Op_UNARY_NEGATIVE             Opcode = 11
	Op_UNARY_NOT                  Opcode = 12
	Op_013                        Opcode = 13
	Op_014                        Opcode = 14
	Op_UNARY_INVERT               Opcode = 15
	Op_BINARY_MATRIX_MULTIPLY     Opcode = 16
	Op_INPLACE_MATRIX_MULTIPLY    Opcode = 17
	Op_018                        Opcode = 18
	Op_BINARY_POWER               Opcode = 19
	Op_BINARY_MULTIPLY            Opcode = 20
	Op_021                        Opcode = 21
	Op_BINARY_MODULO              Opcode = 22
	Op_BINARY_ADD                 Opcode = 23
	Op_BINARY_SUBTRACT            Opcode = 24
	Op_BINARY_SUBSCR              Opcode = 25
	Op_BINARY_FLOOR_DIVIDE        Opcode = 26
	Op_BINARY_TRUE_DIVIDE         Opcode = 27
	Op_INPLACE_FLOOR_DIVIDE       Opcode = 28
	Op_INPLACE_TRUE_DIVIDE        Opcode = 29
	Op_030                        Opcode = 30
	Op_031                        Opcode = 31
	Op_032                        Opcode = 32
	Op_033                        Opcode = 33
	Op_034                        Opcode = 34
	Op_035                        Opcode = 35
	Op_036                        Opcode = 36
	Op_037                        Opcode = 37
	Op_038                        Opcode = 38
	Op_039                        Opcode = 39
	Op_040                        Opcode = 40
	Op_041                        Opcode = 41
	Op_042                        Opcode = 42
	Op_043                        Opcode = 43
	Op_044                        Opcode = 44
	Op_045                        Opcode = 45
	Op_046                        Opcode = 46
	Op_047                        Opcode = 47
	Op_048                        Opcode = 48
	Op_049                        Opcode = 49
	Op_GET_AITER                  Opcode = 50
	Op_GET_ANEXT                  Opcode = 51
	Op_BEFORE_ASYNC_WITH          Opcode = 52
	Op_053                        Opcode = 53
	Op_054                        Opcode = 54
	Op_INPLACE_ADD                Opcode = 55
	Op_INPLACE_SUBTRACT           Opcode = 56
	Op_INPLACE_MULTIPLY           Opcode = 57
	Op_058                        Opcode = 58
	Op_INPLACE_MODULO             Opcode = 59
	Op_STORE_SUBSCR               Opcode = 60
	Op_DELETE_SUBSCR              Opcode = 61
	Op_BINARY_LSHIFT              Opcode = 62
	Op_BINARY_RSHIFT              Opcode = 63
	Op_BINARY_AND                 Opcode = 64
	Op_BINARY_XOR                 Opcode = 65
	Op_BINARY_OR                  Opcode = 66
	Op_INPLACE_POWER              Opcode = 67
	Op_GET_ITER                   Opcode = 68
	Op_GET_YIELD_FROM_ITER        Opcode = 69
	Op_PRINT_EXPR                 Opcode = 70
	Op_LOAD_BUILD_CLASS           Opcode = 71
	Op_YIELD_FROM                 Opcode = 72
	Op_GET_AWAITABLE              Opcode = 73
	Op_074                        Opcode = 74
	Op_INPLACE_LSHIFT             Opcode = 75
	Op_INPLACE_RSHIFT             Opcode = 76
	Op_INPLACE_AND                Opcode = 77
	Op_INPLACE_XOR                Opcode = 78
	Op_INPLACE_OR                 Opcode = 79
	Op_BREAK_LOOP                 Opcode = 80
	Op_WITH_CLEANUP_START         Opcode = 81
	Op_WITH_CLEANUP_FINISH        Opcode = 82
	Op_RETURN_VALUE               Opcode = 83
	Op_IMPORT_STAR                Opcode = 84
	Op_085                        Opcode = 85
	Op_YIELD_VALUE                Opcode = 86
	Op_POP_BLOCK                  Opcode = 87
	Op_END_FINALLY                Opcode = 88
	Op_POP_EXCEPT                 Opcode = 89
	Op_STORE_NAME                 Opcode = 90
	Op_DELETE_NAME                Opcode = 91
	Op_UNPACK_SEQUENCE            Opcode = 92
	Op_FOR_ITER                   Opcode = 93
	Op_UNPACK_EX                  Opcode = 94
	Op_STORE_ATTR                 Opcode = 95
	Op_DELETE_ATTR                Opcode = 96
	Op_STORE_GLOBAL               Opcode = 97
	Op_DELETE_GLOBAL              Opcode = 98
	Op_099                        Opcode = 99
	Op_LOAD_CONST                 Opcode = 100
	Op_LOAD_NAME                  Opcode = 101
	Op_BUILD_TUPLE                Opcode = 102
	Op_BUILD_LIST                 Opcode = 103
	Op_BUILD_SET                  Opcode = 104
	Op_BUILD_MAP                  Opcode = 105
	Op_LOAD_ATTR                  Opcode = 106
	Op_COMPARE_OP                 Opcode = 107
	Op_IMPORT_NAME                Opcode = 108
	Op_IMPORT_FROM                Opcode = 109
	Op_JUMP_FORWARD               Opcode = 110
	Op_JUMP_IF_FALSE_OR_POP       Opcode = 111
	Op_JUMP_IF_TRUE_OR_POP        Opcode = 112
	Op_JUMP_ABSOLUTE              Opcode = 113
	Op_POP_JUMP_IF_FALSE          Opcode = 114
	Op_POP_JUMP_IF_TRUE           Opcode = 115
	Op_LOAD_GLOBAL                Opcode = 116
	Op_117                        Opcode = 117
	Op_118                        Opcode = 118
	Op_CONTINUE_LOOP              Opcode = 119
	Op_SETUP_LOOP                 Opcode = 120
	Op_SETUP_EXCEPT               Opcode = 121
	Op_SETUP_FINALLY              Opcode = 122
	Op_123                        Opcode = 123
	Op_LOAD_FAST                  Opcode = 124
	Op_STORE_FAST                 Opcode = 125
	Op_DELETE_FAST                Opcode = 126
	Op_127                        Opcode = 127
	Op_128                        Opcode = 128
	Op_129                        Opcode = 129
	Op_RAISE_VARARGS              Opcode = 130
	Op_CALL_FUNCTION              Opcode = 131
	Op_MAKE_FUNCTION              Opcode = 132
	Op_BUILD_SLICE                Opcode = 133
	Op_MAKE_CLOSURE               Opcode = 134
	Op_LOAD_CLOSURE               Opcode = 135
	Op_LOAD_DEREF                 Opcode = 136
	Op_STORE_DEREF                Opcode = 137
	Op_DELETE_DEREF               Opcode = 138
	Op_139                        Opcode = 139
	Op_CALL_FUNCTION_VAR          Opcode = 140
	Op_CALL_FUNCTION_KW           Opcode = 141
	Op_CALL_FUNCTION_VAR_KW       Opcode = 142
	Op_SETUP_WITH                 Opcode = 143
	Op_EXTENDED_ARG               Opcode = 144
	Op_LIST_APPEND                Opcode = 145
	Op_SET_ADD                    Opcode = 146
	Op_MAP_ADD                    Opcode = 147
	Op_LOAD_CLASSDEREF            Opcode = 148
	Op_BUILD_LIST_UNPACK          Opcode = 149
	Op_BUILD_MAP_UNPACK           Opcode = 150
	Op_BUILD_MAP_UNPACK_WITH_CALL Opcode = 151
	Op_BUILD_TUPLE_UNPACK         Opcode = 152
	Op_BUILD_SET_UNPACK           Opcode = 153
	Op_SETUP_ASYNC_WITH           Opcode = 154
	Op_155                        Opcode = 155
	Op_156                        Opcode = 156
	Op_157                        Opcode = 157
	Op_158                        Opcode = 158
	Op_159                        Opcode = 159
	Op_160                        Opcode = 160
	Op_161                        Opcode = 161
	Op_162                        Opcode = 162
	Op_163                        Opcode = 163
	Op_164                        Opcode = 164
	Op_165                        Opcode = 165
	Op_166                        Opcode = 166
	Op_167                        Opcode = 167
	Op_168                        Opcode = 168
	Op_169                        Opcode = 169
	Op_170                        Opcode = 170
	Op_171                        Opcode = 171
	Op_172                        Opcode = 172
	Op_173                        Opcode = 173
	Op_174                        Opcode = 174
	Op_175                        Opcode = 175
	Op_176                        Opcode = 176
	Op_177                        Opcode = 177
	Op_178                        Opcode = 178
	Op_179                        Opcode = 179
	Op_180                        Opcode = 180
	Op_181                        Opcode = 181
	Op_182                        Opcode = 182
	Op_183                        Opcode = 183
	Op_184                        Opcode = 184
	Op_185                        Opcode = 185
	Op_186                        Opcode = 186
	Op_187                        Opcode = 187
	Op_188                        Opcode = 188
	Op_189                        Opcode = 189
	Op_190                        Opcode = 190
	Op_191                        Opcode = 191
	Op_192                        Opcode = 192
	Op_193                        Opcode = 193
	Op_194                        Opcode = 194
	Op_195                        Opcode = 195
	Op_196                        Opcode = 196
	Op_197                        Opcode = 197
	Op_198                        Opcode = 198
	Op_199                        Opcode = 199
	Op_200                        Opcode = 200
	Op_201                        Opcode = 201
	Op_202                        Opcode = 202
	Op_203                        Opcode = 203
	Op_204                        Opcode = 204
	Op_205                        Opcode = 205
	Op_206                        Opcode = 206
	Op_207                        Opcode = 207
	Op_208                        Opcode = 208
	Op_209                        Opcode = 209
	Op_210                        Opcode = 210
	Op_211                        Opcode = 211
	Op_212                        Opcode = 212
	Op_213                        Opcode = 213
	Op_214                        Opcode = 214
	Op_215                        Opcode = 215
	Op_216                        Opcode = 216
	Op_217                        Opcode = 217
	Op_218                        Opcode = 218
	Op_219                        Opcode = 219
	Op_220                        Opcode = 220
	Op_221                        Opcode = 221
	Op_222                        Opcode = 222
	Op_223                        Opcode = 223
	Op_224                        Opcode = 224
	Op_225                        Opcode = 225
	Op_226                        Opcode = 226
	Op_227                        Opcode = 227
	Op_228                        Opcode = 228
	Op_229                        Opcode = 229
	Op_230                        Opcode = 230
	Op_231                        Opcode = 231
	Op_232                        Opcode = 232
	Op_233                        Opcode = 233
	Op_234                        Opcode = 234
	Op_235                        Opcode = 235
	Op_236                        Opcode = 236
	Op_237                        Opcode = 237
	Op_238                        Opcode = 238
	Op_239                        Opcode = 239
	Op_240                        Opcode = 240
	Op_241                        Opcode = 241
	Op_242                        Opcode = 242
	Op_243                        Opcode = 243
	Op_244                        Opcode = 244
	Op_245                        Opcode = 245
	Op_246                        Opcode = 246
	Op_247                        Opcode = 247
	Op_248                        Opcode = 248
	Op_249                        Opcode = 249
	Op_250                        Opcode = 250
	Op_251                        Opcode = 251
	Op_252                        Opcode = 252
	Op_253                        Opcode = 253
	Op_254                        Opcode = 254
	Op_255                        Opcode = 255
)

func (op Opcode) String() string {
	switch op {
	case 0:
		return "<0>"
	case 1:
		return "POP_TOP"
	case 2:
		return "ROT_TWO"
	case 3:
		return "ROT_THREE"
	case 4:
		return "DUP_TOP"
	case 5:
		return "DUP_TOP_TWO"
	case 6:
		return "<6>"
	case 7:
		return "<7>"
	case 8:
		return "<8>"
	case 9:
		return "NOP"
	case 10:
		return "UNARY_POSITIVE"
	case 11:
		return "UNARY_NEGATIVE"
	case 12:
		return "UNARY_NOT"
	case 13:
		return "<13>"
	case 14:
		return "<14>"
	case 15:
		return "UNARY_INVERT"
	case 16:
		return "BINARY_MATRIX_MULTIPLY"
	case 17:
		return "INPLACE_MATRIX_MULTIPLY"
	case 18:
		return "<18>"
	case 19:
		return "BINARY_POWER"
	case 20:
		return "BINARY_MULTIPLY"
	case 21:
		return "<21>"
	case 22:
		return "BINARY_MODULO"
	case 23:
		return "BINARY_ADD"
	case 24:
		return "BINARY_SUBTRACT"
	case 25:
		return "BINARY_SUBSCR"
	case 26:
		return "BINARY_FLOOR_DIVIDE"
	case 27:
		return "BINARY_TRUE_DIVIDE"
	case 28:
		return "INPLACE_FLOOR_DIVIDE"
	case 29:
		return "INPLACE_TRUE_DIVIDE"
	case 30:
		return "<30>"
	case 31:
		return "<31>"
	case 32:
		return "<32>"
	case 33:
		return "<33>"
	case 34:
		return "<34>"
	case 35:
		return "<35>"
	case 36:
		return "<36>"
	case 37:
		return "<37>"
	case 38:
		return "<38>"
	case 39:
		return "<39>"
	case 40:
		return "<40>"
	case 41:
		return "<41>"
	case 42:
		return "<42>"
	case 43:
		return "<43>"
	case 44:
		return "<44>"
	case 45:
		return "<45>"
	case 46:
		return "<46>"
	case 47:
		return "<47>"
	case 48:
		return "<48>"
	case 49:
		return "<49>"
	case 50:
		return "GET_AITER"
	case 51:
		return "GET_ANEXT"
	case 52:
		return "BEFORE_ASYNC_WITH"
	case 53:
		return "<53>"
	case 54:
		return "<54>"
	case 55:
		return "INPLACE_ADD"
	case 56:
		return "INPLACE_SUBTRACT"
	case 57:
		return "INPLACE_MULTIPLY"
	case 58:
		return "<58>"
	case 59:
		return "INPLACE_MODULO"
	case 60:
		return "STORE_SUBSCR"
	case 61:
		return "DELETE_SUBSCR"
	case 62:
		return "BINARY_LSHIFT"
	case 63:
		return "BINARY_RSHIFT"
	case 64:
		return "BINARY_AND"
	case 65:
		return "BINARY_XOR"
	case 66:
		return "BINARY_OR"
	case 67:
		return "INPLACE_POWER"
	case 68:
		return "GET_ITER"
	case 69:
		return "GET_YIELD_FROM_ITER"
	case 70:
		return "PRINT_EXPR"
	case 71:
		return "LOAD_BUILD_CLASS"
	case 72:
		return "YIELD_FROM"
	case 73:
		return "GET_AWAITABLE"
	case 74:
		return "<74>"
	case 75:
		return "INPLACE_LSHIFT"
	case 76:
		return "INPLACE_RSHIFT"
	case 77:
		return "INPLACE_AND"
	case 78:
		return "INPLACE_XOR"
	case 79:
		return "INPLACE_OR"
	case 80:
		return "BREAK_LOOP"
	case 81:
		return "WITH_CLEANUP_START"
	case 82:
		return "WITH_CLEANUP_FINISH"
	case 83:
		return "RETURN_VALUE"
	case 84:
		return "IMPORT_STAR"
	case 85:
		return "<85>"
	case 86:
		return "YIELD_VALUE"
	case 87:
		return "POP_BLOCK"
	case 88:
		return "END_FINALLY"
	case 89:
		return "POP_EXCEPT"
	case 90:
		return "STORE_NAME"
	case 91:
		return "DELETE_NAME"
	case 92:
		return "UNPACK_SEQUENCE"
	case 93:
		return "FOR_ITER"
	case 94:
		return "UNPACK_EX"
	case 95:
		return "STORE_ATTR"
	case 96:
		return "DELETE_ATTR"
	case 97:
		return "STORE_GLOBAL"
	case 98:
		return "DELETE_GLOBAL"
	case 99:
		return "<99>"
	case 100:
		return "LOAD_CONST"
	case 101:
		return "LOAD_NAME"
	case 102:
		return "BUILD_TUPLE"
	case 103:
		return "BUILD_LIST"
	case 104:
		return "BUILD_SET"
	case 105:
		return "BUILD_MAP"
	case 106:
		return "LOAD_ATTR"
	case 107:
		return "COMPARE_OP"
	case 108:
		return "IMPORT_NAME"
	case 109:
		return "IMPORT_FROM"
	case 110:
		return "JUMP_FORWARD"
	case 111:
		return "JUMP_IF_FALSE_OR_POP"
	case 112:
		return "JUMP_IF_TRUE_OR_POP"
	case 113:
		return "JUMP_ABSOLUTE"
	case 114:
		return "POP_JUMP_IF_FALSE"
	case 115:
		return "POP_JUMP_IF_TRUE"
	case 116:
		return "LOAD_GLOBAL"
	case 117:
		return "<117>"
	case 118:
		return "<118>"
	case 119:
		return "CONTINUE_LOOP"
	case 120:
		return "SETUP_LOOP"
	case 121:
		return "SETUP_EXCEPT"
	case 122:
		return "SETUP_FINALLY"
	case 123:
		return "<123>"
	case 124:
		return "LOAD_FAST"
	case 125:
		return "STORE_FAST"
	case 126:
		return "DELETE_FAST"
	case 127:
		return "<127>"
	case 128:
		return "<128>"
	case 129:
		return "<129>"
	case 130:
		return "RAISE_VARARGS"
	case 131:
		return "CALL_FUNCTION"
	case 132:
		return "MAKE_FUNCTION"
	case 133:
		return "BUILD_SLICE"
	case 134:
		return "MAKE_CLOSURE"
	case 135:
		return "LOAD_CLOSURE"
	case 136:
		return "LOAD_DEREF"
	case 137:
		return "STORE_DEREF"
	case 138:
		return "DELETE_DEREF"
	case 139:
		return "<139>"
	case 140:
		return "CALL_FUNCTION_VAR"
	case 141:
		return "CALL_FUNCTION_KW"
	case 142:
		return "CALL_FUNCTION_VAR_KW"
	case 143:
		return "SETUP_WITH"
	case 144:
		return "EXTENDED_ARG"
	case 145:
		return "LIST_APPEND"
	case 146:
		return "SET_ADD"
	case 147:
		return "MAP_ADD"
	case 148:
		return "LOAD_CLASSDEREF"
	case 149:
		return "BUILD_LIST_UNPACK"
	case 150:
		return "BUILD_MAP_UNPACK"
	case 151:
		return "BUILD_MAP_UNPACK_WITH_CALL"
	case 152:
		return "BUILD_TUPLE_UNPACK"
	case 153:
		return "BUILD_SET_UNPACK"
	case 154:
		return "SETUP_ASYNC_WITH"
	case 155:
		return "<155>"
	case 156:
		return "<156>"
	case 157:
		return "<157>"
	case 158:
		return "<158>"
	case 159:
		return "<159>"
	case 160:
		return "<160>"
	case 161:
		return "<161>"
	case 162:
		return "<162>"
	case 163:
		return "<163>"
	case 164:
		return "<164>"
	case 165:
		return "<165>"
	case 166:
		return "<166>"
	case 167:
		return "<167>"
	case 168:
		return "<168>"
	case 169:
		return "<169>"
	case 170:
		return "<170>"
	case 171:
		return "<171>"
	case 172:
		return "<172>"
	case 173:
		return "<173>"
	case 174:
		return "<174>"
	case 175:
		return "<175>"
	case 176:
		return "<176>"
	case 177:
		return "<177>"
	case 178:
		return "<178>"
	case 179:
		return "<179>"
	case 180:
		return "<180>"
	case 181:
		return "<181>"
	case 182:
		return "<182>"
	case 183:
		return "<183>"
	case 184:
		return "<184>"
	case 185:
		return "<185>"
	case 186:
		return "<186>"
	case 187:
		return "<187>"
	case 188:
		return "<188>"
	case 189:
		return "<189>"
	case 190:
		return "<190>"
	case 191:
		return "<191>"
	case 192:
		return "<192>"
	case 193:
		return "<193>"
	case 194:
		return "<194>"
	case 195:
		return "<195>"
	case 196:
		return "<196>"
	case 197:
		return "<197>"
	case 198:
		return "<198>"
	case 199:
		return "<199>"
	case 200:
		return "<200>"
	case 201:
		return "<201>"
	case 202:
		return "<202>"
	case 203:
		return "<203>"
	case 204:
		return "<204>"
	case 205:
		return "<205>"
	case 206:
		return "<206>"
	case 207:
		return "<207>"
	case 208:
		return "<208>"
	case 209:
		return "<209>"
	case 210:
		return "<210>"
	case 211:
		return "<211>"
	case 212:
		return "<212>"
	case 213:
		return "<213>"
	case 214:
		return "<214>"
	case 215:
		return "<215>"
	case 216:
		return "<216>"
	case 217:
		return "<217>"
	case 218:
		return "<218>"
	case 219:
		return "<219>"
	case 220:
		return "<220>"
	case 221:
		return "<221>"
	case 222:
		return "<222>"
	case 223:
		return "<223>"
	case 224:
		return "<224>"
	case 225:
		return "<225>"
	case 226:
		return "<226>"
	case 227:
		return "<227>"
	case 228:
		return "<228>"
	case 229:
		return "<229>"
	case 230:
		return "<230>"
	case 231:
		return "<231>"
	case 232:
		return "<232>"
	case 233:
		return "<233>"
	case 234:
		return "<234>"
	case 235:
		return "<235>"
	case 236:
		return "<236>"
	case 237:
		return "<237>"
	case 238:
		return "<238>"
	case 239:
		return "<239>"
	case 240:
		return "<240>"
	case 241:
		return "<241>"
	case 242:
		return "<242>"
	case 243:
		return "<243>"
	case 244:
		return "<244>"
	case 245:
		return "<245>"
	case 246:
		return "<246>"
	case 247:
		return "<247>"
	case 248:
		return "<248>"
	case 249:
		return "<249>"
	case 250:
		return "<250>"
	case 251:
		return "<251>"
	case 252:
		return "<252>"
	case 253:
		return "<253>"
	case 254:
		return "<254>"
	case 255:
		return "<255>"
	default:
		panic(fmt.Errorf("invalid opcode value %%d", byte(op)))
	}
	return ""
}
