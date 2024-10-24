package documentstore

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type Fields map[string]DocumentField

type DocumentField struct {
	Type  DocumentFieldType
	Value any
}

type Document struct {
	Fields
}

func NewDocument(fields map[string]DocumentField) *Document {
	return &Document{
		Fields: fields,
	}
}
