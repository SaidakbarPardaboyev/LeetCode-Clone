package pkg

import (
	"fmt"
	model "leetcode/models"
	"strings"
)

func GetAllDefaultQueries() (string, string, []string, string, []string, []string, []string) {
	selectQuery := `select
						p.id as problem_id,
						case
							when 'Accepted' = ANY(array_agg(sub.submission_status)) then 'Accepted'
							when count(sub.submission_status) > 0 then 'Attempted'
							else 'NOT-TRIED'
						end as status,
						p.problem_number as problem_number,
						p.title as problem_title,
						a.acceptence as acceptence,
						p.difficulty as difficulty
					from
						problems as p`

	innerJoinQuery := `	left join
							submissions as sub
								on sub.problem_id = p.id and
								$1 = sub.user_id
						left join
							acceptenceRatesOfProblems as a
								on a.problem_id = p.id`

	whereQuery := []string{}
	groupByQuery := []string{"p.problem_number", "p.title", "a.acceptence", "p.id"}
	havingQuery := []string{}
	orderByQuery := []string{"p.problem_number"}

	withQuery := `with acceptenceRatesOfProblems as (
					select
						p.id as problem_id,
						round(count(case when submission_status = 'Accepted' then 1 end)::numeric / count(*) * 100, 2)as acceptence
					from
						problems as p
					inner join
						submissions as s
							on p.id = s.problem_id
					group by
						p.problem_number, p.id
					order by
						p.problem_number
				)`

	return withQuery, selectQuery, whereQuery, innerJoinQuery, groupByQuery, havingQuery, orderByQuery
}

func JoinQueryParts(withQuery string, selectQuery string, whereQuery []string,
	innerJoinQuery string, groupByQuery []string, havingQuery []string,
	orderByQuery []string) string {

	query := ""

	// adding additions table's queries
	query += withQuery

	// join select part
	query += selectQuery

	// add join part
	query += innerJoinQuery

	// add where part if exist
	if len(whereQuery) > 0 {
		query += " where " + strings.Join(whereQuery, " and ")
	}

	// add group by part
	if len(groupByQuery) > 0 {
		query += " group by " + strings.Join(groupByQuery, ", ")
	}

	// add having part
	if len(havingQuery) > 0 {
		query += " having " + strings.Join(havingQuery, ", ")
	}

	// add order by part
	if len(orderByQuery) > 0 {
		query += " order by " + strings.Join(orderByQuery, ", ")
	}

	return query
}

func FilterProblemsBySorting(filter *model.ProblemFilter, innerJoinQuery *string,
	groupByQuery *[]string, orderByQuery *[]string) error {
	// sort by acceptance rate desc
	if *filter.Sorting == "W3sic29ydE9yZGVyIjoiREVTQ0VORElORyIsIm9yZGVyQnkiO"+
		"iJBQ19SQVRFIn1d" {
		*innerJoinQuery += ` inner join
								submissions as s
									on p.id = s.problem_id`
		*groupByQuery = append(*groupByQuery, "p.problem_number")
		*groupByQuery = append(*groupByQuery, "p.title")
		*orderByQuery = append([]string{"round(count(case when s.submission_" +
			"status = 'Accepted' then 1 end)::numeric / count(*) * 100, 2) asc"},
			*orderByQuery...)
		// sort by acceptance rate asc
	} else if *filter.Sorting == "W3sic29ydE9yZGVyIjoiQVNDRU5ESU5HIiwib3JkZX"+
		"JCeSI6IkFDX1JBVEUifV0%3D" {
		*innerJoinQuery += ` inner join
								submissions as s
									on p.id = s.problem_id`
		*groupByQuery = append(*groupByQuery, "p.problem_number")
		*groupByQuery = append(*groupByQuery, "p.title")
		*orderByQuery = append([]string{"round(count(case when s.submission_" +
			"status = 'Accepted' then 1 end)::numeric / count(*) * 100, 2) desc"},
			*orderByQuery...)
		// sort order by from hard to easy
	} else if *filter.Sorting == "W3sic29ydE9yZGVyIjoiREVTQ0VORElORyIsIm9yZG"+
		"VyQnkiOiJESUZGSUNVTFRZIn1d" {
		*orderByQuery = append([]string{"difficulty desc"}, *orderByQuery...)
		// sort order by from easy to hard
	} else if *filter.Sorting == "W3sic29ydE9yZGVyIjoiQVNDRU5ESU5HIiwib3JkZX"+
		"JCeSI6IkRJRkZJQ1VMVFkifV0%3D" {
		*orderByQuery = append([]string{"difficulty asc"}, *orderByQuery...)
		// sort by problem number asc
	} else if *filter.Sorting == "W3sic29ydE9yZGVyIjoiQVNDRU5ESU5HIiwib3JkZX"+
		"JCeSI6IkZST05URU5EX0lEIn1d" {
		*orderByQuery = append([]string{"problem_number asc"},
			(*orderByQuery)[:len(*orderByQuery)-1]...)
		// sort by problem number desc
	} else if *filter.Sorting == "W3sic29ydE9yZGVyIjoiREVTQ0VORElORyIsIm9yZG"+
		"VyQnkiOiJGUk9OVEVORF9JRCJ9XQ%3D%3D" {
		*orderByQuery = append([]string{"problem_number desc"},
			(*orderByQuery)[:len(*orderByQuery)-1]...)
		// not spesific / error sorting query
	} else if *filter.Sorting != "W3t9XQ%3D%3D" {
		return fmt.Errorf("error with giving query for sorting filter")
	}
	return nil
}

