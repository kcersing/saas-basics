// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__f64toa = 464
    _entry__format_significand = 48448
    _entry__format_integer = 3264
    _entry__hm_get = 27952
    _entry__i64toa = 3696
    _entry__u64toa = 3824
    _entry__j2t_fsm_exec = 39088
    _entry__advance_ns = 12720
    _entry__fsm_exec = 19536
    _entry__validate_string = 21728
    _entry__utf8_validate = 23152
    _entry__advance_string = 14944
    _entry__skip_number = 18288
    _entry__j2t_number = 34208
    _entry__vnumber = 15952
    _entry__atof_eisel_lemire64 = 10496
    _entry__atof_native = 12048
    _entry__decimal_to_f64 = 10864
    _entry__right_shift = 49408
    _entry__left_shift = 48912
    _entry__j2t_string = 34800
    _entry__unquote = 6880
    _entry__b64decode = 24448
    _entry__j2t_field_vm = 37456
    _entry__tb_write_default_or_empty = 31248
    _entry__j2t_write_unset_fields = 33344
    _entry__j2t_find_field_key = 36432
    _entry__quote = 5136
    _entry__tb_skip = 47408
    _entry__tb_write_i64 = 29504
    _entry__trie_get = 28768
)

const (
    _stack__f64toa = 80
    _stack__format_significand = 24
    _stack__format_integer = 16
    _stack__hm_get = 16
    _stack__i64toa = 16
    _stack__u64toa = 8
    _stack__j2t_fsm_exec = 592
    _stack__advance_ns = 16
    _stack__fsm_exec = 208
    _stack__validate_string = 120
    _stack__utf8_validate = 32
    _stack__advance_string = 64
    _stack__skip_number = 48
    _stack__j2t_number = 296
    _stack__vnumber = 240
    _stack__atof_eisel_lemire64 = 32
    _stack__atof_native = 136
    _stack__decimal_to_f64 = 80
    _stack__right_shift = 8
    _stack__left_shift = 24
    _stack__j2t_string = 160
    _stack__unquote = 88
    _stack__b64decode = 160
    _stack__j2t_field_vm = 312
    _stack__tb_write_default_or_empty = 56
    _stack__j2t_write_unset_fields = 176
    _stack__j2t_find_field_key = 32
    _stack__quote = 64
    _stack__tb_skip = 48
    _stack__tb_write_i64 = 8
    _stack__trie_get = 32
)

const (
    _size__f64toa = 2800
    _size__format_significand = 464
    _size__format_integer = 432
    _size__hm_get = 464
    _size__i64toa = 48
    _size__u64toa = 1264
    _size__j2t_fsm_exec = 7800
    _size__advance_ns = 704
    _size__fsm_exec = 1416
    _size__validate_string = 1424
    _size__utf8_validate = 416
    _size__advance_string = 960
    _size__skip_number = 956
    _size__j2t_number = 548
    _size__vnumber = 1552
    _size__atof_eisel_lemire64 = 368
    _size__atof_native = 608
    _size__decimal_to_f64 = 1184
    _size__right_shift = 400
    _size__left_shift = 496
    _size__j2t_string = 784
    _size__unquote = 2272
    _size__b64decode = 2832
    _size__j2t_field_vm = 1576
    _size__tb_write_default_or_empty = 1024
    _size__j2t_write_unset_fields = 864
    _size__j2t_find_field_key = 752
    _size__quote = 1728
    _size__tb_skip = 968
    _size__tb_write_i64 = 64
    _size__trie_get = 304
)

