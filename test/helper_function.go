package test

import (
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/testing"
)

func GetRetentionPeriod(t testing.TestingT, dbInstanceID string, awsRegion string) int {
	backup, err := GetRetentionPeriodE(t, dbInstanceID, awsRegion)
	if err != nil {
		t.Fatal(err)
	}

	return int(backup)
}

func GetRetentionPeriodE(t testing.TestingT, dbInstanceID string, awsRegion string) (int64, error) {
	dbInstance, err := aws.GetRdsInstanceDetailsE(t, dbInstanceID, awsRegion)
	if err != nil {
		return -1, err
	}

	return *dbInstance.BackupRetentionPeriod, nil
}
