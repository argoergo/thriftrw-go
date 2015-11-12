package ast

import "fmt"

// Type unifies the different types representing Thrift field types.
type Type interface {
	fieldType()

	fmt.Stringer
}

// BaseTypeID is an identifier for primitive types supported by Thrift.
type BaseTypeID int

//go:generate stringer -type=BaseTypeID

// IDs of the base types supported by Thrift.
const (
	BoolTypeID   BaseTypeID = iota + 1 // bool
	ByteTypeID                         // byte
	I16TypeID                          // i16
	I32TypeID                          // i32
	I64TypeID                          // i64
	DoubleTypeID                       // double
	StringTypeID                       // string
	BinaryTypeID                       // binary
)

// BaseType is a reference to a Thrift base type.
//
// 	bool, byte, i16, i32, i64, double, string, binary
//
// All references to base types in the document may be followed by type
// annotations.
//
// 	bool (go.type = "int")
type BaseType struct {
	// ID of the base type.
	ID BaseTypeID

	// Type annotations associated with this reference.
	Annotations []*Annotation
}

func (BaseType) fieldType() {}

func (bt BaseType) String() string {
	var name string

	switch bt.ID {
	case BoolTypeID:
		name = "bool"
	case ByteTypeID:
		name = "byte"
	case I16TypeID:
		name = "i16"
	case I32TypeID:
		name = "i32"
	case I64TypeID:
		name = "i64"
	case DoubleTypeID:
		name = "double"
	case StringTypeID:
		name = "string"
	case BinaryTypeID:
		name = "binary"
	default:
		panic(fmt.Sprintf("unknown base type %v", bt))
	}

	if s := FormatAnnotations(bt.Annotations); len(s) > 0 {
		name = name + " " + s
	}

	return name
}

// MapType is a reference to a the Thrift map type.
//
// 	map<k, v>
//
// All references to map types may be followed by type annotations.
//
// 	map<string, list<i32>> (java.type = "MultiMap")
type MapType struct {
	KeyType, ValueType Type
	Annotations        []*Annotation
}

func (MapType) fieldType() {}

func (mt MapType) String() string {
	name := fmt.Sprintf("map<%s, %s>", mt.KeyType, mt.ValueType)

	if s := FormatAnnotations(mt.Annotations); len(s) > 0 {
		name = name + " " + s
	}

	return name
}

// ListType is a reference to the Thrift list type.
//
// 	list<a>
//
// All references to list types may be followed by type annotations.
//
// 	list<i64> (cpp.type = "vector")
type ListType struct {
	ValueType   Type
	Annotations []*Annotation
}

func (ListType) fieldType() {}

func (lt ListType) String() string {
	name := fmt.Sprintf("list<%s>", lt.ValueType.String())

	if s := FormatAnnotations(lt.Annotations); len(s) > 0 {
		name = name + " " + s
	}

	return name
}

// SetType is a reference to the Thrift set type.
//
// 	set<a>
//
// All references to set types may be followed by type annotations.
//
// 	set<string> (js.type = "list")
type SetType struct {
	ValueType   Type
	Annotations []*Annotation
}

func (SetType) fieldType() {}

func (st SetType) String() string {
	name := fmt.Sprintf("set<%s>", st.ValueType.String())

	if s := FormatAnnotations(st.Annotations); len(s) > 0 {
		name = name + " " + s
	}

	return name
}

// TypeReference references a user-defined type.
type TypeReference struct {
	Name string
	Line int
}

func (TypeReference) fieldType() {}

func (tr TypeReference) String() string {
	return tr.Name
}
