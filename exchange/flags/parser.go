package flags

const CH_FLAG_HA int32 = 1 << 0

func Parse(flags int32, flag int32) bool {
	return flags&flag != 0
}
