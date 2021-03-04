package timepoints

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	utilserver "github.com/ldsec/medco/connector/util/server"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const sqlDateFormat = "2006-01-02"

func startEvent(patientList []int64, conceptCodes, modifierCodes []string, earliest bool) (map[int64]time.Time, map[int64]struct{}, error) {

	setStrings := make([]string, len(patientList))

	for i, patient := range patientList {
		setStrings[i] = strconv.FormatInt(patient, 10)
	}
	setDefinition := "{" + strings.Join(setStrings, ",") + "}"
	conceptDefinition := "{" + strings.Join(conceptCodes, ",") + "}"
	modifierDefinition := "{" + strings.Join(modifierCodes, ",") + "}"

	description := fmt.Sprintf("get start event (patient list: %s, start concept codes: %s, start modifier codes: %s, begins with earliest occurence: %t): procedure: %s",
		setDefinition, conceptDefinition, modifierDefinition, earliest, "i2b2demodata_i2b2.start_event")

	logrus.Debugf("Retrieving the start event dates for the patients: %s", description)
	row, err := utilserver.I2B2DBConnection.Query("SELECT i2b2demodata_i2b2.start_event($1,$2,$3,$4)", setDefinition, conceptDefinition, modifierDefinition, earliest)
	if err != nil {
		err = errors.Errorf("while calling database for retrieving start event dates: %s; DB operation: %s", err.Error(), description)
		return nil, nil, err
	}

	patientsWithStartEvent := make(map[int64]time.Time, len(patientList))
	patientsWithoutStartEvent := make(map[int64]struct{}, len(patientList))

	var record = new(string)
	for row.Next() {
		err = row.Scan(record)
		if err != nil {
			err = errors.Errorf("while reading database record stream for retrieving start event dates: %s; DB operation: %s", err.Error(), description)
			return nil, nil, err
		}

		recordEntries := strings.Split(strings.Trim(*record, "()"), ",")
		if len(recordEntries) != 2 {
			err = errors.Errorf("while parsing SQL record stream: expected to find 2 items in a string like \"(<integer>,<date>)\" in record %s", *record)
		}
		patientID, err := strconv.ParseInt(recordEntries[0], 10, 64)
		if err != nil {
			err = errors.Errorf("while parsing patient number \"%s\": %s; DB operation: %s", recordEntries[0], err.Error(), description)
			return nil, nil, err
		}
		startDate, err := time.Parse(sqlDateFormat, recordEntries[1])
		if err != nil {
			err = errors.Errorf("while parsing patient number \"%s\": %s; DB operation: %s", recordEntries[1], err.Error(), description)
			return nil, nil, err
		}

		if _, isIn := patientsWithStartEvent[patientID]; isIn {
			err = errors.Errorf("while filling patient-to-start-date map: patient %d already found in map, this is not expected; DB operation: %s", patientID, description)
			return nil, nil, err
		}

		patientsWithStartEvent[patientID] = startDate

	}
	for _, patientNumber := range patientList {
		if _, isIn := patientsWithStartEvent[patientNumber]; !isIn {
			patientsWithoutStartEvent[patientNumber] = struct{}{}
		}
	}

	logrus.Debugf("Successfully found %d patients with start event; DB operation: %s", len(patientsWithStartEvent), description)
	return patientsWithStartEvent, patientsWithoutStartEvent, nil

}

func endEvents(patientWithStartEventList map[int64]time.Time, conceptCodes, modifierCodes []string) (map[int64][]time.Time, error) {
	setStrings := make([]string, 0, len(patientWithStartEventList))

	for patient := range patientWithStartEventList {
		setStrings = append(setStrings, strconv.FormatInt(patient, 10))
	}
	setDefinition := "{" + strings.Join(setStrings, ",") + "}"
	conceptDefinition := "{" + strings.Join(conceptCodes, ",") + "}"
	modifierDefinition := "{" + strings.Join(modifierCodes, ",") + "}"

	description := fmt.Sprintf("get start event (patient list: %s, end concept codes: %s, end modifier codes: %s): procedure: %s",
		setDefinition, conceptDefinition, modifierDefinition, "i2b2demodata_i2b2.end_events")

	logrus.Debugf("Retrieving the end event dates for the patients: %s", description)
	row, err := utilserver.I2B2DBConnection.Query("SELECT i2b2demodata_i2b2.end_events($1,$2,$3)", setDefinition, conceptDefinition, modifierDefinition)
	if err != nil {
		err = errors.Errorf("while calling database for retrieving end event dates: %s; DB operation: %s", err.Error(), description)
		return nil, err
	}

	patientsWithEndEvent := make(map[int64][]time.Time, len(patientWithStartEventList))

	var record = new(string)
	for row.Next() {
		err = row.Scan(record)
		if err != nil {
			err = errors.Errorf("while reading database record stream for retrieving start event dates: %s; DB operation: %s", err.Error(), description)
			return nil, err
		}

		recordEntries := strings.Split(strings.Trim(*record, "()"), ",")
		if len(recordEntries) != 2 {
			err = errors.Errorf("while parsing SQL record stream: expected to find 2 items in a string like \"(<integer>,<date>)\" in record %s", *record)
		}
		patientID, err := strconv.ParseInt(recordEntries[0], 10, 64)
		if err != nil {
			err = errors.Errorf("while parsing patient number \"%s\": %s; DB operation: %s", recordEntries[0], err.Error(), description)
			return nil, err
		}
		endDate, err := time.Parse(sqlDateFormat, recordEntries[1])
		if err != nil {
			err = errors.Errorf("while parsing end date \"%s\": %s; DB operation: %s", recordEntries[1], err.Error(), description)
			return nil, err
		}

		if patientWithStartEventList[patientID].Before(endDate) {

			// here, an aggregate was not performed, so it is expected to find a patient ID more than once

			if _, isIn := patientsWithEndEvent[patientID]; isIn {
				patientsWithEndEvent[patientID] = append(patientsWithEndEvent[patientID], endDate)
			} else {
				patientsWithEndEvent[patientID] = []time.Time{endDate}
			}
		} else {
			logrus.Tracef("dropped end date: end date %s before start date %s; patientID: %d", endDate.Format(sqlDateFormat), patientWithStartEventList[patientID].Format(sqlDateFormat), patientID)
		}

	}
	logrus.Debugf("Successfully found %d patients with end event; DB operation: %s", len(patientsWithEndEvent), description)
	return patientsWithEndEvent, nil

}

