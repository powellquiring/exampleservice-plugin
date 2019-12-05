/**
 * (C) Copyright IBM Corp. 2019.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */


package utils

import (
  // do the table
  "reflect"
  "fmt"
  "strings"
  "strconv"
  "github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
  "github.com/go-openapi/strfmt"
)

// returns true if the value is an "object" type that
// has exactky one property that is an array
func hasExactlyOneArrayProperty(thing reflect.Value) bool {
  // three potential paths here:
  // 1. its a struct
  // 2. its a map
  // 3. it is a type that can't have fields, so return false
  numArrayProps := 0
  if thing.Kind() == reflect.Struct {
    for i := 0; i < thing.NumField(); i++ {
      // !!! doesnt handle pointers to slices, might want to deref
      if thing.Field(i).Kind() == reflect.Slice {
        numArrayProps += 1
      }
    }
  } else if thing.Kind() == reflect.Map {
    iter := thing.MapRange()
    for iter.Next() {
      if iter.Value().Kind() == reflect.Slice {
        numArrayProps += 1
      }
    }
  }

  return numArrayProps == 1
}

// return the dereferenced value if a pointer or interface,
// hand the value back if not
func derefValue(thing reflect.Value) reflect.Value {
  if thing.Kind() == reflect.Interface {
    // interface elements can be pointers
    return derefValue(thing.Elem())
  } else if thing.Kind() == reflect.Ptr {
    return thing.Elem()
  } else {
    return thing
  }
}

// takes the final value that is to be written to the table
// and formats it as a string if possible
func getStringValue(thing reflect.Value) string {
  var result string

  // don't bother with invalid values
  if !thing.IsValid() {
    return "-"
  }

  actualValue := thing.Interface()
  switch thing.Kind() {
  case reflect.String:
    result = thing.String()

  case reflect.Bool:
    result = strconv.FormatBool(actualValue.(bool))

  case reflect.Int64:
    result = strconv.FormatInt(actualValue.(int64), 10)

  case reflect.Float32:
    // FormatFloat must take a float64 as its first value, so typecast is needed
    result = strconv.FormatFloat(float64(actualValue.(float32)), 'g', -1, 32)

  case reflect.Float64:
    result = strconv.FormatFloat(actualValue.(float64), 'g', -1, 64)

  case reflect.Struct:
    if thing.Type().String() == "strfmt.DateTime" {
      result = actualValue.(strfmt.DateTime).String()
    } else {
      // dont display nested objects as table elements
      result = "<Nested Object>"
    }

  case reflect.Slice:
    // print something if an array was returned but is hidden
    // to indicate that there is data there
    if thing.Len() > 0 {
      result = "<Array>"
    } else {
      result = "-"
    }

  default:
    fmt.Println("Type not yet supported: " + thing.Kind().String())
    result = "-"
  }

  return result
}

func addFieldToRow(field reflect.Value, fieldName string, row []string) []string {
  // getting rid of URL strings for now, they clutter up table output
  if strings.ToLower(fieldName) != "url" {
    row = append(row, getStringValue(field))
  } else {
    row = append(row, "-")
  }

  return row
}

/**********  STRUCT STUFF  **********/

// get the property names of a struct
func getTableHeadersFromStruct(thing reflect.Value) []string {
  tableHeaders := make([]string, 0)

  // loop through struct fields
  for i := 0; i < thing.NumField(); i++ {
    fieldName := thing.Type().Field(i).Name

    tableHeaders = append(tableHeaders, fieldName) // this will work fine if Field(i) is a pointer
  }

  return tableHeaders
}

func getValuesFromStruct(thing reflect.Value) []string {
  rowValues := make([]string, 0)

  for i := 0; i < thing.NumField(); i++ {
    // this will almost always be a pointer
    field := derefValue(thing.Field(i))
    fieldName := thing.Type().Field(i).Name

    rowValues = addFieldToRow(field, fieldName, rowValues)
  }

  return rowValues
}

/**********  MAP STUFF  **********/

