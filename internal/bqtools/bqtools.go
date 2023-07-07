package bqtools

import (
	awshelp "cio_objects/internal/aws_help"
	"context"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

func NewBQ() *BQ {
	bq := BQ{
		ctx: context.Background(),
	}
	client, err := bigquery.NewClient(bq.ctx, "ebac-287911", bq.getCreds())
	if err != nil {
		panic(err)
	}
	bq.client = client
	return &bq
}

type BQ struct {
	ctx    context.Context
	client *bigquery.Client
}

func (bq *BQ) Query(query string) (*bigquery.RowIterator, error) {
	q := bq.client.Query(query)
	return q.Read(bq.ctx)
}

func (bq *BQ) Close() {
	bq.client.Close()
}

func (bq *BQ) getCreds() option.ClientOption {
	secret := awshelp.GetSecret("google-big-query-client-credentials")
	data, ok := secret["json"]
	if !ok {
		panic("troubles getting secret")
	}
	return option.WithCredentialsJSON([]byte(data))
}