func FilterProblemsByStatus(filter *model.ProblemFilter, withQuery *string,
	whereQuery *[]string, params *[]interface{}, username string,
	paramCount *int) error {

	if *filter.Status == "NOT_STARTED" {
		*withQuery += fmt.Sprintf(`,  AcceptedProblemsTitle as (
							select 
								distinct problems.title as problem_title 
							from 
								submissions
							inner join
								problems
									on problems.id = submissions.problem_id
							where
								user_id=$%d)`, *paramCount)
		*whereQuery = append(*whereQuery, ` p.title not in (
												select 
													problem_title 
												from 
													AcceptedProblemsTitle
											)`)
		*params = append(*params, username)
		*paramCount++
	} else if *filter.Status == "AC" {
		*withQuery += fmt.Sprintf(`,  AcceptedProblemsTitle as (
							select 
								distinct problems.title as problem_title 
							from 
								submissions
							inner join
								problems
									on problems.id = submissions.problem_id
							where
								user_id=$%d and 
								submission_status='Accepted'
						)`, *paramCount)
		*whereQuery = append(*whereQuery, ` p.title in (
												select 
													problem_title 
												from 
													AcceptedProblemsTitle
											)`)
		*params = append(*params, username)
		*paramCount++
	} else if *filter.Status == "TRIED" {
		*withQuery += fmt.Sprintf(`,  AcceptedProblemsTitle as (
								select 
									distinct problems.title as problem_title 
								from 
									submissions 
								inner join
									problems
										on problems.id = submissions.problem_id
								where
									user_id=$%d
								group by
									problems.title
								having
									not ('Accepted' = ANY(array_agg(submission_status)))
							)`, *paramCount)
		*whereQuery = append(*whereQuery, ` p.title in (
												select 
													distinct problem_title 
												from 
													AcceptedProblemsTitle
											)`)
		*params = append(*params, username)
		*paramCount++
	} else if *filter.Status != "W3t9XQ%3D%3D" {
		return fmt.Errorf("error with giving query for status filter")
	}
	return nil
}

func FilterProblemsByTopicsSlugs(filter *model.ProblemFilter, withQuery *string,
	innerJoinQuery *string, whereQuery *[]string, havingQuery *[]string, params *[]interface{},
	paramCount *int) {

	topics := strings.Split(*filter.TopicsSlugs, "%2C")

	dollars := ""
	for i := 0; i < len(topics); i++ {
		dollars += fmt.Sprintf("$%d", *paramCount) + ", "
		*params = append(*params, topics[i])
		*paramCount++
	}
	dollars = dollars[:len(dollars)-2]
	*withQuery += `,  topicsIds as (
							select
								id as topic_id
							from
								topics
							where
								name in (` + dollars + `)
						)`
	*innerJoinQuery += ` inner join
							topics_problems as werr
								on werr.problem_id = p.id`

	newWhere := "werr.topic_id in (select topic_id from topicsIds)"
	*whereQuery = append([]string{newWhere}, *whereQuery...)

	newHaving := fmt.Sprintf("count(distinct werr.topic_id) = %d", len(topics))
	*havingQuery = append(*havingQuery, newHaving)
}

func FilterProblemsBySearch(filter *model.ProblemFilter,
	whereQuery *[]string, params *[]interface{},
	paramCount *int, orderByQuery *[]string) {

	newWhere := fmt.Sprintf("search_vector @@ plainto_tsquery('english', $%d)", *paramCount)
	*whereQuery = append([]string{newWhere}, *whereQuery...)
	*params = append(*params, *filter.Search)
	*paramCount++

	newOrderBy := fmt.Sprintf("ts_rank(search_vector, plainto_tsquery('english', $%d)) DESC", *paramCount)
	*orderByQuery = append([]string{newOrderBy}, *orderByQuery...)
	*params = append(*params, *filter.Search)
	*paramCount++
}