// get the property names of a map
func getTableHeadersFromMap(mapThing reflect.Value) []string {
  tableHeaders := make([]string, 0)

  // loop through map fields
  iter := mapThing.MapRange()
  for iter.Next() {
    tableHeaders = append(tableHeaders, iter.Key().String())
  }

  return tableHeaders
}

func getValuesFromMap(mapThing reflect.Value) []string {
  rowValues := make([]string, 0)

  iter := mapThing.MapRange()
  for iter.Next() {
    // this will almost always be a pointer
    field := derefValue(iter.Value())
    rowValues = addFieldToRow(field, iter.Key().String(), rowValues)
  }

  return rowValues
}

/**********  SINGLE ARRAY STUFF - ONLY STRUCTS  **********/

// for the case when a struct has a single array property.
// this is the common pattern for list operations
// get the keys of the struct to be used as table headers
// if not a struct, use the last segment of the jmespath
// note: the table headers returned here will also include
// the non-array top-level fields on the struct
func getExplodedTableHeaders(thing reflect.Value, lastQuerySegment string) []string {
  tableHeaders := make([]string, 0)
  var theArrayProperty reflect.Value
  var explodedArrayHeaders []string

    // !!! assuming thing is a struct
    for i := 0; i < thing.NumField(); i++ {
      field := derefValue(thing.Field(i))

      if (field.Kind() == reflect.Slice) {
        // this should only happen once
        theArrayProperty = field;
      } else {
        // add non-array headers first
        fieldName := thing.Type().Field(i).Name
        tableHeaders = append(tableHeaders, fieldName)
      }
    }

    // check the type of the items of the array
    switch theArrayProperty.Type().Elem().Kind() {
    case reflect.Struct:
      explodedArrayHeaders = getTableHeadersFromStruct(derefValue(reflect.New(theArrayProperty.Type().Elem())))

    case reflect.Map:
      explodedArrayHeaders = getTableHeadersFromMap(derefValue(reflect.New(theArrayProperty.Type().Elem())))

    default:
      // probably a list of primitives, etc. so needs only only header
      explodedArrayHeaders = []string{lastQuerySegment}
    }

  tableHeaders = append(tableHeaders, explodedArrayHeaders...)

  return tableHeaders
}

// for the case when a struct has a single array property.
// this is the common pattern for list operations
// get the values for each pass of the array, merges them
// with the other non-array values at the top-level
// return as an array of rows represented by arrays of strings
func getExplodedArrayValues(thing reflect.Value) [][]string {
  allRows := make([][]string, 0)
  otherPropRows := make([]string, 0) // for the values other than the array

  var theArrayProperty reflect.Value

  // find the array property and get the other, non-array, top-level values
  // !!! assuming structs here
  for i := 0; i < thing.NumField(); i++ {
    field := derefValue(thing.Field(i))
    fieldName := thing.Type().Field(i).Name

    if (field.Kind() == reflect.Slice) {
      theArrayProperty = field;
    } else {
      otherPropRows = addFieldToRow(field, fieldName, otherPropRows)
    }
  }

  // cycle through the array and get the values for each element
  // if a struct, will add multiple columns, if not, will add just one
  // !!! we aren't handling maps yet

  rowValues := make([]string, 0)
  for i := 0; i < theArrayProperty.Len(); i++ {
    if theArrayProperty.Type().Elem().Kind() == reflect.Struct {
      // !!! assuming the elements are only structs right now
      structElement := derefValue(theArrayProperty.Index(i))

      // then, collect the values into rows and add them to the table
      // making sure to merge them with the other top level property values
      rowValues = getValuesFromStruct(structElement)
      rowValues = append(otherPropRows, rowValues...)
    } else {
      // assuming this is a primitive and not a map or interface
      rowValues = append(otherPropRows, getStringValue(theArrayProperty.Index(i)))
    }

    allRows = append(allRows, rowValues)
    rowValues = nil // clear the array
  }

  // if array has no values, we still want to print the other top level values
  if theArrayProperty.Len() == 0 {
    rowValues = append(rowValues, otherPropRows...)
    allRows = append(allRows, rowValues)
  }

  return allRows
}

