// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"reflect"
	"regexp"
)

const (
	requestTagKey       = "request"
	multipartFormTagKey = "multipartForm"
	formTagKey          = "form"
)

var (
	jsonEncodingRegex       = regexp.MustCompile(`(application|text)\/.*?\+*json.*`)
	multipartEncodingRegex  = regexp.MustCompile(`multipart\/.*`)
	urlEncodedEncodingRegex = regexp.MustCompile(`application\/x-www-form-urlencoded.*`)
)

func SerializeRequestBody(_ context.Context, request interface{}, nullable, optional bool, requestFieldName, serializationMethod, tag string) (io.Reader, string, error) {
	bodyReader, contentType, err := serializeRequestBody(request, nullable, optional, requestFieldName, serializationMethod, tag)
	if err != nil {
		return nil, "", fmt.Errorf("error serializing request body: %w", err)
	}

	if bodyReader == nil && !optional {
		return nil, "", fmt.Errorf("request body is required")
	}

	return bodyReader, contentType, nil
}

func serializeRequestBody(request interface{}, nullable, optional bool, requestFieldName, serializationMethod, tag string) (io.Reader, string, error) {
	requestStructType := reflect.TypeOf(request)
	requestValType := reflect.ValueOf(request)

	if isNil(requestStructType, requestValType) {
		if !nullable && optional {
			return nil, "", nil
		}

		return serializeContentType(requestFieldName, SerializationMethodToContentType[serializationMethod], requestValType, tag)
	}

	if requestStructType.Kind() == reflect.Pointer {
		requestStructType = requestStructType.Elem()
		requestValType = requestValType.Elem()
	}

	if requestStructType.Kind() != reflect.Struct {
		return serializeContentType(requestFieldName, SerializationMethodToContentType[serializationMethod], requestValType, tag)
	}

	requestField, ok := requestStructType.FieldByName(requestFieldName)

	if ok {
		tag := getRequestTag(requestField)
		if tag != nil {
			// request object (non-flattened)
			requestVal := requestValType.FieldByName(requestFieldName)
			if isNil(requestField.Type, requestVal) {
				if !nullable && optional {
					return nil, "", nil
				}

				return serializeContentType(requestFieldName, tag.MediaType, requestVal, string(requestField.Tag))
			}

			return serializeContentType(requestFieldName, tag.MediaType, requestVal, string(requestField.Tag))
		}
	}

	// flattened request object
	return serializeContentType(requestFieldName, SerializationMethodToContentType[serializationMethod], reflect.ValueOf(request), tag)
}

func serializeContentType(fieldName string, mediaType string, val reflect.Value, tag string) (io.Reader, string, error) {
	buf := &bytes.Buffer{}

	if isNil(val.Type(), val) {
		// TODO: what does a null mean for other content types? Just returning an empty buffer for now
		if jsonEncodingRegex.MatchString(mediaType) {
			if _, err := buf.Write([]byte("null")); err != nil {
				return nil, "", err
			}
		}

		return buf, mediaType, nil
	}

	switch {
	case jsonEncodingRegex.MatchString(mediaType):
		data, err := MarshalJSON(val.Interface(), reflect.StructTag(tag), true)
		if err != nil {
			return nil, "", err
		}

		if _, err := buf.Write(data); err != nil {
			return nil, "", err
		}
	case multipartEncodingRegex.MatchString(mediaType):
		var err error
		mediaType, err = encodeMultipartFormData(buf, val.Interface())
		if err != nil {
			return nil, "", err
		}
	case urlEncodedEncodingRegex.MatchString(mediaType):
		if err := encodeFormData(fieldName, buf, val.Interface()); err != nil {
			return nil, "", err
		}
	case val.Type().Implements(reflect.TypeOf((*io.Reader)(nil)).Elem()):
		return val.Interface().(io.Reader), mediaType, nil
	default:
		val = reflect.Indirect(val)

		switch {
		case val.Type().Kind() == reflect.String:
			if _, err := buf.WriteString(valToString(val.Interface())); err != nil {
				return nil, "", err
			}
		case reflect.TypeOf(val.Interface()) == reflect.TypeOf([]byte(nil)):
			if _, err := buf.Write(val.Interface().([]byte)); err != nil {
				return nil, "", err
			}
		default:
			return nil, "", fmt.Errorf("invalid request body type %s for mediaType %s", val.Type(), mediaType)
		}
	}

	return buf, mediaType, nil
}

func encodeMultipartFormData(w io.Writer, data interface{}) (string, error) {
	requestStructType := reflect.TypeOf(data)
	requestValType := reflect.ValueOf(data)

	if requestStructType.Kind() == reflect.Pointer {
		requestStructType = requestStructType.Elem()
		requestValType = requestValType.Elem()
	}

	writer := multipart.NewWriter(w)

	for i := 0; i < requestStructType.NumField(); i++ {
		field := requestStructType.Field(i)
		fieldType := field.Type
		valType := requestValType.Field(i)

		if isNil(fieldType, valType) {
			continue
		}

		if fieldType.Kind() == reflect.Pointer {
			fieldType = fieldType.Elem()
			valType = valType.Elem()
		}

		tag := parseMultipartFormTag(field)
		if tag.File {
			if err := encodeMultipartFormDataFile(writer, fieldType, valType); err != nil {
				writer.Close()
				return "", err
			}
		} else if tag.JSON {
			jw, err := writer.CreateFormField(tag.Name)
			if err != nil {
				writer.Close()
				return "", err
			}
			d, err := MarshalJSON(valType.Interface(), field.Tag, true)
			if err != nil {
				writer.Close()
				return "", err
			}
			if _, err := jw.Write(d); err != nil {
				writer.Close()
				return "", err
			}
		} else {
			switch fieldType.Kind() {
			case reflect.Slice, reflect.Array:
				values := parseDelimitedArray(true, valType, ",")
				for _, v := range values {
					if err := writer.WriteField(tag.Name+"[]", v); err != nil {
						writer.Close()
						return "", err
					}
				}
			default:
				if err := writer.WriteField(tag.Name, valToString(valType.Interface())); err != nil {
					writer.Close()
					return "", err
				}
			}
		}
	}

	if err := writer.Close(); err != nil {
		return "", err
	}

	return writer.FormDataContentType(), nil
}

