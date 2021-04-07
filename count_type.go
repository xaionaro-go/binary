// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package binary

import (
	"fmt"
)

// CountType defines the size of a slice item count field.
type CountType uint

const (
	// CountTypeUvarint will use binary.ReadUvarint to read a slice count.
	CountTypeUvarint = CountType(iota)

	// CountTypeUint8 will use just read one byte as a slice count.
	CountTypeUint8

	// CountTypeLEUint16 will use binary.LittleEndian.Uint16 to read a slice count.
	CountTypeLEUint16

	// CountTypeLEUint32 will use binary.LittleEndian.Uint32 to read a slice count.
	CountTypeLEUint32

	// CountTypeLEUint64 will use binary.LittleEndian.Uint64 to read a slice count.
	CountTypeLEUint64
)

func (t CountType) String() string {
	switch t {
	case CountTypeUvarint:
		return "uvarint"
	case CountTypeUint8:
		return "uint8"
	case CountTypeLEUint16:
		return "uint16_le"
	case CountTypeLEUint32:
		return "uint32_le"
	case CountTypeLEUint64:
		return "uint64_le"
	}

	return fmt.Sprintf("CountType_%d", t)
}

func parseCountType(s string) CountType {
	switch s {
	case "", "uvarint":
		return CountTypeUvarint
	case "uint8":
		return CountTypeUint8
	case "uint16_le":
		return CountTypeLEUint16
	case "uint32_le":
		return CountTypeLEUint32
	case "uint64_le":
		return CountTypeLEUint64
	}

	// if unable to parse, then:
	return CountTypeUvarint
}
