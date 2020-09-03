package survivalserver

import (
	"database/sql"
	"os/exec"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

/**
*  TestSql is aimed at unit tests run at development time. It is not expected to succeed in CI or other automated run env
 */
func TestSql(t *testing.T) {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=i2b2medcosrv0 sslmode=disable")
	logrus.Info(err)
	if err != nil {
		t.Error("Error at creating database", err)
	}
	err = db.Ping()
	if err != nil {
		t.Skip("Unable to connect database for testing", err)
	}

	res, err := db.Exec("CREATE SCHEMA survival_test")
	logrus.Info(res, err)
	path, err := exec.LookPath("./load_test_data.sh")
	if err != nil {
		t.Error("Error when finding ./load_test_data.sh", err)
	}
	logrus.Info("Found " + path)
	command := exec.Command("./load_test_data.sh")
	err = command.Start()
	if err != nil {
		t.Error("Error at starting command", err)
	}
	logrus.Info("Started data loading")
	err = command.Wait()
	if err != nil {
		t.Error("Command terminated with error", err)
	}
	count := new(int)
	err = db.QueryRow(`SELECT COUNT(*) FROM survival_test.observation_fact`).Scan(count)
	if err != nil {
		t.Error(err)
	}
	logrus.Infof("Found %d rows", *count)
	if *count != 393 {
		t.Errorf("Rows number in test observation_fact failed. Expected 393, got %d", *count)
	}
	patient := new(int)
	date := new(string)
	rows, err := db.Query(testSql1, "NDC:00015345620", "@", "{483573, 483574, 483575}")
	if err != nil {
		t.Error(err)
	}
	for rows.Next() {
		rows.Scan(patient, date)
		logrus.Info(*patient, *date)

	}
	timePoint := new(int)
	rows, err = db.Query(testSql3, "NDC:00015345620", "@", "{483573, 483574, 483575}", "DEM|DATE:death", "@")
	if err != nil {
		t.Error(err)
	}
	for rows.Next() {
		rows.Scan(timePoint, count)
		logrus.Info(*timePoint, *count)

	}

	rows, err = db.Query(testSql5, "NDC:00015345620", "@", "{483573, 483574, 483575}", "DEM|DATE:death", "@")
	if err != nil {
		t.Error(err)
	}
	for rows.Next() {
		rows.Scan(timePoint, count)
		logrus.Info(*timePoint, *count)
	}

	eventCount := count
	censoringCount := new(int)

	rows, err = db.Query(testSql6, "NDC:00015345620", "@", "{483573, 483574, 483575}", "DEM|DATE:death", "@")
	if err != nil {
		t.Error(err)
	}

	for rows.Next() {
		rows.Scan(timePoint, eventCount, censoringCount)
		logrus.Info(*timePoint, *eventCount, *censoringCount)
	}

	res, err = db.Exec("DROP SCHEMA survival_test CASCADE")
	logrus.Info(res, err)

	t.Skip("Unable to connect database", err)
}

func TestBuildTimePoint(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=i2b2medcosrv0 sslmode=disable")
	logrus.Info(err)
	if err != nil {
		t.Error("Error at creating database connection", err)
	}
	err = db.Ping()
	if err != nil {
		t.Skip("Unable to connect database for testing", err)
	}
	timePoint, err := BuildTimePoints(db, bigList,
		`A168`,
		"start_date", "@", `A125`, "start_date", "126:1")
	if err != nil {
		t.Error("Test Failed", err)
	}
	logrus.Info(timePoint)
}

func TestGetPatient(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=medcoconnectorsrv0 sslmode=disable")
	logrus.Info(err)
	if err != nil {
		t.Error("Error at creating database connection", err)
	}
	err = db.Ping()
	if err != nil {
		t.Skip("Unable to connect database for testing", err)
	}
	list, err := GetPatientList(db, int64(-1), "test")
	if err != nil {
		t.Error(err)
	}
	if len(list) != 228 {
		t.Errorf("expected number of patients: 228. Got %d", len(list))
	}
}

func TestBackSlashFormat(t *testing.T) {
	res := backSlashFormat([][]string{{"/TABLE/Item/"}})
	if res[0][0] != `\\TABLE\Item\` {
		t.Error(`Expected \\TABLE\Item\, got ` + res[0][0])
	}

}

func TestExpansion(t *testing.T) {
	res := Expansion(timePoints, 15)
	if reflect.DeepEqual(res, timePointsExpanded) {
		t.Error("The expansion list is not correct", res)
	}

}

const testSql1 string = `
SELECT patient_num,start_date 
FROM survival_test.observation_fact
WHERE concept_cd = $1 and modifier_cd = $2 and patient_num = ANY($3::integer[])
`
const testSql2 string = `
SELECT patient_num,end_date
FROM survival_test.observation_fact
WHERE concept_cd = $4 and modifier_cd = $5 and patient_num = ANY($3::integer[])
`
const testSql3 string = `
SELECT DATE_PART('day',end_date::timestamp - start_date::timestamp) AS timepoint, COUNT(*) AS event_count
FROM (` + testSql1 + `) AS x
INNER JOIN  (` + testSql2 + `) AS y
ON x.patient_num = y.patient_num
GROUP BY timepoint
`

const testSql4 string = `
SELECT patient_num, MAX(end_date) AS end_date
FROM survival_test.observation_fact
WHERE patient_num = ANY($2::integer[]) AND patient_num NOT IN (SELECT patient_num FROM (` + testSql2 + `) AS patients_with_events)
GROUP BY patient_num
`

const testSql5 string = `
SELECT DATE_PART('day', end_date::timestamp - start_date::timestamp) AS timepoint, COUNT(*) AS censoring_count
FROM (` + testSql4 + `) AS x
INNER JOIN  (` + testSql1 + `) AS y
ON x.patient_num = y.patient_num
GROUP BY timepoint
`

const testSql6 string = `
SELECT COALESCE(x.timepoint,y.timepoint) AS timepoint , COALESCE(event_count,0) AS event_count, COALESCE(censoring_count,0) AS censoring_count FROM (` + testSql3 + `) AS x  FULL JOIN (` + testSql5 + `) AS y
ON x.timepoint = y.timepoint
`

var bigList = []int64{1000001364, 1000001364, 1000001363, 1000001363, 1000001362, 1000001362, 1000001361, 1000001361, 1000001360, 1000001360, 1000001359, 1000001359, 1000001358, 1000001358, 1000001357, 1000001357, 1000001356, 1000001356, 1000001355, 1000001355, 1000001354, 1000001354, 1000001353, 1000001353, 1000001352, 1000001352, 1000001351, 1000001351, 1000001350, 1000001350, 1000001349, 1000001349, 1000001348, 1000001348, 1000001347, 1000001347, 1000001346, 1000001346, 1000001345, 1000001345, 1000001344, 1000001344, 1000001343, 1000001343, 1000001342, 1000001342, 1000001341, 1000001341, 1000001340, 1000001340, 1000001339, 1000001339, 1000001338, 1000001338, 1000001337, 1000001337, 1000001336, 1000001336, 1000001335, 1000001335, 1000001334, 1000001334, 1000001333, 1000001333, 1000001332, 1000001332, 1000001331, 1000001331, 1000001330, 1000001330, 1000001329, 1000001329, 1000001328, 1000001328, 1000001327, 1000001327, 1000001326, 1000001326, 1000001325, 1000001325, 1000001324, 1000001324, 1000001323, 1000001323, 1000001322, 1000001322, 1000001321, 1000001321, 1000001320, 1000001320, 1000001319, 1000001319, 1000001318, 1000001318, 1000001317, 1000001317, 1000001316, 1000001316, 1000001315, 1000001315, 1000001314, 1000001314, 1000001313, 1000001313, 1000001312, 1000001312, 1000001311, 1000001311, 1000001310, 1000001310, 1000001309, 1000001309, 1000001308, 1000001308, 1000001307, 1000001307, 1000001306, 1000001306, 1000001305, 1000001305, 1000001304, 1000001304, 1000001303, 1000001303, 1000001302, 1000001302, 1000001301, 1000001301, 1000001300, 1000001300, 1000001299, 1000001299, 1000001298, 1000001298, 1000001297, 1000001297, 1000001296, 1000001296, 1000001295, 1000001295, 1000001294, 1000001294, 1000001293, 1000001293, 1000001292, 1000001292, 1000001291, 1000001291, 1000001290, 1000001290, 1000001289, 1000001289, 1000001288, 1000001288, 1000001287, 1000001287, 1000001286, 1000001286, 1000001285, 1000001285, 1000001284, 1000001284, 1000001283, 1000001283, 1000001282, 1000001282, 1000001281, 1000001281, 1000001280, 1000001280, 1000001279, 1000001279, 1000001278, 1000001278, 1000001277, 1000001277, 1000001276, 1000001276, 1000001275, 1000001275, 1000001274, 1000001274, 1000001273, 1000001273, 1000001272, 1000001272, 1000001271, 1000001271, 1000001270, 1000001270, 1000001269, 1000001269, 1000001268, 1000001268, 1000001267, 1000001267, 1000001266, 1000001266, 1000001265, 1000001265, 1000001264, 1000001264, 1000001263, 1000001263, 1000001262, 1000001262, 1000001261, 1000001261, 1000001260, 1000001260, 1000001259, 1000001259, 1000001258, 1000001258, 1000001257, 1000001257, 1000001256, 1000001256, 1000001255, 1000001255, 1000001254, 1000001254, 1000001253, 1000001253, 1000001252, 1000001252, 1000001251, 1000001251, 1000001250, 1000001250, 1000001249, 1000001249, 1000001248, 1000001248, 1000001247, 1000001247, 1000001246, 1000001246, 1000001245, 1000001245, 1000001244, 1000001244, 1000001243, 1000001243, 1000001242, 1000001242, 1000001241, 1000001241, 1000001240, 1000001240, 1000001239, 1000001239, 1000001238, 1000001238, 1000001237, 1000001237, 1000001236, 1000001236, 1000001235, 1000001235, 1000001234, 1000001234, 1000001233, 1000001233, 1000001232, 1000001232, 1000001231, 1000001231, 1000001230, 1000001230, 1000001229, 1000001229, 1000001228, 1000001228, 1000001227, 1000001227, 1000001226, 1000001226, 1000001225, 1000001225, 1000001224, 1000001224, 1000001223, 1000001223, 1000001222, 1000001222, 1000001221, 1000001221, 1000001220, 1000001220, 1000001219, 1000001219, 1000001218, 1000001218, 1000001217, 1000001217, 1000001216, 1000001216, 1000001215, 1000001215, 1000001214, 1000001214, 1000001213, 1000001213, 1000001212, 1000001212, 1000001211, 1000001211, 1000001210, 1000001210, 1000001209, 1000001209, 1000001208, 1000001208, 1000001207, 1000001207, 1000001206, 1000001206, 1000001205, 1000001205, 1000001204, 1000001204, 1000001203, 1000001203, 1000001202, 1000001202, 1000001201, 1000001201, 1000001200, 1000001200, 1000001199, 1000001199, 1000001198, 1000001198, 1000001197, 1000001197, 1000001196, 1000001196, 1000001195, 1000001195, 1000001194, 1000001194, 1000001193, 1000001193, 1000001192, 1000001192, 1000001191, 1000001191, 1000001190, 1000001190, 1000001189, 1000001189, 1000001188, 1000001188, 1000001187, 1000001187, 1000001186, 1000001186, 1000001185, 1000001185, 1000001184, 1000001184, 1000001183, 1000001183, 1000001182, 1000001182, 1000001181, 1000001181, 1000001180, 1000001180, 1000001179, 1000001179, 1000001178, 1000001178, 1000001177, 1000001177, 1000001176, 1000001176, 1000001175, 1000001175, 1000001174, 1000001174, 1000001173, 1000001173, 1000001172, 1000001172, 1000001171, 1000001171, 1000001170, 1000001170, 1000001169, 1000001169, 1000001168, 1000001168, 1000001167, 1000001167, 1000001166, 1000001166, 1000001165, 1000001165, 1000001164, 1000001164, 1000001163, 1000001163, 1000001162, 1000001162, 1000001161, 1000001161, 1000001160, 1000001160, 1000001159, 1000001159, 1000001158, 1000001158, 1000001157, 1000001157, 1000001156, 1000001156, 1000001155, 1000001155, 1000001154, 1000001154, 1000001153, 1000001153, 1000001152, 1000001152, 1000001151, 1000001151, 1000001150, 1000001150, 1000001149, 1000001149, 1000001148, 1000001148, 1000001147, 1000001147, 1000001146, 1000001146, 1000001145, 1000001145, 1000001144, 1000001144, 1000001143, 1000001143, 1000001142, 1000001142, 1000001141, 1000001141, 1000001140, 1000001140, 1000001139, 1000001139, 1000001138, 1000001138, 1000001137, 1000001137}

var timePoints = []SqlTimePoint{{1, 1, 0}, {5, 1, 0}, {13, 2, 5}}
var timePointsExpanded = []SqlTimePoint{{0, 0, 0}, {1, 1, 0}, {2, 0, 0}, {3, 0, 0}, {4, 0, 0}, {5, 1, 0}, {4, 0, 0}, {6, 0, 0}, {7, 0, 0}, {8, 0, 0}, {9, 0, 0}, {10, 0, 0}, {11, 0, 0}, {12, 0, 0}, {13, 2, 5}, {14, 0, 0}}
