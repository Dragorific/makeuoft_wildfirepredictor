package elasticsearch

import (
	"github.com/dragorific/makeuoft_wildfirepredictor/setup"
	"github.com/olivere/elastic/v7"
)

//ExistsByID looks through elasticsearch to see if the document with the following ID exists
func ExistsByID(s *setup.State, index string, searchTerm string) bool {
	client, ctx := s.Elastic, s.Ctx
	termQuery := elastic.NewTermQuery("_id", searchTerm)
	result, err := client.Search().Index(index).Query(termQuery).Do(ctx)
	if err != nil {
		s.Log.Error("error searching for document ", err)
	}
	if result.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}