var (
    _pcsp__f64toa = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {2740, 56},
        {2744, 48},
        {2745, 40},
        {2747, 32},
        {2749, 24},
        {2751, 16},
        {2753, 8},
        {2754, 0},
        {2792, 56},
    }
    _pcsp__format_significand = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {452, 24},
        {453, 16},
        {455, 8},
        {457, 0},
    }
    _pcsp__format_integer = [][2]uint32{
        {1, 0},
        {4, 8},
        {412, 16},
        {413, 8},
        {414, 0},
        {423, 16},
        {424, 8},
        {426, 0},
    }
    _pcsp__hm_get = [][2]uint32{
        {1, 0},
        {4, 8},
        {387, 16},
        {388, 8},
        {390, 0},
        {450, 16},
        {451, 8},
        {452, 0},
        {458, 16},
        {459, 8},
        {461, 0},
    }
    _pcsp__i64toa = [][2]uint32{
        {14, 0},
        {34, 8},
        {36, 0},
    }
    _pcsp__u64toa = [][2]uint32{
        {1, 0},
        {161, 8},
        {162, 0},
        {457, 8},
        {458, 0},
        {772, 8},
        {773, 0},
        {1249, 8},
        {1251, 0},
    }
    _pcsp__j2t_fsm_exec = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {6416, 280},
        {6423, 48},
        {6424, 40},
        {6426, 32},
        {6428, 24},
        {6430, 16},
        {6432, 8},
        {6433, 0},
        {7800, 280},
    }
    _pcsp__advance_ns = [][2]uint32{
        {1, 0},
        {4, 8},
        {671, 16},
        {672, 8},
        {673, 0},
        {697, 16},
        {698, 8},
        {700, 0},
    }
    _pcsp__fsm_exec = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1220, 88},
        {1224, 48},
        {1225, 40},
        {1227, 32},
        {1229, 24},
        {1231, 16},
        {1233, 8},
        {1234, 0},
        {1416, 88},
    }
    _pcsp__validate_string = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {698, 88},
        {702, 48},
        {703, 40},
        {705, 32},
        {707, 24},
        {709, 16},
        {711, 8},
        {712, 0},
        {1409, 88},
    }
    _pcsp__utf8_validate = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {409, 32},
        {410, 24},
        {412, 16},
        {414, 8},
        {416, 0},
    }
    _pcsp__advance_string = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {576, 64},
        {580, 48},
        {581, 40},
        {583, 32},
        {585, 24},
        {587, 16},
        {589, 8},
        {590, 0},
        {955, 64},
    }
    _pcsp__skip_number = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {881, 48},
        {882, 40},
        {884, 32},
        {886, 24},
        {888, 16},
        {890, 8},
        {891, 0},
        {956, 48},
    }
    _pcsp__j2t_number = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {531, 56},
        {535, 48},
        {536, 40},
        {538, 32},
        {540, 24},
        {542, 16},
        {544, 8},
        {548, 0},
    }
    _pcsp__vnumber = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {803, 104},
        {807, 48},
        {808, 40},
        {810, 32},
        {812, 24},
        {814, 16},
        {816, 8},
        {817, 0},
        {1551, 104},
    }
    _pcsp__atof_eisel_lemire64 = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {292, 32},
        {293, 24},
        {295, 16},
        {297, 8},
        {298, 0},
        {362, 32},
    }
    _pcsp__atof_native = [][2]uint32{
        {1, 0},
        {4, 8},
        {587, 56},
        {591, 8},
        {593, 0},
    }
    _pcsp__decimal_to_f64 = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1144, 56},
        {1148, 48},
        {1149, 40},
        {1151, 32},
        {1153, 24},
        {1155, 16},
        {1157, 8},
        {1158, 0},
        {1169, 56},
    }
    _pcsp__right_shift = [][2]uint32{
        {1, 0},
        {318, 8},
        {319, 0},
        {387, 8},
        {388, 0},
        {396, 8},
        {398, 0},
    }
    _pcsp__left_shift = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {363, 24},
        {364, 16},
        {366, 8},
        {367, 0},
        {470, 24},
        {471, 16},
        {473, 8},
        {474, 0},
        {486, 24},
    }
    _pcsp__j2t_string = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {605, 72},
        {609, 48},
        {610, 40},
        {612, 32},
        {614, 24},
        {616, 16},
        {618, 8},
        {619, 0},
        {782, 72},
    }
    _pcsp__unquote = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1684, 88},
        {1688, 48},
        {1689, 40},
        {1691, 32},
        {1693, 24},
        {1695, 16},
        {1697, 8},
        {1698, 0},
        {2270, 88},
    }
    _pcsp__b64decode = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {2813, 160},
        {2817, 48},
        {2818, 40},
        {2820, 32},
        {2822, 24},
        {2824, 16},
        {2826, 8},
        {2828, 0},
    }
    _pcsp__j2t_field_vm = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {343, 72},
        {347, 48},
        {348, 40},
        {350, 32},
        {352, 24},
        {354, 16},
        {356, 8},
        {357, 0},
        {387, 72},
        {391, 48},
        {392, 40},
        {394, 32},
        {396, 24},
        {398, 16},
        {400, 8},
        {401, 0},
        {1576, 72},
    }
    _pcsp__tb_write_default_or_empty = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {960, 56},
        {964, 48},
        {965, 40},
        {967, 32},
        {969, 24},
        {971, 16},
        {973, 8},
        {974, 0},
        {1024, 56},
    }
    _pcsp__j2t_write_unset_fields = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {835, 120},
        {839, 48},
        {840, 40},
        {842, 32},
        {844, 24},
        {846, 16},
        {848, 8},
        {850, 0},
    }
    _pcsp__j2t_find_field_key = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {721, 32},
        {722, 24},
        {724, 16},
        {726, 8},
        {727, 0},
        {738, 32},
    }
    _pcsp__quote = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1681, 64},
        {1685, 48},
        {1686, 40},
        {1688, 32},
        {1690, 24},
        {1692, 16},
        {1694, 8},
        {1695, 0},
        {1722, 64},
    }
    _pcsp__tb_skip = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {944, 48},
        {945, 40},
        {947, 32},
        {949, 24},
        {951, 16},
        {953, 8},
        {954, 0},
        {968, 48},
    }
    _pcsp__tb_write_i64 = [][2]uint32{
        {1, 0},
        {33, 8},
        {34, 0},
        {51, 8},
        {53, 0},
    }
    _pcsp__trie_get = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {286, 32},
        {287, 24},
        {289, 16},
        {291, 8},
        {293, 0},
    }
)

