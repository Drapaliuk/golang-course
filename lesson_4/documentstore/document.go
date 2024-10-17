package documentstore

type DocumentFieldType string

var store = []Document{}

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type DocumentField struct {
	Type DocumentFieldType
}

type Document struct {
	key    string
	Fields map[string]DocumentField
}

func NewDocument(key string) Document {
	return Document{
		key: key,
	}
}