func encodeMultipartFormDataFile(w *multipart.Writer, fieldType reflect.Type, valType reflect.Value) error {
	if fieldType.Kind() != reflect.Struct {
		return fmt.Errorf("invalid type %s for multipart/form-data file", valType.Type())
	}

	var fieldName string
	var fileName string
	var reader io.Reader

	for i := 0; i < fieldType.NumField(); i++ {
		field := fieldType.Field(i)
		val := valType.Field(i)

		tag := parseMultipartFormTag(field)
		if !tag.Content && tag.Name == "" {
			continue
		}

		if tag.Content && val.CanInterface() {
			if reflect.TypeOf(val.Interface()) == reflect.TypeOf([]byte(nil)) {
				reader = bytes.NewReader(val.Interface().([]byte))
			} else if reflect.TypeOf(val.Interface()).Implements(reflect.TypeOf((*io.Reader)(nil)).Elem()) {
				reader = val.Interface().(io.Reader)
			}
		} else {
			fieldName = tag.Name
			fileName = val.String()
		}
	}

	if fieldName == "" || fileName == "" || reader == nil {
		return fmt.Errorf("invalid multipart/form-data file")
	}

	fw, err := w.CreateFormFile(fieldName, fileName)
	if err != nil {
		return err
	}
	if _, err := io.Copy(fw, reader); err != nil {
		return err
	}

	return nil
}

func encodeFormData(fieldName string, w io.Writer, data interface{}) error {
	requestType := reflect.TypeOf(data)
	requestValType := reflect.ValueOf(data)

	if requestType.Kind() == reflect.Pointer {
		requestType = requestType.Elem()
		requestValType = requestValType.Elem()
	}

	dataValues := url.Values{}

	switch requestType.Kind() {
	case reflect.Struct:
		for i := 0; i < requestType.NumField(); i++ {
			field := requestType.Field(i)
			fieldType := field.Type
			valType := requestValType.Field(i)

			if isNil(fieldType, valType) {
				continue
			}

			if fieldType.Kind() == reflect.Pointer {
				fieldType = fieldType.Elem()
				valType = valType.Elem()
			}

			tag := parseFormTag(field)
			if tag.JSON {
				data, err := MarshalJSON(valType.Interface(), field.Tag, true)
				if err != nil {
					return err
				}
				dataValues.Set(tag.Name, string(data))
			} else {
				switch tag.Style {
				// TODO: support other styles
				case "form":
					values := populateForm(tag.Name, tag.Explode, fieldType, valType, ",", func(sf reflect.StructField) string {
						tag := parseFormTag(field)
						if tag == nil {
							return ""
						}

						return tag.Name
					})
					for k, v := range values {
						for _, vv := range v {
							dataValues.Add(k, vv)
						}
					}
				}
			}
		}
	case reflect.Map:
		for _, k := range requestValType.MapKeys() {
			v := requestValType.MapIndex(k)
			dataValues.Set(fmt.Sprintf("%v", k.Interface()), valToString(v.Interface()))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < requestValType.Len(); i++ {
			v := requestValType.Index(i)
			dataValues.Set(fieldName, valToString(v.Interface()))
		}
	}

	if _, err := w.Write([]byte(dataValues.Encode())); err != nil {
		return err
	}

	return nil
}

type requestTag struct {
	MediaType string
}

func getRequestTag(field reflect.StructField) *requestTag {
	// example `request:"mediaType=multipart/form-data"`
	values := parseStructTag(requestTagKey, field)
	if values == nil {
		return nil
	}

	tag := &requestTag{
		MediaType: "application/octet-stream",
	}

	for k, v := range values {
		switch k {
		case "mediaType":
			tag.MediaType = v
		}
	}

	return tag
}

type multipartFormTag struct {
	File    bool
	Content bool
	JSON    bool
	Name    string
}

func parseMultipartFormTag(field reflect.StructField) *multipartFormTag {
	// example `multipartForm:"name=file"`
	values := parseStructTag(multipartFormTagKey, field)

	tag := &multipartFormTag{}

	for k, v := range values {
		switch k {
		case "file":
			tag.File = v == "true"
		case "content":
			tag.Content = v == "true"
		case "name":
			tag.Name = v
		case "json":
			tag.JSON = v == "true"
		}
	}

	return tag
}

type formTag struct {
	Name    string
	JSON    bool
	Style   string
	Explode bool
}

func parseFormTag(field reflect.StructField) *formTag {
	// example `form:"name=propName,style=spaceDelimited,explode"`
	values := parseStructTag(formTagKey, field)

	tag := &formTag{
		Style:   "form",
		Explode: true,
	}

	for k, v := range values {
		switch k {
		case "name":
			tag.Name = v
		case "json":
			tag.JSON = v == "true"
		case "style":
			tag.Style = v
		case "explode":
			tag.Explode = v == "true"
		}
	}

	return tag
}
