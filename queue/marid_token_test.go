package queue

var mockAssumeRoleResult1 = AssumeRoleResult{
	Credentials: Credentials{
		AccessKeyId:      "accessKeyId1",
		SecretAccessKey:  "secretAccessKey1",
		SessionToken:     "sessionToken1",
		ExpireTimeMillis: 123456789123,
	},
}

var mockAssumeRoleResult2 = AssumeRoleResult{
	Credentials: Credentials{
		AccessKeyId:      "accessKeyId2",
		SecretAccessKey:  "secretAccessKey2",
		SessionToken:     "sessionToken2",
		ExpireTimeMillis: 123456789123,
	},
}

var mockQueueUrl1 = "https://sqs.us-west-2.amazonaws.com/255452344566/oec-test-1-2"
var mockQueueUrl2 = "https://sqs.us-east-2.amazonaws.com/255452344566/oec-test-1-2"

var mockQueueConf1 = QueueConfiguration{
	SuccessRefreshPeriodInSeconds: 60,
	ErrorRefreshPeriodInSeconds:   60,
	Region:                        "us-west-2",
	QueueUrl:                      mockQueueUrl1,
}

var mockQueueConf2 = QueueConfiguration{
	SuccessRefreshPeriodInSeconds: 60,
	ErrorRefreshPeriodInSeconds:   60,
	Region:                        "us-east-2",
	QueueUrl:                      mockQueueUrl2,
}

var mockOECMetadata1 = OECMetadata{
	AssumeRoleResult:   mockAssumeRoleResult1,
	QueueConfiguration: mockQueueConf1,
}

var mockOECMetadata2 = OECMetadata{
	AssumeRoleResult:   mockAssumeRoleResult2,
	QueueConfiguration: mockQueueConf2,
}

var mockOECMetadataWithEmptyAssumeRoleResult1 = OECMetadata{
	AssumeRoleResult:   AssumeRoleResult{},
	QueueConfiguration: mockQueueConf1,
}

var mockOECMetadataWithEmptyAssumeRoleResult2 = OECMetadata{
	AssumeRoleResult:   AssumeRoleResult{},
	QueueConfiguration: mockQueueConf2,
}

var mockToken = OECToken{
	"12345",
	[]OECMetadata{
		mockOECMetadata1,
		mockOECMetadata2,
	},
}

var mockTokenWithEmptyAssumeRoleResult = OECToken{
	"54321",
	[]OECMetadata{
		mockOECMetadataWithEmptyAssumeRoleResult1,
		mockOECMetadataWithEmptyAssumeRoleResult2,
	},
}

var mockEmptyToken = OECToken{}
