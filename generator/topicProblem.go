package generator

import (
	"database/sql"
	"leetcode/model"
	"leetcode/storage/postgres"
)

var leetcodeToTopic = map[string][]string{
	"f28e1b09-0a02-48d9-8299-a03c2d1239cb": {"eebc91d3-a74b-4a09-8471-13dea1fea60b"}, // Problem 1
	"063084ad-35c5-41d4-9296-e82252dd625a": {"41cff13b-985d-48e1-8ccd-b6f94bdd1516"}, // Problem 2
	"2a4913b1-f159-4a28-bc62-2c75661bdd23": {"488e46b2-0878-4d24-9cd5-e504624ce482"}, // Problem 3
	"32915be3-277d-46d6-85c7-06d3c6ed73a4": {"e6b1a6aa-d515-47e0-9b3f-0fbf62528342"}, // Problem 4
	"ac76b3f3-4738-4510-b90a-a9b9ae693d8b": {"de688496-fcab-4d20-803e-79f9b49f787d"}, // Problem 5
	"c81c3b88-6937-47cc-9a8f-32f195911209": {"8f442c42-fa5a-44c3-a39a-4a2b1abeb4b0"}, // Problem 6
	"f63436c1-164a-4df9-8c04-a71064fc7adf": {"e323486a-4cf7-4e88-8cc2-69a12a697ed7"}, // Problem 7
	"7ddb1fbf-72ef-4bab-9a10-5a891a4f64a8": {"e22400e6-647f-49f8-ab06-d04089a9fd1f"}, // Problem 8
	"25bcf03e-34d4-424d-9c02-14fb53548482": {"e9fcc81c-eb2e-414c-93db-d915bf422302"}, // Problem 9
	"f374d1ce-6a7c-420a-b0a2-8921568601bc": {"f897345e-5947-4c1c-bd2e-e9577a2b2e51"}, // Problem 10
	"f5878f4b-6e39-42d8-97bf-5a14e2503709": {"f366f62a-4f32-40a8-be48-39e007abc7c4"}, // Problem 11
	"b63d89da-0abd-432b-a88d-b7b880f2462c": {"936bbd7a-bef8-47f5-b1ab-b712ef788d40"}, // Problem 12
	"5b72e78a-0ac6-43a7-8db6-c42eec96589b": {"ddd8acc5-4826-4c44-be3b-a697a9fc1692"}, // Problem 13
	"390fe73e-d5ac-4cf2-90fc-2716199c3bf6": {"6c3c987c-e929-463b-beee-584b0adab7b6"}, // Problem 14
	"b67f41da-6045-4c4d-94cd-3c0b48a49010": {"d24fdaf7-5a14-4755-85a3-04c763643f7c"}, // Problem 15
	"5f969f7a-133a-4393-998d-505099756ca6": {"89d865af-ccef-45db-b06e-c7682edf4fae"}, // Problem 16
	"ff72554b-b8af-41a7-9f98-6c6e8b98a2a0": {"ad421f9d-d16a-40e5-a2de-c9a3f67bb9bf"}, // Problem 17
	"52262b0a-da1b-4068-b126-7a7f54b679f9": {"a41f6449-9eae-4b40-93be-1d4673bfe8d8"}, // Problem 18
	"32a812b3-043d-4bef-9209-4e6322c1009c": {"e117a7f9-1b08-48d7-aed1-d3d2dc912bc8"}, // Problem 19
	"f6469230-24bf-4eba-be30-209f2c91fb2b": {"9458acaa-21b5-40a2-971a-1c8e7a1a61a4"}, // Problem 20
	"2d920ae0-1c24-4185-8909-e97cc279c062": {"a49a313f-f4a5-4b10-8655-7a37fc9a8991"}, // Problem 21
	"4fbfc094-9208-49a3-8cd4-06df5f520b9d": {"3ba64ef0-1cc2-4d24-8e67-b495a81930c3"}, // Problem 22
	"4fd8b5dd-8523-49e8-baeb-db58bda6dc74": {"838c70f7-b69e-49c8-9991-270e7f84f287"}, // Problem 23
	"c593252e-b0ec-49f6-90ad-14500be81d2f": {"89f3a476-6d04-4bc7-a6d7-0883f06cff50"}, // Problem 24
	"c17d12e4-5715-48b2-8c87-f6f6f6a19517": {"e6b1a6aa-d515-47e0-9b3f-0fbf62528342", "489b842f-2253-4a9f-8e4e-13e9e512eefd"}, // Problem 25
	"243cbadd-c430-4d3e-aadf-faad1ce5efa5": {"489b842f-2253-4a9f-8e4e-13e9e512eefd"}, // Problem 35
	"31c0640c-5c01-44b6-8c63-9a1b3e3b4f36": {"f366f62a-4f32-40a8-be48-39e007abc7c4"}, // Problem 42
	"35b972d3-0f3e-4033-9e87-89414a2d055a": {"489b842f-2253-4a9f-8e4e-13e9e512eefd"}, // Problem 26
	"ca9c8452-b6a2-485d-a7a9-81f31a690d0f": {"488e46b2-0878-4d24-9cd5-e504624ce482"}, // Problem 27
	"9b33f48f-5907-4329-bb87-b6806831835a": {"f897345e-5947-4c1c-bd2e-e9577a2b2e51"}, // Problem 28
	"2c276238-adf7-487c-8629-cacff7d538cb": {"936bbd7a-bef8-47f5-b1ab-b712ef788d40"}, // Problem 29
	"c529b2d3-e926-4542-be55-48cff766e17a": {"489b842f-2253-4a9f-8e4e-13e9e512eefd"}, // Problem 30
	"fe366593-dd4b-4dda-94e3-73fe08cb27af": {"6c3c987c-e929-463b-beee-584b0adab7b6"}, // Problem 31
	"c316a90d-5f6c-4f39-9263-663f27f75de5": {"d24fdaf7-5a14-4755-85a3-04c763643f7c"}, // Problem 32
	"15946a7b-07b1-4583-a1f3-e98e7917ec14": {"a49a313f-f4a5-4b10-8655-7a37fc9a8991"}, // Problem 33
	"3043420c-de59-47f8-9e66-3d4868df2e67": {"f366f62a-4f32-40a8-be48-39e007abc7c4"}, // Problem 34
	"4677c4c7-3eff-4e39-bf13-c8847c988042": {"41a27fcc-94f9-46a9-9264-e33661e2302a"}, // Problem 43
	"d6c82754-627b-4ebd-acaf-d6a89b558d03": {"5dff77b8-7216-43df-8250-2053a2cc9467"}, // Problem 44
	"494555ce-81f4-4d8a-8a9d-5b8c7c53646c": {"1298bded-ed28-46eb-8e68-be43b4b50556"}, // Problem 45
	"94df6b20-da96-4c0c-b09a-e7d21a2c610f": {"bc790005-2000-4872-bba8-72a6f63205a6"}, // Problem 46
	"2e5c0693-63bb-44a2-b782-eb49b6b73617": {"daae8d80-8184-49ed-906e-84eb376bd518"}, // Problem 47
	"26af5a5d-d976-4940-abbd-091be5340c6e": {"5dff77b8-7216-43df-8250-2053a2cc9467"}, // Problem 48
	"0b46934d-8328-4aa8-afb9-d9f2a8dc28b8": {"f897345e-5947-4c1c-bd2e-e9577a2b2e51"}, // Problem 49
	"96da33a1-379e-4081-87dd-be23f126d768": {"a41f6449-9eae-4b40-93be-1d4673bfe8d8"}, // Problem 50
	"6b98e203-e766-421c-b29a-6288adff31f1": {"41a27fcc-94f9-46a9-9264-e33661e2302a"}, // Problem 51
	"0e64337c-e2eb-4395-a172-a5e1d36d5e25": {"3ba64ef0-1cc2-4d24-8e67-b495a81930c3"}, // Problem 52
	"b861f053-2fdb-4061-990a-5eb605e9002c": {"f366f62a-4f32-40a8-be48-39e007abc7c4"}, // Problem 53
	"fed94de7-0c0a-4a25-94b2-09bdda50cf4e": {"41cff13b-985d-48e1-8ccd-b6f94bdd1516"}, // Problem 54
	"307b5177-580f-415b-b03a-c71d67fcf14c": {"a41f6449-9eae-4b40-93be-1d4673bfe8d8"}, // Problem 55
	"10b89afb-1b6b-46c4-95f5-55c93fd5ca4d": {"d24fdaf7-5a14-4755-85a3-04c763643f7c"}, // Problem 56
	"fdda2423-fb58-4500-8116-5432c4bd9206": {"41a27fcc-94f9-46a9-9264-e33661e2302a"}, // Problem 57
	"da1545ff-92f2-4124-8d6b-bb5f87d16a85": {"89f3a476-6d04-4bc7-a6d7-0883f06cff50"}, // Problem 58
	"7a6e75ec-f04b-4459-932e-d57d8d9342e8": {"e117a7f9-1b08-48d7-aed1-d3d2dc912bc8"}, // Problem 59
	"2c6a6d1b-f1d8-4466-b1cb-dc042b3f3da9": {"e117a7f9-1b08-48d7-aed1-d3d2dc912bc8"}, // Problem 60
}

func InsertTopicProblems(db *sql.DB) {

	tp := postgres.NewTopicProblemRepo(db)
	for provlemId, topics := range leetcodeToTopic {
		for _, topicId := range topics{
			topicProblem := model.TopicProblem{TopicId: topicId, ProblemId: provlemId}
			tp.CreateTopicProblem(topicProblem)
		}
	}
}