func censoringEvents(patientWithStartEventList map[int64]time.Time, patientWithoutEndEvent map[int64]struct{}, endConceptCodes []string, endModifierCodes []string) (map[int64]time.Time, map[int64]struct{}, error) {
	setStrings := make([]string, 0, len(patientWithoutEndEvent))

	for patient := range patientWithoutEndEvent {
		setStrings = append(setStrings, strconv.FormatInt(patient, 10))
	}
	setDefinition := "{" + strings.Join(setStrings, ",") + "}"
	conceptDefinition := "{" + strings.Join(endConceptCodes, ",") + "}"
	modifierDefinition := "{" + strings.Join(endModifierCodes, ",") + "}"

	description := fmt.Sprintf("get start event (patient list: %s, start concept codes: %s, start modifier codes: %s): procedure: %s",
		setDefinition, conceptDefinition, modifierDefinition, "i2b2demodata_i2b2.censoring_event")

	logrus.Debugf("Retrieving the censoring event dates for the patients: %s", description)
	row, err := utilserver.I2B2DBConnection.Query("SELECT i2b2demodata_i2b2.censoring_event($1,$2,$3)", setDefinition, conceptDefinition, modifierDefinition)
	if err != nil {
		err = errors.Errorf("while calling database for retrieving right censoring event dates: %s; DB operation: %s", err.Error(), description)
		return nil, nil, err
	}

	patientsWithCensoringEvent := make(map[int64]time.Time, len(patientWithoutEndEvent))
	patientsWithoutCensoringEvent := make(map[int64]struct{}, len(patientWithoutEndEvent))

	var record = new(string)
	for row.Next() {
		err = row.Scan(record)
		if err != nil {
			err = errors.Errorf("while reading database record stream for retrieving start event dates: %s; DB operation: %s", err.Error(), description)
			return nil, nil, err
		}

		recordEntries := strings.Split(strings.Trim(*record, "()"), ",")
		if len(recordEntries) != 2 {
			err = errors.Errorf("while parsing SQL record stream: expected to find 2 items in a string like \"(<integer>,<date>)\" in record %s", *record)
		}
		patientID, err := strconv.ParseInt(recordEntries[0], 10, 64)
		if err != nil {
			err = errors.Errorf("while parsing patient number \"%s\": %s; DB operation: %s", recordEntries[0], err.Error(), description)
			return nil, nil, err
		}
		censoringDate, err := time.Parse(sqlDateFormat, recordEntries[1])
		if err != nil {
			err = errors.Errorf("while parsing patient number \"%s\": %s; DB operation: %s", recordEntries[1], err.Error(), description)
			return nil, nil, err
		}

		if _, ok := patientWithStartEventList[patientID]; !ok {
			err = errors.Errorf("while looking for a start date patient %d was not found in sart event map, this is not expected; DB operation: %s", patientID, description)
			return nil, nil, err
		}

		if patientWithStartEventList[patientID].Before(censoringDate) {

			if _, isIn := patientsWithCensoringEvent[patientID]; isIn {
				err = errors.Errorf("while filling patient-to-censoring-date map: patient %d already found in map, this is not expected; DB operation: %s", patientID, description)
				return nil, nil, err
			}

			patientsWithCensoringEvent[patientID] = censoringDate
		} else {
			if _, isIn := patientsWithoutCensoringEvent[patientID]; isIn {
				err = errors.Errorf("while filling patient-without-censoring set: patient %d already found in set, this is not expected; DB operation: %s", patientID, description)
				return nil, nil, err
			}
			patientsWithoutCensoringEvent[patientID] = struct{}{}
		}

	}

	logrus.Debugf("Successfully found %d patients with right censoring event; DB operation: %s", len(patientsWithCensoringEvent), description)
	return patientsWithCensoringEvent, patientsWithoutCensoringEvent, nil
}