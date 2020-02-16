package elasticsearch

import (
	"github.com/dragorific/makeuoft_wildfirepredictor/setup"
	"github.com/olivere/elastic/v7"
)

//GetDocumentByID returns the json value of a document at the given ID
func GetDocumentByID(s *setup.State, index string, searchTerm string) ([]byte, error) {
	client, ctx := s.Elastic, s.Ctx
	termQuery := elastic.NewTermQuery("_id", searchTerm)
	result, err := client.Search().Index(index).Query(termQuery).Do(ctx)
	return result.Hits.Hits[0].Source, err
}
