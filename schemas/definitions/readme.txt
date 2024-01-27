/home/cato/Downloads/jsonschema/go-jsonschema --only-models -p WorkspaceAddOns --resolve-extension json -t -v *.json


Error
calendar_event_object_schema.json:          "$ref": "#definitions/conferenceData"
drive_event_object_schema.json:          "$ref": "#definitions/driveItemMetaData"
drive_event_object_schema.json:            "$ref": "#definitions/driveItemMetaData"
gmailSchema.json:          "$ref": "#definitions/updateDraftActionMarkup"
gmailSchema.json:          "$ref": "#definitions/openCreatedDraftActionMarkup"
gmailSchema.json:          "$ref": "#definitions/updateSubject"
gmailSchema.json:            "$ref": "#definitions/insertContent"


Missing
conferenceDataMarkupSchema.json


Error
o-jsonschema: Loading docs_event_object_schema.json
go-jsonschema: Failed: error parsing from file docs_event_object_schema.json: error parsing JSON file docs_event_object_schema.json: failed to unmarshal JSON: invalid character '"' after object key:value pair


Errors
go-jsonschema: Warning: Multiple types map to the name "AddOnParameters"; declaring duplicate as "AddOnParameters_1" instead
go-jsonschema: Warning: Multiple types map to the name "AddOnParametersParameters"; declaring duplicate as "AddOnParametersParameters_1" instead
go-jsonschema: Warning: Multiple types map to the name "Attendee"; declaring duplicate as "Attendee_1" instead
go-jsonschema: Warning: Multiple types map to the name "CalendarEventObject"; declaring duplicate as "CalendarEventObject_1" instead
go-jsonschema: Warning: Multiple types map to the name "CalendarEventObjectCapabilities"; declaring duplicate as "CalendarEventObjectCapabilities_1" instead
go-jsonschema: Warning: Multiple types map to the name "ConferenceData"; declaring duplicate as "ConferenceData_1" instead
go-jsonschema: Warning: Multiple types map to the name "ConferenceSolution"; declaring duplicate as "ConferenceSolution_1" instead
go-jsonschema: Warning: Multiple types map to the name "ConferenceSolutionKey"; declaring duplicate as "ConferenceSolutionKey_1" instead
go-jsonschema: Warning: Multiple types map to the name "EntryPoint"; declaring duplicate as "EntryPoint_1" instead
go-jsonschema: Warning: Multiple types map to the name "Parameters"; declaring duplicate as "Parameters_1" instead
go-jsonschema: Warning: Multiple types map to the name "CalendarEventObjectOrganizer"; declaring duplicate as "CalendarEventObjectOrganizer_1" instead
go-jsonschema: Warning: Multiple types map to the name "DocsEventObject"; declaring duplicate as "DocsEventObject_1" instead
go-jsonschema: Warning: Multiple types map to the name "DocsEventObjectMatchedUrl"; declaring duplicate as "DocsEventObjectMatchedUrl_1" instead
go-jsonschema: Warning: Multiple types map to the name "DriveEventObject"; declaring duplicate as "DriveEventObject_1" instead
go-jsonschema: Warning: Multiple types map to the name "DriveItemMetaData"; declaring duplicate as "DriveItemMetaData_1" instead
go-jsonschema: Failed: could not generate type for field "sheets": could not follow $ref "./sheets_event_object_schema.json#/definitions/sheetsEventObject" to file "./sheets_event_object_schema.json": error parsing JSON file sheets_event_object_schema.json: failed to unmarshal JSON: invalid character '}' looking for beginning of object key string