var Funcs = []loader.CFunc{
    {"__native_entry__", 0, 67, 0, nil},
    {"_f64toa", _entry__f64toa, _size__f64toa, _stack__f64toa, _pcsp__f64toa},
    {"_format_significand", _entry__format_significand, _size__format_significand, _stack__format_significand, _pcsp__format_significand},
    {"_format_integer", _entry__format_integer, _size__format_integer, _stack__format_integer, _pcsp__format_integer},
    {"_hm_get", _entry__hm_get, _size__hm_get, _stack__hm_get, _pcsp__hm_get},
    {"_i64toa", _entry__i64toa, _size__i64toa, _stack__i64toa, _pcsp__i64toa},
    {"_u64toa", _entry__u64toa, _size__u64toa, _stack__u64toa, _pcsp__u64toa},
    {"_j2t_fsm_exec", _entry__j2t_fsm_exec, _size__j2t_fsm_exec, _stack__j2t_fsm_exec, _pcsp__j2t_fsm_exec},
    {"_advance_ns", _entry__advance_ns, _size__advance_ns, _stack__advance_ns, _pcsp__advance_ns},
    {"_fsm_exec", _entry__fsm_exec, _size__fsm_exec, _stack__fsm_exec, _pcsp__fsm_exec},
    {"_validate_string", _entry__validate_string, _size__validate_string, _stack__validate_string, _pcsp__validate_string},
    {"_utf8_validate", _entry__utf8_validate, _size__utf8_validate, _stack__utf8_validate, _pcsp__utf8_validate},
    {"_advance_string", _entry__advance_string, _size__advance_string, _stack__advance_string, _pcsp__advance_string},
    {"_skip_number", _entry__skip_number, _size__skip_number, _stack__skip_number, _pcsp__skip_number},
    {"_j2t_number", _entry__j2t_number, _size__j2t_number, _stack__j2t_number, _pcsp__j2t_number},
    {"_vnumber", _entry__vnumber, _size__vnumber, _stack__vnumber, _pcsp__vnumber},
    {"_atof_eisel_lemire64", _entry__atof_eisel_lemire64, _size__atof_eisel_lemire64, _stack__atof_eisel_lemire64, _pcsp__atof_eisel_lemire64},
    {"_atof_native", _entry__atof_native, _size__atof_native, _stack__atof_native, _pcsp__atof_native},
    {"_decimal_to_f64", _entry__decimal_to_f64, _size__decimal_to_f64, _stack__decimal_to_f64, _pcsp__decimal_to_f64},
    {"_right_shift", _entry__right_shift, _size__right_shift, _stack__right_shift, _pcsp__right_shift},
    {"_left_shift", _entry__left_shift, _size__left_shift, _stack__left_shift, _pcsp__left_shift},
    {"_j2t_string", _entry__j2t_string, _size__j2t_string, _stack__j2t_string, _pcsp__j2t_string},
    {"_unquote", _entry__unquote, _size__unquote, _stack__unquote, _pcsp__unquote},
    {"_b64decode", _entry__b64decode, _size__b64decode, _stack__b64decode, _pcsp__b64decode},
    {"_j2t_field_vm", _entry__j2t_field_vm, _size__j2t_field_vm, _stack__j2t_field_vm, _pcsp__j2t_field_vm},
    {"_tb_write_default_or_empty", _entry__tb_write_default_or_empty, _size__tb_write_default_or_empty, _stack__tb_write_default_or_empty, _pcsp__tb_write_default_or_empty},
    {"_j2t_write_unset_fields", _entry__j2t_write_unset_fields, _size__j2t_write_unset_fields, _stack__j2t_write_unset_fields, _pcsp__j2t_write_unset_fields},
    {"_j2t_find_field_key", _entry__j2t_find_field_key, _size__j2t_find_field_key, _stack__j2t_find_field_key, _pcsp__j2t_find_field_key},
    {"_quote", _entry__quote, _size__quote, _stack__quote, _pcsp__quote},
    {"_tb_skip", _entry__tb_skip, _size__tb_skip, _stack__tb_skip, _pcsp__tb_skip},
    {"_tb_write_i64", _entry__tb_write_i64, _size__tb_write_i64, _stack__tb_write_i64, _pcsp__tb_write_i64},
    {"_trie_get", _entry__trie_get, _size__trie_get, _stack__trie_get, _pcsp__trie_get},
}