func getLastQuerySegment(query string) string {
  queryArr := strings.Split(query, ".")
  return queryArr[len(queryArr) - 1]
}

func DoTheTable(result interface{}, jmesQuery string) {
  // !!! we could potentially pass this around instead of redefining every time
  ui = terminal.NewStdUI()

  // get last segment of jmes query in case it needs to be
  // used as a column header
  lastQuerySegment := getLastQuerySegment(jmesQuery)

  resultValue := derefValue(reflect.ValueOf(result))

  // if nothing is passed in, there's nothing to do
  if (!resultValue.IsValid()) {
    ui.Say("Nothing to show.")
    return
  }

  var table terminal.Table

  if (resultValue.Kind() == reflect.Slice || resultValue.Kind() == reflect.Array) {
    /********** ARRAY CASE **********/
    arrayElementType := resultValue.Type().Elem().Kind()
    for i := 0; i < resultValue.Len(); i++ {
      switch arrayElementType {
      case reflect.Struct:
        if (i == 0) {
          tableHeaders := getTableHeadersFromStruct(derefValue(resultValue.Index(0)))
          table = ui.Table(tableHeaders)
        }

        structElement := derefValue(resultValue.Index(i))
        rowValues := getValuesFromStruct(structElement)

        table.Add(rowValues...)
        rowValues = nil // clear the array

      case reflect.Map:
        if (i == 0) {
          tableHeaders := getTableHeadersFromMap(derefValue(resultValue.Index(0)))
          table = ui.Table(tableHeaders)
        }

        mapElement := derefValue(resultValue.Index(i))
        rowValues := getValuesFromMap(mapElement)

        table.Add(rowValues...)
        rowValues = nil // clear the array


      // !!! this can be an issue when the type is just "interface"
      // theoretically, we should be derefing it and running it though
      // everything again
      default:
        if i == 0 {
          table = ui.Table([]string{lastQuerySegment})
        }
        listElement := derefValue(resultValue.Index(i).Elem())
        table.Add(getStringValue(listElement))
      }
    }

    if resultValue.Len() != 0 {
      table.Print()
    } else {
      ui.Say("Nothing to show.")
    }

  } else if (hasExactlyOneArrayProperty(resultValue)) {
    /********** SINGLE ARRAY PROPERTY CASE **********/
    // attach the non array elements and print the rest
    // 1. grab the keys of the non-array objects, add to table headers
    // 2. reserve the values of the top-level non-array props, they will be added to each row
    // 3. go through the array property
    //   3a. If primitive, add jmes seg to table headers
    //   3b. If struct, add all prop names to table headers
    // 4. Add rows for each element of the array prop, also remembering to add the other top level guys

    // !!! bad assumption that it's a struct while i work out the code
    tableHeaders := getExplodedTableHeaders(resultValue, lastQuerySegment)
    table := ui.Table(tableHeaders)

    allRows := getExplodedArrayValues(resultValue)
    for _, row := range allRows {
      table.Add(row...)
    }

    table.Print()
  } else if resultValue.Kind() == reflect.Struct {
    /********** STRUCT CASE **********/
    tableHeaders := getTableHeadersFromStruct(resultValue)
    rowValues := getValuesFromStruct(resultValue)

    table := ui.Table(tableHeaders)
    table.Add(rowValues...)
    table.Print()
  } else if resultValue.Kind() == reflect.Map {
    /********** MAP CASE **********/
    tableHeaders := getTableHeadersFromMap(resultValue)
    rowValues := getValuesFromMap(resultValue)

    table := ui.Table(tableHeaders)
    table.Add(rowValues...)
    table.Print()
  } else {
    /********** BASE CASE - TREAT AS SINGLE VALUE **********/
    table = ui.Table([]string{lastQuerySegment})
    singleValue := derefValue(resultValue)
    table.Add(getStringValue(singleValue))
    table.Print()
  }
}
