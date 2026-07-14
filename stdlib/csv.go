package stdlib

import (
	"encoding/csv"
	"errors"
	"strings"
	"github.com/2dprototype/tender"
)

// CSVModule exports the functions for encoding/decoding CSV strings.
var csvModule = map[string]tender.Object{
	"decode": &tender.NativeFunction{Name: "decode", Value: csvDecode},
	"encode": &tender.NativeFunction{Name: "encode", Value: csvEncode},
}

// csvDecode decodes a CSV string into a array of arrays (rows and columns).
func csvDecode(args ...tender.Object) (tender.Object, error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	csvString, ok := tender.ToString(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "1st",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}

	reader := csv.NewReader(strings.NewReader(csvString))
	records, err := reader.ReadAll()
	if err != nil {
		return wrapError(err), nil
	}

	// Convert to tender-compatible format (array of arrays).
	var tenderRecords tender.Array
	for _, record := range records {
		var tenderRow tender.Array
		for _, field := range record {
			tenderRow.Value = append(tenderRow.Value, &tender.String{Value: field})
		}
		tenderRecords.Value = append(tenderRecords.Value, &tender.Array{Value: tenderRow.Value})
	}

	return &tender.Array{Value: tenderRecords.Value}, nil
}

// csvEncode encodes a array of arrays (rows and columns) into a CSV string.
func csvEncode(args ...tender.Object) (tender.Object, error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	data, ok := args[0].(*tender.Array)
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "1st",
			Expected: "array",
			Found:    args[0].TypeName(),
		}
	}

	var sb strings.Builder
	writer := csv.NewWriter(&sb)

	for _, row := range data.Value {
		rowarray, ok := row.(*tender.Array)
		if !ok {
			return wrapError(errors.New("data should be a array of arrays")), nil
		}

		var record []string
		for _, field := range rowarray.Value {
			fieldStr, ok := tender.ToString(field)
			if !ok {
				return wrapError(errors.New("all fields must be strings")), nil
			}
			record = append(record, fieldStr)
		}

		if err := writer.Write(record); err != nil {
			return wrapError(err), nil
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return wrapError(err), nil
	}

	return &tender.String{Value: sb.String()}, nil
}
